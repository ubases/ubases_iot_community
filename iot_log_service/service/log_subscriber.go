package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_log_service/config"
	"cloud_platform/iot_log_service/rpc/rpcclient"
	models "cloud_platform/iot_model/ch_log/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type LogSubscriber struct {
	suber      *jetstream.JSPullSubscriber
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewBuildSubscriber() (*LogSubscriber, error) {
	suber, err := jetstream.NewJSPullSubscriber("iot_log_service", iotconst.NATS_STREAM_APP, iotconst.NATS_SUBJECT_RECORDS, connerrhandler, config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	Concurrent := 1
	return &LogSubscriber{suber, Concurrent, ctx, cancel}, nil
}

func connerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
	}
}

func (bs LogSubscriber) Run() {
	// 从nats消息队列拉取数据，将日志写入clickhouse数据库
	for {
		if bs.ctx.Err() != nil {
			break
		}
		msgList, err := bs.suber.FetchMessageEx(100)
		if err != nil {
			if errors.Is(err, nats.ErrConnectionClosed) {
				iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
				time.Sleep(3 * time.Second)
			} else if !errors.Is(err, nats.ErrTimeout) {
				iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
			}
			continue
		}
		alr := []models.AppLogRecords{}
		for _, v := range msgList {
			switch v.Subject {
			case iotconst.NATS_SUBJECT_RECORDS:
				al := models.AppLogRecords{}
				err = json.Unmarshal(v.Data, &al)
				if err != nil {
					iotlogger.LogHelper.Errorf("解析json信息失败,内容[%s],错误:%s", string(v.Data), err.Error())
					continue
				}
				alr = append(alr, al)
				// 注册接口，异步聚合用户和app数据到t_iot_log_app_user
				if al.EventName == iotconst.APP_EVENT_REGISTER {
					aggregationUserData(al.Account, al.AppKey, al.TenantId, al.EventName, al.RegionServerId)
				}
				// 登录接口异步查询是否存在用户聚合数据，没有则重新聚合用户数据到t_iot_log_app_user， 为以前注册过的用户做兼容，并更新最后登录时间
				if al.EventName == iotconst.APP_EVENT_LOGIN {
					aggregationUserData(al.Account, al.AppKey, al.TenantId, al.EventName, al.RegionServerId)
				}
				// 注销账号，删除t_iot_log_app_user和t_iot_log_app_records信息
				if al.EventName == iotconst.APP_URI_CANCEL_ACCOUNT {
					aggregationUserData(al.Account, al.AppKey, al.TenantId, al.EventName, al.RegionServerId)
				}
				iotlogger.LogHelper.Debugf("接收APP日志记录:%v", al)
			}
		}
		if len(alr) != 0 {
			if err := CreateAppLogRecords(alr); err != nil {
				iotlogger.LogHelper.Errorf("批量插入APP日志记录失败: %v", err)
				continue
			}
		}
	}
}

