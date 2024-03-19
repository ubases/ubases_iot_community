package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotstrings"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OemAppIntroduceService struct {
	Ctx context.Context
}

func (s OemAppIntroduceService) SetContext(ctx context.Context) OemAppIntroduceService {
	s.Ctx = ctx
	return s
}

func (s OemAppIntroduceService) OemAppIntroduceCheckVersion(req entitys.OemAppIntroduceVersionReq) (bool, error) {
	res, err := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		AppId:       iotutil.ToInt64(req.AppId),
		ContentType: req.ContentType,
	})

	if err != nil {
		return false, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return false, errors.New(res.Message)
	}
	//没有数据. 可以添加此版本的数据
	if len(res.Data) <= 0 {
		return true, nil
	}

	result := true
	for _, v := range res.Data {
		if !iotstrings.VersionCompared(req.Version, v.Version) {
			result = false
			break
		}
	}
	return result, nil
}
func (s OemAppIntroduceService) OemAppIntroduceAdd(req entitys.OemAppIntroduceSaveReq) (string, error) {
	id := iotutil.GetNextSeqInt64()

	resApp, errApp := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: iotutil.ToInt64(req.AppId),
	})
	if errApp != nil {
		return "", errApp
	}
	res, err := rpc.ClientOemAppIntroduceService.Create(s.Ctx, &protosService.OemAppIntroduce{
		Id:          id,
		AppId:       iotutil.ToInt64(req.AppId),
		Version:     req.Version,
		Status:      2, //1 已启用, 2 未启用,3 已失效
		Lang:        req.Lang,
		ContentType: req.ContentType,
		ContentUrl:  "",
		Content:     req.Content,
		AppKey:      resApp.Data[0].AppKey,
		VoiceCode:   req.VioceCode,
		Abstract:    req.Abstract,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(id), nil
}

// 创建
func (s OemAppIntroduceService) GetOemAppIntroduce(appId string, version string, contentType int32, lang string) (*protosService.OemAppIntroduce, error) {
	res, err := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		AppId:       iotutil.ToInt64(appId),
		Version:     version,
		Lang:        lang,
		ContentType: contentType,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	//没有数据. 可以添加此版本的数据
	if len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data[0], nil
}

// 根据appid和版本和类型 获取多个文档记录
func (s OemAppIntroduceService) GetOemAppIntroduceListBySel(appId string, version string, contentType int32) ([]*protosService.OemAppIntroduce, error) {
	res, err := rpc.ClientOemAppIntroduceService.Lists(s.Ctx, &protosService.OemAppIntroduceListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.OemAppIntroduce{
			AppId:       iotutil.ToInt64(appId),
			Version:     version,
			ContentType: contentType,
		},
	})
	// res,err := rpc.ClientOemAppIntroduceService.Lists(s.Ctx,&protosService.OemAppIntroduceFilter{
	// 	AppId: iotutil.ToInt64(appId),
	// 	Version: version,
	// 	ContentType: contentType,
	// })
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	//没有数据. 可以添加此版本的数据
	if len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data, nil
}

// 把状态等于启用的版本状态修改为 未启用.
func (s OemAppIntroduceService) UpdateStatusDisable(appId string, contentType int32) error {
	res, err := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		AppId:       iotutil.ToInt64(appId),
		ContentType: contentType,
		Status:      1,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return errors.New(res.Message)
	}
	//没有数据. 可以添加此版本的数据
	// if len(res.Data) <=0 {
	// 	return errors.New("参数错误,未找到记录")
	// }

	for _, v := range res.Data {
		res, err := rpc.ClientOemAppIntroduceService.UpdateFields(s.Ctx, &protosService.OemAppIntroduceUpdateFieldsRequest{
			Fields: []string{"status"},
			Data: &protosService.OemAppIntroduce{
				Id:     v.Id,
				Status: 3,
			},
		})
		if err != nil {
			return err
		}
		if res.Code != 200 && res.Message != "record not found" {
			return errors.New(res.Message)
		}
	}
	return nil
}

