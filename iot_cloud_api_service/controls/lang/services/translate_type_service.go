package services

import (
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type LangTranslateTypeService struct {
	Ctx context.Context
}

func (s LangTranslateTypeService) SetContext(ctx context.Context) LangTranslateTypeService {
	s.Ctx = ctx
	return s
}

// 固件详细
func (s LangTranslateTypeService) GetLangTranslateTypeDetail(id string) (*entitys.LangTranslateTypeEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientLangTranslateTypeService.FindById(s.Ctx, &protosService.LangTranslateTypeFilter{Id: rid})
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
	return entitys.LangTranslateType_pb2e(data), err
}

// QueryLangTranslateTypeList 固件列表
func (s LangTranslateTypeService) QueryLangTranslateTypeList(filter entitys.LangTranslateTypeQuery) ([]*entitys.LangTranslateTypeEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientLangTranslateTypeService.Lists(s.Ctx, &protosService.LangTranslateTypeListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query:     &protosService.LangTranslateType{},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.LangTranslateTypeEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.LangTranslateType_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddLangTranslateType 新增固件
func (s LangTranslateTypeService) AddLangTranslateType(req entitys.LangTranslateTypeEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.LangTranslateType_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientLangTranslateTypeService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件
func (s LangTranslateTypeService) UpdateLangTranslateType(req entitys.LangTranslateTypeEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientLangTranslateTypeService.Update(s.Ctx, entitys.LangTranslateType_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除固件
func (s LangTranslateTypeService) DeleteLangTranslateType(req entitys.LangTranslateTypeFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientLangTranslateTypeService.Delete(s.Ctx, &protosService.LangTranslateType{
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
