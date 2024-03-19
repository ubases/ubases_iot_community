package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/app/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
	"strings"
	"time"
)

type UserController struct {
}

var Usercontroller UserController

// @Summary 区域列表
// @Description
// @Tags 注销
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/regionList [get]
func (UserController) RegionList(c *gin.Context) {
	lang := c.GetHeader("lang")
	ip := c.ClientIP()
	resp, msg := RegionList(lang, ip)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

func RegionList(lang, ip string) ([]*entitys.SysRegionServerEntitysList, error) {
	//geo, err := iotutil.Geoip(ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode)
	geo, err := controls.Geoip(ip)
	country := geo.Country
	result := make([]*entitys.SysRegionServerEntitysList, 0)
	//增加状态判断查询（用于启用关闭区域操作）
	rep, err := rpc.SysRegionServerService.Lists(context.Background(), &protosService.SysRegionServerListRequest{
		Query: &protosService.SysRegionServer{Enabled: 1},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, err
	}
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_CONFIG_REGION_SERVER).Result()
	if err != nil {
		langMap = make(map[string]string)
		//异常不处理
		iotlogger.LogHelper.Errorf("字典缓存读取异常, %s", err.Error())
	}
	haveDefaultData := false
	for _, v := range rep.Data {
		item := entitys.SysRegionServer_pb2e(v, lang)
		fKey := fmt.Sprintf("%s_%d_name", lang, item.Id)
		if country != "" && v.Describe == country {
			item.IsDefault = 1
			haveDefaultData = true
		}
		item.Describe = iotutil.MapGetStringVal(langMap[fKey], item.Describe)
		result = append(result, item)
	}
	if !haveDefaultData {
		result[0].IsDefault = 1
	}
	return result, nil
}

// @Summary 注销账号
// @Description
// @Tags 注销
// @Accept application/json
// @Param data body entitys.AppCancelAccount true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/appCancelAccount [post]
func (UserController) CancelAccount(c *gin.Context) {
	bm := entitys.AppCancelAccount{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if strings.TrimSpace(bm.Account) == "" || strings.TrimSpace(bm.Code) == "" || strings.TrimSpace(bm.AppKey) == "" || strings.TrimSpace(bm.TenantId) == "" {
		iotgin.ResErrParams(c)
		return
	}
	appKey := bm.AppKey           // controls.GetAppKey(c)
	tenantId := bm.TenantId       // controls.GetTenantId(c)
	regionServerId := bm.RegionId // controls.GetRegionInt(c)
	ctx := controls.WithUserContextV2(c, 0, regionServerId, appKey, tenantId)
	code, msg := CancelAccount(ctx, true, bm, appKey, tenantId, regionServerId, c.ClientIP())
	if code != 0 {
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// CancelAccount 注销用户操作
func CancelAccount(ctx context.Context, isVerifyCode bool, cancelAccountParam entitys.AppCancelAccount, appKey, tenantId string, regionServerId int64, ip string) (code int, msg string) {
	account := cancelAccountParam.Account
	var err error
	if isVerifyCode {
		verificationCode := cancelAccountParam.Code
		_, _, err := iotutil.CheckUserName(account)
		if err != nil {
			return -1, err.Error()
		}
		resp := iotredis.GetClient().Get(ctx, cached.APP+"_"+appKey+"_"+account+"_4")
		if resp.Val() != verificationCode {
			iotlogger.LogHelper.Errorf("验证码有误[cached:%s-req:%s]", resp.Val(), verificationCode)
			return -1, "验证码有误"
		}
	}
	userInfo, err := rpc.UcUserService.Lists(ctx, &protosService.UcUserListRequest{
		Query: &protosService.UcUser{
			//Id:             userId,
			UserName:       account,
			Status:         1,
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("QueryQuestionTypeList error,%s", err.Error())
		return -1, err.Error()
	}
	if userInfo.Code != 200 || userInfo.Data == nil {
		return ioterrs.ERROR_NOT_BELONG_TO_USER.Code, ioterrs.ERROR_NOT_BELONG_TO_USER.Msg
	}
	userId := userInfo.Data[0].Id
	_, err = rpc.UcUserService.Update(ctx, &protosService.UcUser{
		Id:         userInfo.Data[0].Id,
		Status:     iotutil.ToInt32(iotconst.AccountAlreadyCancel),
		CancelTime: time.Now().Unix(),
	})
	if err != nil {
		return -1, "CancelAccount error"
	}
	//新增协程清理用户数据
	go func() {
		defer iotutil.PanicHandler()
		ucHomeUsersResponse, err := rpc.UcHomeUserService.Lists(ctx, &protosService.UcHomeUserListRequest{
			Query: &protosService.UcHomeUser{UserId: userId},
		})
		if err != nil {
			return
		}

		//geo, err := iotutil.Geoip(ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode) //根据ip获取位置信息
		geo, err := controls.Geoip(ip)
		if err != nil {
			logger.Errorf("get address by ip[%s], error:%s", ip, err.Error())
		}

		//删除家庭用户信息
		_, err = rpc.UcHomeUserService.Delete(context.Background(), &protosService.UcHomeUser{
			UserId: userId,
		})
		if err != nil {
			iotlogger.LogHelper.Error("CancelAccount error")
			return
		}

		//用户所在的家庭，并且当前用户为家庭所有者，则需要进行家庭数据清理
		for _, v := range ucHomeUsersResponse.Data {
			//owner-所有者, admin-管理者, member-成员
			//家庭所有者才进行数据清理
			if v.RoleType == iotconst.HOME_USER_ROLE_1 {
				_, err = rpc.UcHomeService.ChangeAllUserDefaultHomeId(ctx, &protosService.UcHome{
					Id:       v.HomeId,
					Lat:      geo.Lat,
					Lng:      geo.Lng,
					Country:  geo.Country,
					Province: geo.Province,
					City:     geo.City,
					District: geo.District,
				})
				if err != nil {
					iotlogger.LogHelper.Error("CancelAccount error")
					return
				}
				_, err = rpc.IotDeviceHomeService.Delete(context.Background(), &protosService.IotDeviceHome{
					HomeId: v.HomeId,
				})
				if err != nil {
					iotlogger.LogHelper.Error("CancelAccount error")
					return
				}

				_, err = rpc.IotDeviceGroupService.Delete(context.Background(), &protosService.IotDeviceGroup{
					HomeId: v.HomeId,
				})
			}
		}
		//todo 清空共享设备和群组信息
		_, err = rpc.IotDeviceSharedService.Delete(context.Background(), &protosService.IotDeviceShared{
			BelongUserId: userId,
		})

		//todo 家庭信息、设备信息、第三方登录相关信息需要删除
		_, err = rpc.ClientUcUserThirdService.Delete(context.Background(), &protosService.UcUserThird{
			UserId: userId,
		})
		if err != nil {
			iotlogger.LogHelper.Error("CancelAccount error")
			return
		}
	}()
	//清除token
	controls.ClearTokenByUserId(userId)
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+account+"_4")
	return 0, ""
}

// @Summary 发送验证码
// @Description
// @Tags 注销
// @Accept application/json
// @Param data body entitys.SendCodeParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/sendVerityCode [post]
func (UserController) SendVerityCode(c *gin.Context) {
	bm := entitys.SendCodeParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if strings.TrimSpace(bm.Account) == "" {
		iotgin.ResBusinessP(c, "账号为空")
		return
	}
	if strings.TrimSpace(bm.AppKey) == "" {
		iotgin.ResBusinessP(c, "AppKey为空")
		return
	}
	lang := c.Request.Header.Get("lang")
	if iotutil.IsEmail(bm.Account) {
		lang := c.GetHeader("lang")
		appKey := c.Request.Header.Get("appKey")
		// 发送邮件
		_, code, msg := SendEmailCode(bm.Account, 4, appKey, lang)
		if code != 0 {
			iotgin.ResBusinessP(c, msg)
			return
		}
	} else {
		var phone = bm.Account
		if strings.TrimSpace(bm.AreaPhoneNumber) == "" {
			iotgin.ResBusinessP(c, "区域手机号为空")
			return
		}
		if iotutil.CheckAllPhone(bm.AreaPhoneNumber, bm.Account) == false {
			iotgin.ResBusinessP(c, "手机号码不合法")
			return
		}
		var phoneType int32
		if bm.AreaPhoneNumber == "86" {
			phoneType = 1
		} else if bm.AreaPhoneNumber == "1" {
			phoneType = 3
			phone = bm.AreaPhoneNumber + bm.Account
		} else {
			iotgin.ResBusinessP(c, "手机号码不合法")
			return
		}
		// 发送验证码
		_, code, msg := SendSmsCode(lang, phone, bm.Account, phoneType, 4, bm.AppKey)
		if code != 0 {
			iotgin.ResBusinessP(c, msg)
			return
		}
	}
	iotgin.ResSuccessMsg(c)
}

// SendSmsCode 发送短信验证码
func SendSmsCode(lang, areaPhone, phone string, phoneType, smsType int32, appKey string) (smsCode string, code int, msg string) {
	smsCode = iotutil.Getcode()
	res := iotredis.GetClient().Set(context.Background(), cached.APP+"_"+appKey+"_"+phone+"_"+iotutil.ToString(smsType), smsCode, 600*time.Second) //有效期10分钟
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendVerityCode,缓存smsCodeInt失败:%s", res.Err().Error())
		return "", -1, res.Err().Error()
	}
	_, err := rpc.ClientSmsService.SendSMSVerifyCode(context.Background(), &protosService.SendSMSVerifyCodeRequest{
		PhoneNumber: areaPhone,
		UserName:    areaPhone,
		Code:        smsCode,
		Lang:        lang,
		TplType:     smsType,
		PhoneType:   phoneType,
	})
	if err != nil {
		iotlogger.LogHelper.Error("发送短信验证码失败，原因:%s", err.Error())
		return "", -1, err.Error()
	}
	return smsCode, 0, ""
}

func SendEmailCode(email string, emailType int32, appKey, lang string) (emailCode string, code int, msg string) {
	emailCode = iotutil.Getcode()
	iotlogger.LogHelper.Info(emailCode)
	res := iotredis.GetClient().Set(context.Background(), cached.APP+"_"+appKey+"_"+email+"_"+iotutil.ToString(emailType), emailCode, 600*time.Second) //有效期10分钟
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendEmail,缓存emailCode失败:%s", res.Err().Error())
		return "", -1, res.Err().Error()
	}
	_, err := rpc.ClientEmailService.SendEmailUserCode(context.Background(), &protosService.SendEmailUserCodeRequest{
		Email:   email,
		Code:    emailCode,
		Lang:    lang,
		TplType: emailType,
	})
	if err != nil {
		iotlogger.LogHelper.Error("发送邮件失败，原因:%s", err.Error())
		return "", -1, err.Error()
	}
	return emailCode, 0, ""
}

// @Summary 获取APP信息
// @Description
// @Tags 注销
// @Accept application/json
// @Param appKey query string true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/getAppInfo [post]
func (UserController) GetAppInfo(c *gin.Context) {
	appKey := c.Query("appKey")
	resp, err := QueryAppInfo(context.Background(), appKey)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

func QueryAppInfo(ctx context.Context, appKey string) (map[string]string, error) {
	res, err := rpc.ClientOemAppService.Find(ctx, &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	resMap := make(map[string]string)
	resMap["nameEn"] = res.Data[0].NameEn
	resMap["nameZh"] = res.Data[0].Name
	resIcon, _ := rpc.ClientOemAppUiConfigService.Find(context.Background(), &protosService.OemAppUiConfigFilter{
		AppId:   res.Data[0].Id,
		Version: res.Data[0].Version,
	})
	if resIcon != nil && len(resIcon.Data) > 0 {
		resMap["iconUrl"] = resIcon.Data[0].IconUrl
	}
	return resMap, nil
}
