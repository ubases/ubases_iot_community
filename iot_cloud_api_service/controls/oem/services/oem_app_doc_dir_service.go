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

	"go-micro.dev/v4/logger"
)

type OemAppDocDirService struct {
	Ctx context.Context
}

func (s OemAppDocDirService) SetContext(ctx context.Context) OemAppDocDirService {
	s.Ctx = ctx
	return s
}

func (s OemAppDocDirService) CreateDir(req entitys.OemAppDocDirSaveReq) (string, error) {
	var dirId int64
	if req.Id != "" && req.Id != "0" {
		dirId = iotutil.ToInt64(req.Id)
	} else {
		dirId = iotutil.GetNextSeqInt64()
	}
	reqDir := &protosService.OemAppDocDir{
		Id:       dirId,
		DocId:    iotutil.ToInt64(req.DocId),
		DirName:  s.GetDefaultLangDirLang(req.DocDirLangs).DirName,
		DirImg:   req.DirImg,
		ParentId: iotutil.ToInt64(req.ParentId),
		Sort:     req.Sort,
	}
	res, err := rpc.ClientOemAppDocDirService.Create(s.Ctx, reqDir)
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

func (s OemAppDocDirService) UpdateDir(req entitys.OemAppDocDirSaveReq) (string, error) {
	res, err := rpc.ClientOemAppDocDirService.UpdateFields(s.Ctx, &protosService.OemAppDocDirUpdateFieldsRequest{
		Fields: []string{"dir_name", "dir_img", "sort"},
		Data: &protosService.OemAppDocDir{
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

func (s OemAppDocDirService) DetailDir(dirId string) (*entitys.OemAppDocDirSaveReq, error) {
	res, err := rpc.ClientOemAppDocDirService.FindById(s.Ctx, &protosService.OemAppDocDirFilter{
		Id: iotutil.ToInt64(dirId),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	var rs = entitys.OemAppDocDirSaveReq{}
	rs.DocDirLangs = make([]*entitys.OemAppDocDirLang, 0)
	if res.Data != nil && len(res.Data) > 0 {
		tmp := res.Data[0]
		rs.DirImg = tmp.DirImg
		rs.DocId = iotutil.ToString(tmp.DocId)
		rs.ParentId = iotutil.ToString(tmp.ParentId)
		rs.Sort = tmp.Sort
	}

	//获取支持的语种
	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	//去翻译表找语种翻译数据
	resTran, errTran := s.TranslateGet(iotconst.LANG_T_OEM_APP_DOC_DIR, iotutil.ToInt64(dirId), "")
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

func (s OemAppDocDirService) DeleteDir(dirId string) (string, error) {

	resDir, errDir := rpc.ClientOemAppDocDirService.Find(s.Ctx, &protosService.OemAppDocDirFilter{
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

	resSet, errSet := rpc.ClientOemAppEntrySetingService.Find(s.Ctx, &protosService.OemAppEntrySetingFilter{
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

	res, err := rpc.ClientOemAppDocDirService.DeleteById(s.Ctx, &protosService.OemAppDocDir{
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

func (s OemAppDocDirService) ListDir(docId string, lang string) ([]*entitys.OemAppDocDirListRes, error) {

	res, err := rpc.ClientOemAppDocDirService.Lists(s.Ctx, &protosService.OemAppDocDirListRequest{
		Page:     1,
		PageSize: 99999999,
		Query: &protosService.OemAppDocDir{
			DocId: iotutil.ToInt64(docId),
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	// for i := range res.Data {
	// 	reqDir := &protosService.OemAppDocDirListRequest{
	// 		Page:     1,
	// 		PageSize: 99999999,
	// 		Query: &protosService.OemAppDocDir{
	// 			ParentId: res.Data[i].Id,
	// 		},
	// 	}
	// 	resp, err := rpc.ClientOemAppDocDirService.Lists(s.Ctx, reqDir)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if res.Code != 200 {
	// 		return nil, errors.New(res.Message)
	// 	}
	// 	res.Data = append(res.Data, resp.Data...)
	// }

	var rs = make([]*entitys.OemAppDocDirListRes, 0)

	if res.Data != nil {
		for _, v := range res.Data {
			//dirName 需要根据lang进行翻译之后赋值
			dirName := v.DirName
			resTran, errTran := s.TranslateGet(iotconst.LANG_T_OEM_APP_DOC_DIR, v.Id, lang)
			if errTran != nil {
				logger.Error(errTran.Error())
			}
			if len(resTran) > 0 {
				dirName = resTran[0].DirName
			}
			rs = append(rs, &entitys.OemAppDocDirListRes{
				Id:         iotutil.ToString(v.Id),
				ParentId:   iotutil.ToString(v.ParentId),
				DirImg:     v.DirImg,
				DocDirName: dirName,
				DocId:      docId,
				Sort:       v.Sort,
			})

		}
	}
	return rs, nil

}

// 获取中文目录名称
func (s OemAppDocDirService) GetDefaultLangDirLang(req []*entitys.OemAppDocDirLang) entitys.OemAppDocDirLang {
	rs := entitys.OemAppDocDirLang{}
	for _, v := range req {
		if v.Lang == "zh" {
			rs = *v
			break
		}
	}
	return rs
}

func (s OemAppDocDirService) ConvertLangTranslateEntity(req entitys.OemAppDocDirSaveReq) (*langEntitys.LangTranslateEntitys, error) {
	var rs = langEntitys.LangTranslateEntitys{}
	rs.SourceTable = iotconst.LANG_T_OEM_APP_DOC_DIR
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
		rs.TranslateList = append(rs.TranslateList, langEntitys.BatchSaveTranslateItem{
			Lang:       v.Lang,
			FieldName:  "dir_name",
			FieldType:  0,
			FieldValue: v.DirName,
		})
	}
	return &rs, nil
}

// TranslateSave 翻译保存（业务翻译保存）
func (s OemAppDocDirService) TranslateSave(req langEntitys.LangTranslateEntitys) (string, error) {
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
func (s OemAppDocDirService) TranslateGet(tabName string, dirId int64, lang string) ([]*entitys.OemAppDocDirLang, error) {
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
