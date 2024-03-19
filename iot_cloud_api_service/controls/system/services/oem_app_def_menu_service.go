package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OemAppDefMenuService struct {
	Ctx context.Context
}

func (s OemAppDefMenuService) SetContext(ctx context.Context) OemAppDefMenuService {
	s.Ctx = ctx
	return s
}

// 新增OEMAPP默认菜单
func (s OemAppDefMenuService) AddOemAppDefMenu(req entitys.OemAppDefMenuEntitys) (string, error) {
	req.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	protos := entitys.OemAppDefMenu_e2pb(&req)
	res, err := rpc.ClientOemAppDefMenuService.Create(s.Ctx, protos)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 修改OEMAPP默认菜单
func (s OemAppDefMenuService) EditOemAppDefMenu(req entitys.OemAppDefMenuEntitys) (string, error) {
	protos := entitys.OemAppDefMenu_e2pb(&req)
	res, err := rpc.ClientOemAppDefMenuService.Update(s.Ctx, protos)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// OEMAPP默认菜单列表
func (s OemAppDefMenuService) QueryOemAppDefMenuList(filter entitys.OemAppDefMenuQuery) ([]*entitys.OemAppDefMenuEntitys, int64, error) {
	rep, err := rpc.ClientOemAppDefMenuService.Lists(s.Ctx, &protosService.OemAppDefMenuListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.OemAppDefMenu{
			Name: filter.Name,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return nil, 0, errors.New("未查询到数据")
	}
	var resultList = []*entitys.OemAppDefMenuEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OemAppDefMenu_pb2e(item))
	}
	return resultList, rep.Total, err
}
