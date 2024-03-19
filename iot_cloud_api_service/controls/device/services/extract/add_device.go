package extract

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
)

// CreateAndBindDeviceTriad  添加APP账号并设置虚拟设备
func CreateAndBindDeviceTriad(ctx context.Context, req entitys.AddAppAccountEntity, triadReq *entitys.GenerateDeviceTriad) error {
	//查询用户信息
	userRes, err := rpc.UcUserService.Find(ctx, &protosService.UcUserFilter{
		UserName:       req.Account,
		AppKey:         req.AppKey,
		RegionServerId: req.RegionServerId, //默认中国地区
	})
	if err != nil {
		return err
	}
	if userRes.Code != 200 {
		return errors.New("该账号尚未在APP中注册，err：" + userRes.Message)
	}
	user := userRes.Data[0]
	homeId, err := iotutil.ToInt64AndErr(user.DefaultHomeId)
	if err != nil {
		return err
	}
	//查询产品信息
	proRes, err := rpc.ClientOpmProductService.FindById(ctx, &protosService.OpmProductFilter{
		Id: req.ProductId,
	})
	if err != nil {
		return err
	}
	if proRes.Code != 200 {
		return errors.New(proRes.Message)
	}
	pro := proRes.Data[0]

	//面板预览的时候需要前置判断，是否已经创建过预览产品，如果创建过则不需要再次创建
	if req.CheckRepeat {
		dhRes, err := rpc.ClientIotDeviceHome.Find(ctx, &protosService.IotDeviceHomeFilter{
			HomeId:       homeId,
			AppPanelType: req.AppPanelType,
			PanelId:      req.PanelId,
		})
		if err == nil {
			if dhRes.Code == 200 {
				iotlogger.LogHelper.Infof("已有为用户创建预览产品，无需重复创建")
				return nil
			}
		}
	}

	//是否启用传入的设备名称
	deviceName := pro.Name
	if req.DeviceName != "" {
		deviceName = req.DeviceName
	}

	appInfo, err := GetAppInfo(ctx, req.AppKey)
	if err != nil {
		return err
	}
	generate := &protosService.IotDeviceTriadGenerateRequest{
		ProductId:  pro.Id,
		ProductKey: pro.ProductKey,
		UseType:    iotconst.Use_Type_Device_Real_Test,
		IsTest:     1,
		BindInfo: &protosService.BindTestAccountRequest{
			AppKey:       req.AppKey,
			UserId:       user.Id,
			ProductKey:   pro.ProductKey,
			DeviceId:     req.DeviceId,
			HomeId:       homeId,
			ProductId:    pro.Id,
			ProductName:  deviceName,
			UserAccount:  req.Account,
			AppName:      appInfo.Name,
			TenantId:     req.TenantId,
			PanelId:      req.PanelId,
			AppPanelType: req.AppPanelType,
		},
	}
	rep, err := rpc.ClientIotDeviceServer.CreateAndBindDeviceTriad(ctx, generate)
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	clearHomeCached(context.Background(), homeId)
	return nil
}

// 清理用户的缓存
func clearHomeCached(ctx context.Context, homeId int64) error {
	defer iotutil.PanicHandler("clearHomeCached", homeId)
	// 删除家庭详情缓存
	keys := []string{}
	resp, err := rpc.UcHomeUserService.Lists(ctx, &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: homeId,
		},
	})
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Message)
	}
	for _, u := range resp.Data {
		keys = append(keys,
			fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(u.UserId)),
			fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(u.HomeId), iotutil.ToString(u.UserId)))
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

// 清理用户的缓存
func SetUserDeviceCached(ctx context.Context, req entitys.GenerateDeviceTriad) error {
	defer iotutil.PanicHandler(req)
	//查询用户信息
	userRes, err := rpc.UcUserService.Find(ctx, &protosService.UcUserFilter{
		UserName:       req.UserAccount,
		TenantId:       req.TenantId,
		AppKey:         req.AppKey,
		RegionServerId: req.RegionServerId, //默认中国地区
	})
	if err != nil {
		return err
	}
	if userRes.Code != 200 {
		return errors.New("该账号尚未在APP中注册")
	}
	user := userRes.Data[0]
	homeId, err := iotutil.ToInt64AndErr(user.DefaultHomeId)
	if err != nil {
		return err
	}
	// 删除家庭详情缓存
	keys := []string{}
	resp, err := rpc.UcHomeUserService.Lists(ctx, &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: homeId,
		},
	})
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Message)
	}
	for _, u := range resp.Data {
		keys = append(keys,
			fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(u.UserId)),
			fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(u.HomeId), iotutil.ToString(u.UserId)))
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
