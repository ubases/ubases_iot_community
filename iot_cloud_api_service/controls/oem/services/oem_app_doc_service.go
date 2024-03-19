package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	entitys2 "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppDocService struct {
	Ctx context.Context
}

func (s OemAppDocService) SetContext(ctx context.Context) OemAppDocService {
	s.Ctx = ctx
	return s
}

// 获取空间下的所有app[新增和修改的时候使用]
func (s OemAppDocService) GetApps(tenantId string, docId string) ([]*entitys.OemAppDocApp, error) {
	res, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.OemApp{
			TenantId: tenantId,
		},
	})

	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	//获取到空间下所有已经写过文档的app记录
	resRela, errReal := s.GetDocAppRelationAll(tenantId)
	if errReal != nil {
		return nil, errReal
	}
	var rs = make([]*entitys.OemAppDocApp, 0)

	//编辑的时候. 需要把自己的app数据加入.
	appKeyMaps, _ := GetAppMaps(s.Ctx, tenantId)
	if docId != "" {
		detail, err := s.DetailDoc(docId)
		if err != nil {
			return nil, err
		}
		if detail != nil && len(detail.Apps) > 0 {
			for _, d := range detail.Apps {
				tmp := entitys.OemAppDocApp{
					AppId:   d.AppId,
					AppName: appKeyMaps[d.AppKey],
					AppKey:  d.AppKey,
				}
				rs = append(rs, &tmp)
			}
		}
	}

	if res.Data != nil && len(res.Data) > 0 {
		for _, v := range res.Data {
			//默认该app没有关联过文档.
			isExists := false
			//循环判断, 如果app关联过文档则 不能在提供前端选择
			for _, r := range resRela {
				if r.AppId == iotutil.ToString(v.Id) {
					isExists = true
					break
				}
			}
			//没有关联过文档的app才显示到前端
			if !isExists {
				rs = append(rs, &entitys.OemAppDocApp{
					AppId:   iotutil.ToString(v.Id),
					AppName: appKeyMaps[v.AppKey],
					AppKey:  v.AppKey,
				})
			}
		}
	}
	return rs, nil
}

// 获取空间下的所有app
func (s OemAppDocService) GetTenantApps(tenantId string) ([]*entitys.OemAppDocApp, error) {
	res, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Query: &protosService.OemApp{
			TenantId: tenantId,
		},
	})

	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	var rs = make([]*entitys.OemAppDocApp, 0)
	for _, v := range res.Data {

		rs = append(rs, &entitys.OemAppDocApp{
			AppId:              iotutil.ToString(v.Id),
			AppName:            v.Name,
			AppKey:             v.AppKey,
			AppTemplateId:      v.AppTemplateId,
			AppTemplateType:    v.AppType,
			AppTemplateVersion: v.AppTemplateVersion,
		})
	}
	return rs, nil
}

//获取公版语种[暂时先从字典获取.]
// func (s OemAppDocService) GetPubLangs() ([]*entitys.OemAppDocLang,error){
// 	//language_type
// 	langs := GetBaseDataValue("language_type",s.Ctx)
// 	var rs  = make([]*entitys.OemAppDocLang,0)
// 	for k,v := range langs{
// 		rs = append(rs, &entitys.OemAppDocLang{
// 			Lang: iotutil.ToString(v),
// 			LangName:k  ,
// 		})
// 	}
// 	//排序
// 	s.SortOemAppDocLang(rs)
// 	return rs, nil
// }

// 获取公版语种[暂时先从字典获取.]
func (s OemAppDocService) GetPubLangs() ([]*entitys.OemAppDocLang, error) {
	res, err := rpc.ClientSysAppEntryService.EntryLangsDiy(s.Ctx, &protosService.SysAppEntryFilter{})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	//language_type
	langs := GetBaseDataValue("language_type", s.Ctx)
	var rs = make([]*entitys.OemAppDocLang, 0)
	for k, v := range langs {
		rs = append(rs, &entitys.OemAppDocLang{
			Lang:     iotutil.ToString(v),
			LangName: k,
		})
	}

	var list = make([]*entitys.OemAppDocLang, 0)
	//获取到语种名称
	for _, v := range res.Data {
		for _, vv := range rs {
			if v == vv.Lang {
				list = append(list, &entitys.OemAppDocLang{
					Lang:     iotutil.ToString(vv.Lang),
					LangName: vv.LangName,
				})
				break
			}
		}
	}

	//排序
	s.SortOemAppDocLang(list)
	return list, nil
}

