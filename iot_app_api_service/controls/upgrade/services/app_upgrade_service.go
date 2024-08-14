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
func (s AppUpgradeService) GetLatestApp(userId int64, appKey string, lang string, os int32) (rets *entitys.AppUpgradeVo, err error) {
	ret, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	theAppVer := ret.Data[0].Version
	switch os {
	case 1:
		theAppVer = ret.Data[0].IosVersion
	case 2:
		theAppVer = ret.Data[0].AndroidInterVersion
	case 3:
		theAppVer = ret.Data[0].AndroidOuterVersion
	}
	res := &entitys.AppUpgradeVo{
		Id:      ret.Data[0].Id,
		Name:    ret.Data[0].Name,
		Version: iotutil.IfStringEmpty(theAppVer, ret.Data[0].Version),
	}

	//获取APP升级版本信息
	resp, err := rpc.ClientOemAppCustomRecordService.GetLastVersion(context.Background(), &protosService.OemAppCustomRecordFilter{
		AppId: ret.Data[0].Id, Os: os, Status: 1,
	})
	if err == nil && len(resp.Data) > 0 {
		res.UpgradePrompt = iotutil.IfString(lang == "en", resp.Data[0].RemindDescEn, resp.Data[0].RemindDesc)
		remindMode := int(resp.Data[0].RemindMode)
		switch remindMode {
		case 1:
			res.UpgradeMode = 0
		case 2:
			res.UpgradeMode = 2
		case 3:
			res.UpgradeMode = 1
		}
		//res.Url = resp.Data[0].PkgUrl
		funcCfg, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
			AppId:   ret.Data[0].Id,
			Version: ret.Data[0].Version,
		})
		if err != nil {
			return nil, err
		}
		if funcCfg.Code != 200 && funcCfg.Message != "record not found" {
			return nil, errors.New(funcCfg.Message)
		}
		var autoUpgradeUrl entitys.OemAppAutoUpgrade
		iotutil.JsonToStruct(funcCfg.Data[0].AutoUpgrade, &autoUpgradeUrl)
		switch os {
		case 1:
			res.Url = autoUpgradeUrl.IosAddr
		case 2:
			res.Url = autoUpgradeUrl.AndroidInterAddr
		case 3:
			res.Url = autoUpgradeUrl.AndroidOuterAddr
		default:
			res.Url = ""
			res.UpgradeMode = 0
		}
	}

	//协议更新情况
	userRes, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id:     userId,
		AppKey: appKey,
	})
	if err == nil {
		res.AgreementRemind = userRes.Data[0].AgreementFlag
	}
	return res, nil
}

// SetAgreementAgree 设置用户协议同意
func (s AppUpgradeService) SetAgreementAgree(userId int64) error {
	ret, err := rpc.TUcUserService.UpdateFields(context.Background(), &protosService.UcUserUpdateFieldsRequest{
		Fields: []string{"agreement_flag"},
		Data: &protosService.UcUser{
			Id: userId, AgreementFlag: 2,
		},
	})
	if err != nil {
		return err
	}
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return nil
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
