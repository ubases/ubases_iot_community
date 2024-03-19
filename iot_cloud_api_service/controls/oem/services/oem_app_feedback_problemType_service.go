package services

import (
	langEntitys "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppFeedbackProblemTypeService struct {
	Ctx context.Context
}

func (s OemAppFeedbackProblemTypeService) SetContext(ctx context.Context) OemAppFeedbackProblemTypeService {
	s.Ctx = ctx
	return s
}

func (s OemAppFeedbackProblemTypeService) OemAppFeedbackProblemTypeAdd(req entitys.OemFeedbackTypeSaveReq) (string, error) {
	id := iotutil.GetNextSeqInt64()
	resApp, errApp := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: iotutil.ToInt64(req.AppId),
	})
	if errApp != nil {
		return "", errApp
	}
	res, err := rpc.ClientFeedbackTypeService.Create(s.Ctx, &protosService.OemFeedbackType{
		Id:   id,
		Name: req.Name,
		//Name:  	     s.GetDefaultLangDirLang(req.DocDirLangs).DirName,
		AppId:     iotutil.ToInt64(req.AppId),
		AppKey:    resApp.Data[0].AppKey,
		Sort:      req.Sort,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//剩下的插入到翻译表去.
	req.Id = iotutil.ToString(id)

	//获取翻译参数
	reqLang, errLang := s.ConvertLangTranslateEntity(req)
	if errLang != nil {
		return "", errLang
	}
	//保存翻译
	resTran, errTran := s.TranslateSave(*reqLang)
	if errTran != nil {
		return resTran, errTran
	}

	return iotutil.ToString(id), nil
}

// 修改
func (s OemAppFeedbackProblemTypeService) OemAppFeedbackProblemTypeUpdate(req entitys.OemFeedbackTypeSaveReq) (string, error) {
	resFind, errFind := rpc.ClientFeedbackTypeService.Update(s.Ctx, &protosService.OemFeedbackType{
		Id:   iotutil.ToInt64(req.Id),
		Name: req.Name,
		Sort: req.Sort,
	})
	if errFind != nil {
		return "", errFind
	}
	if resFind.Code != 200 {
		return "", errors.New(resFind.Message)
	}

	//剩下的去修改翻译表

	//获取翻译参数
	reqLang, errLang := s.ConvertLangTranslateEntity(req)
	if errLang != nil {
		return "", errLang
	}
	//保存翻译
	resTran, errTran := s.TranslateSave(*reqLang)
	if errTran != nil {
		return resTran, errTran
	}
	return "success", nil
}

// 获取列表.
func (s OemAppFeedbackProblemTypeService) OemAppFeedbackProblemTypeList(req entitys.OemFeedbackTypeEntitys) ([]*entitys.OemFeedbackTypeListRes, int64, error) {
	var resultList = []*entitys.OemFeedbackTypeListRes{}
	appId, _ := iotutil.ToInt64AndErr(req.AppId)
	res, err := rpc.ClientFeedbackTypeService.Lists(s.Ctx, &protosService.OemFeedbackTypeListRequest{
		Query:     &protosService.OemFeedbackType{AppId: appId},
		OrderKey:  "sort",
		OrderDesc: "asc",
	})
	if err != nil {
		return resultList, 0, nil
	}
	if res.Code != 200 {
		return resultList, 0, nil
	}

	if res.Data == nil || len(res.Data) == 0 {
		return resultList, 0, nil
	}

	for _, v := range res.Data {
		resultList = append(resultList, &entitys.OemFeedbackTypeListRes{
			Id:        iotutil.ToString(v.Id),
			UpdatedAt: v.UpdatedAt.AsTime().Unix(),
			Name:      v.Name,
			Sort:      v.Sort,
			AppId:     iotutil.ToString(v.AppId),
		})
	}
	return resultList, res.Total, nil
}

// 获取详细.
func (s OemAppFeedbackProblemTypeService) OemAppFeedbackProblemTypeDetail(id int64) (*entitys.OemFeedbackTypeEntitys, error) {

	res, err := rpc.ClientFeedbackTypeService.FindById(s.Ctx, &protosService.OemFeedbackTypeFilter{
		Id: id,
	})
	if err != nil {
		return nil, nil
	}
	if res.Code != 200 {
		return nil, nil
	}

	var rs = entitys.OemFeedbackTypeEntitys{}
	rs.DocDirLangs = make([]*entitys.OemAppDocDirLang, 0)
	if res.Data != nil && len(res.Data) > 0 {
		tmp := res.Data[0]
		rs.Id = tmp.Id
		rs.Name = tmp.Name
		rs.AppId = iotutil.ToString(tmp.AppId)
		rs.Sort = tmp.Sort
	}
	//获取支持的语种
	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	//去翻译表找语种翻译数据
	resTran, errTran := s.TranslateGet(iotconst.LANG_T_OEM_FEEDBACK_TYPE, iotutil.ToInt64(id), "")
	if errTran != nil {
		return nil, errTran
	}

	//to map
	docDirMap := make(map[string]*entitys.OemAppDocDirLang)
	for _, d := range resTran {
		docDirMap[d.Lang] = d
	}

	//处理排序
	var results []*entitys.OemAppDocDirLang
	for _, langType := range langTypes {
		if langType.Code == "zh" {
			continue
		}
		if v, ok := docDirMap[langType.Code]; ok {
			results = append(results, &entitys.OemAppDocDirLang{
				DirName:  v.DirName,
				Lang:     langType.Code,
				LangName: langType.Name,
			})
		} else {

			results = append(results, &entitys.OemAppDocDirLang{
				DirName:  "",
				Lang:     langType.Code,
				LangName: langType.Name,
			})
		}
	}
	rs.DocDirLangs = results

	return &rs, nil
}

// 删除
func (s OemAppFeedbackProblemTypeService) DeleteOemAppFeedbackProblemType(id int64) (*protosService.OemAppIntroduce, error) {
	res, err := rpc.ClientFeedbackTypeService.Delete(s.Ctx, &protosService.OemFeedbackType{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	return nil, nil
}

// 获取中文目录名称
func (s OemAppFeedbackProblemTypeService) GetDefaultLangDirLang(req []*entitys.OemAppDocDirLang) entitys.OemAppDocDirLang {
	rs := entitys.OemAppDocDirLang{}
	for _, v := range req {
		if v.Lang == "zh" {
			rs = *v
			break
		}
	}
	return rs
}

func (s OemAppFeedbackProblemTypeService) ConvertLangTranslateEntity(req entitys.OemFeedbackTypeSaveReq) (*langEntitys.LangTranslateEntitys, error) {
	var rs = langEntitys.LangTranslateEntitys{}
	rs.SourceTable = iotconst.LANG_T_OEM_FEEDBACK_TYPE
	rs.SourceRowId = req.Id

	for _, v := range req.DocDirLangs {
		rs.TranslateList = append(rs.TranslateList, langEntitys.BatchSaveTranslateItem{
			Lang:       v.Lang,
			FieldName:  "name",
			FieldType:  0,
			FieldValue: v.DirName,
		})
	}
	return &rs, nil
}

// TranslateSave 翻译保存（业务翻译保存）
func (s OemAppFeedbackProblemTypeService) TranslateSave(req langEntitys.LangTranslateEntitys) (string, error) {
	translateList := []*protosService.BatchSaveTranslateItem{}
	for _, item := range req.TranslateList {
		translateItem := protosService.BatchSaveTranslateItem{
			Lang:       item.Lang,
			FieldName:  item.FieldName,
			FieldType:  item.FieldType,
			FieldValue: item.FieldValue,
		}
		if item.Id != "" {
			translateItem.Id = iotutil.ToInt64(item.Id)
		}
		translateList = append(translateList, &translateItem)
	}
	ret, err := rpc.ClientLangTranslateService.BatchCreate(s.Ctx, &protosService.BatchSaveTranslate{
		SourceRowId:  req.SourceRowId,
		SourceTable:  req.SourceTable,
		List:         translateList,
		PlatformType: int32(iotconst.OPEN_USER),
	})
	if err != nil {
		return "", err
	}
	if ret.Code != 200 {
		return "", errors.New(ret.Message)
	}
	return "success", nil
}

// TranslateGet 获取指定的翻译详细（业务翻译获取）
func (s OemAppFeedbackProblemTypeService) TranslateGet(tabName string, dirId int64, lang string) ([]*entitys.OemAppDocDirLang, error) {
	ret, err := rpc.ClientLangTranslateService.Lists(context.Background(), &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable: tabName,
			SourceRowId: iotutil.ToString(dirId),
		},
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	var res []*entitys.OemAppDocDirLang
	for _, data := range ret.Data {
		if lang != "" {
			if data.Lang == lang {
				res = append(res, &entitys.OemAppDocDirLang{
					Lang:    data.Lang,
					DirName: data.FieldValue,
				})
			}
		} else {
			res = append(res, &entitys.OemAppDocDirLang{
				Lang:    data.Lang,
				DirName: data.FieldValue,
			})
		}

	}
	return res, nil
}
