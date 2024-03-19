package services

import (
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type TplTestcaseTemplateService struct {
}

// GetTplTestcaseTemplateDetail 测试用例模板详细
func (s TplTestcaseTemplateService) GetTplTestcaseTemplateDetail(id string) (*entitys.TplTestcaseTemplateEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientTestcaseTemplateService.FindById(context.Background(), &protosService.TplTestcaseTemplateFilter{Id: rid})
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
	return entitys.TplTestcaseTemplate_pb2e(data), err
}

// QueryTplTestcaseTemplateList 测试用例模板列表
func (s TplTestcaseTemplateService) QueryTplTestcaseTemplateList(filter entitys.TplTestcaseTemplateQuery) ([]*entitys.TplTestcaseTemplateEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	var (
		status        int32 = -1
		productId     int64 = 0
		productTypeId int64 = 0
	)
	if filter.Query == nil {
		filter.Query = new(entitys.TplTestcaseTemplateFilter)
	}
	if filter.Query.Status != nil {
		status = *filter.Query.Status
	}
	if filter.Query.ProductId != "" {
		productId = iotutil.ToInt64(filter.Query.ProductId)
	}
	if filter.Query.ProductTypeId != "" {
		productTypeId = iotutil.ToInt64(filter.Query.ProductTypeId)
	}
	var queryObj = &protosService.TplTestcaseTemplate{
		ProductId:     productId,
		ProductTypeId: productTypeId,
		Status:        status,
		TplName:       filter.Query.TplName,
	}
	rep, err := rpc.ClientTestcaseTemplateService.Lists(context.Background(), &protosService.TplTestcaseTemplateListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderKey:  filter.SortField,
		OrderDesc: filter.Sort,
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.TplTestcaseTemplateEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.TplTestcaseTemplate_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddTplTestcaseTemplate 新增测试用例模板
func (s TplTestcaseTemplateService) AddTplTestcaseTemplate(req entitys.TplTestcaseTemplateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	//基础参数验证
	saveObj := entitys.TplTestcaseTemplate_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientTestcaseTemplateService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateTplTestcaseTemplate 修改测试用例模板
func (s TplTestcaseTemplateService) UpdateTplTestcaseTemplate(req entitys.TplTestcaseTemplateEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientTestcaseTemplateService.Update(context.Background(), entitys.TplTestcaseTemplate_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteTplTestcaseTemplate 删除测试用例模板
func (s TplTestcaseTemplateService) DeleteTplTestcaseTemplate(req entitys.TplTestcaseTemplateFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientTestcaseTemplateService.Delete(context.Background(), &protosService.TplTestcaseTemplate{
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

// SetStatusTplTestcaseTemplate 禁用/启用测试用例模板
func (s TplTestcaseTemplateService) SetStatusTplTestcaseTemplate(req entitys.TplTestcaseTemplateFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	if req.Status == nil {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientTestcaseTemplateService.UpdateFields(context.Background(), &protosService.TplTestcaseTemplateUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.TplTestcaseTemplate{
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