// 获取文档支持的语种
func (s OemAppDocService) GetDocSupportLangs(docId string) ([]*entitys.OemAppDocLang, error) {
	res, err := rpc.ClientOemAppDocService.FindById(s.Ctx, &protosService.OemAppDocFilter{
		Id: iotutil.ToInt64(docId),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var rs = make([]*entitys.OemAppDocLang, 0)
	if res.Data != nil && len(res.Data) > 0 && res.Data[0].Langs != "" {
		iotutil.JsonToStruct(res.Data[0].Langs, &rs)
	}

	s.SortOemAppDocLang(rs)

	return rs, nil
}

// 文档详细
func (s OemAppDocService) DetailDoc(docId string) (*entitys.OemAppDocSaveReq, error) {
	if docId == "" {
		return nil, errors.New("参数错误")
	}

	res, err := rpc.ClientOemAppDocService.FindById(s.Ctx, &protosService.OemAppDocFilter{
		Id: iotutil.ToInt64(docId),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	var rs = entitys.OemAppDocSaveReq{}
	if res.Data != nil && len(res.Data) > 0 {
		rs = *entitys.OemAppDoc_pb2e(res.Data[0])
	}
	return &rs, nil
}

func (s OemAppDocService) CreateDoc(req entitys.OemAppDocSaveReq) (string, error) {
	newHelpId := iotutil.GetNextSeqInt64()
	//判断是否继承
	if req.IsSucceedPUbDoc == 1 && len(req.SucceedPubDoc) > 0 {
		// diyCtx := s.CopyNewContext(s.Ctx)
		// go func() {
		// 	err := s.SucceedPubDoc(diyCtx, req.SucceedPubDoc, req.Id, iotutil.ToString(docId))
		// 	if err != nil {
		// 		logger.Error(err)
		// 	}
		// }()

		// 先通过helpId查询相应的目录列表, 然后级联复制目录及词条
		respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				HelpId:   iotutil.ToInt64(req.Id),
				ParentId: -1,
			},
		})
		if err != nil {
			return "", err
		}
		if len(respDir.Data) != 0 {
			if err := CopyAppEntryByRecurse(respDir.Data, newHelpId, 0); err != nil {
				return "", err
			}
		}
	}
	err := s.SaveDocAppRelation(newHelpId, false, req.Apps)
	if err != nil {
		return "", err
	}
	res, err := rpc.ClientOemAppDocService.Create(s.Ctx, &protosService.OemAppDoc{
		Id:              newHelpId,
		Name:            req.Name,
		Apps:            iotutil.ToString(req.Apps),
		Langs:           iotutil.ToString(req.Langs),
		RemainLang:      req.RemainLang,
		IsSucceedPubDoc: int32(req.IsSucceedPUbDoc),
		SucceedPubDoc:   iotutil.ToString(req.SucceedPubDoc),
		HelpCenterName:  req.HelpCenterName,
		Version:         req.Version,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

// 判断语种是否包含在集合中.
func (s OemAppDocService) ExistsLangs(lang string, langs []entitys.OemAppDocLang) bool {
	var b bool
	b = false
	for _, v := range langs {
		if v.Lang == lang {
			b = true
			break
		}
	}
	return b
}

// 继承公版文档执行[后续考虑批量入库优化]
func (s OemAppDocService) SucceedPubDoc(ctxDiy context.Context, langs []entitys.OemAppDocLang, helpId, docId string) error {
	resDir, errDir := rpc.ClientSysAppDocDirService.Lists(ctxDiy, &protosService.SysAppDocDirListRequest{
		Page:     1,
		PageSize: 99999999999,
		Query: &protosService.SysAppDocDir{
			HelpId: iotutil.ToInt64(helpId),
		},
	})
	if errDir != nil && errDir.Error() != "record not found" {
		return nil
	}
	//公版目录
	logger.Info(resDir)

	resSeting, errSeting := rpc.ClientSysAppEntrySetingService.Lists(ctxDiy, &protosService.SysAppEntrySetingListRequest{
		Page:     1,
		PageSize: 99999999999,
		Query: &protosService.SysAppEntrySeting{
			IsEnable: 1,
		},
	})
	if errSeting != nil && errSeting.Error() != "record not found" {
		return nil
	}
	//公版词条设置
	logger.Info(resSeting)

	resEntry, errEntry := rpc.ClientSysAppEntryService.Lists(ctxDiy, &protosService.SysAppEntryListRequest{
		Page:     1,
		PageSize: 99999999999,
		Query:    &protosService.SysAppEntry{},
	})

	if errEntry != nil && errEntry.Error() != "record not found" {
		return nil
	}
	//公版词条
	logger.Info(resEntry)
	listPubDir := resDir.Data
	listPubSeting := resSeting.Data
	listPubEntry := resEntry.Data

	//ids[oldid] = newid 直接一次性把id生成放到集合.
	var ids = make(map[string]string, 0)
	//生成id替换集合
	for _, v := range listPubDir {
		newId := iotutil.ToString(iotutil.GetNextSeqInt64())
		ids[iotutil.ToString(v.Id)] = newId
	}
	for _, v := range listPubSeting {
		newId := iotutil.ToString(iotutil.GetNextSeqInt64())
		ids[iotutil.ToString(v.Id)] = newId
	}
	for _, v := range listPubEntry {
		newId := iotutil.ToString(iotutil.GetNextSeqInt64())
		ids[iotutil.ToString(v.Id)] = newId
	}

	dirSer := OemAppDocDirService{}
	dirSer.Ctx = ctxDiy

	entrySer := OemAppEntryService{}
	entrySer.Ctx = ctxDiy

	for _, dir := range listPubDir {
		//获取原来目录的翻译
		langsTran, errLangs := dirSer.TranslateGet("t_sys_app_doc_dir", dir.Id, "")
		if errLangs != nil {
			logger.Error(errLangs)
		}
		//过滤一下翻译
		var filterLangs = make([]*entitys.OemAppDocDirLang, 0)
		for _, v := range langsTran {
			if s.ExistsLangs(v.Lang, langs) {
				filterLangs = append(filterLangs, v)
			}
		}

		pid := iotutil.ToString(dir.ParentId)
		if pid != "0" {
			pid = ids[pid]
		}

		dirSer.CreateDir(entitys.OemAppDocDirSaveReq{
			DocId:       docId,
			DirImg:      dir.DirImg,
			ParentId:    pid,
			DocDirLangs: filterLangs,
			Id:          ids[iotutil.ToString(dir.Id)],
		})
	}

	for _, seting := range listPubSeting {
		id := ids[iotutil.ToString(seting.Id)]
		if id != "" {
			entrySer.EntrySetingCreate(entitys.OemAppEntrySetingSaveReq{
				SetingId: ids[iotutil.ToString(seting.Id)],
				Sort:     int64(seting.Sort),
				IsEnable: int64(seting.IsEnable),
				IsNormal: int64(seting.IsNormal),
				DirId:    ids[iotutil.ToString(seting.DirId)],
			})
		}
	}

	for _, entry := range listPubEntry {
		//继承的语种才加入
		if s.ExistsLangs(entry.Lang, langs) {

			id := ids[iotutil.ToString(entry.Id)]
			setingId := ids[iotutil.ToString(entry.SetingId)]
			//防止公版数据为删除干净导致错误.
			if id != "" && setingId != "" {
				entrySer.EntryCreate(&protosService.OemAppEntry{
					Id:        iotutil.ToInt64(id),
					Lang:      entry.Lang,
					Title:     entry.Title,
					Content:   entry.Content,
					SetingId:  iotutil.ToInt64(setingId),
					UpdatedAt: timestamppb.New(time.Now()),
				})
			}
		}
	}

	return nil

}

// 复制一个新的上下文. 通常在携程中使用.
func (s *OemAppDocService) CopyNewContext(old context.Context) context.Context {

	userid, _ := metadata.Get(old, "userid")
	tenandId, _ := metadata.Get(old, "tenantid")
	token, _ := metadata.Get(old, "token")
	//新建一个ctx 在携程中使用.
	ctxDiy := metadata.NewContext(context.Background(), map[string]string{
		"userid":   userid,
		"tenantid": tenandId,
		"token":    token,
	})
	return ctxDiy
}

// 更新文档
func (s OemAppDocService) UpdateDoc(req entitys.OemAppDocSaveReq) (string, error) {
	//先保存关系.
	errRela := s.SaveDocAppRelation(iotutil.ToInt64(req.Id), true, req.Apps)
	if errRela != nil {
		return "", errRela
	}
	//修改字段
	res, err := rpc.ClientOemAppDocService.UpdateFields(s.Ctx, &protosService.OemAppDocUpdateFieldsRequest{
		Fields: []string{"name", "apps", "langs", "remain_lang", "is_succeed_pub_doc", "succeed_pub_doc", "updated_by", "updated_at", "help_id", "version"},
		Data: &protosService.OemAppDoc{
			Id:              iotutil.ToInt64(req.Id),
			Name:            req.Name,
			Apps:            iotutil.ToString(req.Apps),
			Langs:           iotutil.ToString(req.Langs),
			RemainLang:      req.RemainLang,
			IsSucceedPubDoc: int32(req.IsSucceedPUbDoc),
			SucceedPubDoc:   iotutil.ToString(req.SucceedPubDoc),
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

// 更新文档
func (s OemAppDocService) DeleteDoc(req entitys.OemAppDocSaveReq) error {
	// 先通过helpId查询相应的目录列表, 然后级联删除目录及词条
	respDir, err := rpc.ClientOemAppDocDirService.Lists(s.Ctx, &protosService.OemAppDocDirListRequest{
		Query: &protosService.OemAppDocDir{
			DocId:    iotutil.ToInt64(req.Id),
			ParentId: -1,
		},
	})
	if err != nil {
		return err
	}
	if len(respDir.Data) != 0 {
		if err := DelAppEntryByRecurse(respDir.Data); err != nil {
			return err
		}
	}
	_, err = rpc.ClientOemAppDocRelationService.Delete(s.Ctx, &protosService.OemAppDocRelation{
		DocId: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	_, err = rpc.ClientOemAppDocService.DeleteById(s.Ctx, &protosService.OemAppDoc{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	return nil
}

// 根据appid , 查出改appid关联在哪个文档下
func (s OemAppDocService) AppByDoc(appId int64) ([]*protosService.OemAppDocRelation, error) {
	res, err := rpc.ClientOemAppDocRelationService.Lists(s.Ctx, &protosService.OemAppDocRelationListRequest{
		Page:     1,
		PageSize: 99999999,
		Query: &protosService.OemAppDocRelation{
			AppId: appId,
		},
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res.Data, nil
}

// 获取空间下所有已经设置过文档的应用app
func (s OemAppDocService) GetDocAppRelationAll(tenantId string) ([]*entitys.OemAppDocApp, error) {
	res, err := rpc.ClientOemAppDocService.Lists(s.Ctx, &protosService.OemAppDocListRequest{
		Query: &protosService.OemAppDoc{
			TenantId: tenantId,
		},
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	var rs = make([]*entitys.OemAppDocApp, 0)
	if res.Data != nil && len(res.Data) > 0 {
		for _, v := range res.Data {
			apps := make([]*entitys.OemAppDocApp, 0)
			iotutil.JsonToStruct(v.Apps, &apps)
			rs = append(rs, apps...)
		}
	}
	return rs, nil
}

// 保存文档与app关系表
func (s OemAppDocService) SaveDocAppRelation(docId int64, isUpdate bool, apps []entitys.OemAppDocApp) error {
	var rs error
	//
	if isUpdate {
		resDel, errDel := rpc.ClientOemAppDocRelationService.Delete(s.Ctx, &protosService.OemAppDocRelation{
			DocId: docId,
		})
		if errDel != nil {
			return errDel
		}
		if resDel.Code != 200 {
			return errors.New(resDel.Message)
		}
	}

	for _, v := range apps {
		rid := iotutil.GetNextSeqInt64()
		_, errFor := rpc.ClientOemAppDocRelationService.Create(s.Ctx, &protosService.OemAppDocRelation{
			Id:     rid,
			DocId:  docId,
			AppId:  iotutil.ToInt64(v.AppId),
			AppKey: v.AppKey,
		})
		if errFor != nil {
			rs = errFor
			break
		}
	}
	return rs
}

// 获取公版数据.
func (s OemAppDocService) GetDocPublic(tenantId string) ([]*entitys.OmeAppDocListRes, error) {
	data := []*entitys.OmeAppDocListRes{}
	// 从字段中获取app模板类型列表
	respDict, err := rpc.TConfigDictDataServerService.Lists(s.Ctx, &protosService.ConfigDictDataListRequest{
		Query: &protosService.ConfigDictData{
			DictType: "app_template_type",
		},
	})
	if err != nil {
		return nil, err
	}
	for i := range respDict.Data {
		resp, err := rpc.ClientSysAppHelpCenterService.Lists(s.Ctx, &protosService.SysAppHelpCenterListRequest{
			Page:      1,
			PageSize:  1,
			OrderKey:  "version",
			OrderDesc: "desc",
			Query: &protosService.SysAppHelpCenter{
				TemplateType: iotutil.ToInt32(respDict.Data[i].DictValue),
			},
		})
		if err != nil {
			return nil, err
		}

		langs, err := s.GetPubLangs()
		if err != nil {
			logger.Error(err.Error())
		}

		for i := range resp.Data {
			count, err := s.GetPubDocByAnswerCount(resp.Data[i].Id)
			if err != nil {
				return nil, err
			}
			pubDoc := &entitys.OmeAppDocListRes{
				Id:          iotutil.ToString(resp.Data[i].Id),
				Name:        resp.Data[i].Name,
				Version:     resp.Data[i].Version,
				AnswerCount: count,
				IsPub:       1,
			}
			//获取空间下所有app
			apps, _ := s.GetTenantApps(tenantId)

			//获取空间下所有文档关联的所有app
			docApps, _ := s.GetDocAppRelationAll(tenantId)

			for _, v := range apps {
				isExists := false
				//判断app是否已经有关联文档. 如果关联了文档. 则不添加到公版app
				for _, vv := range docApps {
					if v.AppKey == vv.AppKey {
						isExists = true
						break
					}
				}
				if !isExists {
					pubDoc.Apps = append(pubDoc.Apps, *v)
				}
			}
			for _, v := range langs {
				pubDoc.Langs = append(pubDoc.Langs, entitys.OemAppDocLang{
					Lang:     v.Lang,
					LangName: v.LangName,
				})
			}

			data = append(data, pubDoc)
		}
	}

	return data, nil
}

// 获取公版数据.
func (s OemAppDocService) GetDocPublicExt(tenantId string) ([]*entitys.OmeAppDocListRes, error) {
	data := make([]*entitys.OmeAppDocListRes, 0)
	//获取帮助中心数据
	helpRes, err := rpc.ClientSysAppHelpCenterService.Lists(s.Ctx, &protosService.SysAppHelpCenterListRequest{
		OrderKey:  "version",
		OrderDesc: "desc",
		Query:     &protosService.SysAppHelpCenter{
			//TemplateType: iotutil.ToInt32(respDict.Data[i].DictValue),
		},
	})
	if err != nil {
		return nil, err
	}
	helpMap := make(map[string][]*protosService.SysAppHelpCenter)
	for _, d := range helpRes.Data {
		k := fmt.Sprintf("%v_%v", d.TemplateType, d.Version)
		if _, ok := helpMap[k]; !ok {
			helpMap[k] = make([]*protosService.SysAppHelpCenter, 0)
		}
		helpMap[k] = append(helpMap[k], d)
	}
	//获取空间下所有app
	apps, _ := s.GetTenantApps(tenantId)

	//获取空间下所有文档关联的所有app
	docApps, _ := s.GetDocAppRelationAll(tenantId)

	for _, heplList := range helpMap {
		for _, item := range heplList {
			count, err := s.GetPubDocByAnswerCount(item.Id)
			if err != nil {
				return nil, err
			}
			pubDoc := &entitys.OmeAppDocListRes{
				Id:          iotutil.ToString(item.Id),
				Name:        item.Name,
				Version:     item.Version,
				AnswerCount: count,
				IsPub:       1,
			}

			for _, v := range apps {
				//判断APP是否与模板一直，不一致的app不需要计入到类型语言中
				if v.AppTemplateType == item.TemplateType && v.AppTemplateVersion == item.Version {
					isExists := false
					//判断app是否已经有关联文档. 如果关联了文档. 则不添加到公版app
					for _, vv := range docApps {
						if v.AppKey == vv.AppKey {
							isExists = true
							break
						}
					}
					if !isExists {
						pubDoc.Apps = append(pubDoc.Apps, *v)
					}
				}
			}

			//没有匹配到APP的公版语言包，将不显示到列表中
			if len(pubDoc.Apps) == 0 {
				continue
			}
			//获取公版支持的语言类型
			langs, err := s.GetPubLangs()
			if err != nil {
				logger.Error(err.Error())
			}
			for _, v := range langs {
				pubDoc.Langs = append(pubDoc.Langs, entitys.OemAppDocLang{
					Lang:     v.Lang,
					LangName: v.LangName,
				})
			}

			data = append(data, pubDoc)
		}
	}

	return data, nil
}

// 获取公版数据.
func (s OemAppDocService) GetHelpCenterListForOpen(tenantId string, req entitys2.SysAppHelpCenterQuery) ([]*entitys2.SysAppHelpCenterEntitys, error) {
	data := make([]*entitys2.SysAppHelpCenterEntitys, 0)
	//获取帮助中心数据
	helpRes, err := rpc.ClientSysAppHelpCenterService.Lists(s.Ctx, &protosService.SysAppHelpCenterListRequest{
		OrderKey:  "version",
		OrderDesc: "desc",
		Query:     &protosService.SysAppHelpCenter{},
	})
	if err != nil {
		return nil, err
	}
	helpMap := make(map[string][]*protosService.SysAppHelpCenter)
	for _, d := range helpRes.Data {
		k := fmt.Sprintf("%v_%v", d.TemplateType, d.Version)
		if _, ok := helpMap[k]; !ok {
			helpMap[k] = make([]*protosService.SysAppHelpCenter, 0)
		}
		helpMap[k] = append(helpMap[k], d)
	}
	langs, err := s.GetPubLangs()
	if err != nil {
		logger.Error(err.Error())
	}
	//获取空间下所有app
	apps, _ := s.GetTenantApps(tenantId)
	appMaps := make(map[string]int32)
	for _, app := range apps {
		appMaps[fmt.Sprintf("%v_%v", app.AppTemplateType, app.AppTemplateVersion)] = 1
	}
	for _, heplList := range helpMap {
		for _, item := range heplList {
			pubDoc := entitys2.SysAppHelpCenter_pb2e(item)
			if _, ok := appMaps[fmt.Sprintf("%v_%v", item.TemplateType, item.Version)]; ok {
				data = append(data, pubDoc)
			}
			for _, v := range langs {
				pubDoc.Langs = append(pubDoc.Langs, entitys2.TempOemAppDocLang{
					Lang:     v.Lang,
					LangName: v.LangName,
				})
			}
		}
	}
	return data, nil
}

func (s OemAppDocService) DocList(tenantId string) ([]*entitys.OmeAppDocListRes, error) {
	res, err := rpc.ClientOemAppDocService.Lists(s.Ctx, &protosService.OemAppDocListRequest{
		Query: &protosService.OemAppDoc{
			TenantId: tenantId,
		},
	})

	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	rs := make([]*entitys.OmeAppDocListRes, 0)

	//pubDocs, err := s.GetDocPublic(tenantId)
	pubDocs, err := s.GetDocPublicExt(tenantId)
	if err != nil {
		return nil, err
	}
	rs = append(rs, pubDocs...)

	if res.Data != nil && len(res.Data) > 0 {
		for _, v := range res.Data {
			t := entitys.OemAppDoc_pb2eList(v)
			count, err := s.GetDocByAnswerCount(v.Id)
			if err != nil {
				return nil, err
			}
			t.AnswerCount = count
			t.IsPub = 2
			rs = append(rs, t)
		}
	}
	return rs, nil
}

// 后续需要优化. 数据量上来会拖慢查询.
func (s OemAppDocService) GetDocByAnswerCount(docId int64) (int64, error) {
	// 查询每个帮助中心下的词条数
	var count int64
	respDir, err := rpc.ClientOemAppDocDirService.Lists(s.Ctx, &protosService.OemAppDocDirListRequest{
		Query: &protosService.OemAppDocDir{
			DocId:    iotutil.ToInt64(docId),
			ParentId: -1,
		},
	})
	if err != nil {
		return count, err
	}
	if len(respDir.Data) != 0 {
		c, err := GetAppEntryCountForOpen(respDir.Data)
		if err != nil {
			return count, err
		}
		count += int64(c)
	}
	return count, nil
}

// 后续需要优化. 数据量上来会拖慢查询.
func (s OemAppDocService) GetPubDocByAnswerCount(docId int64) (int64, error) {
	// 查询每个帮助中心下的词条数
	var count int64
	respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
		Query: &protosService.SysAppDocDir{
			HelpId:   docId,
			ParentId: -1,
		},
	})
	if err != nil {
		return count, err
	}
	if len(respDir.Data) != 0 {
		c, err := GetAppEntryCount(respDir.Data)
		if err != nil {
			return count, err
		}
		count += int64(c)
	}
	return count, nil
}

// 语种进行升序
func (s *OemAppDocService) SortOemAppDocLang(list []*entitys.OemAppDocLang) {
	sort.Slice(list, func(i, j int) bool { // asc
		return list[i].Lang > list[j].Lang
	})
}

// 递归获取app帮助中心下文档数
func GetAppEntryCount(data []*protosService.SysAppDocDir) (int, error) {
	var count int
	ctx := context.Background()
	for i := range data {
		respEntrySeting, err := rpc.ClientSysAppEntrySetingService.Lists(ctx, &protosService.SysAppEntrySetingListRequest{
			Query: &protosService.SysAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		count += len(respEntrySeting.Data)
		respDir, err := rpc.ClientSysAppDocDirService.Lists(ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		if len(respDir.Data) != 0 {
			c, err := GetAppEntryCount(respDir.Data)
			if err != nil {
				return 0, err
			}
			count += c
		}
	}
	return count, nil
}

// 递归获取app帮助中心下文档数(开放平台)
func GetAppEntryCountForOpen(data []*protosService.OemAppDocDir) (int, error) {
	var count int
	ctx := context.Background()
	for i := range data {
		respEntrySeting, err := rpc.ClientOemAppEntrySetingService.Lists(ctx, &protosService.OemAppEntrySetingListRequest{
			Query: &protosService.OemAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		count += len(respEntrySeting.Data)
		respDir, err := rpc.ClientOemAppDocDirService.Lists(ctx, &protosService.OemAppDocDirListRequest{
			Query: &protosService.OemAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		if len(respDir.Data) != 0 {
			c, err := GetAppEntryCountForOpen(respDir.Data)
			if err != nil {
				return 0, err
			}
			count += c
		}
	}
	return count, nil
}

// 递归复制app帮助中心下目录和文档
func CopyAppEntryByRecurse(data []*protosService.SysAppDocDir, newHelpId, parentId int64) error {
	ctx := context.Background()
	dirs := []*protosService.OemAppDocDir{}
	for i := range data {
		// 先查询app目录设置id关联表
		respEntrySeting, err := rpc.ClientSysAppEntrySetingService.Lists(ctx, &protosService.SysAppEntrySetingListRequest{
			Query: &protosService.SysAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		// 遍历关联表列表，通过settingId来复制文档
		dirId := iotutil.GetNextSeqInt64()
		entrySets := []*protosService.OemAppEntrySeting{}
		for j := range respEntrySeting.Data {
			// 先通过setintId来获取多语言文档列表
			respEntry, err := rpc.ClientSysAppEntryService.Lists(ctx, &protosService.SysAppEntryListRequest{
				Query: &protosService.SysAppEntry{
					SetingId: respEntrySeting.Data[j].Id,
				},
			})
			if err != nil {
				return err
			}
			// 重置文档id和setingId, 并新建文档记录
			setingId := iotutil.GetNextSeqInt64()
			entrys := []*protosService.OemAppEntry{}
			for k := range respEntry.Data {
				reqEntry := &protosService.OemAppEntry{
					Id:        iotutil.GetNextSeqInt64(),
					Lang:      respEntry.Data[k].Lang,
					Title:     respEntry.Data[k].Title,
					Content:   respEntry.Data[k].Content,
					SetingId:  setingId,
					UpdatedAt: respEntry.Data[k].UpdatedAt,
				}
				entrys = append(entrys, reqEntry)
			}
			// 批量创建文档记录
			if len(entrys) != 0 {
				_, err := rpc.ClientOemAppEntryService.CreateBatch(ctx, &protosService.OemAppEntryBatchRequest{
					OemAppEntrys: entrys,
				})
				if err != nil {
					return err
				}
			}
			reqEntrySet := &protosService.OemAppEntrySeting{
				Id:       setingId,
				DirId:    dirId,
				Sort:     respEntrySeting.Data[j].Sort,
				IsEnable: respEntrySeting.Data[j].IsEnable,
				IsNormal: respEntrySeting.Data[j].IsNormal,
			}
			entrySets = append(entrySets, reqEntrySet)
		}
		// 批量创建app文档setingId关联记录
		if len(entrySets) != 0 {
			_, err = rpc.ClientOemAppEntrySetingService.CreateBatch(ctx, &protosService.OemAppEntrySetingBatchRequest{
				OemAppEntrySetings: entrySets,
			})
			if err != nil {
				return err
			}
		}

		// 考虑存在多层子目录，需级联查询并复制
		respDir, err := rpc.ClientSysAppDocDirService.Lists(ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		if len(respDir.Data) != 0 {
			err = CopyAppEntryByRecurse(respDir.Data, newHelpId, dirId)
			if err != nil {
				return err
			}
		}

		dirs = append(dirs, &protosService.OemAppDocDir{
			Id:       dirId,
			ParentId: parentId,
			DirName:  data[i].DirName,
			DirImg:   data[i].DirImg,
			DocId:    newHelpId,
		})
	}

	// 批量创建目录
	if len(dirs) != 0 {
		_, err := rpc.ClientOemAppDocDirService.CreateBatch(ctx, &protosService.OemAppDocDirBatchRequest{
			OemAppDocDirs: dirs,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// 递归删除app帮助中心下目录和文档
func DelAppEntryByRecurse(data []*protosService.OemAppDocDir) error {
	ctx := context.Background()
	for i := range data {
		// 先查询app目录设置id关联表
		respEntrySeting, err := rpc.ClientOemAppEntrySetingService.Lists(ctx, &protosService.OemAppEntrySetingListRequest{
			Query: &protosService.OemAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		// 遍历关联表列表，通过settingId来删除文档
		for j := range respEntrySeting.Data {
			// 先删除关联表的目录和settingId之间的关联关系
			_, err = rpc.ClientOemAppEntrySetingService.Delete(ctx, &protosService.OemAppEntrySeting{
				Id: respEntrySeting.Data[j].Id,
			})
			if err != nil {
				return err
			}
			// 然后通过settingId来删除文档
			_, err = rpc.ClientOemAppEntryService.Delete(ctx, &protosService.OemAppEntry{
				SetingId: respEntrySeting.Data[j].Id,
			})
			if err != nil {
				return err
			}
		}
		// 考虑存在多层子目录，需级联查询并删除
		respDir, err := rpc.ClientOemAppDocDirService.Lists(ctx, &protosService.OemAppDocDirListRequest{
			Query: &protosService.OemAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		if len(respDir.Data) != 0 {
			err = DelAppEntryByRecurse(respDir.Data)
			if err != nil {
				return err
			}
		}
		// 删除目录
		_, err = rpc.ClientOemAppDocDirService.Delete(ctx, &protosService.OemAppDocDir{
			Id: data[i].Id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
