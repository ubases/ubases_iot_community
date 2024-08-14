package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/mitchellh/mapstructure"
)

type NetworkGuideStepService struct {
}

// CreateNetworkGuideStep create one record
func (s NetworkGuideStepService) CreateNetworkGuideStep(req *entitys.PmNetworkGuideStepEntitys) (ret int64, err error) {
	var (
		data = protosService.PmNetworkGuideStep{}
		now  = timestamppb.New(time.Now())
	)
	mapstructure.WeakDecode(req, &data)
	//参数填充
	data.CreatedAt = now
	data.UpdatedAt = now
	res, err := rpc.ClientNetworkGuideStepService.Create(context.Background(), &data)

	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Message)
	}
	//设置上传图片对应业务是否成功
	urls := []string{}
	if data.ImageUrl != "" {
		urls = append(urls, data.ImageUrl)
	}
	if data.VideoUrl != "" {
		urls = append(urls, data.VideoUrl)
	}
	commonGlobal.SetAttachmentStatus(model.TableNameTOpmNetworkGuide, iotutil.ToString(req.Id), urls...)
	return 0, err
}

// UpdateNetworkGuideStep edit NetworkGuideStep one record
func (s NetworkGuideStepService) UpdateNetworkGuideStep(req *entitys.PmNetworkGuideStepEntitys) (err error) {
	var (
		data = protosService.PmNetworkGuideStep{}
		now  = timestamppb.New(time.Now())
	)
	mapstructure.WeakDecode(req, &data)
	data.UpdatedAt = now

	ret, err := rpc.ClientNetworkGuideStepService.Update(context.Background(), &data)
	if err != nil {
		return err
	}
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Message)
	}
	urls := []string{}
	if data.ImageUrl != "" {
		urls = append(urls, data.ImageUrl)
	}
	if data.VideoUrl != "" {
		urls = append(urls, data.VideoUrl)
	}
	commonGlobal.SetAttachmentStatus(model.TableNameTOpmNetworkGuide, iotutil.ToString(req.Id), urls...)
	return nil
}

// GetNetworkGuideStepBatch get NetworkGuideStep list  data
func (s NetworkGuideStepService) GetNetworkGuideStepList(filter *entitys.PmNetworkGuideStepQuery) (rets []*entitys.PmNetworkGuideStepEntitys, total int64, err error) {
	var (
		queryObj = &protosService.PmNetworkGuideStep{
			NetworkGuideId: filter.Query.NetworkGuideId,
		}
	)

	ret, err := rpc.ClientNetworkGuideStepService.Lists(context.Background(), &protosService.PmNetworkGuideStepListRequest{
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
		rets = append(rets, &entitys.PmNetworkGuideStepEntitys{
			Id:             data.Id,
			NetworkGuideId: data.NetworkGuideId,
			Instruction:    data.Instruction,
			InstructionEn:  data.InstructionEn,
			ImageUrl:       data.ImageUrl,
			VideoUrl:       data.VideoUrl,
			Sort:           data.Sort,
			CreatedBy:      data.CreatedBy,
			UpdatedAt:      data.UpdatedAt.AsTime(),
		})
	}

	return rets, ret.Total, nil
}

// GetNetworkGuideStep get NetworkGuideStep one record
func (s NetworkGuideStepService) GetNetworkGuideStep(id string) (res *entitys.PmNetworkGuideStepFilter, err error) {
	var (
		ret *protosService.PmNetworkGuideStepResponse
	)
	ret, err = rpc.ClientNetworkGuideStepService.FindById(context.Background(), &protosService.PmNetworkGuideStepFilter{
		Id: iotutil.ToInt64(id),
	})
	if err = mapstructure.WeakDecode(ret.GetData(), &res); err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return res, nil
}

// DelNetworkGuideStep delete NetworkGuideStep one record
func (s NetworkGuideStepService) DelNetworkGuideStep(id int64) (err error) {
	var (
		data = protosService.PmNetworkGuideStep{}
	)
	data.Id = id
	ret, err := rpc.ClientNetworkGuideStepService.DeleteById(context.Background(), &data)
	if err != nil {
		return err
	}
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}
