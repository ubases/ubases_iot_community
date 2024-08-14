package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductManualService struct {
	Ctx context.Context
}

func (s ProductManualService) SetContext(ctx context.Context) ProductManualService {
	s.Ctx = ctx
	return s
}

func (s ProductManualService) CreateProductManual(obj *entitys.OpmProductManualEntitys) error {
	obj.TenantId, _ = metadata.Get(s.Ctx, "tenantId")
	req := entitys.OpmProductManual_e2pb(obj)
	req.Id = iotutil.GetNextSeqInt64()
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err := rpc.ClientOpmProductManualService.Create(s.Ctx, req)
	if err != nil {
		return err
	}
	if req.FileUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmProductManual, iotutil.ToString(req.Id), req.FileUrl)
	}
	return nil
}

func (s ProductManualService) UpdateProductManual(obj *entitys.OpmProductManualEntitys) error {
	req := entitys.OpmProductManual_e2pb(obj)
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err := rpc.ClientOpmProductManualService.Update(s.Ctx, req)
	if err != nil {
		return err
	}
	if req.FileUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmProductManual, iotutil.ToString(req.Id), req.FileUrl)
	}
	return nil
}

func (s ProductManualService) DeleteProductManual(obj *entitys.OpmProductManualEntitys) error {
	_, err := rpc.ClientOpmProductManualService.DeleteById(s.Ctx, &protosService.OpmProductManual{
		Id: obj.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s ProductManualService) GetProductManual(productKey string) (*entitys.OpmProductManualEntitys, error) {
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

func (s ProductManualService) GetProductManualList(obj *entitys.OpmProductManualQuery) ([]*entitys.OpmProductManualEntitys, int64, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	req := &protosService.OpmProductManualListRequest{
		Page:     int64(obj.Page),
		PageSize: int64(obj.Limit),
		Query: &protosService.OpmProductManual{
			TenantId: tenantId,
		},
		OrderKey:  "updated_at",
		OrderDesc: "desc",
	}
	resp, err := rpc.ClientOpmProductManualService.Lists(s.Ctx, req)
	if err != nil {
		return nil, 0, err
	}
	items := []*entitys.OpmProductManualEntitys{}
	for i := range resp.Data {
		item := entitys.OpmProductManual_pb2e(resp.Data[i])
		items = append(items, item)
	}
	return items, resp.Total, nil
}
