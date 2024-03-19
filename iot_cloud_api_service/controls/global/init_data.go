package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"encoding/json"
)

func RefreshDictCache() {
	go CacheDict()
}

// CacheDict 初始化缓存
func CacheDict() {
	defer iotutil.PanicHandler()
	var filter entitys.BaseDataQuery = entitys.BaseDataQuery{}
	res, err := services.BaseDataService{}.QueryBaseDataList(filter)
	if err != nil {
		return
	}
	result := make(map[string]map[string]interface{}, 0)
	for _, item := range res {
		dictType := iotutil.ToString(item.DictType)
		if _, ok := result[dictType]; !ok {
			result[dictType] = map[string]interface{}{}
		}
		result[dictType][item.DictValue] = item.DictLabel
	}
	for k, m := range result {
		iotredis.GetClient().Set(context.Background(), persist.GetRedisKey(iotconst.DICT_DATA, k), m, 0)
		//iotredis.GetClient().Set(context.Background(), k, iotutil.ToString(m), 0)
	}
}

type DictTempData struct {
	data map[string]interface{}
}

func (s *DictTempData) GetData() map[string]interface{} {
	return s.data
}

func (s *DictTempData) GetDictByCode(code string) (*DictTempData, error) {
	if c := iotredis.GetClient().Get(context.Background(), persist.GetRedisKey(iotconst.DICT_DATA, code)); c.Err() == nil {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(c.Val()), &data)
		if err != nil {
			goto reload
		}
		if data == nil || len(data) == 0 {
			goto reload
		}
		s.data = data
		return s, nil
	}
reload:
	s.data = services.BaseDataService{}.GetDictByType(code)
	c := iotredis.GetClient().Set(context.Background(), persist.GetRedisKey(iotconst.DICT_DATA, code), iotutil.ToString(s.data), 0)
	if c.Err() != nil {
		return s, c.Err()
	}
	return s, nil
}

func (s *DictTempData) Value(k int32) string {
	val, ok := s.data[iotutil.ToString(k)]
	if ok {
		return val.(string)
	}
	return ""
}

func (s *DictTempData) ValueStr(k string) string {
	val, ok := s.data[k]
	if ok {
		return val.(string)
	}
	return ""
}
