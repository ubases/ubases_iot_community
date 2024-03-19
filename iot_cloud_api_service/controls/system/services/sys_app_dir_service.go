package services

import (
	//langEntitys "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	langEnt "cloud_platform/iot_cloud_api_service/controls/system/entitys"

	"go-micro.dev/v4/logger"
)

type SysAppDocDirService struct {
	Ctx context.Context
}

func (s SysAppDocDirService) SetContext(ctx context.Context) SysAppDocDirService {
	s.Ctx = ctx
	return s
}

func (s SysAppDocDirService) CreateDir(req entitys.SysAppDocDirSaveReq) (string, error) {
	dirId := iotutil.GetNextSeqInt64()
	reqDir := &protosService.SysAppDocDir{
		Id:       dirId,
		HelpId:   iotutil.ToInt64(req.HelpId),
		DirName:  s.GetDefaultLangDirLang(req.DocDirLangs).DirName,
		DirImg:   req.DirImg,
		ParentId: iotutil.ToInt64(req.ParentId),
		Sort:     req.Sort,
	}
	// if len(req.HelpId) != 0 {
	// 	reqDir.HelpId = iotutil.ToInt64(req.HelpId)
	// }
	res, err := rpc.ClientSysAppDocDirService.Create(s.Ctx, reqDir)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//剩下的插入到翻译表去.
	req.Id = iotutil.ToString(dirId)

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

func (s SysAppDocDirService) UpdateDir(req entitys.SysAppDocDirSaveReq) (string, error) {
	res, err := rpc.ClientSysAppDocDirService.UpdateFields(s.Ctx, &protosService.SysAppDocDirUpdateFieldsRequest{
		Fields: []string{"dir_name", "dir_img", "sort"},
		Data: &protosService.SysAppDocDir{
			Id:      iotutil.ToInt64(req.Id),
			DirName: s.GetDefaultLangDirLang(req.DocDirLangs).DirName,
			DirImg:  req.DirImg,
			Sort:    req.Sort,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
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

func (s SysAppDocDirService) DetailDir(dirId string) (*entitys.SysAppDocDirSaveReq, error) {
	res, err := rpc.ClientSysAppDocDirService.FindById(s.Ctx, &protosService.SysAppDocDirFilter{
		Id: iotutil.ToInt64(dirId),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var rs = entitys.SysAppDocDirSaveReq{}
	rs.DocDirLangs = make([]entitys.SysAppDocDirLang, 0)
	if res.Data != nil && len(res.Data) > 0 {
		tmp := res.Data[0]
		rs.DirImg = tmp.DirImg
		//rs.DocId = iotutil.ToString(tmp.DocId)
		rs.ParentId = iotutil.ToString(tmp.ParentId)
		rs.HelpId = iotutil.ToString(tmp.HelpId)
		rs.Sort = tmp.Sort
		// rs.DocDirLangs = append(rs.DocDirLangs, entitys.SysAppDocDirLang{
		// 	DirName: tmp.DirName,
		// 	Lang: "zh",
		// 	LangName: "简体中文",
		// })
	}
	//去翻译表找语种翻译数据
	resTran, errTran := s.TranslateGet(iotutil.ToInt64(dirId), "")
	if errTran != nil {
		return nil, errTran
	}
	//rs.DocDirLangs = append(rs.DocDirLangs, &resTran)
	for _, v := range resTran {
		rs.DocDirLangs = append(rs.DocDirLangs, *v)
	}
	rs.Id = dirId

	return &rs, nil
}

func (s SysAppDocDirService) DeleteDir(dirId string) (string, error) {

	resDir, errDir := rpc.ClientSysAppDocDirService.Find(s.Ctx, &protosService.SysAppDocDirFilter{
		ParentId: iotutil.ToInt64(dirId),
	})
	if errDir != nil && errDir.Error() != "record not found" {
		return "", errDir
	}
	if resDir.Code != 200 && resDir.Message != "record not found" {
		return "", errors.New(resDir.Message)
	}
	if resDir.Data != nil && len(resDir.Data) > 0 {
		return "", errors.New("目录下存在子目录,不可删除")
	}

	resSet, errSet := rpc.ClientSysAppEntrySetingService.Find(s.Ctx, &protosService.SysAppEntrySetingFilter{
		DirId: iotutil.ToInt64(dirId),
	})
	if errSet != nil && errSet.Error() != "record not found" {
		return "", errSet
	}
	if resSet.Code != 200 && resSet.Message != "record not found" {
		return "", errors.New(resSet.Message)
	}
	if resSet.Data != nil && len(resSet.Data) > 0 {
		return "", errors.New("目录下有关联词条,不可删除")
	}

	res, err := rpc.ClientSysAppDocDirService.DeleteById(s.Ctx, &protosService.SysAppDocDir{
		Id: iotutil.ToInt64(dirId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

func (s SysAppDocDirService) ListDir(lang, helpId string) ([]*entitys.SysAppDocDirListRes, error) {
	req := &protosService.SysAppDocDirListRequest{
		Page:     1,
		PageSize: 99999999,
		Query: &protosService.SysAppDocDir{
			//DocId: iotutil.ToInt64(docId),
			HelpId: iotutil.ToInt64(helpId),
		},
	}
	// if len(helpId) != 0 {
	// 	req.Query.HelpId = iotutil.ToInt64(helpId)
	// }
	res, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, req)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	// for i := range res.Data {
	// 	reqDir := &protosService.SysAppDocDirListRequest{
	// 		Page:     1,
	// 		PageSize: 99999999,
	// 		Query: &protosService.SysAppDocDir{
	// 			ParentId: res.Data[i].Id,
	// 		},
	// 	}
	// 	resp, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, reqDir)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if res.Code != 200 {
	// 		return nil, errors.New(res.Message)
	// 	}
	// 	res.Data = append(res.Data, resp.Data...)
	// }

	var rs = make([]*entitys.SysAppDocDirListRes, 0)

	if res.Data != nil {
		for _, v := range res.Data {
			//dirName 需要根据lang进行翻译之后赋值
			dirName := v.DirName
			resTran, errTran := s.TranslateGet(v.Id, lang)
			if errTran != nil {
				logger.Error(errTran.Error())
			}
			if len(resTran) > 0 {
				dirName = resTran[0].DirName
			}
			rs = append(rs, &entitys.SysAppDocDirListRes{
				Id:         iotutil.ToString(v.Id),
				ParentId:   iotutil.ToString(v.ParentId),
				DirImg:     v.DirImg,
				DocDirName: dirName,
				Sort:       v.Sort,
				//DocId: docId,
			})

		}
	}
	return rs, nil

}

// 获取中文目录名称
func (s SysAppDocDirService) GetDefaultLangDirLang(req []entitys.SysAppDocDirLang) entitys.SysAppDocDirLang {
	rs := entitys.SysAppDocDirLang{}
	for _, v := range req {
		if v.Lang == "zh" {
			rs = v
			break
		}
	}
	return rs
}

func (s SysAppDocDirService) ConvertLangTranslateEntity(req entitys.SysAppDocDirSaveReq) (*langEnt.LangTranslateEntitys, error) {
	var rs = langEnt.LangTranslateEntitys{}
	rs.SourceTable = "t_sys_app_doc_dir"
	rs.SourceRowId = req.Id
	// Id            string                   `json:"id"`
	// SourceTable   string                   `json:"sourceTable,omitempty"`
	// SourceRowId   string                   `json:"sourceRowId,omitempty"`
	// Lang          string                   `json:"lang"`
	// FieldName     string                   `json:"fieldName"`
	// FieldType     int32                    `json:"fieldType"`
	// FieldValue    string                   `json:"fieldValue"`
	// TranslateList []BatchSaveTranslateItem `json:"translateList,omitempty"`
	// type BatchSaveTranslateItem struct {
	// 	Id         string `json:"id,omitempty"`
	// 	Lang       string `json:"lang,omitempty"`
	// 	FieldName  string `json:"fieldName,omitempty"`
	// 	FieldType  int32  `json:"fieldType,omitempty"`
	// 	FieldValue string `json:"fieldValue,omitempty"`
	// }
	for _, v := range req.DocDirLangs {
		rs.TranslateList = append(rs.TranslateList, langEnt.BatchSaveTranslateItem{
			Lang:       v.Lang,
			FieldName:  "dir_name",
			FieldType:  0,
			FieldValue: v.DirName,
		})
	}
	return &rs, nil
}

// TranslateSave 翻译保存（业务翻译保存）
func (s SysAppDocDirService) TranslateSave(req langEnt.LangTranslateEntitys) (string, error) {
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
	ret, err := rpc.ClientLangTranslateService.BatchCreate(context.Background(), &protosService.BatchSaveTranslate{
		SourceRowId: req.SourceRowId,
		SourceTable: req.SourceTable,
		List:        translateList,
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
func (s SysAppDocDirService) TranslateGet(dirId int64, lang string) ([]*entitys.SysAppDocDirLang, error) {

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
