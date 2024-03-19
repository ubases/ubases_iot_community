package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmDocumentsService struct {
	Ctx context.Context
}

func (s OpmDocumentsService) SetContext(ctx context.Context) OpmDocumentsService {
	s.Ctx = ctx
	return s
}

// 详细
func (s OpmDocumentsService) GetOpmDocumentsDetail(id string) (*entitys.OpmDocumentsEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmDocumentsService.FindById(s.Ctx, &protosService.OpmDocumentsFilter{Id: rid})
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
	return entitys.OpmDocuments_pb2e(data), err
}

// QueryOpmDocumentsList 列表
func (s OpmDocumentsService) QueryOpmDocumentsList(filter entitys.OpmDocumentsQuery) ([]*entitys.OpmDocumentsEntitys, int64, error) {
	var queryObj = filter.OpmDocumentsQuery_e2pb()
	rep, err := rpc.ClientOpmDocumentsService.Lists(s.Ctx, &protosService.OpmDocumentsListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderDesc: "desc",
		OrderKey:  "doc_code",
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmDocumentsEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmDocuments_pb2e(item))
	}
	return resultList, rep.Total, err
}

func (s OpmDocumentsService) QueryOpmDocumentsListByProductId(productId string) ([]*entitys.OpmDocumentsEntitys, int64, error) {
	rep, err := rpc.ClientOpmDocumentsService.Lists(s.Ctx, &protosService.OpmDocumentsListRequest{
		Query: &protosService.OpmDocuments{OriginId: iotutil.ToInt64(productId)},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = make([]*entitys.OpmDocumentsEntitys, 0)
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmDocuments_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmDocuments 新增
func (s OpmDocumentsService) AddOpmDocuments(req entitys.OpmDocumentsEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmDocuments_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientOpmDocumentsService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改
func (s OpmDocumentsService) UpdateOpmDocuments(req entitys.OpmDocumentsEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpmDocumentsService.UpdateAll(s.Ctx, entitys.OpmDocuments_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除
func (s OpmDocumentsService) DeleteOpmDocuments(req entitys.OpmDocumentsFilter) error {
	rep, err := rpc.ClientOpmDocumentsService.Delete(s.Ctx, &protosService.OpmDocuments{
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
