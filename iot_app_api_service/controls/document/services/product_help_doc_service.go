package services

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

type ProductHelpDocService struct {
	Ctx context.Context
}

func (s ProductHelpDocService) SetContext(ctx context.Context) ProductHelpDocService {
	s.Ctx = ctx
	return s
}

func (s ProductHelpDocService) GetProductHelpDoc(productKey, tenantId, lang string) ([]*entitys.ProductHelpDocEntitys, error) {
	result := []*entitys.ProductHelpDocEntitys{}
	req, err := rpc.ProductHelpDocService.Lists(s.Ctx, &protosService.ProductHelpDocListRequest{Query: &protosService.ProductHelpDoc{
		TenantId:   tenantId,
		ProductKey: productKey,
		Lang:       lang,
		Status:     1,
	}})
	if err != nil {
		return result, err
	}
	if len(req.Data) == 0 {
		return result, nil
	}
	for _, productHelpDoc := range req.Data {
		result = append(result, entitys.ProductHelpDoc_pb2e(productHelpDoc))
	}
	return result, err
}
