package iotstruct

import "html/template"

type AppPanelPage struct {
	JsonContent  *AppPanelPageContent    `json:"jsonContent"`
	PopupContent []AppPanelComponentItem `json:"popupContent"`
}

type AppPanelPageContent struct {
	Type          string                  `json:"type"`
	VariableName  string                  `json:"variableName"`
	Name          string                  `json:"name"`
	Style         map[string]interface{}  `json:"style"`
	ComponentList []AppPanelComponentItem `json:"componentList"`
}

// 面板组件详情
type AppPanelComponentItem struct {
	Id           string                 `json:"id"` //组件Id，用于弹框效果
	Type         string                 `json:"type"`
	VariableName string                 `json:"variableName"`
	Name         string                 `json:"name"`
	Style        map[string]interface{} `json:"style"`
	IsLock       bool                   `json:"isLock"`
	AutoHeight   bool                   `json:"autoHeight"`
	//FunctionSet      string                    `json:"functionSet"`
	FunctionDpId           interface{}              `json:"functionDpId"`
	FunctionData           *FunctionData            `json:"functionData"`           //功能
	FunctionDatas          []*FunctionData          `json:"functionDatas"`          //功能集合（用于预约功能，需要显示的功能）
	StyleLinkageList       []StyleLinkageItem       `json:"styleLinkageList"`       //规则约束
	InteractionList        []InteractionItem        `json:"interactionList"`        //跳转弹框
	FunctionValuePropsList []FunctionValuePropsItem `json:"functionValuePropsList"` //跳转弹框
	FunctionPropsList      []FunctionPropsItem      `json:"functionPropsList"`      //跳转弹框
	Action                 []ComponentAction        `json:"action"`
	Logic                  []Component              `json:"logic"`
	ComponentList          []AppPanelComponentItem  `json:"componentList"`
	//ComponentLists         [][]AppPanelComponentItem `json:"componentLists"`
}

// 功能属性
type FunctionValuePropsItem struct {
	FunctionDpId     interface{}            `json:"functionDpId"`
	FunctionData     *FunctionData          `json:"functionData"`
	FunctionValue    interface{}            `json:"functionValue"`
	FunctionName     map[string]interface{} `json:"functionName"`
	StyleLinkageList []StyleLinkageItem     `json:"styleLinkageList"`
}

type FunctionValuePropsItemExt struct {
	FunctionDpId     interface{}        `json:"functionDpId"`
	FunctionData     *FunctionData      `json:"functionData"`
	FunctionValue    interface{}        `json:"functionValue"`
	FunctionName     string             `json:"functionName"`
	StyleLinkageList []StyleLinkageItem `json:"styleLinkageList"`
}

// 功能值属性
type FunctionPropsItem struct {
	FunctionDpId     interface{}            `json:"functionDpId"`
	FunctionData     *FunctionData          `json:"functionData"`
	FunctionName     map[string]interface{} `json:"functionName"`
	StyleLinkageList []StyleLinkageItem     `json:"styleLinkageList"`
}
type FunctionPropsItemExt struct {
	FunctionDpId     interface{}        `json:"functionDpId"`
	FunctionData     *FunctionData      `json:"functionData"`
	FunctionName     string             `json:"functionName"`
	StyleLinkageList []StyleLinkageItem `json:"styleLinkageList"`
}

// 样式连接规则
type StyleLinkageItem struct {
	RuleName  string    `json:"ruleName"`
	IfSpecs   IfSpecs   `json:"ifSpecs"`
	ThenSpecs ThenSpecs `json:"thenSpecs"`
}

type IfSpecs struct {
	DpID    interface{} `json:"dpId"`
	Operate interface{} `json:"operate"`
	Value   interface{} `json:"value"`
}

type ThenSpecs struct {
	DpID    interface{} `json:"dpId"`
	Operate interface{} `json:"operate"`
}

// 页面跳转弹框定义
type InteractionItem struct {
	Function interface{} `json:"function"`
	Action   interface{} `json:"action,omitempty"`
	Page     string      `json:"page,omitempty"`
	Popup    string      `json:"popup,omitempty"`
	DpID     interface{} `json:"dpId,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	LinkUrl  string      `json:"linkUrl,omitempty"`
}

// 功能数据，与物模型相关
type FunctionData struct {
	Dpid          interface{} `json:"dpid"`
	DataType      string      `json:"dataType"`
	Name          string      `json:"name"`
	RwFlag        string      `json:"rwFlag"`
	DataSpecs     string      `json:"dataSpecs"`
	DataSpecsList string      `json:"dataSpecsList"`
	Required      int         `json:"required"`
	Value         interface{} `json:"value"`
	Identifier    string      `json:"identifier"`
	DefaultVal    string      `json:"defaultVal"`
}

type Style struct {
	Left              int                    `json:"left"`
	Top               int                    `json:"top"`
	Width             int                    `json:"width"`
	Height            int                    `json:"height"`
	IsDeviceName      int                    `json:"isDeviceName"`
	HeaderTitle       map[string]interface{} `json:"headerTitle"`
	BackgroundColor   string                 `json:"backgroundColor"`
	BackgroundOpacity int                    `json:"backgroundOpacity"`
	FontColor         string                 `json:"fontColor"`
	FontOpacity       int                    `json:"fontOpacity"`
	BorderColor       string                 `json:"borderColor"`
	BorderOpacity     int                    `json:"borderOpacity"`
	BorderWidth       int                    `json:"borderWidth"`
	BorderRadius      int                    `json:"borderRadius"`
	IsScroll          bool                   `json:"isScroll"`
	ScrollDirection   string                 `json:"scrollDirection"`
	Text              map[string]interface{} `json:"text"`
	BtnColor          string                 `json:"btnColor"`
	BtnOpacity        int                    `json:"btnOpacity"`
	ImgURL            string                 `json:"imgUrl"`
	ImgName           string                 `json:"imgName"`
	FontSize          int                    `json:"fontSize"`
	TextAlign         string                 `json:"textAlign"`
	IsLockHeight      bool                   `json:"isLockHeight"`
	IsLockWidth       bool                   `json:"isLockWidth"`
	OpenColor         string                 `json:"openColor"`
	OpenOpacity       int                    `json:"openOpacity"`
	CloseColor        string                 `json:"closeColor"`
	CloseOpacity      int                    `json:"closeOpacity"`
	SwitchSize        int                    `json:"switchSize"`
}

type ComponentAction struct {
	ActionType  string `json:"actionType"`
	Target      string `json:"target"`
	TargetValue string `json:"targetValue"`
}

type Component struct {
	Name      string           `json:"name"`
	Condition []LogicCondition `json:"condition"`
	Result    []LogicResult    `json:"result"`
}

type LogicCondition struct {
	Func      string `json:"func"`
	Operator  string `json:"operator"`
	FuncValue string `json:"funcValue"`
}

type LogicResult struct {
	StyleKey   string `json:"styleKey"`
	StyleValue string `json:"styleValue"`
}

// 转换后的内容对象
type ContentObj struct {
	PageName            string
	PageIdentify        string
	VueComponentContent template.HTML
	VuePopupContent     template.HTML
	VueDataContent      template.JS
	VueMethodContent    template.JS
	VueCallbackContent  template.JS
	VueFuncDefine       template.JS
}

type TagConfig struct {
	TagName string
	Attrs   map[string]string
}
