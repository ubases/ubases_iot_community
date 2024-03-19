package services

import (
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type PmFirmwareSettingService struct {
}

// 固件设置详细
func (s PmFirmwareSettingService) GetPmFirmwareSettingDetail(id string) (*entitys.PmFirmwareSettingEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientFirmwareSettingService.FindById(context.Background(), &protosService.PmFirmwareSettingFilter{Id: rid})
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
	return entitys.PmFirmwareSetting_pb2e(data), err
}

// QueryPmFirmwareSettingList 固件设置列表
func (s PmFirmwareSettingService) QueryPmFirmwareSettingList(filter entitys.PmFirmwareSettingQuery) ([]*entitys.PmFirmwareSettingEntitys, int64, error) {
	rep, err := rpc.ClientFirmwareSettingService.Lists(context.Background(), &protosService.PmFirmwareSettingListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.PmFirmwareSetting{
			ProductModel: filter.Query.ProductModel,
			ModuleId:     filter.Query.ModuleId,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.PmFirmwareSettingEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.PmFirmwareSetting_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddPmFirmwareSetting 新增固件设置
func (s PmFirmwareSettingService) AddPmFirmwareSetting(req entitys.PmFirmwareSettingEntitys) (string, error) {
	saveObj := entitys.PmFirmwareSetting_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	//saveObj.CreatedAt = timestamppb.Now()
	//saveObj.UpdatedAt = timestamppb.Now()
	res, err := rpc.ClientFirmwareSettingService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件设置
func (s PmFirmwareSettingService) UpdatePmFirmwareSetting(req entitys.PmFirmwareSettingEntitys) (string, error) {
	//req.UpdatedAt = time.Now()
	res, err := rpc.ClientFirmwareSettingService.Update(context.Background(), entitys.PmFirmwareSetting_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除固件设置
func (s PmFirmwareSettingService) DeletePmFirmwareSetting(req entitys.PmFirmwareSettingFilter) error {
	rep, err := rpc.ClientFirmwareSettingService.Delete(context.Background(), &protosService.PmFirmwareSetting{
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
