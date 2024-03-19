package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpenCompanyConnectService struct {
	Ctx context.Context
}

func (s OpenCompanyConnectService) SetContext(ctx context.Context) OpenCompanyConnectService {
	s.Ctx = ctx
	return s
}

func (s OpenCompanyConnectService) AddConnect(req entitys.OpenCompanyConnect) (string, error) {

	reqpb := entitys.OpenCompanyConnect_e2pb(&req)
	res, err := rpc.ClientOpenCompanyConnectService.Create(s.Ctx, reqpb)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "ok", nil
}

func (s OpenCompanyConnectService) UpdateConnect(req entitys.OpenCompanyConnect) (string, error) {
	reqpb := entitys.OpenCompanyConnect_e2pb(&req)
	res, err := rpc.ClientOpenCompanyConnectService.UpdateFields(s.Ctx, &protosService.OpenCompanyConnectUpdateFieldsRequest{
		Fields: []string{"name", "account", "phone", "address", "job"},
		Data:   reqpb,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "ok", nil
}

func (s OpenCompanyConnectService) DeleteConnect(id int64) (string, error) {
	res, err := rpc.ClientOpenCompanyConnectService.DeleteById(s.Ctx, &protosService.OpenCompanyConnect{
		Id: id,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "ok", nil
}
