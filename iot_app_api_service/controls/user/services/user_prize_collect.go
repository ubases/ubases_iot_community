package services

import (
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"

	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppPrizeCollectService struct {
	Ctx context.Context
}

func (s AppPrizeCollectService) SetContext(ctx context.Context) AppPrizeCollectService {
	s.Ctx = ctx
	return s
}

func (s AppPrizeCollectService) AddPrizeCollect(req entitys.UcUserPrizeCollectEntitys, userId int64) error {
	appKey, _ := metadata.Get(s.Ctx, "appKey")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	saveObj := entitys.UcUserPrizeCollect_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.TenantId = tenantId
	saveObj.AppKey = appKey
	saveObj.UserId = userId
	saveObj.CreatedAt = timestamppb.Now()
	ctx := context.Background()
	res, err := rpc.ClientUcUserPrizeCollectService.Create(ctx, saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}
