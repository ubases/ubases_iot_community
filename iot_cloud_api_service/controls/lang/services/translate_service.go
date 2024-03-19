package services

import (
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type LangTranslateService struct {
	Ctx context.Context
}

func (s LangTranslateService) SetContext(ctx context.Context) LangTranslateService {
	s.Ctx = ctx
	return s
}

// 翻译详细
func (s LangTranslateService) GetLangTranslateDetail(id string) (*entitys.LangTranslateEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientLangTranslateService.FindById(s.Ctx, &protosService.LangTranslateFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]
	return entitys.LangTranslate_pb2e(data), err
}

// QueryLangTranslateList 翻译列表
func (s LangTranslateService) QueryLangTranslateList(filter entitys.LangTranslateQuery) ([]*entitys.LangTranslateGroupEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientLangTranslateService.GroupLists(s.Ctx, &protosService.LangTranslateListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.LangTranslate{
			SourceTable:  filter.SourceTable,
			SourceRowId:  filter.SourceRowId,
			PlatformType: filter.PlatformType,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.LangTranslateGroupEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, &entitys.LangTranslateGroupEntitys{
			LangKey:     item.SourceRowId,
			SourceTable: item.SourceTable,
			SourceRowId: item.SourceRowId,
			FieldName:   item.FieldName,
			Name:        item.Zh,
			NameEn:      item.En,
		})
	}
	return resultList, rep.Total, err
}

// AddLangTranslate 新增翻译
func (s LangTranslateService) AddLangTranslate(req entitys.LangTranslateEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.LangTranslate_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientLangTranslateService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改翻译
func (s LangTranslateService) UpdateLangTranslate(req entitys.LangTranslateEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientLangTranslateService.Update(s.Ctx, entitys.LangTranslate_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除翻译
func (s LangTranslateService) DeleteLangTranslate(req entitys.LangTranslateFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientLangTranslateService.Delete(s.Ctx, &protosService.LangTranslate{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// QueryLangList 获取翻译列表（多个sourceRowId）
func (s LangTranslateService) QueryLangList(sourceTable string, sourceRowIdList []string) (map[string][]*entitys.LangTranslateEntitys, error) {
	rep, err := rpc.ClientLangTranslateService.Lists(s.Ctx, &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable:     sourceTable,
			SourceRowIdList: sourceRowIdList,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	resList := map[string][]*entitys.LangTranslateEntitys{}
	for _, item := range rep.Data {
		if _, ok := resList[item.SourceRowId]; !ok {
			resList[item.SourceRowId] = []*entitys.LangTranslateEntitys{}
		}
		resList[item.SourceRowId] = append(resList[item.SourceRowId], &entitys.LangTranslateEntitys{
			Lang:       item.Lang,
			FieldName:  item.FieldName,
			FieldType:  item.FieldType,
			FieldValue: item.FieldValue,
		})
	}
	return resList, err
}

// BatchInsert 批量插入翻译
func (s LangTranslateService) BatchInsert(req entitys.LangTranslateEntitys) error {
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
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return nil
}
