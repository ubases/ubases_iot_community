package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/services/openData"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmProductAppRelationService struct {
	Ctx context.Context
}

func (s OpmProductAppRelationService) SetContext(ctx context.Context) OpmProductAppRelationService {
	s.Ctx = ctx
	return s
}

// QueryList 产品APP关联列表
func (s OpmProductAppRelationService) QueryList(productId, appKey, tenantId string) ([]*entitys.OpmProductAppRelationEntitys, error) {
	productIdInt, _ := iotutil.ToInt64AndErr(productId)
	rep, err := rpc.ClientProductAppRelationService.Lists(s.Ctx, &protosService.OpmProductAppRelationListRequest{
		Query: &protosService.OpmProductAppRelationFilter{
			ProductId: productIdInt,
			AppKey:    appKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}

	//查询关联APP信息
	appMap, _ := openData.GetAppMaps(s.Ctx, tenantId)
	resultList := make([]*entitys.OpmProductAppRelationEntitys, 0)
	for _, item := range rep.Data {
		item.AppName = iotutil.IfString(item.AppName == "", item.AppKey, item.AppName)
		if v, ok := appMap[item.AppKey]; ok {
			item.AppName = v
		}
		resultList = append(resultList, entitys.OpmProductAppRelation_pb2e(item))
	}
	return resultList, err
}

// QueryByIds 产品APP关联列表
func (s OpmProductAppRelationService) ProductRelationMap(productIds []int64, tenantId string) (res map[int64][]*iotstruct.KeyValue, resName map[int64][]string, err error) {
	res = make(map[int64][]*iotstruct.KeyValue)
	resName = make(map[int64][]string)
	rep, err := rpc.ClientProductAppRelationService.Lists(s.Ctx, &protosService.OpmProductAppRelationListRequest{
		Query: &protosService.OpmProductAppRelationFilter{
			ProductIds: productIds,
		},
	})
	if err != nil {
		return
	}
	if rep.Code != 200 {
		err = errors.New(rep.Message)
		return
	}
	appMap, _ := openData.GetAppMaps(s.Ctx, tenantId)
	for _, d := range rep.Data {
		d.AppName = iotutil.IfString(d.AppName == "", d.AppKey, d.AppName)
		if v, ok := appMap[d.AppKey]; ok {
			d.AppName = v
		}
		res[d.ProductId] = append(res[d.ProductId], &iotstruct.KeyValue{
			Key:   d.AppKey,
			Value: d.AppName,
		})
		resName[d.ProductId] = append(resName[d.ProductId], d.AppName)
	}
	return
}

func (s OpmProductAppRelationService) AppRelationMap(appKeys []string) (res map[string][]*iotstruct.KeyValue, resName map[string][]string, err error) {
	res = make(map[string][]*iotstruct.KeyValue)
	resName = make(map[string][]string)
	rep, err := rpc.ClientProductAppRelationService.Lists(s.Ctx, &protosService.OpmProductAppRelationListRequest{
		Query: &protosService.OpmProductAppRelationFilter{
			AppKeys: appKeys,
		},
	})
	if err != nil {
		return
	}
	if rep.Code != 200 {
		err = errors.New(rep.Message)
		return
	}
	for _, d := range rep.Data {
		d.ProductName = iotutil.IfString(d.ProductName == "", d.ProductKey, d.ProductName)
		res[d.AppKey] = append(res[d.AppKey], &iotstruct.KeyValue{
			Key:   iotutil.ToString(d.ProductId),
			Value: d.ProductName,
		})
		resName[d.AppKey] = append(resName[d.AppKey], d.ProductName)
	}
	return
}

// Add 新增产品APP关联
func (s OpmProductAppRelationService) BindRelation(req entitys.OpmProductAppRelationEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmProductAppRelation_e2pb(&req)
	res, err := rpc.ClientProductAppRelationService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "", err
}