func aggregationUserData(account, appKey, tenanntId, eventName string, regionServerId int64) {
	switch eventName {
	case iotconst.APP_EVENT_REGISTER:
		// 先查询用户信息存不存在, 若不存在则创建，若存在，则更新login_time
		userInfo := &models.AppLogUser{
			Account:        account,
			AppKey:         appKey,
			TenantId:       tenanntId,
			RegionServerId: regionServerId,
		}
		appUser1, err := GetAppLogUser(userInfo)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			iotlogger.LogHelper.Errorf("查询app log user信息失败: %v", err)
			return
		}
		if appUser1.Id == 0 {
			appUser, err := newAppLogUser(account, appKey, tenanntId, regionServerId)
			if err != nil {
				iotlogger.LogHelper.Errorf("聚合app log user信息失败: %v", err)
			}
			if err := CreateAppLogUser(appUser); err != nil {
				iotlogger.LogHelper.Errorf("创建app log user信息失败: %v", err)
				return
			}
		}
	case iotconst.APP_EVENT_LOGIN:
		// 先查询用户信息存不存在, 若不存在则创建，若存在，则更新login_time
		userInfo := &models.AppLogUser{
			Account:        account,
			AppKey:         appKey,
			TenantId:       tenanntId,
			RegionServerId: regionServerId,
		}
		appUser1, err := GetAppLogUser(userInfo)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			iotlogger.LogHelper.Errorf("查询app log user信息失败: %v", err)
			return
		}
		iotlogger.LogHelper.Helper.Debugf("get user Info: %v", appUser1)
		appUser, err := newAppLogUser(account, appKey, tenanntId, regionServerId)
		if err != nil {
			iotlogger.LogHelper.Errorf("聚合app log user信息失败: %v", err)
			return
		}
		if appUser1.Id == 0 {
			if err := CreateAppLogUser(appUser); err != nil {
				iotlogger.LogHelper.Errorf("创建app log user信息失败: %v", err)
				return
			}
		} else {
			if err := UpdateAppLogUser(appUser); err != nil {
				iotlogger.LogHelper.Errorf("更新app log user信息失败: %v", err)
				return
			}
		}
	case iotconst.APP_EVENT_CANCEL_ACCOUNT:
		// // 删除用户有待考虑，让数据自动过期？ 要考虑性能，及用户注销后，用户操作的追踪
		// userInfo := &AppLogUser{Account: account}
		// if err := DeleteAppLogUser(userInfo); err != nil {
		// 	iotlogger.LogHelper.Errorf("删除app log user信息失败: %v", err)
		// 	return
		// }
		// // 删除记录有待考虑，让数据自动过期? 要考虑性能，及用户注销后，用户操作的追踪
		// recordsInfo := &AppLogRecords{Account: account}
		// if err := DeleteAppLogRecords(recordsInfo); err != nil {
		// 	iotlogger.LogHelper.Errorf("删除app log records信息失败: %v", err)
		// 	return
		// }
	}
}

func newAppLogUser(account, appKey, tenanntId string, regionServerId int64) (*models.AppLogUser, error) {
	// 查询用户信息
	reqUser := &protosService.UcUserFilter{
		AppKey:   appKey,
		TenantId: tenanntId,
		Status:   1,
	}
	if strings.Contains(account, "@") {
		reqUser.Email = account
	} else {
		reqUser.Phone = account
	}
	respUser, err := rpcclient.ClientUcUserService.Find(context.Background(), reqUser)
	if err != nil {
		return nil, err
	}
	if respUser.Code != 200 {
		return nil, errors.New(respUser.Message)
	}
	if respUser.Data == nil || len(respUser.Data) == 0 {
		return nil, errors.New("user data is nil or length equal 0")
	}
	iotlogger.LogHelper.Helper.Debugf("account %s, appKey: %s, tenantId: %s, region: %s", account, appKey, tenanntId, respUser.Data[0].RegisterRegion)
	// 根据appkey查询app信息
	reqApp := &protosService.OemAppFilter{
		AppKey: respUser.Data[0].AppKey,
	}
	respApp, err := rpcclient.ClientOemAppService.Find(context.Background(), reqApp)
	if err != nil {
		return nil, err
	}
	if respApp.Code != 200 {
		return nil, errors.New(respApp.Message)
	}
	if respApp.Data == nil || len(respApp.Data) == 0 {
		return nil, errors.New("app data is nil or length equal 0")
	}
	appUser := &models.AppLogUser{
		Id:             respUser.Data[0].Id,
		Account:        account,
		AppKey:         appKey,
		TenantId:       tenanntId,
		AppName:        respApp.Data[0].Name,
		Region:         respUser.Data[0].RegisterRegion,
		RegionServerId: regionServerId,
		LoginTime:      time.Now(),
		CreatedAt:      time.Now(),
	}
	return appUser, nil
}

func (bs LogSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}
