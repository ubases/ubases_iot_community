package scene_executor

import (
	"bytes"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	context2 "context"
	"text/template"
	"time"
)

type RedisTempSet struct {
	RedisKey string
	Key      string
	Value    interface{}
	Duration int32
}

// 转换键值对
func setKeyValue(redisKey, key string, value interface{}, duration int32) *RedisTempSet {
	return &RedisTempSet{
		RedisKey: redisKey,
		Key:      key,
		Value:    value,
		Duration: duration,
	}
}

// 设置条件的当前redis缓存
func setRedisBeforeData(redisSaveQueue *[]*RedisTempSet) {
	for _, r := range *redisSaveQueue {
		if r.Duration == 0 {
			if r.Key == "" {
				iotredis.GetClient().Set(context2.Background(), r.RedisKey, r.Value, 0)
			} else {
				iotredis.GetClient().HSet(context2.Background(), r.RedisKey, r.Key, r.Value)
			}
		} else {
			iotredis.GetClient().Set(context2.Background(), r.RedisKey, r.Value, time.Duration(r.Duration)*time.Minute)
		}
	}
}

func GetTestIntelligenceId(id string) bool {
	deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), "test_intelligence_ids")
	if deviceCmd.Err() != nil {
		iotlogger.LogHelper.Error(deviceCmd.Err())
		return false
	}
	_, ok := deviceCmd.Val()[id]
	return ok
}

// setRuleParams 设备规则的参数
func setRuleParams(templateContent string, data interface{}) (string, error) {
	tmp, err := template.New("Email").Parse(templateContent)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
