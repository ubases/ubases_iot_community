package services

import (
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type TplDocumentTemplateService struct {
}

// GetTplDocumentTemplateDetail 测试用例模板详细
func (s TplDocumentTemplateService) GetTplDocumentTemplateDetail(id string) (*entitys.TplDocumentTemplateEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientDocumentTemplateService.FindById(context.Background(), &protosService.TplDocumentTemplateFilter{Id: rid})
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
	return entitys.TplDocumentTemplate_pb2e(data), err
}

// QueryTplDocumentTemplateList 测试用例模板列表
func (s TplDocumentTemplateService) QueryTplDocumentTemplateList(filter entitys.TplDocumentTemplateQuery) ([]*entitys.TplDocumentTemplateEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}

	// 状态字段查询的方案
	var status int32 = -1
	if filter.Query == nil {
		filter.Query = new(entitys.TplDocumentTemplateFilter)
	}
	if filter.Query.Status != nil {
		status = *filter.Query.Status
	}
	rep, err := rpc.ClientDocumentTemplateService.Lists(context.Background(), &protosService.TplDocumentTemplateListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query: &protosService.TplDocumentTemplate{
			Status:  status,
			TplName: filter.Query.TplName,
			Lang:    filter.Query.Lang,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.TplDocumentTemplateEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.TplDocumentTemplate_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddTplDocumentTemplate 新增测试用例模板
func (s TplDocumentTemplateService) AddTplDocumentTemplate(req entitys.TplDocumentTemplateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}

	saveObj := entitys.TplDocumentTemplate_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientDocumentTemplateService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateTplDocumentTemplate 修改测试用例模板
func (s TplDocumentTemplateService) UpdateTplDocumentTemplate(req entitys.TplDocumentTemplateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}

	res, err := rpc.ClientDocumentTemplateService.Update(context.Background(), entitys.TplDocumentTemplate_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteTplDocumentTemplate 删除测试用例模板
func (s TplDocumentTemplateService) DeleteTplDocumentTemplate(req entitys.TplDocumentTemplateFilter) error {
	rep, err := rpc.ClientDocumentTemplateService.Delete(context.Background(), &protosService.TplDocumentTemplate{
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

// SetStatusTplDocumentTemplate 禁用/启用测试用例模板
func (s TplDocumentTemplateService) SetStatusTplDocumentTemplate(req entitys.TplDocumentTemplateFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	if req.Status == nil {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientDocumentTemplateService.UpdateFields(context.Background(), &protosService.TplDocumentTemplateUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.TplDocumentTemplate{
			Id:     iotutil.ToInt64(req.Id),
			Status: *req.Status,
		},
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
