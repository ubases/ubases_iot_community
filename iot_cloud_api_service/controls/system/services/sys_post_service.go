package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type SysPostService struct {
}

// 岗位详细
func (s SysPostService) GetSysPostDetail(id string) (*entitys.SysPostEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientSysPostService.FindById(context.Background(), &protosService.SysPostFilter{PostId: rid})
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
	return entitys.SysPost_pb2e(data), err
}

// QuerySysPostList 岗位列表
func (s SysPostService) QuerySysPostList(filter entitys.SysPostQuery) ([]*entitys.SysPostEntitys, int64, error) {
	rep, err := rpc.ClientSysPostService.Lists(context.Background(), &protosService.SysPostListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.SysPost{
			PostName: filter.PostName,
			PostCode: filter.PostCode,
			Status:   filter.Status,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var resultList = []*entitys.SysPostEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysPost_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddSysPost 新增岗位
func (s SysPostService) AddSysPost(req entitys.SysPostEntitys) (string, error) {
	req.PostId = iotutil.ToString(iotutil.GetNextSeqInt64())
	res, err := rpc.ClientSysPostService.Create(context.Background(), entitys.SysPost_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.PostId), err
}

// 修改岗位
func (s SysPostService) UpdateSysPost(req entitys.SysPostEntitys) (string, error) {
	res, err := rpc.ClientSysPostService.UpdateAll(context.Background(), entitys.SysPost_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.PostId), err
}

// 删除岗位
func (s SysPostService) DeleteSysPost(req entitys.DeleteCommonQuery) error {
	var err error
	for _, id := range req.Ids {
		res, errDel := rpc.ClientSysPostService.Delete(context.Background(), &protosService.SysPost{
			PostId: iotutil.ToInt64(id),
		})
		if errDel != nil {
			err = errDel
			break
		}
		if res.Code != 200 {
			err = errors.New(res.Message)
			break
		}
	}
	return err
}
