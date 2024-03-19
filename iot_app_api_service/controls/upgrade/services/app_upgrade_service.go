package services

import (
	"cloud_platform/iot_app_api_service/controls/upgrade/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type AppUpgradeService struct {
}

// GetLatestApp list  data
func (s AppUpgradeService) GetLatestApp(appKey string) (rets *entitys.AppUpgradeVo, err error) {
	ret, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return &entitys.AppUpgradeVo{
		Id:      ret.Data[0].Id,
		Name:    ret.Data[0].Name,
		Version: iotutil.IfStringEmpty(ret.Data[0].AndroidInterVersion, ret.Data[0].Version),
	}, nil
}

// 获取APP自动更新功能配置
func (s AppUpgradeService) GetFunctionConfigAutoUpgrade(req entitys.OemAppCommonReq) (*entitys.OemAppAutoUpgradeServiceRes, error) {
	ret, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: req.AppKey,
	})

	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	res, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
		AppId:   ret.Data[0].Id,
		Version: ret.Data[0].Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}

	var data entitys.OemAppAutoUpgradeServiceRes
	if len(res.Data) != 0 {
		if res.Data[0].AutoUpgrade != "" {
			iotutil.JsonToStruct(res.Data[0].AutoUpgrade, &data.AutoUpgrade)
		}
		data.Id = iotutil.ToString(res.Data[0].Id)
	}
	return &data, err
}
