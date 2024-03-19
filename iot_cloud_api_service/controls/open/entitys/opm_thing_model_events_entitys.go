// Code generated by sgen.exe,2022-04-29 15:04:32. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type OpmThingModelEventsEntitys struct {
	Id            int64  `json:"id"`
	ModelId       int64  `json:"modelId"`
	ProductId     string `json:"productId"`
	CreateTs      string `json:"createTs"`
	Identifier    string `json:"identifier"`
	EventName     string `json:"eventName"`
	EventType     string `json:"eventType"`
	EventTypeDesc string `json:"eventTypeDesc"`
	Outputdata    string `json:"outputdata"`
	Required      int32  `json:"required"`
	Custom        int32  `json:"custom"`
	Desc          string `json:"desc"`
	Extension     string `json:"extension"`
}

// 查询条件
type OpmThingModelEventsQuery struct {
	Page      uint64                     `json:"page,omitempty"`
	Limit     uint64                     `json:"limit,omitempty"`
	Sort      string                     `json:"sort,omitempty"`
	SortField string                     `json:"sortField,omitempty"`
	SearchKey string                     `json:"searchKey,omitempty"`
	Query     *OpmThingModelEventsFilter `json:"query,omitempty"`
}

// OpmThingModelEventsFilter，查询条件，字段请根据需要自行增减
type OpmThingModelEventsFilter struct {
	Id         int64  `json:"id,omitempty"`
	ModelId    int64  `json:"modelId,omitempty"`
	ProductId  string `json:"productId,omitempty"`
	CreateTs   string `json:"createTs,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	EventName  string `json:"eventName,omitempty"`
	EventType  string `json:"eventType,omitempty"`
	Outputdata string `json:"outputdata,omitempty"`
	Required   int32  `json:"required,omitempty"`
	Custom     int32  `json:"custom,omitempty"`
	Extension  string `json:"extension,omitempty"`
}

// 实体转pb对象
func OpmThingModelEvents_e2pb(src *OpmThingModelEventsEntitys) *proto.OpmThingModelEvents {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmThingModelEvents{
		Id:         src.Id,
		ModelId:    src.ModelId,
		ProductId:  src.ProductId,
		CreateTs:   src.CreateTs,
		Identifier: src.Identifier,
		EventName:  src.EventName,
		EventType:  src.EventType,
		Outputdata: src.Outputdata,
		Required:   src.Required,
		Custom:     src.Custom,
		Desc:       src.Desc,
		//Extension:  src.Extension,
	}
	return &pbObj
}

// pb对象转实体
func OpmThingModelEvents_pb2e(src *proto.OpmThingModelEvents) *OpmThingModelEventsEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmThingModelEventsEntitys{
		Id:            src.Id,
		ModelId:       src.ModelId,
		ProductId:     src.ProductId,
		CreateTs:      src.CreateTs,
		Identifier:    src.Identifier,
		EventName:     src.EventName,
		EventType:     src.EventType,
		EventTypeDesc: changeEventType(src.EventType),
		Outputdata:    src.Outputdata,
		Required:      src.Required,
		Custom:        src.Custom,
		Desc:          src.Desc,
		//Extension:     src.Extension,
	}
	return &entitysObj
}

func changeEventType(eventType string) (res string) {
	switch eventType {
	case "INFO_EVENT_TYPE":
		res = "事件类型：信息"
	case "ALERT_EVENT_TYPE":
		res = "事件类型：告警"
	case "ERROR_EVENT_TYPE":
		res = "事件类型：故障"
	}
	return
}
