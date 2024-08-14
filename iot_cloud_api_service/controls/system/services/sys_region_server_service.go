package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type SysRegionServerService struct {
	Ctx context.Context
}

func (s SysRegionServerService) SetContext(ctx context.Context) SysRegionServerService {
	s.Ctx = ctx
	return s
}

// 详细
func (s SysRegionServerService) GetSysRegionServerDetail(id string) (*entitys.SysRegionServerEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.SysRegionServerService.FindById(s.Ctx, &protosService.SysRegionServerFilter{Id: rid})
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
	return entitys.SysRegionServer_pb2e(data), err
}

// QuerySysRegionServerList 列表
func (s SysRegionServerService) QuerySysRegionServerList(filter entitys.SysRegionServerQuery) ([]*entitys.SysRegionServerEntitys, int64, error) {
	var queryObj = filter.SysRegionServerQuery_e2pb()
	rep, err := rpc.SysRegionServerService.Lists(s.Ctx, &protosService.SysRegionServerListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderKey:  "sort",
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = make([]*entitys.SysRegionServerEntitys, 0)
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysRegionServer_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddSysRegionServer 新增
func (s SysRegionServerService) AddSysRegionServer(req entitys.SysRegionServerEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.SysRegionServer_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.SysRegionServerService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改
func (s SysRegionServerService) UpdateSysRegionServer(req entitys.SysRegionServerEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.SysRegionServerService.Update(s.Ctx, entitys.SysRegionServer_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除
func (s SysRegionServerService) DeleteSysRegionServer(req entitys.SysRegionServerFilter) error {
	rep, err := rpc.SysRegionServerService.DeleteById(s.Ctx, &protosService.SysRegionServer{
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

// SetStatusSysRegionServer 禁用/启用
func (s SysRegionServerService) SetStatusSysRegionServer(req entitys.SysRegionServerFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.SysRegionServerService.Update(context.Background(), &protosService.SysRegionServer{
		Id:      iotutil.ToInt64(req.Id),
		Enabled: req.Enabled,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
