// Code generated by sgen.exe,2022-05-16 15:51:16. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

type TConfigDictData struct {
	DictCode  int64  `gorm:"column:dict_code;primaryKey;autoIncrement:true" json:"id,string,omitempty"` // 字典编码
	DictLabel string `gorm:"column:dict_label;default:''" json:"name,omitempty"`                        // 字典标签
	DictValue string `gorm:"column:dict_value;default:''" json:"value,omitempty"`                       // 字典键值
	DictType  string `gorm:"column:dict_type;default:''" json:"code,omitempty"`                         // 字典类型
	Listimg   string `gorm:"column:listimg" json:"listImg,omitempty"`                                   // 图片
}

type DictListParam struct {
	DictTypeList []string `json:"paramList"` //
}

func ConfigDictData_pb2db(src *proto.ConfigDictData) *TConfigDictData {
	if src == nil {
		return nil
	}
	dbObj := TConfigDictData{
		DictCode:  src.DictCode,
		DictLabel: src.DictLabel,
		DictValue: src.DictValue,
		DictType:  src.DictType,
		Listimg:   src.Listimg,
	}
	return &dbObj
}

func ConfigDictData_db2pb(src *TConfigDictData) *proto.ConfigDictData {
	if src == nil {
		return nil
	}
	pbObj := proto.ConfigDictData{
		DictCode:  src.DictCode,
		DictLabel: src.DictLabel,
		DictValue: src.DictValue,
		DictType:  src.DictType,
		Listimg:   src.Listimg,
	}
	return &pbObj
}

type WeatherData struct {
	Weather  string  `json:"weather"`  // 天气
	Humidity string  `json:"humidity"` // 湿度
	Temp     string  `json:"temp"`     // 温度
	Quality  string  `json:"quality"`  // 质量
	PicType  string  `json:"picType"`  // 图片类型
	Pm25     float64 `json:"pm25"`     // PM25
	Pm10     float64 `json:"pm10"`     // PM10
	Aqi      float64 `json:"aqi"`      // AQI
}
