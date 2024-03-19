package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpenUserCompanyService struct {
	Ctx context.Context
}

func (s OpenUserCompanyService) SetContext(ctx context.Context) OpenUserCompanyService {
	s.Ctx = ctx
	return s
}

// 用户授权
func (s OpenUserCompanyService) UserCompanyAuth(req entitys.OpenUserCompanyAuthReq) (string, error) {
	res, err := rpc.ClientOpenUserCompanyService.UserCompanyAuth(s.Ctx, &protosService.OpenUserCompanyAuthRequest{
		UserName: req.UserName,
		Remark:   req.Remark,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	return "ok", nil
}

// 用户授权备注修改
func (s OpenUserCompanyService) UserCompanyUpdateReamk(req entitys.OpenUserCompanyAuthRemarkReq) (string, error) {
	res, err := rpc.ClientOpenUserCompanyService.UpdateFields(s.Ctx, &protosService.OpenUserCompanyUpdateFieldsRequest{
		Fields: []string{"remark"},
		Data: &protosService.OpenUserCompany{
			Id:     iotutil.ToInt64(req.Id),
			Remark: req.Remark,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	return "ok", nil
}

// 用户授权删除
func (s OpenUserCompanyService) UserCompanyDelete(id string) (string, error) {
	res, err := rpc.ClientOpenUserCompanyService.DeleteById(s.Ctx, &protosService.OpenUserCompany{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	return "ok", nil
}

// 授权列表
func (s OpenUserCompanyService) UserCompanyAuthList(req entitys.OpenUserCompanyAuthListReq) ([]*entitys.OpenUserCompanyAuthListRes, error) {
	res, err := rpc.ClientOpenUserCompanyService.UserCompanyAuthList(s.Ctx, &protosService.OpenUserCompanyAuthListRequest{
		AccountType: req.AccountType,
		Account:     req.Account,
		RoleId:      req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.List) == 0 {
		return nil, nil
	}
	var resList []*entitys.OpenUserCompanyAuthListRes
	for _, v := range res.List {

		entity := entitys.OpenUserCompanyAuthListRes_e2pb(v)
		resList = append(resList, entity)
	}
	return resList, nil
}
