package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BaseDataService struct {
}

func (s BaseDataService) GetBaseDataDetail(id string) (*protosService.ConfigDictDataResponse, error) {
	rid := iotutil.ToInt64(id)
	ret, err := rpc.TConfigDictDataServerService.FindById(context.Background(), &protosService.ConfigDictDataFilter{DictCode: rid})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, err
}

func (s BaseDataService) QueryBaseDataList(filter entitys.BaseDataQuery) ([]*protosService.ConfigDictData, error) {
	QueryObj := protosService.ConfigDictData{
		DictLabel: filter.DictLabel,
		DictType:  filter.DictType,
		DictValue: filter.DictValue,
	}
	ret, err := rpc.TConfigDictDataServerService.Lists(context.Background(), &protosService.ConfigDictDataListRequest{
		Query: &QueryObj,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret.Data, err
}

// TConfigDictDataRequest
func (s BaseDataService) AddBaseData(req entitys.BaseData) error {
	ret, err := rpc.TConfigDictDataServerService.Create(context.Background(), &protosService.ConfigDictData{
		DictCode:    iotutil.GetNextSeqInt64(),
		DictSort:    req.DictSort,
		DictLabel:   req.DictLabel,
		DictValue:   iotutil.ToString(req.DictValue),
		DictType:    req.DictType,
		CssClass:    req.CSSClass,
		ListClass:   req.ListClass,
		IsDefault:   req.IsDefault,
		Status:      req.Status,
		ValueType:   req.ValueType,
		Remark:      req.Remark,
		Pinyin:      req.Pinyin,
		Firstletter: req.Firstletter,
		Listimg:     req.Listimg,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(err.Error())
	}
	//清理缓存
	iotredis.GetClient().Del(context.Background(), fmt.Sprintf("dict_data_%v", req.DictType))
	return nil
}

func (s BaseDataService) UpdateBaseData(req entitys.BaseData) error {
	ret, err := rpc.TConfigDictDataServerService.Update(context.Background(), &protosService.ConfigDictData{
		DictCode:  iotutil.ToInt64(req.DictId),
		DictSort:  req.DictSort,
		DictLabel: req.DictLabel,
		DictValue: iotutil.ToString(req.DictValue),
		DictType:  req.DictType,
		CssClass:  req.CSSClass,
		ListClass: req.ListClass,
		IsDefault: req.IsDefault,
		Status:    req.Status,
		ValueType: req.ValueType,
		//UpdatedBy:   iotutil.ToInt64(req.UpdateBy),
		Remark:      req.Remark,
		Pinyin:      req.Pinyin,
		Firstletter: req.Firstletter,
		Listimg:     req.Listimg,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	//清理缓存
	iotredis.GetClient().Del(context.Background(), fmt.Sprintf("dict_data_%v", req.DictType))
	return nil
}

func (s BaseDataService) DeleteBaseData(id string) error {
	ret, err := rpc.TConfigDictDataServerService.Delete(context.Background(), &protosService.ConfigDictData{
		DictCode: iotutil.ToInt64(id),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) GetBaseDataTypeDetail(id string) (*protosService.ConfigDictTypeResponse, error) {
	rid := iotutil.ToInt64(id)
	ret, err := rpc.TConfigDictTypeServerService.FindById(context.Background(), &protosService.ConfigDictTypeFilter{DictId: rid})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, err
}

func (s BaseDataService) QueryBaseDataTypeList(filter entitys.BaseDataTypeQuery) (*protosService.ConfigDictTypeResponse, error) {
	QueryObj := protosService.ConfigDictType{
		DictName: filter.DictName,
		DictType: filter.DictType,
	}
	ret, err := rpc.TConfigDictTypeServerService.Lists(context.Background(), &protosService.ConfigDictTypeListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query:    &QueryObj,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, nil
}

// TConfigDictDataRequest
func (s BaseDataService) AddBaseDataType(req entitys.BaseDataType) error {
	ret, err := rpc.TConfigDictTypeServerService.Create(context.Background(), &protosService.ConfigDictType{
		DictName:  req.DictName,
		DictType:  req.DictType,
		ValueType: req.ValueType,
		IsSystem:  req.IsSystem,
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) UpdateBaseDataType(req entitys.BaseDataType) error {
	ret, err := rpc.TConfigDictTypeServerService.Update(context.Background(), &protosService.ConfigDictType{
		DictId:    iotutil.ToInt64(req.DictID),
		DictName:  req.DictName,
		DictType:  req.DictType,
		ValueType: req.ValueType,
		Status:    req.Status,
		//IsSystem:  req.IsSystem,
		Remark:    req.Remark,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
		//DeletedAt: "",
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) DeleteBaseDataType(id string) error {
	ret, err := rpc.TConfigDictTypeServerService.DeleteById(context.Background(), &protosService.ConfigDictType{
		DictId: iotutil.ToInt64(id),
	})
	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) GetTConfigTranslateDetail(code string) (*protosService.ConfigTranslateResponse, error) {
	ret, err := rpc.TConfigTranslateServerService.Find(context.Background(), &protosService.ConfigTranslateFilter{Code: code})
	//fmt.Println(rep1)
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, err
}

// TConfigDictDataRequest
func (s BaseDataService) AddTConfigTranslate(req entitys.TranslateParam) (int64, error) {
	id := iotutil.GetNextSeqInt64()
	ret, err := rpc.TConfigTranslateServerService.Create(context.Background(), &protosService.ConfigTranslate{
		Id:        id,
		Code:      req.Code,
		En:        req.En,
		Zh:        req.Zh,
		Jp:        req.Jp,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})
	if err != nil {
		return 0, err
	}
	if ret.Code != 200 {
		return 0, errors.New(ret.Message)
	}
	return id, err
}

// TConfigDictDataRequest
func (s BaseDataService) UpdateTConfigTranslate(req entitys.TranslateParam) error {
	ret, err := rpc.TConfigTranslateServerService.Update(context.Background(), &protosService.ConfigTranslate{
		Id:        iotutil.ToInt64(req.ID),
		Code:      req.Code,
		En:        req.En,
		Zh:        req.Zh,
		Jp:        req.Jp,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})

	if err != nil {
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	return err
}

func (s BaseDataService) QueryTranslateLanguageList() (*protosService.ConfigTranslateLanguageResponse, error) {
	QueryObj := protosService.ConfigTranslateLanguage{
		//Title: "1",
	}
	ret, err := rpc.TConfigTranslateLanguageServerService.Lists(context.Background(), &protosService.ConfigTranslateLanguageListRequest{
		Query: &QueryObj,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret, err
}

func (s BaseDataService) GetLangType() []*entitys.DictKeyVal {
	list, err := s.QueryBaseDataList(entitys.BaseDataQuery{
		DictType: "language_type",
	})
	if err != nil {
		return nil
	}
	var resLangType []*entitys.DictKeyVal
	for _, data := range list {
		resLangType = append(resLangType, &entitys.DictKeyVal{Name: data.DictLabel, Code: data.DictValue})
	}
	return resLangType
}

func (s BaseDataService) GetDictList(dictCode string, dictValue string) []*entitys.DictKeyVal {
	list, err := s.QueryBaseDataList(entitys.BaseDataQuery{
		DictType:  dictCode,
		DictValue: dictValue,
	})
	if err != nil {
		return nil
	}
	var resLangType []*entitys.DictKeyVal
	for _, data := range list {
		resLangType = append(resLangType, &entitys.DictKeyVal{Name: data.DictLabel, Code: data.DictValue})
	}
	return resLangType
}

// 通过语言字典名称获取字典列表数据
func (s BaseDataService) GetDictByType(dictType string) map[string]interface{} {
	list, err := s.QueryBaseDataList(entitys.BaseDataQuery{
		DictType: dictType,
	})
	if err != nil {
		return nil
	}
	resLangType := make(map[string]interface{})
	for _, data := range list {
		resLangType[data.DictValue] = data.DictLabel
	}
	return resLangType
}
