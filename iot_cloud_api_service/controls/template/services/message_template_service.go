package services

import (
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type MpMessageTemplateService struct {
}

// GetMpMessageTemplateDetail 消息模板详细
func (s MpMessageTemplateService) GetMpMessageTemplateDetail(id string) (*entitys.MpMessageTemplateEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientMessageTemplateService.FindById(context.Background(), &protosService.MpMessageTemplateFilter{Id: rid})
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
	return entitys.MpMessageTemplate_pb2e(data), err
}

// QueryMpMessageTemplateList 消息模板列表
func (s MpMessageTemplateService) QueryMpMessageTemplateList(filter entitys.MpMessageTemplateQuery) ([]*entitys.MpMessageTemplateEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	// 状态字段查询的方案
	if filter.Query == nil {
		filter.Query = new(entitys.MpMessageTemplateFilter)
	}
	rep, err := rpc.ClientMessageTemplateService.Lists(context.Background(), &protosService.MpMessageTemplateListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query: &protosService.MpMessageTemplate{
			TplCode:     filter.Query.TplCode,
			TplName:     filter.Query.TplName,
			PushType:    filter.Query.PushType,
			MessageType: filter.Query.MessageType,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.MpMessageTemplateEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.MpMessageTemplate_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddMpMessageTemplate 新增消息模板
func (s MpMessageTemplateService) AddMpMessageTemplate(req entitys.MpMessageTemplateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.MpMessageTemplate_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientMessageTemplateService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateMpMessageTemplate 修改消息模板
func (s MpMessageTemplateService) UpdateMpMessageTemplate(req entitys.MpMessageTemplateEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientMessageTemplateService.Update(context.Background(), entitys.MpMessageTemplate_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteMpMessageTemplate 删除消息模板
func (s MpMessageTemplateService) DeleteMpMessageTemplate(req entitys.MpMessageTemplateFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientMessageTemplateService.Delete(context.Background(), &protosService.MpMessageTemplate{
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
