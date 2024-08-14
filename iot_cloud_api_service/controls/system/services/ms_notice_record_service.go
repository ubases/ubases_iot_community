package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type MsNoticeRecordService struct {
	Ctx context.Context
}

func (s MsNoticeRecordService) SetContext(ctx context.Context) MsNoticeRecordService {
	s.Ctx = ctx
	return s
}

// 详细
func (s MsNoticeRecordService) GetMsNoticeRecordDetail(id string) (*entitys.MsNoticeRecordEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientMsNoticeRecordService.FindById(s.Ctx, &protosService.MsNoticeRecordFilter{Id: rid})
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
	return entitys.MsNoticeRecord_pb2e(data), err
}

// QueryMsNoticeRecordList 列表
func (s MsNoticeRecordService) QueryMsNoticeRecordList(filter entitys.MsNoticeRecordQuery) ([]*entitys.MsNoticeRecordEntitys, int64, error) {
	rep, err := rpc.ClientMsNoticeRecordService.Lists(s.Ctx, &protosService.MsNoticeRecordListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query: &protosService.MsNoticeRecordFilter{
			Id:              filter.Query.Id,
			Lang:            filter.Query.Lang,
			Platform:        filter.Query.Platform,
			Account:         filter.Query.Account,
			Method:          filter.Query.Method,
			Methods:         filter.Query.Methods,
			SmsSupplier:     filter.Query.SmsSupplier,
			Type:            filter.Query.Type,
			ThirdparyCode:   filter.Query.ThirdparyCode,
			NoticeTempateId: filter.Query.NoticeTempateId,
			Title:           filter.Query.Title,
			Content:         filter.Query.Content,
			AppKey:          filter.Query.AppKey,
			TenantId:        filter.Query.TenantId,
			Status:          filter.Query.Status,
			BeginTime:       filter.Query.BeginTime,
			EndTime:         filter.Query.EndTime,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = make([]*entitys.MsNoticeRecordEntitys, 0)
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.MsNoticeRecord_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddMsNoticeRecord 新增
func (s MsNoticeRecordService) AddMsNoticeRecord(req entitys.MsNoticeRecordEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.MsNoticeRecord_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientMsNoticeRecordService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改
func (s MsNoticeRecordService) UpdateMsNoticeRecord(req entitys.MsNoticeRecordEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientMsNoticeRecordService.Update(s.Ctx, entitys.MsNoticeRecord_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除
func (s MsNoticeRecordService) DeleteMsNoticeRecord(req entitys.MsNoticeRecordFilter) error {
	rep, err := rpc.ClientMsNoticeRecordService.DeleteById(s.Ctx, &protosService.MsNoticeRecord{
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
