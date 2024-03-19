package services

import (
	"cloud_platform/iot_demo_api_service/controls/system/entitys"
	"cloud_platform/iot_demo_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BaseDataService struct {
}

func (s BaseDataService) GetBaseDataTypeDetail(id string) (*protosService.ConfigDictTypeResponse, error) {
	rid := iotutil.ToInt64(id)
	ret, err := rpc.TConfigDictTypeServerService.FindById(context.Background(), &protosService.ConfigDictTypeFilter{DictId: rid})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, err
}

func (s BaseDataService) QueryBaseDataTypeList(filter entitys.BaseDataTypeQuery) (*protosService.ConfigDictTypeResponse, error) {
	QueryObj := protosService.ConfigDictType{
		DictName: filter.DictName,
		DictType: filter.DictType,
	}
	ret, err := rpc.TConfigDictTypeServerService.Lists(context.Background(), &protosService.ConfigDictTypeListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query:    &QueryObj,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, nil
}

// TConfigDictDataRequest
func (s BaseDataService) AddBaseDataType(req entitys.BaseDataType) error {
	ret, err := rpc.TConfigDictTypeServerService.Create(context.Background(), &protosService.ConfigDictType{
		DictName:  req.DictName,
		DictType:  req.DictType,
		ValueType: req.ValueType,
		IsSystem:  req.IsSystem,
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) UpdateBaseDataType(req entitys.BaseDataType) error {
	ret, err := rpc.TConfigDictTypeServerService.Update(context.Background(), &protosService.ConfigDictType{
		DictId:    iotutil.ToInt64(req.DictID),
		DictName:  req.DictName,
		DictType:  req.DictType,
		ValueType: req.ValueType,
		Status:    req.Status,
		//IsSystem:  req.IsSystem,
		Remark:    req.Remark,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
		//DeletedAt: "",
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) DeleteBaseDataType(id string) error {
	ret, err := rpc.TConfigDictTypeServerService.DeleteById(context.Background(), &protosService.ConfigDictType{
		DictId: iotutil.ToInt64(id),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}