// 修改
func (s OemAppIntroduceService) OemAppIntroduceUpdate(req entitys.OemAppIntroduceSaveReq) (string, error) {
	resFind, errFind := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		AppId:       iotutil.ToInt64(req.AppId),
		Version:     req.Version,
		ContentType: req.ContentType,
		Lang:        req.Lang,
		VoiceCode:   req.VioceCode,
	})
	if errFind != nil {
		return "", errFind
	}
	if resFind.Code != 200 && resFind.Message != "record not found" {
		return "", errors.New(resFind.Message)
	}
	//没有找到记录. 需要去新增
	if len(resFind.Data) == 0 {
		resAdd, errAdd := s.OemAppIntroduceAdd(entitys.OemAppIntroduceSaveReq{
			AppId:       req.AppId,
			ContentType: req.ContentType,
			Content:     req.Content,
			Lang:        req.Lang,
			Version:     req.Version,
			VioceCode:   req.VioceCode,
			Abstract:    req.Abstract,
		})
		return resAdd, errAdd
	} else {
		var fields = []string{"content"}
		if req.ContentType == 4 {
			fields = append(fields, "abstract")
		}
		res, err := rpc.ClientOemAppIntroduceService.UpdateFields(s.Ctx, &protosService.OemAppIntroduceUpdateFieldsRequest{
			Fields: fields,
			Data: &protosService.OemAppIntroduce{
				Id:       resFind.Data[0].Id,
				Content:  req.Content,
				Abstract: req.Abstract,
			},
		})
		if err != nil {
			return "", err
		}
		if res.Code != 200 && res.Message != "record not found" {
			return "", errors.New(res.Message)
		}
		return "success", nil
	}
}

