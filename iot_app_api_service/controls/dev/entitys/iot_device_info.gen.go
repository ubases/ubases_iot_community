// Code generated by sgen.exe,2022-04-21 14:54:14. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package entitys

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_device/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TIotDeviceInfo struct {
	Id               int64   `gorm:"column:id;primaryKey" json:"id,string"`              // 主键ID
	Did              string  `gorm:"column:did;not null" json:"did"`                     // 设备唯一ID（14位 1~9 A~Z随机）
	ProductId        int64   `gorm:"column:product_id;not null" json:"productId,string"` // 产品ID(t_cloud_product.id)
	OnlineStatus     int32   `gorm:"column:online_status" json:"onlineStatus"`           // 在线状态（0 在线 1 不在线）
	DeviceName       string  `gorm:"column:device_name" json:"deviceName"`               // 设备名称
	DeviceNature     string  `gorm:"column:device_nature" json:"deviceNature"`           // 设备性质
	Sn               string  `gorm:"column:sn" json:"sn"`                                // 设备SN
	BatchId          int64   `gorm:"column:batch_id" json:"batchId"`                     // 批次ID(t_cloud_device_batch.id)
	GroupId          int64   `gorm:"column:group_id" json:"groupId"`                     // 设备组ID（t_cloud_device_group.id）
	DeviceModel      string  `gorm:"column:device_model" json:"deviceModel"`             // 设备型号
	UserName         string  `gorm:"column:user_name" json:"userName"`                   // 用户名
	Passward         string  `gorm:"column:passward" json:"passward"`                    // 设备密码
	Salt             string  `gorm:"column:salt" json:"salt"`                            // 盐值
	DeviceSecretHttp string  `gorm:"column:device_secret_http" json:"deviceSecretHttp"`  // 设备密钥（http）
	DeviceSecretMqtt string  `gorm:"column:device_secret_mqtt" json:"deviceSecretMqtt"`  // 设备密钥（mqtt）
	IpAddress        string  `gorm:"column:ip_address" json:"ipAddress"`                 // ip地址
	Lat              float64 `gorm:"column:lat" json:"lat"`                              // 纬度
	Lng              float64 `gorm:"column:lng" json:"lng"`                              // 经度
	Country          string  `gorm:"column:country" json:"country"`                      // 国家编码
	Province         string  `gorm:"column:province" json:"province"`                    // 省份编码
	City             string  `gorm:"column:city" json:"city"`                            // 城市编码
	District         string  `gorm:"column:district" json:"district"`                    // 地区编码
	MacAddress       string  `gorm:"column:mac_address" json:"macAddress"`               // mac地址
	DeviceVersion    string  `gorm:"column:device_version" json:"deviceVersion"`         // 设备版本
	ActiveStatus     string  `gorm:"column:active_status" json:"activeStatus"`           // 激活状态[0:未激活,1:已激活]
}

type CurrentIotDeviceInfo struct {
	AccessToken      string                 `protobuf:"bytes,101,opt,name=accessToken,proto3" json:"accessToken"` //
	Batch            string                 `protobuf:"bytes,102,opt,name=batch,proto3" json:"batch"`             //
	DevId            string                 `protobuf:"bytes,103,opt,name=devId,proto3" json:"devId"`             //
	HomeId           string                 `protobuf:"bytes,104,opt,name=homeId,proto3" json:"homeId"`           //
	Model            string                 `protobuf:"bytes,105,opt,name=model,proto3" json:"productKey"`        //
	Name             string                 `protobuf:"bytes,106,opt,name=name,proto3" json:"name"`               //
	NetworkMode      string                 `protobuf:"bytes,107,opt,name=networkMode,proto3" json:"networkMode"` //
	Pic              string                 `protobuf:"bytes,108,opt,name=pic,proto3" json:"pic"`                 //
	RoomId           string                 `protobuf:"bytes,109,opt,name=roomId,proto3" json:"roomId"`           //
	Secretkey        string                 `protobuf:"bytes,110,opt,name=secretkey,proto3" json:"secretKey"`     //
	Ssid             string                 `protobuf:"bytes,111,opt,name=ssid,proto3" json:"ssid"`               //
	State            int32                  `protobuf:"varint,112,opt,name=state,proto3" json:"state"`            //
	Switch           int32                  `protobuf:"varint,113,opt,name=switch,proto3" json:"switch"`          //
	ProductId        string                 `json:"productId"`                                                    //
	MqttServer       string                 `json:"mqttServer"`                                                   //
	RoomName         string                 `json:"roomName"`                                                     //
	ProductName      string                 `json:"productName"`                                                  //
	HomeRoomList     []HomeRoomParam        `json:"homeRoomList"`                                                 //
	DeviceStatus     map[string]interface{} `json:"deviceStatus"`                                                 //
	Props            map[string]TslInfo     `json:"tsl"`                                                          //
	DeviceVersion    string                 `json:"version"`                                                      //
	DeviceMcuVersion string                 `json:"mcuVersion"`                                                   //
	DeviceType       int32                  `json:"devType" 													//1-用户设备，2-共享设备 `
	BelongUserName   string                 `json:"belongUserName"` //用户昵称
	ReceiveShareId   string                 `json:"receiveShareId"` //接受共享id
	IsShowImg        bool                   `json:"isShowImg"`      //显示图片
	PanelProImg      string                 `json:"panelProImg"`    //面板产品图片
	StyleLinkage     map[string]interface{} `json:"styleLinkage"`   //面板交互样式
}

