package services

import (
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

type AppControlService struct {
}

func (s AppControlService) CheckControlPageVersion(productId, panelId, appPanelType string) (*entitys.CheckControlPageVersionResp, error) {
	var (
		result          entitys.CheckControlPageVersionResp
		panelIdInt      int64
		appPanelTypeInt int32
	)
	//TODO 此处加换成需要使用面板Id作为Key，否则不好清理
	//err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.CONTROL_PANEL_IS_UPDATE, iotutil.ToInt64(productId)), &result)
	//if err == nil {
	//	return &result, nil
	//}
	if panelId != "" {
		panelIdInt, _ = iotutil.ToInt64AndErr(panelId)
		appPanelTypeInt, _ = iotutil.ToInt32Err(appPanelType)
	}

	ProductId, _ := iotutil.ToInt64AndErr(productId)
	rep, err := rpc.ProductService.ControlPanelsUrlAndMd5(context.Background(), &protosService.ControlPanelsUrlAndMd5Request{
		ProductId:    ProductId,
		PanelId:      panelIdInt,
		AppPanelType: appPanelTypeInt,
	})
	if err != nil {
		return nil, err
	}
	result.Md5 = rep.ControlpageMd5
	result.Url = rep.Url
	//if err := cached.RedisStore.Set(persist.GetRedisKey(iotconst.CONTROL_PANEL_IS_UPDATE, iotutil.ToInt64(productId)), result, 600*time.Second); err != nil {
	//	return nil, err
	//}
	return &result, nil
}
