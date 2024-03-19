package services

import (
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/mitchellh/mapstructure"
)

type NetworkGuideService struct {
}

// CreateNetworkGuide create one record
func (s NetworkGuideService) CreateNetworkGuide(req *entitys.PmNetworkGuideEntitys) (ret int64, err error) {
	var (
		data = protosService.PmNetworkGuide{}
		now  = timestamppb.New(time.Now())
	)
	mapstructure.WeakDecode(req, &data)
	//参数填充
	data.CreatedAt = now
	data.UpdatedAt = now
	res, err := rpc.ClientNetworkGuideService.Create(context.Background(), &data)

	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Message)
	}
	return 0, err
}

// UpdateNetworkGuide edit NetworkGuide one record
func (s NetworkGuideService) UpdateNetworkGuide(req *entitys.PmNetworkGuideEntitys) (err error) {
	var (
		data = protosService.PmNetworkGuide{}
		now  = timestamppb.New(time.Now())
	)
	mapstructure.WeakDecode(req, &data)
	data.UpdatedAt = now

	res, err := rpc.ClientNetworkGuideService.Update(context.Background(), &data)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

// GetNetworkGuideBatch get NetworkGuide list  data
func (s NetworkGuideService) GetNetworkGuideList(filter *entitys.PmNetworkGuideQuery) (rets []*entitys.PmNetworkGuideEntitys, total int64, err error) {
	var (
		queryObj = &protosService.PmNetworkGuide{
			ProductId: filter.Query.ProductId,
		}
	)

	ret, err := rpc.ClientNetworkGuideService.Lists(context.Background(), &protosService.PmNetworkGuideListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query:    queryObj,
	})

	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Message)
	}
	for _, data := range ret.GetData() {
		rets = append(rets, &entitys.PmNetworkGuideEntitys{
			Id:            data.Id,
			ProductId:     data.ProductId,
			ProductTypeId: data.ProductTypeId,
			Type:          data.Type,
			CreatedBy:     data.CreatedBy,
			UpdatedAt:     data.UpdatedAt.AsTime(),
		})
	}
	return rets, ret.Total, nil
}

// GetNetworkGuide get NetworkGuide one record
func (s NetworkGuideService) GetNetworkGuide(id string) (res *entitys.PmNetworkGuideFilter, err error) {
	var (
		ret *protosService.PmNetworkGuideResponse
	)
	ret, err = rpc.ClientNetworkGuideService.FindById(context.Background(), &protosService.PmNetworkGuideFilter{
		Id: iotutil.ToInt64(id),
	})
	if err = mapstructure.WeakDecode(ret.GetData(), &res); err != nil {
		return
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return res, err
}

// DelNetworkGuide delete NetworkGuide one record
func (s NetworkGuideService) DelNetworkGuide(id int64) (err error) {
	var (
		data = protosService.PmNetworkGuide{}
	)
	data.Id = id
	ret, err := rpc.ClientNetworkGuideService.DeleteById(context.Background(), &data)
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}
