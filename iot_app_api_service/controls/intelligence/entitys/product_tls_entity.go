package entitys

import (
	productEntity "cloud_platform/iot_app_api_service/controls/product/entitys"
)

type ProductThingsModel struct {
	Services     []*ThingModelServices                      `json:"actions"`
	Events       []*ThingModelEvents                        `json:"events"`
	Properties   []*ThingModelProperties                    `json:"attrs"`
	StyleLinkage map[string]interface{}                     `json:"styleLinkage"` //面板交互样式
	Rules        []productEntity.ThingModelRuleItemResponse `json:"rules,omitempty"`
}

type ThingModelEvents struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	EventType  string `json:"eventType"`
	OutputData string `json:"outputData"`
}

type ThingModelProperties struct {
	Name          string                  `json:"name"`
	Identifier    string                  `json:"identifier"`     //标识符类型原来的pid
	DpId          string                  `json:"dpId,omitempty"` //标识符类型原来的pid
	DataType      string                  `json:"dataType"`
	DataSpecs     string                  `json:"dataSpecs"`
	DataSpecsList string                  `json:"dataSpecsList"`
	RwFlag        string                  `json:"rwFlag"`
	DefaultVal    interface{}             `json:"defaultVal"`
	Childrens     []*ThingModelProperties `json:"childrens"`
	Sort          int32                   `json:"sort"`
}

type ThingModelServices struct {
	Name         string `json:"name"`
	Identifier   string `json:"identifier"`
	ServiceName  string `json:"serviceName"`
	InputParams  string `json:"inputParams"`
	OutputParams string `json:"outputParams"`
	CallType     int32  `json:"callType"`
}
