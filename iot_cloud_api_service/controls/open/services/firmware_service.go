package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmFirmwareService struct {
	Ctx context.Context
}

func (s OpmFirmwareService) SetContext(ctx context.Context) OpmFirmwareService {
	s.Ctx = ctx
	return s
}

// 固件详细
func (s OpmFirmwareService) GetOpmFirmwareDetail(id string) (*entitys.OpmFirmwareEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientFirmwareService.FindById(s.Ctx, &protosService.OpmFirmwareFilter{Id: rid})
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
	return entitys.OpmFirmware_pb2e(data), err
}

// QueryOpmFirmwareList 固件列表
func (s OpmFirmwareService) QueryOpmFirmwareList(filter entitys.OpmFirmwareQuery) ([]*entitys.OpmFirmwareEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	queryParams := entitys.OpmFirmwareFilter_e2pb(filter.Query)
	queryParams.IsQueryValidVersion = true //只获取最新启用的的版本
	rep, err := rpc.ClientFirmwareService.Lists(s.Ctx, &protosService.OpmFirmwareListRequest{
		Page:      filter.Page,
		PageSize:  filter.Limit,
		SearchKey: filter.SearchKey,
		Query:     queryParams,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	resultList := make([]*entitys.OpmFirmwareEntitys, 0)
	for _, item := range rep.Data {
		//if item.Version == "" {
		//	continue
		//}
		resultList = append(resultList, entitys.OpmFirmware_pb2e(item))

	}
	return resultList, rep.Total, err
}

// AddOpmFirmware 新增固件
func (s OpmFirmwareService) AddOpmFirmware(req entitys.OpmFirmwareEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmFirmware_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientFirmwareService.CreateAndInitVersion(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//将固件关联到产品
	if req.ProductId != 0 {
		svr := OpmProductService{}.SetContext(s.Ctx)
		svr.SaveOpenProductAndModuleRelation(entitys.OpmProductModuleRelationEntitys{
			ProductId:  req.ProductId,
			FirmwareId: saveObj.Id,
			IsCustom:   1,
		})
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件
func (s OpmFirmwareService) UpdateOpmFirmware(req entitys.OpmFirmwareEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientFirmwareService.Update(s.Ctx, entitys.OpmFirmware_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除固件
func (s OpmFirmwareService) DeleteOpmFirmware(req entitys.OpmFirmwareFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientFirmwareService.Delete(s.Ctx, &protosService.OpmFirmware{
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

// SetStatusOpmFirmware 禁用/启用固件
func (s OpmFirmwareService) SetStatusOpmFirmware(req entitys.OpmFirmwareFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == nil {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientFirmwareService.UpdateFields(s.Ctx, &protosService.OpmFirmwareUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OpmFirmware{
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
