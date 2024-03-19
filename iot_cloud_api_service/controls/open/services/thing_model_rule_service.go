package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmThingModelRuleService struct {
	Ctx context.Context
}

func (s OpmThingModelRuleService) SetContext(ctx context.Context) OpmThingModelRuleService {
	s.Ctx = ctx
	return s
}

// 详细
func (s OpmThingModelRuleService) GetOpmThingModelRuleDetail(id string) (*entitys.OpmThingModelRuleEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmThingModelRuleService.FindById(s.Ctx, &protosService.OpmThingModelRuleFilter{Id: rid})
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
	return entitys.OpmThingModelRule_pb2e(data), err
}

// QueryOpmThingModelRuleList 列表
func (s OpmThingModelRuleService) QueryOpmThingModelRuleList(filter entitys.OpmThingModelRuleQuery) ([]*entitys.OpmThingModelRuleEntitys, int64, error) {
	var queryObj = filter.OpmThingModelRuleQuery_e2pb()
	rep, err := rpc.ClientOpmThingModelRuleService.Lists(s.Ctx, &protosService.OpmThingModelRuleListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderDesc: "desc",
		OrderKey:  "updated_at",
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmThingModelRuleEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmThingModelRule_pb2e(item))
	}
	return resultList, rep.Total, err
}

func (s OpmThingModelRuleService) QueryOpmThingModelRuleListByProductId(productId string, dataOrigin int32) ([]*entitys.OpmThingModelRuleEntitys, int64, error) {
	rep, err := rpc.ClientOpmThingModelRuleService.Lists(s.Ctx, &protosService.OpmThingModelRuleListRequest{
		Query: &protosService.OpmThingModelRule{ProductId: productId, DataOrigin: dataOrigin},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmThingModelRuleEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmThingModelRule_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmThingModelRule 新增
func (s OpmThingModelRuleService) AddOpmThingModelRule(req entitys.OpmThingModelRuleEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmThingModelRule_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientOpmThingModelRuleService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改
func (s OpmThingModelRuleService) UpdateOpmThingModelRule(req entitys.OpmThingModelRuleEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpmThingModelRuleService.UpdateAll(s.Ctx, entitys.OpmThingModelRule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除
func (s OpmThingModelRuleService) DeleteOpmThingModelRule(req entitys.OpmThingModelRuleFilter) error {
	rep, err := rpc.ClientOpmThingModelRuleService.Delete(s.Ctx, &protosService.OpmThingModelRule{
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

// SetStatusOpmThingModelRule 禁用/启用
func (s OpmThingModelRuleService) SetStatusOpmThingModelRule(req entitys.OpmThingModelRuleFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientOpmThingModelRuleService.UpdateStatus(context.Background(), &protosService.OpmThingModelRule{
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
