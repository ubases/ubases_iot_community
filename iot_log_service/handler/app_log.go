package handler

import (
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_log_service/convert"
	"cloud_platform/iot_log_service/service"
	"cloud_platform/iot_model/ch_log/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"

	goerrors "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppLogHandler struct{}

// 创建AppLogUser
func (al *AppLogHandler) CreateAppLogUser(ctx context.Context, req *proto.AppLogUser, resp *proto.AppLogCommonResponse) error {
	alData := convert.AppLogUserPb2Db(req)
	err := service.CreateAppLogUser(alData)
	if err != nil {
		CommonResponse(resp, err)
		return err
	}
	CommonResponse(resp, err)
	return nil
}

// 更新AppLogUser
func (al *AppLogHandler) UpdateAppLogUser(ctx context.Context, req *proto.AppLogUser, resp *proto.AppLogCommonResponse) error {
	alData := convert.AppLogUserPb2Db(req)
	err := service.UpdateAppLogUser(alData)
	if err != nil {
		CommonResponse(resp, err)
		return err
	}
	CommonResponse(resp, err)
	return nil
}

// 删除AppLogUser
func (al *AppLogHandler) DeleteAppLogUser(ctx context.Context, req *proto.AppLogUser, resp *proto.AppLogCommonResponse) error {
	alData := convert.AppLogUserPb2Db(req)
	err := service.DeleteAppLogUser(alData)
	if err != nil {
		CommonResponse(resp, err)
		return err
	}
	CommonResponse(resp, err)
	return nil
}

// 获取AppLogUser
func (al *AppLogHandler) GetAppLogUser(ctx context.Context, req *proto.AppLogUser, resp *proto.AppLogUser) error {
	reqUser := model.AppLogUser{
		Account: req.Account,
	}
	alData, err := service.GetAppLogUser(&reqUser)
	if err != nil {
		return err
	}
	resp.Id = alData.Id
	resp.Account = alData.Account
	resp.AppName = alData.AppName
	resp.Region = alData.Region
	resp.LoginTime = timestamppb.New(alData.LoginTime)
	resp.CreatedAt = timestamppb.New(alData.CreatedAt)
	return nil
}

// 获取AppLogUserList
func (al *AppLogHandler) GetAppLogUserList(ctx context.Context, req *proto.AppLogUserListReq, resp *proto.AppLogUserListResp) error {
	alr, total, err := service.GetAppLogUserList(req.Query.Account, req.Query.AppName, req.Page, req.Limit, req.Query.RegionServerId)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBAppLogUserList)
	}
	list := []*proto.AppLogUser{}
	for i := range alr {
		list = append(list, convert.AppLogUserDb2Pb(&alr[i]))
	}
	data := &proto.AppLogUserData{
		List:  list,
		Total: total,
	}
	resp.Code = 200
	resp.Msg = "Success"
	resp.Data = data
	return nil
}

// 获取AppLogUserList
func (al *AppLogHandler) GetAppLogRecordsList(ctx context.Context, req *proto.AppLogRecordsListReq, resp *proto.AppLogRecordsListResp) error {
	alr, total, err := service.GetAppLogRecordsList(req.Account, req.AppKey, req.TenantId, req.Query.StartTime, req.Query.EndTime, req.Query.EventName, req.Query.LogType, req.Page, req.Limit, req.RegionServerId)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBAppLogRecordsList)
	}
	list := []*proto.AppLogRecords{}
	for i := range alr {
		list = append(list, convert.AppLogRecordsDb2Pb(&alr[i]))
	}
	data := &proto.AppLogRecordsData{
		List:  list,
		Total: total,
	}
	resp.Code = 200
	resp.Msg = "Success"
	resp.Data = data
	return nil
}

func CommonResponse(resp *proto.AppLogCommonResponse, err error) {
	if err != nil {
		resp.Code = 400
		resp.Msg = err.Error()
	} else {
		resp.Code = 200
		resp.Msg = "Success"
	}
}
