package services

import (
	"cloud_platform/iot_cloud_api_service/controls/config/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type ConfigPlatformService struct {
}

// GetConfigPlatformDetail 平台配置项详细
func (s ConfigPlatformService) GetConfigPlatformDetail(id string) (*entitys.ConfigPlatformEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientConfigPlatformService.FindById(context.Background(), &protosService.ConfigPlatformFilter{Id: rid})
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
	return entitys.ConfigPlatform_pb2e(data), err
}

// QueryConfigPlatformList 平台配置项列表
func (s ConfigPlatformService) QueryConfigPlatformList(filter entitys.ConfigPlatformQuery) ([]*entitys.ConfigPlatformEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	// 状态字段查询的方案
	if filter.Query == nil {
		filter.Query = new(entitys.ConfigPlatformFilter)
	}
	rep, err := rpc.ClientConfigPlatformService.Lists(context.Background(), &protosService.ConfigPlatformListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query:     &protosService.ConfigPlatform{},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.ConfigPlatformEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.ConfigPlatform_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddConfigPlatform 新增平台配置项
func (s ConfigPlatformService) AddConfigPlatform(req entitys.ConfigPlatformEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.ConfigPlatform_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientConfigPlatformService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateConfigPlatform 修改平台配置项
func (s ConfigPlatformService) UpdateConfigPlatform(req entitys.ConfigPlatformEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}

	res, err := rpc.ClientConfigPlatformService.Update(context.Background(), entitys.ConfigPlatform_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteConfigPlatform 删除平台配置项
func (s ConfigPlatformService) DeleteConfigPlatform(req entitys.ConfigPlatformFilter) error {
	rep, err := rpc.ClientConfigPlatformService.Delete(context.Background(), &protosService.ConfigPlatform{
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
