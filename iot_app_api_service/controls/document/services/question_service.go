package services

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type QuestionService struct {
}

func (s QuestionService) QueryQuestionTop5(appKey, lang string) ([]*entitys.QuestionRes, error) {
	var rs = make([]*entitys.QuestionRes, 0)
	relation, err := s.getIsRelationEntry(appKey)
	if err != nil {
		return nil, err
	}
	if len(relation) == 0 { //公版帮助中心数据
		res, err := s.getPublicEntry(&protosService.SysAppEntryListDiyRequqest{
			DirId:    0,
			Lang:     lang,
			IsNormal: 1, //是否设置为常见 1-常见  2-不常见
			IsEnable: 1, //是否启用 1-启用  2-不启用
			Page:     1,
			PageSize: 5,
		})
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		rs, err := s.getOemAppEntry(&protosService.OemAppEntryListDiyRequqest{
			DocId:    iotutil.ToInt64(relation[0].DocId),
			Lang:     lang,
			IsNormal: 1, //是否设置为常见 1-常见  2-不常见
			IsEnable: 1, //是否启用 1-启用  2-不启用
			Page:     1,
			PageSize: 5,
		})
		if err != nil {
			return nil, err
		}
		return rs, nil
	}
	return rs, nil
}

// QueryQuestionList 查询帮助文档列表和产品文档列表
func (s QuestionService) QueryQuestionList(req entitys.QuestionQuery, appKey, lang, tenantId string) ([]*entitys.QuestionRes, error) {
	var rs = make([]*entitys.QuestionRes, 0)
	//如果productKey=“”则显示帮助文档，否则显示产品文档
	if req.Model == "" {
		typeId, _ := iotutil.ToInt64AndErr(req.TypeId)
		relation, err := s.getIsRelationEntry(appKey)
		if err != nil {
			return nil, err
		}
		if len(relation) == 0 { //公版帮助中心数据
			rs, err := s.getPublicEntry(&protosService.SysAppEntryListDiyRequqest{
				DirId:    typeId,
				Lang:     lang,
				Title:    req.Title,
				IsNormal: 0, //是否设置为常见 1-常见  2-不常见
				IsEnable: 1, //是否启用 1-启用  2-不启用
			})
			if err != nil {
				return nil, err
			}
			return rs, nil
		} else {
			rs, err := s.getOemAppEntry(&protosService.OemAppEntryListDiyRequqest{
				DirId:    typeId,
				DocId:    iotutil.ToInt64(relation[0].DocId),
				Lang:     lang,
				Title:    req.Title,
				IsNormal: 0, //是否设置为常见 1-常见  2-不常见
				IsEnable: 1, //是否启用 1-启用  2-不启用
			})
			if err != nil {
				return nil, err
			}
			return rs, nil
		}
	} else {
		//读取产品文档
		if req.Model == "" {
			return nil, errors.New("productKey为空")
		}
		res, err := rpc.ProductHelpDocService.Lists(context.Background(), &protosService.ProductHelpDocListRequest{Query: &protosService.ProductHelpDoc{
			TenantId:   tenantId,
			ProductKey: req.Model,
			Lang:       lang,
			Title:      req.Title,
			Status:     1,
		}})
		if err != nil {
			return rs, err
		}
		if len(res.Data) == 0 {
			return rs, nil
		}
		for _, productHelpDoc := range res.Data {
			rs = append(rs, &entitys.QuestionRes{
				SetingId: iotutil.ToString(productHelpDoc.Id),
				Title:    productHelpDoc.Title,
			})
		}
	}
	return rs, nil
}

func (s QuestionService) getIsRelationEntry(appKey string) ([]*protosService.OemAppDocRelation, error) {
	appDocRelationRep, err := rpc.ClientOemAppDocRelationService.Find(context.Background(), &protosService.OemAppDocRelationFilter{
		AppKey: appKey,
	})
	if err != nil && err.Error() != "record not found" {
		iotlogger.LogHelper.Errorf("getIsRelationEntry error,%s", err.Error())
		return nil, err
	}
	if appDocRelationRep.Code != 200 && appDocRelationRep.Message != "record not found" {
		return nil, errors.New(appDocRelationRep.Message)
	}
	return appDocRelationRep.Data, nil
}

