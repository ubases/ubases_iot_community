package services

import (
	"cloud_platform/iot_cloud_api_service/controls/config/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type SysAreaService struct {
}

// GetSysAreaDetail 测试用例模板详细
func (s SysAreaService) GetSysAreaDetail(id string) (*entitys.SysAreaEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientAreaService.FindById(context.Background(), &protosService.SysAreaFilter{Id: rid})
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
	return entitys.SysArea_pb2e(data), err
}

// QuerySysAreaList 测试用例模板列表
func (s SysAreaService) QuerySysAreaList(filter entitys.SysAreaQuery) ([]*entitys.SysAreaEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	var pidInt64 int64 = -1
	if filter.Query.Pid != nil {
		pidInt64 = *filter.Query.Pid
	}
	// 状态字段查询的方案
	if filter.Query == nil {
		filter.Query = new(entitys.SysAreaFilter)
	}
	rep, err := rpc.ClientAreaService.Lists(context.Background(), &protosService.SysAreaListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query: &protosService.SysArea{
			Pid:       pidInt64,
			ShowChild: filter.Query.ShowChild,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.SysAreaEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysArea_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddSysArea 新增测试用例模板
func (s SysAreaService) AddSysArea(req entitys.SysAreaEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}

	saveObj := entitys.SysArea_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientAreaService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateSysArea 修改测试用例模板
func (s SysAreaService) UpdateSysArea(req entitys.SysAreaEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}

	res, err := rpc.ClientAreaService.Update(context.Background(), entitys.SysArea_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteSysArea 删除测试用例模板
func (s SysAreaService) DeleteSysArea(req entitys.SysAreaFilter) error {
	rep, err := rpc.ClientAreaService.Delete(context.Background(), &protosService.SysArea{
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

// QueryRegionList 获取区域列表
func (s SysAreaService) QueryRegionList() ([]*protosService.SysRegionServer, error) {
	rep, err := rpc.SysRegionServerService.Lists(context.Background(), &protosService.SysRegionServerListRequest{
		Query: &protosService.SysRegionServer{},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	return rep.Data, nil
}
func (s SysAreaService) GetRegionMap() (map[int64]*protosService.SysRegionServer, error) {
	res := make(map[int64]*protosService.SysRegionServer)
	list, err := s.QueryRegionList()
	if err != nil {
		return res, err
	}
	for _, server := range list {
		res[server.Id] = server
	}
	return res, nil
}
