package services

import (
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type LangResourcesService struct {
	Ctx context.Context
}

func (s LangResourcesService) SetContext(ctx context.Context) LangResourcesService {
	s.Ctx = ctx
	return s
}

// 翻译详细
func (s LangResourcesService) GetLangResourcesDetail(id string) (*entitys.LangResourcesEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientLangResourcesService.FindById(s.Ctx, &protosService.LangResourcesFilter{Id: rid})
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
	return entitys.LangResources_pb2e(data), err
}

// QueryLangResourcesList 翻译列表
func (s LangResourcesService) QueryLangResourcesList(filter entitys.LangResourcesQuery) ([]*entitys.LangResourcesEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientLangResourcesService.Lists(s.Ctx, &protosService.LangResourcesListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query:     &protosService.LangResources{},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.LangResourcesEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.LangResources_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddLangResources 新增翻译
func (s LangResourcesService) AddLangResources(req entitys.LangResourcesEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.LangResources_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()

	res, err := rpc.ClientLangResourcesService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改翻译
func (s LangResourcesService) UpdateLangResources(req entitys.LangResourcesEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientLangResourcesService.Update(s.Ctx, entitys.LangResources_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除翻译
func (s LangResourcesService) DeleteLangResources(req entitys.LangResourcesFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientLangResourcesService.Delete(s.Ctx, &protosService.LangResources{
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
