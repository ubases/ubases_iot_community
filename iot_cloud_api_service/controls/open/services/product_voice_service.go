package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
)

type ProductViceService struct {
	Ctx context.Context
}

func (s ProductViceService) SetContext(ctx context.Context) ProductViceService {
	s.Ctx = ctx
	return s
}
func (s ProductViceService) GetVoicePublishRecord(voiceNo string, productKey string) ([]*entitys.VoicePublishRecordRes, error) {
	res, err := rpc.ClientOpmVoicePublishRecordService.Lists(s.Ctx, &protosService.OpmVoicePublishRecordListRequest{
		Page:     1,
		PageSize: 100000000,
		Query: &protosService.OpmVoicePublishRecord{
			ProductKey: productKey,
			VoiceNo:    voiceNo,
		},
		OrderDesc: "desc",
		OrderKey:  "created_at",
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var list = make([]*entitys.VoicePublishRecordRes, 0)
	for _, v := range res.Data {
		var arr = make([]string, 0)
		if v.AttrJson != "" {
			iotutil.JsonToStruct(v.AttrJson, &arr)
		}
		list = append(list, &entitys.VoicePublishRecordRes{
			ProductName: v.ProductName,
			VoiceName:   v.VoiceName,
			AttrText:    arr,
			CreatedAt:   v.CreatedAt.AsTime().Unix(),
			Id:          iotutil.ToString(v.Id),
		})
	}
	return list, nil
}

// 获取语控的操作指引
func (s ProductViceService) GetVoiceDoc(voiceNo string) (string, error) {
	res, err := rpc.ClientOpmVoiceService.Find(s.Ctx, &protosService.OpmVoiceFilter{
		VoiceNo: voiceNo,
	})
	if err != nil && err.Error() != "record not found" {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	doc := ""
	if res.Data != nil && len(res.Data) > 0 {
		doc = res.Data[0].VoiceDoc
	}
	return doc, nil
}

// 产品语控配置发布
func (s ProductViceService) Publish(req protosService.OpmVoiceProductPublishReq) (string, error) {
	dataRes, err := s.GetDetail(protosService.OpmVoiceProductDetailReq{
		ProductKey: req.ProductKey,
		VoiceNo:    req.VoiceNo,
	})
	if err != nil {
		return "", err
	}
	//TODO 发布之前可以对details进行检查

	res, err := rpc.ClientOpmProductVoiceService.Publish(s.Ctx, &req)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//刷新产品更新时间
	if req.ProductKey != "" {
		rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{ProductKey: req.ProductKey})
		cached.ClearCachedByKeys(context.Background(), []string{
			persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, req.VoiceNo, req.ProductKey),
			persist.GetRedisKey(iotconst.VOICE_PRODUCT_SHILL_CACHED, req.VoiceNo, dataRes.VoiceSkill),
		}...)
	}
	return "success", nil
}

// 保存产品语控配置
func (s ProductViceService) Save(req protosService.OpmVoiceProductSaveReq) (string, error) {
	// dpid去重
	checked := map[int32]struct{}{}
	for i := range req.AttrList {
		if _, ok := checked[req.AttrList[i].AttrDpid]; ok {
			return "", fmt.Errorf("dpid：%v 功能点：%v 映射重复", req.AttrList[i].AttrDpid, req.AttrList[i].FunName)
		} else {
			checked[req.AttrList[i].AttrDpid] = struct{}{}
		}
	}
	res, err := rpc.ClientOpmProductVoiceService.Save(s.Ctx, &req)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//刷新产品更新时间
	if req.ProductKey != "" {
		cached.ClearCachedByKeys(context.Background(), []string{
			persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, req.VoiceNo, req.ProductKey),
			persist.GetRedisKey(iotconst.VOICE_PRODUCT_SHILL_CACHED, req.VoiceNo, req.VoiceSkill),
		}...)
		rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{ProductKey: req.ProductKey})
	}

	return res.Data, nil
}

// 语气产品语控列表
func (s ProductViceService) GetList(req protosService.OpmVoiceProductListReq) ([]*protosService.OpmVoiceProductItem, error) {
	res, err := rpc.ClientOpmProductVoiceService.GetList(s.Ctx, &req)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var resList = make([]*protosService.OpmVoiceProductItem, 0)
	if len(res.Data) > 0 {
		resList = res.Data
	}
	return resList, nil
}

// 获取详细
func (s ProductViceService) GetDetail(req protosService.OpmVoiceProductDetailReq) (*protosService.OpmVoiceProductDetailRes, error) {
	res, err := rpc.ClientOpmProductVoiceService.GetDetail(s.Ctx, &req)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res, nil
}
