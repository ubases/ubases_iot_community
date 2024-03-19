package services

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type IntroduceService struct {
	Ctx context.Context
}

func (s IntroduceService) SetContext(ctx context.Context) IntroduceService {
	s.Ctx = ctx
	return s
}

// GetIntroduceDetail 获取文档详情
//func (s IntroduceService) GetIntroduceDetail(id int64) (*entitys.CmsIntroduceEntitys, error) {
//	if id == 0 {
//		return nil, errors.New("id not found")
//	}
//	req, err := rpc.IntroduceService.FindById(s.Ctx, &protosService.CmsIntroduceFilter{Id: id})
//	if err != nil {
//		return nil, err
//	}
//	if req.Code != 200 {
//		return nil, errors.New(req.Message)
//	}
//	if len(req.Data) == 0 {
//		return nil, errors.New("未获取到协议内容")
//	}
//	var data = req.Data[0]
//	return entitys.CmsIntroduceDetail_pb2e(data), err
//}

// GetAppIntroduceByCode 通过编码集合获取文档
//func (s IntroduceService) GetAppIntroduceByCode(codes []string) (result []*entitys.CmsIntroduceEntitys, err error) {
//	if codes == nil || len(codes) == 0 {
//		return nil, errors.New("codes not found")
//	}
//
//	for _, code := range codes {
//		req, err := rpc.IntroduceService.Find(s.Ctx, &protosService.CmsIntroduceFilter{Code: code})
//		if err != nil {
//			return nil, err
//		}
//		if req.Code != 200 && req.Message != "record not found" {
//			return nil, errors.New(req.Message)
//		} else if req.Code != 200 && req.Message == "record not found" {
//			continue
//		}
//		if len(req.Data) == 0 {
//			return nil, errors.New("未获取到协议内容")
//		}
//		var data = req.Data[0]
//		if data.Id != 0 {
//			result = append(result, entitys.CmsIntroduce_pb2e(data))
//		}
//	}
//	return result, nil
//}

// GetIntroduceDetailByCode 通过编码获取文档详情
func (s IntroduceService) GetIntroduceDetailByCode(appKey string, code string, lang string) (*entitys.CmsIntroduceEntitys, error) {
	ct := 0
	title := ""
	if code == "UserAgreement" {
		ct = 1
		title = "用户协议"
	}
	if code == "PrivacyPolicy" {
		ct = 2
		title = "隐私政策"
	}
	if code == "AboutUs" {
		ct = 3
		title = "关于我们"
	}

	req, err := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		Status:      1,
		Lang:        lang,
		ContentType: int32(ct),
		AppKey:      appKey,
	})

	//req, err := rpc.IntroduceService.Find(s.Ctx, &protosService.CmsIntroduceFilter{Code: code, Status: 1})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 && req.Message != "record not found" {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("未获取到协议内容")
	}
	var data = req.Data[0]
	res := &entitys.CmsIntroduceEntitys{
		Id:          data.Id,
		Title:       title,
		ContentMode: data.ContentType,
		Content:     data.Content,
		ContentUrl:  data.ContentUrl,
		Lang:        data.Lang,
		Status:      data.Status,
	}
	return res, err
}
