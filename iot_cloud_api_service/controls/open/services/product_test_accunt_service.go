package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmProductTestAccountService struct {
	Ctx context.Context
}

func (s OpmProductTestAccountService) SetContext(ctx context.Context) OpmProductTestAccountService {
	s.Ctx = ctx
	return s
}

// QueryOpmProductTestAccountList 列表
func (s OpmProductTestAccountService) QueryOpmProductTestAccountList(filter entitys.OpmProductTestAccountQuery) ([]*entitys.OpmProductTestAccountEntitys, int64, error) {
	queryObj := &protosService.OpmProductTestAccountFilter{
		Account:   filter.Query.Account,
		ProductId: filter.Query.ProductId,
	}
	rep, err := rpc.ClientProductTestAccountService.Lists(s.Ctx, &protosService.OpmProductTestAccountListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = make([]*entitys.OpmProductTestAccountEntitys, 0)
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmProductTestAccount_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmProductTestAccount 新增
func (s OpmProductTestAccountService) AddOpmProductTestAccount(req entitys.OpmProductTestAccountEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmProductTestAccount_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	if req.RegionServerId == 0 {
		//return "", errors.New("服务器Id不能为空")
		req.RegionServerId = 1 //默认是区域1
	}
	if req.AppKey == "" {
		return "", errors.New("APPKey不能为空")
	}
	var phone, email string
	if iotutil.IsEmail(req.Account) {
		email = req.Account
	} else if iotutil.IsPhone(req.Account) {
		phone = req.Account
	} else {
		return "", errors.New("账号错误")
	}
	//获取用户Id
	userReq, err := rpc.UcUserService.Find(context.Background(), &protosService.UcUserFilter{
		Phone:          phone,
		Email:          email,
		TenantId:       req.TenantId,
		RegionServerId: req.RegionServerId,
		AppKey:         req.AppKey,
	})
	if err != nil {
		return "", err
	}
	if len(userReq.Data) == 0 {
		return "", errors.New("当前账号未注册该APP")
	}
	saveObj.UserId = userReq.Data[0].Id
	res, err := rpc.ClientProductTestAccountService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// DeleteOpmProductTestAccount 刪除
func (s OpmProductTestAccountService) DeleteOpmProductTestAccount(req entitys.OpmProductTestAccountEntitys) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	res, err := rpc.ClientProductTestAccountService.DeleteById(s.Ctx, &protosService.OpmProductTestAccount{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}