// OEM的帮助文档
func (s QuestionService) getOemAppEntry(request *protosService.OemAppEntryListDiyRequqest) ([]*entitys.QuestionRes, error) {
	var rs = make([]*entitys.QuestionRes, 0)
	res, err := rpc.ClientOemAppEntryService.ListDiy(context.Background(), request)
	if err != nil {
		iotlogger.LogHelper.Errorf("QueryQuestionList error,%s", err.Error())
		return rs, nil
	}
	if res.Code != 200 {
		return rs, nil
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return rs, nil
	}
	for _, v := range res.Data {
		rs = append(rs, &entitys.QuestionRes{
			SetingId:  iotutil.ToString(v.SetingId),
			DocId:     iotutil.ToString(v.DocId),
			DirId:     iotutil.ToString(v.DirId),
			Lang:      v.Lang,
			Title:     v.Title,
			IsNormal:  v.IsNormal,
			IsEnable:  v.IsEnable,
			DirName:   v.DirName,
			Sort:      v.Sort,
			UpdatedAt: v.UpdatedAt.AsTime().Unix(),
		})
	}
	return rs, nil
}

// 获取公版帮助文档
func (s QuestionService) getPublicEntry(request *protosService.SysAppEntryListDiyRequqest) ([]*entitys.QuestionRes, error) {
	res, err := rpc.ClientSysAppEntryService.ListDiy(context.Background(), request)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var rs = make([]*entitys.QuestionRes, 0)
	if res.Data == nil || len(res.Data) <= 0 {
		return rs, nil
	}
	for _, v := range res.Data {
		rs = append(rs, &entitys.QuestionRes{
			SetingId:  iotutil.ToString(v.SetingId),
			DirId:     iotutil.ToString(v.DirId),
			Lang:      v.Lang,
			Title:     v.Title,
			IsNormal:  v.IsNormal,
			IsEnable:  v.IsEnable,
			DirName:   v.DirName,
			Sort:      v.Sort,
			UpdatedAt: v.UpdatedAt.AsTime().Unix(),
		})
	}
	return rs, nil
}

func (s QuestionService) GetQuestionDetail(id string, queryType int32, appKey, lang string) (*entitys.QuestionItem, error) {
	result := &entitys.QuestionItem{}
	if queryType != 2 { //常见问题
		relation, err := s.getIsRelationEntry(appKey)
		if err != nil {
			return nil, err
		}
		if len(relation) == 0 { //公版帮助中心数据
			return s.getPublicEntryById(id, lang)
		} else {
			//如何判断是公版
			qId := iotutil.ToInt64(id)
			res, err := rpc.ClientOemAppEntryService.Find(context.Background(), &protosService.OemAppEntryFilter{
				Lang:     lang,
				SetingId: qId,
			})
			if err != nil {
				iotlogger.LogHelper.Errorf("GetQuestionDetail error,%s", err.Error())
				return result, err
			}
			if res.Code != 200 {
				return result, nil
			}
			if res.Data == nil || len(res.Data) <= 0 {
				return result, nil
			}
			questionInfo := res.Data[0]
			result = &entitys.QuestionItem{
				Id:      questionInfo.Id,
				Lang:    questionInfo.Lang,
				Title:   questionInfo.Title,
				Content: questionInfo.Content,
			}
		}
	} else { //产品帮助中心详细
		res, err := rpc.ProductHelpDocService.FindById(context.Background(), &protosService.ProductHelpDocFilter{Id: iotutil.ToInt64(id)})
		if err != nil {
			return result, err
		}
		if len(res.Data) == 0 {
			return result, nil
		}
		productHelpDocInfo := res.Data[0]

		result = &entitys.QuestionItem{
			Id:      productHelpDocInfo.Id,
			Lang:    productHelpDocInfo.Lang,
			Title:   productHelpDocInfo.Title,
			Content: productHelpDocInfo.Content,
		}
	}
	return result, nil
}

// 获取公版文档详情
func (s QuestionService) getPublicEntryById(id, lang string) (*entitys.QuestionItem, error) {
	settingId, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		return nil, err
	}
	res, err := rpc.ClientSysAppEntryService.Find(context.Background(), &protosService.SysAppEntryFilter{
		Lang:     lang,
		SetingId: settingId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, nil
	}
	questionInfo := res.Data[0]
	result := &entitys.QuestionItem{
		Id:      questionInfo.Id,
		Lang:    questionInfo.Lang,
		Title:   questionInfo.Title,
		Content: questionInfo.Content,
	}
	return result, nil
}
