/**
 * @Author: hogan
 * @Date: 2021/11/16 14:44
 */
package pushModel

import (
	"cloud_platform/iot_common/iotutil"
)

/*
消息的结构
{
	"userid":"用户",
	"homeid":"家庭",
	"type":"alarm:消息类型（告警消息）home:家庭消息 notice通知消息",
	"objecttype":"1设备/2智能/3群组",
	"objectid":"设备id/智能id/群组id",
	"content":"消息内容",
	"childrens":[
		"","",""
	],
	"isread":false,
}
*/

type TagType int32

const (
	ALIAS TagType = 0
	TAGS  TagType = 1
	REGID TagType = 2
)

// 消息结构
type MessageModel struct {
	Type        string   `json:"type"`        //alarm:消息类型（告警消息）home:家庭消息 notice通知消息
	ChildType   string   `json:"childtype"`   //
	ObjectType  string   `json:"objecttype"`  //1设备/2智能/3群组
	ObjectId    string   `json:"objectid"`    //设备id/智能id/群组id
	Model       string   `json:"model"`       //产品Model
	Title       string   `json:"title"`       //标题
	Devids      string   `json:"devids"`      //设备Id列表
	Extands     string   `json:"extands"`     //扩展数据类型
	Content     string   `json:"content"`     //内容
	Url         string   `json:"url"`         //外链地址
	Params      []string `json:"params"`      //翻译参数
	IsRead      bool     `json:"isread"`      //是否已读
	UnSetExpire bool     `json:"unsetexpire"` //不设置超时
	IsPublic    string   `json:"ispublic"`    //0 普通消息 1 公共消息
	Result      string   `json:"result"`      //结果标识
	HomeId      string   `json:"homeid"`      //家庭编号
	UserId      []string `json:"userid"`      //用户编号
	CreateTime  int64    `json:"createtime"`  //创建时间
}

// 消息结构接收数据
type MessageRequestModel struct {
	MessageId   string                 `json:"messageId"` //消息id
	Type        string                 `json:"type"`      //消息类型 device、alrame、notice
	ChildType   string                 `json:"childtype"`
	ObjectType  string                 `json:"objecttype"`
	ObjectId    string                 `json:"objectid"`
	Model       string                 `json:"model"`
	DevImg      string                 `json:"devimg"`
	HomeName    string                 `json:"homename"`
	Title       string                 `json:"title"`
	Content     string                 `json:"content"`
	Devids      string                 `json:"devids"`
	Extands     map[string]interface{} `json:"extands"`
	Url         string                 `json:"url"` //外链地址
	Params      map[string]string      `json:"params"`
	IsRead      bool                   `json:"isread"`
	UnSetExpire bool                   `json:"unsetexpire"` //不设置超时
	IsPublic    int32                  `json:"ispublic"`    //0 普通消息 1 公共消息
	TagType     int32                  `json:"tagtype"`     //标签类型 0：alias 、1：tag、 2：regid
	HomeId      string                 `json:"homeid"`
	UserId      string                 `json:"userid"`
	UserIds     []string               `json:"userids"`
	Result      string                 `json:"result"`
	CreateTime  int64                  `json:"createtime"`
	ExpireTime  int64                  `json:"expiretime"` //消息过期时间
	CreatedAt   int64                  `json:"created_at" bson:"created_at"`
	CreatorId   string                 `json:"creator_id"`
	CreatorName string                 `json:"creator_name"`
	AppKey      string                 `json:"appKey"`   // APP Key
	TenantId    string                 `json:"tenantId"` // 开发者租户编号
	TplCode     string                 `json:"tplCode"`
}

type MessageResponseModel struct {
	Type       string `json:"type"` //消息类型 device、alrame、notice
	ChildType  string `json:"childtype"`
	ObjectType string `json:"objecttype"`
	ObjectId   string `json:"objectid"`
	Model      string `json:"model"`
	DevImg     string `json:"devimg"`
	HomeName   string `json:"homename"`
	Title      string `json:"title"`
	Devids     string `json:"devids"`
	Content    string `json:"content"`
	Url        string `json:"url"` //外链地址
	HomeId     string `json:"homeid"`
	CreateTime int64  `json:"createtime"`
	ExpireTime int64  `json:"expiretime"` //消息过期时间
}

// 消息发送实体
type MessageSendModel struct {
	MessageId  string `json:"messageId"` //消息id
	Type       string `json:"msgType"`   //消息类型 device、alrame、notice
	ChildType  string `json:"childType"`
	ObjectId   string `json:"objectId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Url        string `json:"url"`        //外链地址
	Sound      string `json:"sound"`      //声音
	CreateTime int64  `json:"createTime"` //
	ExpireTime int64  `json:"expireTime"` //消息过期时间
}

func (s MessageSendModel) SetModel(model MessageRequestModel) MessageSendModel {
	s.Type = model.Type
	s.ChildType = model.ChildType
	s.ObjectId = model.ObjectId
	s.Title = model.Title
	s.Content = model.Content
	s.CreateTime = model.CreateTime
	s.ExpireTime = model.ExpireTime
	s.MessageId = model.MessageId

	//扩展信息特殊处理
	extands := model.Extands
	if extands != nil {
		cueSound, ok := extands["cueSound"]
		if ok && cueSound != "" {
			s.Sound = iotutil.ToString(cueSound)
		}
	}
	return s
}

// 消息目标
type MessageTarget struct {
	IsPublic   bool            `json:"isPublic"`
	Tags       []string        `json:"tags"`
	Alias      []string        `json:"alias"`
	RegIds     []string        `json:"regIds"`
	PushTokens []PushTokenItem `json:"pushTokens"` //推送tokens数组，包含推送平台，推送目标token
}

type PushTokenItem struct {
	Lang            string `gorm:"column:lang" json:"lang"`                                  //所属语言
	UserId          int64  `gorm:"column:user_id" json:"userId"`                             // 用户Id
	AppPushId       string `gorm:"column:app_push_id;not null" json:"appPushId"`             // APP生成的手机唯一Id
	AppKey          string `gorm:"column:app_key" json:"appKey"`                             // APP Key
	TenantId        string `gorm:"column:tenant_id" json:"tenantId"`                         // 开发者租户编号
	RegionId        int64  `gorm:"column:region_id" json:"regionId"`                         // 区域服务Id
	AppToken        string `gorm:"column:app_token;not null" json:"appToken"`                // APP生成的推送Token
	AppPushPlatform string `gorm:"column:app_push_platform;not null" json:"appPushPlatform"` // APP推送注册平台 ios、android、huawei
	AppPacketName   string `gorm:"column:app_packet_name" json:"appPacketName"`              // APP包名
}