type TslInfo struct {
	DpId          int32       `json:"dpid"`
	DataType      string      `json:"dataType"`
	Name          string      `json:"name"`
	RwFlag        string      `json:"rwFlag,omitempty"`
	DataSpecs     string      `json:"dataSpecs"`
	DataSpecsList string      `json:"dataSpecsList"`
	Required      int32       `json:"required,omitempty"`
	Value         string      `json:"value"`
	Identifier    string      `json:"identifier"`
	DefaultVal    interface{} `json:"defaultVal"`
}

type HomeRoomParam struct {
	RoomId string `json:"roomId"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Sort   int32  `json:"sort"`
}

func IotDeviceInfo_pb2db(src *proto.IotDeviceInfo) *TIotDeviceInfo {
	if src == nil {
		return nil
	}
	dbObj := TIotDeviceInfo{
		Id:               src.Id,
		Did:              src.Did,
		ProductId:        src.ProductId,
		OnlineStatus:     src.OnlineStatus,
		DeviceName:       src.DeviceName,
		DeviceNature:     src.DeviceNature,
		Sn:               src.Sn,
		BatchId:          src.BatchId,
		GroupId:          src.GroupId,
		DeviceModel:      src.DeviceModel,
		UserName:         src.UserName,
		Passward:         src.Passward,
		Salt:             src.Salt,
		DeviceSecretHttp: src.DeviceSecretHttp,
		DeviceSecretMqtt: src.DeviceSecretMqtt,
		IpAddress:        src.IpAddress,
		Lat:              src.Lat,
		Lng:              src.Lng,
		Country:          src.Country,
		Province:         src.Province,
		City:             src.City,
		District:         src.District,
		MacAddress:       src.MacAddress,
		DeviceVersion:    src.DeviceVersion,
		ActiveStatus:     src.ActiveStatus,
	}
	return &dbObj
}

func IotDeviceInfo_db2pb(src *model.TIotDeviceInfo) *proto.IotDeviceInfo {
	if src == nil {
		return nil
	}
	pbObj := proto.IotDeviceInfo{
		Id:               iotutil.ToInt64(src.Id),
		Did:              src.Did,
		ProductId:        iotutil.ToInt64(src.ProductId),
		OnlineStatus:     src.OnlineStatus,
		DeviceName:       src.DeviceName,
		DeviceNature:     src.DeviceNature,
		Sn:               src.Sn,
		BatchId:          src.BatchId,
		GroupId:          src.GroupId,
		DeviceModel:      src.DeviceModel,
		UserName:         src.UserName,
		Passward:         src.Passward,
		Salt:             src.Salt,
		DeviceSecretHttp: src.DeviceSecretHttp,
		DeviceSecretMqtt: src.DeviceSecretMqtt,
		IpAddress:        src.IpAddress,
		Lat:              src.Lat,
		Lng:              src.Lng,
		Country:          src.Country,
		Province:         src.Province,
		City:             src.City,
		District:         src.District,
		ActivatedTime:    timestamppb.New(src.ActivatedTime),
		MacAddress:       src.MacAddress,
		DeviceVersion:    src.DeviceVersion,
		ActiveStatus:     src.ActiveStatus,
		CreatedBy:        src.CreatedBy,
		UpdatedBy:        src.UpdatedBy,
		CreatedAt:        timestamppb.New(src.CreatedAt),
		UpdatedAt:        timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

func IotDevInfo_db2pb(src *model.TIotDeviceInfo) *proto.IotDevInfo {
	if src == nil {
		return nil
	}
	pbObj := proto.IotDevInfo{
		Id:               iotutil.ToInt64(src.Id),
		Did:              src.Did,
		ProductId:        iotutil.ToInt64(src.ProductId),
		OnlineStatus:     src.OnlineStatus,
		DeviceName:       src.DeviceName,
		DeviceNature:     src.DeviceNature,
		Sn:               src.Sn,
		BatchId:          src.BatchId,
		GroupId:          src.GroupId,
		DeviceModel:      src.DeviceModel,
		UserName:         src.UserName,
		Passward:         src.Passward,
		Salt:             src.Salt,
		DeviceSecretHttp: src.DeviceSecretHttp,
		DeviceSecretMqtt: src.DeviceSecretMqtt,
		IpAddress:        src.IpAddress,
		Lat:              src.Lat,
		Lng:              src.Lng,
		Country:          src.Country,
		Province:         src.Province,
		City:             src.City,
		District:         src.District,
		ActivatedTime:    timestamppb.New(src.ActivatedTime),
		MacAddress:       src.MacAddress,
		DeviceVersion:    src.DeviceVersion,
		ActiveStatus:     src.ActiveStatus,
	}
	return &pbObj
}

func CurrentDeviceInfo_pb2db(src *proto.CurrentIotDeviceInfo) *CurrentIotDeviceInfo {
	if src == nil {
		return nil
	}
	homeRoomList := []HomeRoomParam{}
	for _, v := range src.HomeRoomList {
		homeRoomList = append(homeRoomList, HomeRoomParam{
			RoomId: iotutil.ToString(v.RoomId),
			Name:   v.Name,
			Icon:   v.Icon,
			Sort:   v.Sort,
		})
	}
	var newDeviceStatus map[string]interface{}
	_ = json.Unmarshal(src.DeviceStatus, &newDeviceStatus)
	dbObj := CurrentIotDeviceInfo{
		AccessToken:      src.AccessToken,
		Batch:            src.Batch,
		DevId:            src.DevId,
		HomeId:           src.HomeId,
		Model:            src.Model,
		Name:             src.Name,
		NetworkMode:      src.NetworkMode,
		Pic:              src.Pic,
		RoomId:           src.RoomId,
		Secretkey:        src.Secretkey,
		Ssid:             src.Ssid,
		State:            src.State,
		Switch:           src.Switch,
		ProductId:        iotutil.ToString(src.ProductId),
		MqttServer:       src.MqttServer,
		RoomName:         src.RoomName,
		HomeRoomList:     homeRoomList,
		DeviceStatus:     newDeviceStatus,
		DeviceVersion:    src.DeviceVersion,
		DeviceMcuVersion: src.DeviceMcuVersion,
		IsShowImg:        iotutil.IntToBoolean(src.IsShowImg),
		PanelProImg:      src.PanelProImg,
	}
	if src.StyleLinkage != "" {
		dbObj.StyleLinkage, _ = iotutil.JsonToMapErr(src.StyleLinkage)
	}
	return &dbObj
}

type CheckDeviceParam struct {
	Token string `json:"token"` // token
}

type RemoveDevFilter struct {
	HomeId    string   `json:"homeId"`    // 家庭id
	RoomId    string   `json:"roomId"`    // 房间id
	DevId     string   `json:"devId"`     // 设备id
	DevIdList []string `json:"devIdList"` // 设备id集合
}

type UpdateDevFilter struct {
	HomeId  string `json:"homeId"`  // 家庭id
	RoomId  string `json:"roomId"`  // 房间id
	DevId   string `json:"devId"`   // 设备id
	DevName string `json:"devName"` // 设备名称
}

type AddDevFilter struct {
	HomeId string `json:"homeId"` // 家庭id
	RoomId string `json:"roomId"` // 房间id
	DevId  string `json:"devId"`  // 设备id
	Sort   int32  `json:"sort"`   // 排序
}

// 设备列表传输对象
type SharedDeviceListEntityDto struct {
	DevId    string `json:"devId"`      //设备编号
	Model    string `json:"productKey"` //产品Model
	RoomId   string `json:"roomId"`     //所属房间
	Name     string `json:"name"`       //设备名称
	Time     int64  `json:"time"`       //共享时间
	Pic      string `json:"pic"`        //设备图片
	RoomName string `json:"roomName"`
	//Server ServerSimpleEntity `json:"server"`	//控制单元地址
	MqttServer string `json:"mqttServer"`
	Type       int32  `json:"type"` //1-已共享 2-未共享
}

// 设备已共享设备数据表
type SharedDevice struct {
	DevName      string `json:"devname"`      //设备名称
	DevId        string `json:"devid"`        //设备ID
	RoomId       string `json:"roomid"`       //所属房间id
	Model        string `json:"model"`        //设备model
	Time         int64  `json:"time"`         //共享时间
	UserPhone    string `json:"userphone"`    //共享用户手机
	UserEmail    string `json:"useremail"`    //共享用户邮箱
	HomeId       string `json:"homeid"`       //设备归属家庭
	UserId       string `json:"userid"`       //共享用户ID
	UserPhoto    string `json:"userphoto"`    //用户图像
	ServerId     string `json:"serverid"`     //控制单元地址
	UserNickname string `json:"usernickname"` //共享用户昵称
}

// 共享用户列表
type SharedUserDto struct {
	Id           string `json:"id"`           //共享id
	Time         int64  `json:"time"`         //共享时间
	UserId       string `json:"userId"`       //共享用户ID
	UserNickname string `json:"userNickname"` //共享用户昵称
	Type         int32  `json:"type"`         //1-已接受邀请 2-等待接受
	Photo        string `json:"photo"`        //用户头像
}

type CancelShare struct {
	Id   string `json:"id"`   //id
	Type int32  `json:"type"` //类型  1-取消已接受邀请，2-取消等待接受
}

type CancelReceiveShared struct {
	Id string `json:"id"` //id
}

type Addshared struct {
	HomeId     string `json:"homeId"`     //家庭id
	DevId      string `json:"devId"`      //设备id
	RoomId     string `json:"roomId"`     //房间id
	ProductKey string `json:"productKey"` //设备productKey
	DevName    string `json:"devName"`    //设备名称
	Account    string `json:"account"`    //用户信息
	Type       int32  `json:"type"`       //用户类型 1-手机,2-邮箱
}

// 设备共享接收数据表
type ReceiveShared struct {
	DevName        string `json:"devName"`        //设备名称
	DevId          string `json:"devId"`          //设备ID
	Roomid         string `json:"roomId"`         //房间id
	Model          string `json:"productKey"`     //设备model
	Time           int64  `json:"time"`           //共享时间
	UserPhone      string `json:"userPhone"`      //共享用户手机
	UserEmail      string `json:"userEmail"`      //共享用户邮箱
	HomeId         string `json:"homeId"`         //设备归属家庭
	UserId         string `json:"userId"`         //共享用户ID
	UserPhoto      string `json:"userPhoto"`      //用户图像
	UserNickname   string `json:"userNickname"`   //共享用户昵称
	IsAgree        int32  `json:"isAgree"`        //是否同意   1-同意，0-未同意
	BelongUserId   string `json:"belongUserid"`   //设备归属用户ID
	BelongUserName string `json:"belongUsername"` //设备归属用户名称
}

func IotDeviceShared_pb2db(src *proto.IotDeviceShared) *model.TIotDeviceShared {
	if src == nil {
		return nil
	}
	dbObj := model.TIotDeviceShared{
		Id:         src.Id,
		CustomName: src.CustomName,
		UserId:     src.UserId,
		Phone:      src.Phone,
		Email:      src.Email,
		DeviceId:   src.DeviceId,
		HomeId:     src.HomeId,
		Photo:      src.Photo,
		Sid:        src.Sid,
		SharedTime: src.SharedTime.AsTime(),
		CreatedAt:  src.CreatedAt.AsTime(),
		UpdatedAt:  src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func IotDeviceShared_db2pb(src *model.TIotDeviceShared) *proto.IotDeviceShared {
	if src == nil {
		return nil
	}
	pbObj := proto.IotDeviceShared{
		Id:         src.Id,
		CustomName: src.CustomName,
		UserId:     src.UserId,
		Phone:      src.Phone,
		Email:      src.Email,
		DeviceId:   src.DeviceId,
		HomeId:     src.HomeId,
		Photo:      src.Photo,
		Sid:        src.Sid,
		SharedTime: timestamppb.New(src.SharedTime),
		CreatedAt:  timestamppb.New(src.CreatedAt),
		UpdatedAt:  timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

func IotDeviceShareReceive_pb2db(src *proto.IotDeviceShareReceive) *model.TIotDeviceShareReceive {
	if src == nil {
		return nil
	}
	dbObj := model.TIotDeviceShareReceive{
		Id:           src.Id,
		CustomName:   src.CustomName,
		UserId:       src.UserId,
		Phone:        src.Phone,
		Email:        src.Email,
		DeviceId:     src.DeviceId,
		BelongUserId: src.BelongUserId,
		HomeId:       src.HomeId,
		RoomId:       src.RoomId,
		ProductKey:   src.ProductKey,
		Photo:        src.Photo,
		IsAgree:      src.IsAgree,
		SharedTime:   src.SharedTime.AsTime(),
		CreatedAt:    src.CreatedAt.AsTime(),
		UpdatedAt:    src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func IotDeviceShareReceive_db2pb(src *model.TIotDeviceShareReceive) *proto.IotDeviceShareReceive {
	if src == nil {
		return nil
	}
	pbObj := proto.IotDeviceShareReceive{
		Id:           src.Id,
		CustomName:   src.CustomName,
		UserId:       src.UserId,
		Phone:        src.Phone,
		Email:        src.Email,
		DeviceId:     src.DeviceId,
		BelongUserId: src.BelongUserId,
		HomeId:       src.HomeId,
		RoomId:       src.RoomId,
		ProductKey:   src.ProductKey,
		Photo:        src.Photo,
		IsAgree:      src.IsAgree,
		SharedTime:   timestamppb.New(src.SharedTime),
		CreatedAt:    timestamppb.New(src.CreatedAt),
		UpdatedAt:    timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// 接收共享设备列表
type ReceiveSharedDto struct {
	Id             string `json:"id"`             //接收共享设备ID
	DevName        string `json:"devName"`        //设备名称
	ProductKey     string `json:"productKey"`     //产品key
	DeviceImg      string `json:"deviceImg"`      //设备图片
	BelongUserName string `json:"belongUserName"` //用户昵称
	Status         int32  `json:"status"`         //1 待同意 ，2 已同意 ，3 已过期
}

type DeviceGroupListDto struct {
	Id         string `json:"id"`         //编号
	DevId      string `json:"devId"`      //设备id
	DevName    string `json:"devName"`    //设备名称
	ProductKey string `json:"productKey"` //产品productKey
	Pic        string `json:"pic"`        //设备图片
}

type UpsertGroup struct {
	GroupId    string    `json:"groupId"`    //群组id
	GroupName  string    `json:"groupName"`  //群组名称
	RoomId     string    `json:"roomId"`     //房间id
	RoomName   string    `json:"roomName"`   //房间名称
	HomeId     string    `json:"homeId"`     //家庭id
	DeviceList []DevInfo `json:"deviceList"` //设备列表
}

type DevInfo struct {
	Devid      string `json:"devId"`      //设备id
	Id         string `json:"id"`         //编号
	ProductKey string `json:"productKey"` //产品Key
}

type DeviceGroupList struct {
	Id         string `json:"id"`         //id
	GroupId    string `json:"groupId"`    //群组编号
	DevName    string `json:"devName"`    //设备名称
	DevId      string `json:"devId"`      //设备id
	JoinTime   int64  `json:"joinTime"`   //加入群组时间
	ProductKey string `json:"productKey"` //产品model
	UserId     string `json:"userId"`     //操作用户编号
	HomeId     string `json:"homeId"`     //家庭编号
	Pic        string `json:"pic"`        //设备图片
}

func UpsertGroup_db2pb(src UpsertGroup) *proto.UpsertGroupRequest {
	//if src == nil {
	//	return nil
	//}

	pbObj := proto.UpsertGroupRequest{
		IotDeviceGroupReq: &proto.IotDeviceGroupReq{
			GroupId:   src.GroupId,
			GroupName: src.GroupName,
			RoomId:    src.RoomId,
			RoomName:  src.RoomName,
		},
	}
	if src.DeviceList != nil {
		deviceList := []*proto.DeviceData{}
		for _, v := range src.DeviceList {
			deviceList = append(deviceList, &proto.DeviceData{
				DevId:      v.Devid,
				Id:         v.Id,
				ProductKey: v.ProductKey,
			})
		}
		pbObj.IotDeviceGroupReq.DeviceList = deviceList
	}

	return &pbObj
}

type DeviceGroup struct {
	Id            string `json:"id"`                      //群组编号
	Name          string `json:"name"`                    //群组名称
	RoomId        string `json:"roomId"`                  //房间id
	RoomName      string `json:"roomName"`                //房间名称
	HomeId        string `json:"homeId"`                  //家庭编号
	Time          int64  `json:"time"`                    //创建时间
	UserId        string `json:"userId"`                  //用户编号
	DevCount      int    `json:"devCount"`                //设备数量
	Pic           string `json:"pic"`                     //群组图片
	ProductKey    string `json:"productKey"`              //productKey
	ProductTypeId string `json:"productTypeId,omitempty"` //产品分类id
	//GroupId  string `json:"groupId"`  	//群组id
}

func DevGroupInfo_pb2db(src []*proto.IotDeviceGroupList, devKeyValue map[string]*proto.IotDeviceInfo, productKeyValue map[string]string) []DeviceGroupListDto {
	list := []DeviceGroupListDto{}
	for _, v := range src {
		list = append(list, DeviceGroupListDto{
			Id:         iotutil.ToString(v.Id),
			DevId:      v.DevId,
			DevName:    devKeyValue[v.DevId].DeviceName,
			ProductKey: devKeyValue[v.DevId].ProductKey,
			Pic:        productKeyValue[devKeyValue[v.DevId].ProductKey],
		})
	}
	return list
}

type GroupTslInfo struct {
	//DevId          string                 `protobuf:"bytes,103,opt,name=devId,proto3" json:"devId"`             //
	//HomeId         string                 `protobuf:"bytes,104,opt,name=homeId,proto3" json:"homeId"`           //
	//DeviceStatus   map[string]interface{} `json:"deviceStatus"`                                                 //
	Props []TslInfo `json:"tsl"` //
}

type GroupExecute struct {
	GroupId   string      `json:"groupId"`   //群组id
	DevIdList []string    `json:"devIdList"` //设备列表
	Dpid      string      `json:"dpid"`      //指令键
	Value     interface{} `json:"value"`     //指令值
}

type DictKeyVal struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LangTranslateEntitys struct {
	Id            string                   `json:"id"`
	SourceTable   string                   `json:"sourceTable,omitempty"`
	SourceRowId   string                   `json:"sourceRowId,omitempty"`
	Lang          string                   `json:"lang"`
	FieldName     string                   `json:"fieldName"`
	FieldType     int32                    `json:"fieldType"`
	FieldValue    string                   `json:"fieldValue"`
	PlatformType  int32                    `json:"platformType"` //平台类型 =1 云管平台 =2 开发平台
	TranslateList []BatchSaveTranslateItem `json:"translateList,omitempty"`
	Sort          int                      `json:"sort"`
}

type BatchSaveTranslateItem struct {
	Id         string `json:"id,omitempty"`
	Lang       string `json:"lang,omitempty"`
	FieldName  string `json:"fieldName,omitempty"`
	FieldType  int32  `json:"fieldType,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}

type ProductThingsModel struct {
	Services   []*ThingModelServices   `json:"actions"`
	Events     []*ThingModelEvents     `json:"events"`
	Properties []*ThingModelProperties `json:"attrs"`
}

type ThingModelEvents struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	EventType  string `json:"eventType"`
	OutputData string `json:"outputData"`
}

type ThingModelProperties struct {
	Name          string `json:"name"`
	Identifier    string `json:"identifier"` //标识符类型原来的pid
	DataType      string `json:"dataType"`
	DataSpecs     string `json:"dataSpecs"`
	DataSpecsList string `json:"dataSpecsList"`
}

type ThingModelServices struct {
	Name         string `json:"name"`
	Identifier   string `json:"identifier"`
	ServiceName  string `json:"serviceName"`
	InputParams  string `json:"inputParams"`
	OutputParams string `json:"outputParams"`
	CallType     int32  `json:"callType"`
}

// 家庭房间列表
type UcHomeRoomList struct {
	RoomId      string `json:"roomId"`
	Name        string `json:"name"`
	Sort        int32  `json:"sort"`
	Icon        string `json:"icon"`
	DeviceCount int32  `json:"deviceCount"`
}
