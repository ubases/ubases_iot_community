package services

import (
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	goerrors "go-micro.dev/v4/errors"
)

type ProductManualService struct {
	Ctx context.Context
}

func (s ProductManualService) SetContext(ctx context.Context) ProductManualService {
	s.Ctx = ctx
	return s
}

func (s ProductManualService) GetProductManual(productKey string) (*entitys.OpmProductManualEntitys, error) {
	rep, err := rpc.ClientOpmDocumentsService.Lists(s.Ctx, &protosService.OpmDocumentsListRequest{
		Query: &protosService.OpmDocuments{OriginKey: productKey, DocCodes: "product_manual_doc"},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, errors.New("未获取到产品说明书")
	}
	data := entitys.OpmProductManual_pb2eV2(rep.Data[0])
	return data, nil
}

// GetProductManualOld 原始方法，后来改位document
func (s ProductManualService) GetProductManualOld(productKey string) (*entitys.OpmProductManualEntitys, error) {
	req := &protosService.OpmProductManualFilter{
		ProductKey: productKey,
	}
	resp, err := rpc.ClientOpmProductManualService.Find(s.Ctx, req)
	if err != nil && goerrors.FromError(err).GetDetail() == ioterrs.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	data := entitys.OpmProductManual_pb2e(resp.Data[0])
	return data, nil
}