// 获取详细.
func (s OemAppIntroduceService) OemAppIntroduceDetail(req entitys.OemAppIntroduceDetailReq) (*entitys.OemAppIntroduceDetailRes, error) {

	res, err := rpc.ClientOemAppIntroduceService.Find(s.Ctx, &protosService.OemAppIntroduceFilter{
		AppId:       iotutil.ToInt64(req.AppId),
		Version:     req.Version,
		Lang:        req.Lang,
		ContentType: req.ContentType,
		VoiceCode:   req.VioceCode,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var result = entitys.OemAppIntroduceDetailRes{}
	if len(res.Data) <= 0 {
		result.AppId = req.AppId
		result.Content = ""
		result.Version = req.Version
		result.Lang = req.Lang
		result.ContentType = req.ContentType
		result.VioceCode = req.VioceCode
		result.Abstract = ""
	} else {
		result.AppId = iotutil.ToString(res.Data[0].AppId)
		result.Content = res.Data[0].Content
		result.Version = res.Data[0].Version
		result.Lang = res.Data[0].Lang
		result.ContentType = res.Data[0].ContentType
		result.VioceCode = res.Data[0].VoiceCode
		result.Abstract = res.Data[0].Abstract
	}
	return &result, nil
}

func (s OemAppIntroduceService) OemAppIntroduceStatusEnable(req entitys.OemAppIntroduceStatusReq) (string, error) {
	appId, err := iotutil.ToInt64AndErr(req.AppId)
	if err != nil {
		return "", errors.New("appId error")
	}
	//启用文档
	res, err := rpc.ClientOemAppIntroduceService.Enable(s.Ctx, &protosService.OemAppIntroduce{
		AppId:       appId,
		ContentType: req.ContentType,
		Version:     req.Version,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//
	////把目前已经启动的状态变更为未启用.(只能有一个版本属于启用的状态)
	//errDis := s.UpdateStatusDisable(req.AppId, req.ContentType)
	//if errDis != nil {
	//	return "", errDis
	//}
	////把该版本下所有的语种记录都改为已启用.
	//d, errSel := s.GetOemAppIntroduceListBySel(req.AppId, req.Version, req.ContentType)
	//if errSel != nil {
	//	return "", errSel
	//}
	//for _, v := range d {
	//	res, err := rpc.ClientOemAppIntroduceService.UpdateFields(s.Ctx, &protosService.OemAppIntroduceUpdateFieldsRequest{
	//		Fields: []string{"status"},
	//		Data: &protosService.OemAppIntroduce{
	//			Id:     v.Id,
	//			Status: 1,
	//		},
	//	})
	//	if err != nil {
	//		return "", err
	//	}
	//	if res.Code != 200 && res.Message != "record not found" {
	//		return "", errors.New(res.Message)
	//	}
	//}
	return "success", nil

}

// 获取协议文档链接的列表
func (s OemAppIntroduceService) OemAppIntroduceDetailById(id string) (*entitys.OemAppIntroduceEntitys, error) {
	res, err := rpc.ClientOemAppIntroduceService.FindById(s.Ctx, &protosService.OemAppIntroduceFilter{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if len(res.Data) > 0 {
		result := entitys.OemAppIntroduce_pb2e(res.Data[0])
		return result, nil
	}
	return nil, nil
}

// 获取协议文档链接的列表
func (s OemAppIntroduceService) OemAppIntroduceLinkList(req entitys.OemAppIntroduceListReq) ([]*entitys.OemAppIntroduceLinkRes, error) {
	res, err := rpc.ClientOemAppIntroduceService.Lists(s.Ctx, &protosService.OemAppIntroduceListRequest{
		Page:     1,
		PageSize: 10000,
		Query: &protosService.OemAppIntroduce{
			AppId:       iotutil.ToInt64(req.AppId),
			ContentType: req.ContentType,
			Status:      1,
		},
	})
	if err != nil {
		return nil, err
	}

	var result = make([]*entitys.OemAppIntroduceLinkRes, 0)
	mp := GetBaseDataValue("oem_app_package_domain", s.Ctx)
	domain := iotutil.ToString(mp[GetOemAppEnv()])

	for _, v := range res.Data {
		var tmp = entitys.OemAppIntroduceLinkRes{}
		tmp.Lang = v.Lang
		tmp.Url = domain + "/app/introduce/detail/" + iotutil.ToString(v.Id) + "/" + v.Lang + ".html"
		result = append(result, &tmp)

	}
	return result, nil
}

// 列表(按语种分组的列表)
func (s OemAppIntroduceService) OemAppIntroduceList(filter entitys.OemAppIntroduceListReq) ([]*entitys.OemAppIntroduceListRes, int64, error) {
	rep, err := rpc.ClientOemAppIntroduceService.IntroduceList(s.Ctx, &protosService.OemAppIntroduceListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.OemAppIntroduce{
			AppId:       iotutil.ToInt64(filter.AppId),
			ContentType: filter.ContentType,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var resultList = []*entitys.OemAppIntroduceListRes{}

	if rep.Data == nil || len(rep.Data) == 0 {

		return resultList, 0, nil
	}

	for _, v := range rep.Data {
		resultList = append(resultList, &entitys.OemAppIntroduceListRes{
			Version:     v.Version,
			CreatedAt:   int32(v.CreatedAt),
			UpdatedAt:   int32(v.UpdatedAt),
			LangCount:   v.LangCount,
			Status:      v.Status,
			AppId:       iotutil.ToString(v.AppId),
			ContentType: v.ContentType,
		})
	}
	return resultList, rep.Total, err
}

// 列表(按语种分组的列表)
func (s OemAppIntroduceService) CheckAppIntroduce(appId int64, contentType int32) (bool, error) {
	rep, err := rpc.ClientOemAppIntroduceService.IntroduceList(s.Ctx, &protosService.OemAppIntroduceListRequest{
		Query: &protosService.OemAppIntroduce{
			AppId:       appId,
			ContentType: contentType,
			Status:      1,
		},
	})
	if err != nil {
		return false, err
	}
	if rep.Code != 200 {
		return false, errors.New(rep.Message)
	}
	return len(rep.Data) > 0, err
}

// 获取协议模板内容
func (s OemAppIntroduceService) OemAppIntroduceTemplateDetailById(id string) (*entitys.OemAppIntroduceDetailRes, error) {
	res, err := rpc.ClientDocumentTemplateService.FindById(s.Ctx, &protosService.TplDocumentTemplateFilter{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if len(res.Data) > 0 {
		var resdata = entitys.OemAppIntroduceDetailRes{}
		resdata.Content = res.Data[0].TplContent
		resdata.Lang = res.Data[0].Lang

		return &resdata, nil
	}
	return &entitys.OemAppIntroduceDetailRes{}, nil
}

// 获取协议模板链接
func (s OemAppIntroduceService) OemAppIntroduceTemplateLink(contentType int) ([]*entitys.OemAppIntroduceLinkRes, error) {
	resFind, errFind := rpc.ClientDocumentTemplateService.Lists(s.Ctx, &protosService.TplDocumentTemplateListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.TplDocumentTemplate{
			TplType: iotutil.ToString(contentType),
			Status:  1,
		},
	})
	if errFind != nil && errFind.Error() != "record not found" {
		return nil, errFind
	}
	if resFind.Code != 200 && resFind.Message != "record not found" {
		return nil, errors.New(resFind.Message)
	}

	var reslist = make([]*entitys.OemAppIntroduceLinkRes, 0)

	mp := GetBaseDataValue("oem_app_package_domain", s.Ctx)
	domain := iotutil.ToString(mp[GetOemAppEnv()])
	for _, v := range resFind.Data {
		var tmp = entitys.OemAppIntroduceLinkRes{}
		tmp.Lang = v.Lang
		// "/v1/platform/web/open/oem/app/introduce/detail/"+iotutil.ToString(r.Id)+"/"+r.Lang+".html"
		tmp.Url = domain + "/app/introduce/template/detail/" + iotutil.ToString(v.Id) + "/" + v.Lang + ".html"
		reslist = append(reslist, &tmp)
	}

	return reslist, nil
}
