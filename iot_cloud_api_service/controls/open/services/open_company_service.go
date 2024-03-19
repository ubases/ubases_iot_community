package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strings"
)

type OpenCompanyService struct {
	Ctx context.Context
}

func (s OpenCompanyService) SetContext(ctx context.Context) OpenCompanyService {
	s.Ctx = ctx
	return s
}

// 获取公司信息
func (s OpenCompanyService) GetCompanyInfo(tenantId string) (*entitys.OpenCompanyInfoRes, error) {
	res, err := rpc.ClientOpenCompanyService.Find(context.Background(), &protosService.OpenCompanyFilter{
		TenantId: tenantId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data != nil && len(res.Data) == 0 {
		return nil, errors.New("数据不存在.")
	}
	info := entitys.OpenCompany_pb2eCompanyInfoRes(res.Data[0])

	resConnect, errConnect := rpc.ClientOpenCompanyConnectService.Lists(context.Background(), &protosService.OpenCompanyConnectListRequest{
		Page:      1,
		PageSize:  1000000,
		OrderKey:  "create_at",
		OrderDesc: "1",
		Query: &protosService.OpenCompanyConnect{
			TenantId: tenantId,
		},
	})
	if errConnect != nil {
		return nil, errConnect
	}
	if resConnect.Code != 200 && resConnect.Message != "record not found" {
		return nil, errors.New(resConnect.Message)
	}
	if resConnect.Data != nil && len(resConnect.Data) > 0 {

		for _, mm := range resConnect.Data {

			tmp := entitys.OpenCompanyConnect{
				Id:      iotutil.ToString(mm.Id),
				Name:    mm.Name,
				Account: mm.Account,
				Phone:   mm.Phone,
				Job:     mm.Job,
				Address: mm.Address,
			}
			info.Connect = append(info.Connect, tmp)
		}
	}
	return info, nil
}

func (s OpenCompanyService) GetBaseInfo(userId int64) (*entitys.OpenCompanyBaseRes, error) {
	res, err := rpc.ClientOpenCompanyService.Find(context.Background(), &protosService.OpenCompanyFilter{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data != nil && len(res.Data) == 0 {
		return nil, errors.New("数据不存在.")
	}
	company := entitys.OpenCompany_pb2eRes(res.Data[0])
	return company, nil
}

// 获取企业认证信息
func (s OpenCompanyService) GetCompanyAuth() (*entitys.OpenCompanyAuthRes, error) {
	res, err := rpc.ClientOpenCompanyService.Find(s.Ctx, &protosService.OpenCompanyFilter{})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data != nil && len(res.Data) == 0 {
		return nil, errors.New("数据不存在.")
	}

	resLogs, errLogs := rpc.ClientOpenCompanyAuthLogsService.Lists(context.Background(), &protosService.OpenCompanyAuthLogsListRequest{
		Page:      1,
		PageSize:  1000000,
		OrderKey:  "created_at",
		OrderDesc: "",
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

	var authlist []*entitys.OpenAuthEntity
	if len(resLogs.Data) > 0 {
		for _, v := range resLogs.Data {
			auth := entitys.OpenAuthEntity{
				Result:    v.AuthResult,
				OpterTime: v.CreatedAt.AsTime().Unix(),
				Why:       v.Why,
			}
			authlist = append(authlist, &auth)
		}

	}
	company := entitys.OpenCompany_pb2eCompanyAuth(res.Data[0])
	company.AuthList = authlist
	if company.Nature == 0 {
		company.Nature = 2 //默认等于民营企业
	}

	return company, nil
}

func (s OpenCompanyService) CompanyAuth(req entitys.OpenCompanyAuthReq) (string, error) {
	res, err := rpc.ClientOpenCompanyService.CompanyAuth(context.Background(), &protosService.OpenCompanyUpdateFieldsRequest{
		Fields: []string{"nature", "name", "phone", "address", "license_no", "license", "legal_person", "apply_person", "idcard", "idcard_front_img", "idcard_after_img"},
		Data: &protosService.OpenCompany{
			Id:             iotutil.ToInt64(req.Id),
			Name:           req.CompanyName,
			Nature:         iotutil.ToString(req.Nature),
			Phone:          req.Phone,
			Address:        req.Address,
			LicenseNo:      req.LicenseNo,
			License:        req.License,
			LegalPerson:    req.LegalPerson,
			ApplyPerson:    req.ApplyPerson,
			Idcard:         req.Idcard,
			IdcardFrontImg: req.IdcardFrontImg,
			IdcardAfterImg: req.IdcardAfterImg,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return req.Id, nil
}

func (s OpenCompanyService) CompanyChangeName(req entitys.OpenCompanyChangeNameReq) (string, error) {
	company, errCompany := rpc.ClientOpenCompanyService.FindById(context.Background(), &protosService.OpenCompanyFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if errCompany != nil {
		return "", errCompany
	}
	if len(company.Data) <= 0 {
		return "", errors.New("未找到公司记录.")
	}
	//1=未提交,  2=认证中,   4=不通过,   3=已认证,   5=已撤销
	tmp := company.Data[0].Status
	if tmp != 1 && tmp != 4 && tmp != 5 {
		return "", errors.New("目前公司状态下, 不可更改名称.")
	}
	res, err := rpc.ClientOpenCompanyService.UpdateFields(context.Background(), &protosService.OpenCompanyUpdateFieldsRequest{
		Fields: []string{"name"},
		Data: &protosService.OpenCompany{
			Id:   iotutil.ToInt64(req.Id),
			Name: req.CompanyName,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return req.Id, nil
}

func (s OpenCompanyService) UpdateBaseInfo(req entitys.OpenCompanyBaseReq, userId int64) (string, error) {

	fileds := []string{"region", "phone", "email", "address"}
	resCompany, errCompany := rpc.ClientOpenCompanyService.FindById(context.Background(), &protosService.OpenCompanyFilter{Id: iotutil.ToInt64(req.Id)})
	if errCompany != nil {
		return "", errCompany
	}
	//企业账号, 未提交状态  才可以修改公司名
	if resCompany.Data[0].AccountType == 1 && resCompany.Data[0].Status == 1 {
		fileds = append(fileds, "name")
	}
	//个人账号可以修改所在企业
	if resCompany.Data[0].AccountType == 2 {
		fileds = append(fileds, "name")
	}

	res, err := rpc.ClientOpenCompanyService.UpdateFields(context.Background(), &protosService.OpenCompanyUpdateFieldsRequest{
		Fields: fileds,
		Data: &protosService.OpenCompany{
			Id:      iotutil.ToInt64(req.Id),
			Region:  strings.Join(req.Area, ","),
			Name:    req.CompanyName,
			Phone:   req.Phone,
			Email:   req.Email,
			Address: req.Address,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return req.Id, nil
}

// GetCompanyList 获取公司列表，通过租户Id集合
func (s OpenCompanyService) GetCompanyList(tenantIds []string) (map[string]*protosService.OpenCompany, error) {
	//TODO 需要增加缓存处理， 通过tenantId作为key进行缓存
	resp, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &protosService.OpenCompanyListRequest{
		Query: &protosService.OpenCompany{TenantIds: tenantIds},
	})
	if err != nil {
		return nil, err
	}
	result := make(map[string]*protosService.OpenCompany, 0)
	for _, item := range resp.Data {
		tenantId := item.TenantId
		if _, ok := result[tenantId]; !ok {
			result[tenantId] = &protosService.OpenCompany{}
		}
		result[tenantId] = item
	}
	return result, nil
}
