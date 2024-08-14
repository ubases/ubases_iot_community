package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmCommunityProductService struct {
	Ctx context.Context
}

func (s OpmCommunityProductService) SetContext(ctx context.Context) OpmCommunityProductService {
	s.Ctx = ctx
	return s
}

// 详细
func (s OpmCommunityProductService) GetOpmCommunityProductDetail(id string) (*entitys.OpmCommunityProductEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmCommunityProductService.FindById(s.Ctx, &protosService.OpmCommunityProductFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]
	return entitys.OpmCommunityProduct_pb2e(data), err
}

// QueryOpmCommunityProductList 列表
func (s OpmCommunityProductService) QueryOpmCommunityProductList(filter entitys.OpmCommunityProductQuery) ([]*entitys.OpmCommunityProductEntitys, int64, error) {
	var queryObj = filter.OpmCommunityProductQuery_e2pb()
	rep, err := rpc.ClientOpmCommunityProductService.Lists(s.Ctx, &protosService.OpmCommunityProductListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderKey:  "sort",
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmCommunityProductEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmCommunityProduct_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmCommunityProduct 新增
func (s OpmCommunityProductService) AddOpmCommunityProduct(req entitys.OpmCommunityProductEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmCommunityProduct_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientOpmCommunityProductService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.ImageUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmCommunityProduct, iotutil.ToString(req.Id), req.ImageUrl)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改
func (s OpmCommunityProductService) UpdateOpmCommunityProduct(req entitys.OpmCommunityProductEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpmCommunityProductService.Update(s.Ctx, entitys.OpmCommunityProduct_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.ImageUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmCommunityProduct, iotutil.ToString(req.Id), req.ImageUrl)
	}
	return iotutil.ToString(req.Id), err
}

// 删除
func (s OpmCommunityProductService) DeleteOpmCommunityProduct(req entitys.OpmCommunityProductFilter) error {
	rep, err := rpc.ClientOpmCommunityProductService.DeleteById(s.Ctx, &protosService.OpmCommunityProduct{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// SetStatusOpmCommunityProduct 禁用/启用
func (s OpmCommunityProductService) SetStatusOpmCommunityProduct(req entitys.OpmCommunityProductFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientOpmCommunityProductService.UpdateStatus(context.Background(), &protosService.OpmCommunityProduct{
		Id:     iotutil.ToInt64(req.Id),
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
