package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpenDevService struct {
	Ctx context.Context
}

func (s OpenDevService) SetContext(ctx context.Context) OpenDevService {
	s.Ctx = ctx
	return s
}

// GetOpenDevList 开发者列表
func (s OpenDevService) GetOpenDevList(req entitys.OpenDevListReq) ([]*entitys.OpenDevListEntityRes, int64, error) {
	res, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &protosService.OpenCompanyListRequest{
		Page:      int64(req.PageNum),
		PageSize:  int64(req.PageSize),
		OrderKey:  "request_auth_at",
		OrderDesc: "desc",
		Query: &protosService.OpenCompany{
			Name:   req.CompanyName,
			Status: req.Status,
		},
	})

	if err != nil {
		return nil, 0, err
	}
	if res.Code != 200 {
		return nil, 0, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return nil, 0, nil
	}

	var resList []*entitys.OpenDevListEntityRes
	for _, v := range res.Data {
		m := entitys.OpenCompanyToListReq(v)
		resList = append(resList, m)
	}
	return resList, res.Total, nil
}

// 获取开发者Map格式数据
func (s OpenDevService) GetOpenDevMap(status int32) (map[string][]*entitys.OpenDevListEntityRes, error) {
	res, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &protosService.OpenCompanyListRequest{
		Query: &protosService.OpenCompany{
			Status: status,
		},
	})

	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return nil, nil
	}

	var resList map[string][]*entitys.OpenDevListEntityRes = make(map[string][]*entitys.OpenDevListEntityRes)
	for _, v := range res.Data {
		m := entitys.OpenCompanyToListReq(v)
		resList[v.TenantId] = append(resList[v.TenantId], m)
	}
	return resList, nil
}

func (s OpenDevService) GetOpenDevDetail(id string) (*entitys.OpenDevDetailRes, error) {
	res, err := rpc.ClientOpenCompanyService.FindById(context.Background(), &protosService.OpenCompanyFilter{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return nil, nil
	}

	resLogs, errLogs := rpc.ClientOpenCompanyAuthLogsService.Lists(context.Background(), &protosService.OpenCompanyAuthLogsListRequest{
		Page:      1,
		PageSize:  1000000,
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.OpenCompanyAuthLogs{
			CompanyId: res.Data[0].Id,
		},
	})
	if errLogs != nil {
		return nil, errLogs
	}
	if resLogs.Code != 200 {
		return nil, errors.New(resLogs.Message)
	}

	var authlist []*entitys.OpenDevAuthEntity
	if len(resLogs.Data) > 0 {
		for _, v := range resLogs.Data {
			auth := entitys.OpenDevAuthEntity{
				Result:    v.AuthResult,
				Opter:     v.AuthName,
				OpterTime: v.CreatedAt.AsTime().Unix(),
				Why:       v.Why,
			}
			authlist = append(authlist, &auth)
		}

	}

	opendev := entitys.OpenDevDetailRes{
		ID:             iotutil.ToString(res.Data[0].Id),
		Nature:         res.Data[0].Nature,
		Status:         res.Data[0].Status,
		LicenseNo:      res.Data[0].LicenseNo,
		License:        res.Data[0].License,
		LegalPerson:    res.Data[0].LegalPerson,
		ApplyPerson:    res.Data[0].ApplyPerson,
		Idcard:         res.Data[0].Idcard,
		IdcardFrontImg: res.Data[0].IdcardFrontImg,
		IdcardAfterImg: res.Data[0].IdcardAfterImg,
		Phone:          res.Data[0].Phone,
		Address:        res.Data[0].Address,
		CompanyName:    res.Data[0].Name,
		AuthList:       authlist,
	}

	return &opendev, nil
}

// 开发者认证审核
func (s OpenDevService) OpenDevAuth(req entitys.OpenDevCompanyAuthReq, userId string) (string, error) {
	var user = SysUserService{}
	resUser, errUser := user.GetUserProfile(userId)
	if errUser != nil {
		return "", errUser
	}

	res, err := rpc.ClientOpenCompanyService.OpenDevCompanyAuth(s.Ctx, &protosService.OpenDevCompanyAuthRequest{
		Status:   req.Status,
		Why:      req.Why,
		Id:       req.Id,
		AuthName: resUser.UserNickname,
	})

	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "ok", nil

}
