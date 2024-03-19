package services

import (
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type MsNoticeTemplateService struct {
}

// GetMsNoticeTemplateDetail 测试用例模板详细
func (s MsNoticeTemplateService) GetMsNoticeTemplateDetail(id string) (*entitys.MsNoticeTemplateEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientNoticeTemplateService.FindById(context.Background(), &protosService.MsNoticeTemplateFilter{Id: rid})
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
	return entitys.MsNoticeTemplate_pb2e(data), err
}

// QueryMsNoticeTemplateList 测试用例模板列表
func (s MsNoticeTemplateService) QueryMsNoticeTemplateList(filter entitys.MsNoticeTemplateQuery) ([]*entitys.MsNoticeTemplateEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	// 状态字段查询的方案
	if filter.Query == nil {
		filter.Query = new(entitys.MsNoticeTemplateFilter)
	}
	rep, err := rpc.ClientNoticeTemplateService.Lists(context.Background(), &protosService.MsNoticeTemplateListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query: &protosService.MsNoticeTemplate{
			TplCode:     filter.Query.TplCode,
			TplName:     filter.Query.TplName,
			TplSubject:  filter.Query.TplSubject,
			SmsSupplier: filter.Query.SmsSupplier,
			Lang:        filter.Query.Lang,
			TplType:     filter.Query.TplType,
			Method:      filter.Query.Method,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.MsNoticeTemplateEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.MsNoticeTemplate_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddMsNoticeTemplate 新增测试用例模板
func (s MsNoticeTemplateService) AddMsNoticeTemplate(req entitys.MsNoticeTemplateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.MsNoticeTemplate_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientNoticeTemplateService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateMsNoticeTemplate 修改测试用例模板
func (s MsNoticeTemplateService) UpdateMsNoticeTemplate(req entitys.MsNoticeTemplateEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientNoticeTemplateService.Update(context.Background(), entitys.MsNoticeTemplate_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteMsNoticeTemplate 删除测试用例模板
func (s MsNoticeTemplateService) DeleteMsNoticeTemplate(req entitys.MsNoticeTemplateFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientNoticeTemplateService.Delete(context.Background(), &protosService.MsNoticeTemplate{
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
