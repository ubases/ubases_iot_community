package services

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"

	"go-micro.dev/v4/metadata"
)

type QuestionTypeService struct {
	Ctx context.Context
}

func (s *QuestionTypeService) SetContext(ctx context.Context) *QuestionTypeService {
	s.Ctx = ctx
	return s
}

func (s QuestionTypeService) QueryQuestionTypeList() ([]*entitys.QuestioTypeItem, error) {
	var resultList = []*entitys.QuestioTypeItem{}
	lang, _ := metadata.Get(s.Ctx, "lang")
	appKey, _ := metadata.Get(s.Ctx, "appKey")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	if appKey == "" {
		iotlogger.LogHelper.Errorf("appKey is empty")
		return nil, errors.New(("参数异常"))
	}
	//获取APP信息
	appInfo, err := s.GetAppInfo(appKey)
	if err != nil {
		return nil, errors.New(("参数异常，AppInfo"))
	}

	appDocRelationRep, err := rpc.ClientOemAppDocRelationService.Find(s.Ctx, &protosService.OemAppDocRelationFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("QueryQuestionTypeList error,%s", err.Error())
		return resultList, nil
	}
	if appDocRelationRep.Code != 200 || len(appDocRelationRep.Data) == 0 { //公版帮助中心数据
		res, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
			Query:    &protosService.SysAppDocDir{AppTemplateType: appInfo.AppType, AppTemplateVersion: appInfo.AppTemplateVersion},
			OrderKey: "sort",
		})
		if err != nil {
			return resultList, nil
		}
		if res.Code != 200 {
			return resultList, nil
		}

		cacheKey := fmt.Sprintf("%s%s", iotconst.HKEY_LANGUAGE_DATA_PREFIX, iotconst.LANG_T_SYS_APP_DOC_DIR)
		langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

		for _, item := range res.Data {
			if item.ParentId != 0 {
				continue
			}
			name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%d_dir_name", lang, item.Id)], item.DirName)
			resultList = append(resultList, &entitys.QuestioTypeItem{
				Id:       item.Id,
				Name:     name,
				ParentId: 0,
				ListImg:  item.DirImg,
				Sort:     item.Sort,
			})
		}
	} else { //开发者自己创建的帮助中心的数据
		docId := appDocRelationRep.Data[0].DocId
		appDocDirRep, err := rpc.ClientOemAppDocDirService.Lists(context.Background(), &protosService.OemAppDocDirListRequest{
			Query: &protosService.OemAppDocDir{
				DocId: docId,
			},
			OrderKey: "sort",
		})
		if err != nil {
			iotlogger.LogHelper.Errorf("QueryQuestionTypeList error,%s", err.Error())
			return resultList, nil
		}
		if appDocDirRep.Code != 200 {
			return resultList, nil
		}

		cacheKey := fmt.Sprintf("%s_%s%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX, iotconst.LANG_T_OEM_APP_DOC_DIR)
		langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
		for _, item := range appDocDirRep.Data {
			if item.ParentId != 0 {
				continue
			}
			langKey := fmt.Sprintf("%s_%d_dir_name", lang, item.Id)
			name := iotutil.MapGetStringVal(langMap[langKey], item.DirName)
			resultList = append(resultList, &entitys.QuestioTypeItem{
				Id:       item.Id,
				Name:     name,
				ParentId: 0,
				ListImg:  item.DirImg,
				Sort:     item.Sort,
			})
		}
	}
	return resultList, err
}

// TranslateGet 获取指定的翻译详细（业务翻译获取）
func (s QuestionTypeService) TranslateGet(dirId int64, lang string) ([]*entitys.SysAppDocDirLang, error) {
	ret, err := rpc.ClientLangTranslateService.Lists(context.Background(), &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable: "t_sys_app_doc_dir",
			SourceRowId: iotutil.ToString(dirId),
		},
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	var res []*entitys.SysAppDocDirLang
	for _, data := range ret.Data {
		if lang != "" {
			if data.Lang == lang {
				res = append(res, &entitys.SysAppDocDirLang{
					Lang:    data.Lang,
					DirName: data.FieldValue,
				})
			}
		} else {
			res = append(res, &entitys.SysAppDocDirLang{
				Lang:    data.Lang,
				DirName: data.FieldValue,
			})
		}
	}
	return res, nil
}

//func (s QuestionTypeService) QueryQuestionTypeListBack(pid string) ([]*entitys.QuestioTypeItem, error) {
//	//默认查询顶级节点
//	if pid == "" {
//		pid = "0"
//	}
//	rep, err := rpc.QuestionTypeService.Lists(context.Background(), &protosService.CmsQuestionTypeListRequest{
//		Query: &protosService.CmsQuestionType{ParentId: pid},
//	})
//	if err != nil {
//		return nil, err
//	}
//	if rep.Code != 200 {
//		return nil, errors.New(rep.Message)
//	}
//	var resultList = []*entitys.QuestioTypeItem{}
//	for _, item := range rep.Data {
//		resultList = append(resultList, &entitys.QuestioTypeItem{
//			Id:       item.Id,
//			Name:     item.Title,
//			ParentId: iotutil.ToInt(pid),
//			ListImg:  item.ImageUrl,
//		})
//	}
//	return resultList, err
//}

// 获取APP信息
func (s QuestionTypeService) GetAppInfo(appKey string) (*protosService.OemApp, error) {
	res, err := rpc.ClientOemAppService.Find(s.Ctx, &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data[0], nil
}
