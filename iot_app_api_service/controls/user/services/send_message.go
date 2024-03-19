package services

import (
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
	"time"

	"go-micro.dev/v4/metadata"

	"github.com/gin-gonic/gin"
)

type MessageAppInfo struct {
	AppKey   string
	TenantId string
	Lang     string
}

func SetAppInfo(c *gin.Context) (res MessageAppInfo) {
	res = MessageAppInfo{}
	defer iotutil.PanicHandler()
	res = MessageAppInfo{
		TenantId: c.Request.Header.Get("tenantId"),
		AppKey:   c.Request.Header.Get("appKey"),
	}
	return res
}
func SetAppInfoByContext(ctx context.Context) (res MessageAppInfo) {
	res = MessageAppInfo{}
	defer iotutil.PanicHandler()
	md, _ := metadata.FromContext(ctx)
	tenantId, _ := md.Get("tenantId")
	appKey, _ := md.Get("appKey")
	lang, _ := md.Get("lang")
	res = MessageAppInfo{
		TenantId: tenantId,
		AppKey:   appKey,
		Lang:     lang,
	}
	return
}

// SendHomeAppMessage 推送消息
// homeId 推送给家庭
// userIds 推送给用户
// tplCode 消息模板，在云管平台消息模板中进行配置
// AddDevice
// subject 推送消息的主题，建议放到常量中
// params 模板对应的动态参数
func SendHomeAppMessage(appInfo MessageAppInfo, pushTo string, homeId int64, userIds []int64, tplCode string, subject string, ChildType int32, params map[string]string) {
	defer iotutil.PanicHandler(appInfo, pushTo, homeId, userIds, tplCode, subject, ChildType, params)
	pushMsg := &protosService.SendMessageRequest{
		TplCode:   tplCode,
		Params:    params,
		TimeUnix:  time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		HomeId:    homeId,
		UserId:    userIds,
		IsPublic:  false,
		PushTo:    pushTo,
		ChildType: ChildType,
		Subject:   subject,
		AppKey:    appInfo.AppKey,
		TenantId:  appInfo.TenantId,
		Lang:      "", //不指定语言则，则全语言推送
	}
	//发送消息  测试消息推送
	//TODO 修改为消息队列推送
	iotlogger.LogHelper.WithTag("method", "SendHomeAppMessage").Info("request" + iotutil.ToString(pushMsg))
	ret, err := rpc.MessageService.SendMessage(context.Background(), pushMsg)
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendHomeAppMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendHomeAppMessage").Error(ret.Message)
		return
	}
}

func SendDeviceAppMessage(childType int32, appKey, tenantId, devId, subject, tplCode string, homeId int64, userId int64, params map[string]string) {
	//发送消息  测试消息推送
	ret, err := rpc.MessageService.SendLaserMessage(context.Background(), &protosService.SendMessageRequest{
		TplCode:     tplCode,
		Params:      params,
		TimeUnix:    time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		SourceTable: model.TableNameTIotDeviceInfo,
		SourceRowId: devId,
		HomeId:      homeId,
		UserId:      []int64{userId},
		IsPublic:    false,
		PushTo:      "device",
		ChildType:   childType,
		Subject:     subject,
		Lang:        "", //不指定语言则，则全语言推送
		AppKey:      appKey,
		TenantId:    tenantId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if ret.Code == 200 {
		fmt.Println(ret.Message)
		return
	}
}

// SendUpdatePasswordMessage 用户消息（user)	修改密码-强制退出	14
func SendUpdatePasswordMessage(appInfo MessageAppInfo, userId int64, account string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendUpdatePasswordMessage start")
	//title := "修改密码"
	//templateC := "账号{.account}修改密码成功"
	var (
		pushTo    string = "home"
		childType int32  = 14
		subject   string = "密码修改"
		tplCode   string = iotconst.APP_MESSAGE_UPDATE_PASSWORD
	)
	SendHomeAppMessage(appInfo, pushTo, 0, []int64{userId}, tplCode, subject, childType, map[string]string{
		"account": account,
	})
}

// SendRemoveDeviceMessage 移除设备	8
// 移除设备	“用户名称”将家庭“家庭名称”中 “产品名称”移除了
func SendRemoveDeviceMessage(appInfo MessageAppInfo, data *protosService.UcHomeDetail, opUserId int64, homeId int64, deviceIds ...string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendRemoveDeviceMessage start")
	if data == nil {
		ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
			HomeId: homeId,
			UnloadSet: &protosService.UnLoadDataSet{
				UnLoadDevices: false, //不加载设备
				UnLoadRooms:   true,  //不加载房间
				UnLoadUsers:   false, //加载用户
			},
		})
		if err != nil {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return
		}
		if ret.Code != 200 {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
			return
		}
		data = ret.Data
	}
	var (
		pushTo     string  = "home"
		childType  int32   = 8
		subject    string  = "移除设备"
		tplCode    string  = iotconst.APP_MESSAGE_REMOVE_DEVICE
		userIds    []int64 = make([]int64, 0)
		homeName   string  = data.Data.Name
		opUserName string  = ""
	)
	for _, user := range data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		if userId == opUserId {
			opUserName = user.NickName
		}
		userIds = append(userIds, userId)
	}

	deviceMap := map[string]string{}
	for _, device := range data.DeviceList {
		deviceMap[device.Data.Did] = device.Data.DeviceName
	}
	for _, devId := range deviceIds {
		deviceName, ok := deviceMap[devId]
		if !ok {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error("未获取到设备名称")
			continue
		}
		SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
			"userName":   opUserName,
			"homeName":   HomeLanguage(appInfo.Lang, homeName),
			"deviceName": deviceName,
		})
	}
}

// SendRemoveHomeMessage 家庭被移除	10
// 家庭被移除	“用户名称”移除了家庭“家庭名称”
func SendRemoveHomeMessage(appInfo MessageAppInfo, retData *protosService.UcHomeDetailResponse, opUserId int64, homeId int64) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendRemoveHomeMessage start")
	//ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
	//	HomeId: homeId,
	//})
	//if err != nil {
	//	iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
	//	return
	//}
	//if ret.Code != 200 {
	//	iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
	//	return
	//}
	var (
		pushTo     string  = "home"
		childType  int32   = 10
		subject    string  = ""
		tplCode    string  = iotconst.APP_MESSAGE_REMOVE_HOME
		userIds    []int64 = make([]int64, 0)
		homeName   string  = retData.Data.Data.Name
		opUserName string  = ""
	)
	for _, user := range retData.Data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		if userId == opUserId {
			opUserName = user.NickName
		}
		userIds = append(userIds, userId)
	}
	subject = opUserName + "移除了家庭" + homeName
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"userName": opUserName,
		"homeName": HomeLanguage(appInfo.Lang, homeName),
	})
}

// SendJoinHomeMessage	12
// 新成员加入	“用户名称”邀请用户“用户名称”加入 “家庭名称”家庭
func SendJoinHomeMessage(appInfo MessageAppInfo, ret *protosService.UcHomeDetailResponse, homeId int64, fromUserId, toUserId int64) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendJoinHomeMessage start")

	var (
		pushTo      string  = "home"
		childType   int32   = 12
		subject     string  = "新成员加入"
		tplCode     string  = iotconst.APP_MESSAGE_JOIN_HOME
		inviteUser  string  = ""
		invitedUser string  = ""
		userIds     []int64 = make([]int64, 0)
		homeName    string  = ret.Data.Data.Name
	)
	for _, user := range ret.Data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		//如果是被邀请者，则不需要推送消息
		if userId == toUserId {
			invitedUser = user.NickName
			continue
		}
		//如果是邀请者，需要记录邀请人名称
		if userId == fromUserId {
			inviteUser = user.NickName
		}
		userIds = append(userIds, userId)
	}
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"inviteUser":  inviteUser,
		"invitedUser": invitedUser,
		"homeName":    HomeLanguage(appInfo.Lang, homeName),
	})
}

// SendQuitHomeMessage	15
// 成员退出家庭	"用户名称"退出了家庭"家庭名称"
func SendQuitHomeMessage(appInfo MessageAppInfo, ret *protosService.UcHomeDetailResponse, homeId int64, invitedUserId int64) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendQuitHomeMessage start")
	var (
		pushTo      string  = "home"
		childType   int32   = 15
		subject     string  = "成员退出家庭"
		tplCode     string  = iotconst.APP_MESSAGE_QUIT_HOME
		invitedUser string  = ""
		userIds     []int64 = make([]int64, 0)
		homeName    string  = ret.Data.Data.Name
	)
	for _, user := range ret.Data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		//给家庭所有者推送消息
		if user.Role == 1 {
			userIds = append(userIds, userId)
			continue
		}
		//退出家庭的成员，需要记录成员的名称，不需要推送消息
		if userId == invitedUserId {
			invitedUser = user.NickName
			continue
		}
	}
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"userName": invitedUser,
		"homeName": HomeLanguage(appInfo.Lang, homeName),
	})
}

// SendRemoveMembersMessage	13
// 成员移除	您被家庭管理员从家庭"家庭名称"中移除
func SendRemoveMembersMessage(appInfo MessageAppInfo, ret *protosService.UcHomeDetailResponse, homeId, invitedUserId int64) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendRemoveMembersMessage start")
	var (
		pushTo    string  = "home"
		childType int32   = 13
		subject   string  = "成员移除"
		tplCode   string  = iotconst.APP_MESSAGE_REMOVE_MEMBERS
		userIds   []int64 = make([]int64, 0)
		homeName  string  = ret.Data.Data.Name
	)
	for _, user := range ret.Data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		//被移除的成员，推送消息
		if userId == invitedUserId {
			userIds = append(userIds, userId)
		}
	}
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"homeName": HomeLanguage(appInfo.Lang, homeName),
	})
}

// SendAddSharedMessage 共享者发送邀请给接受者	23
// 共享者发送邀请给接受者	共享者“{{.userName}}”发送了设备共享邀请
func SendAddSharedMessage(data *protosService.UcHomeDetail, userId, opUserId int64, homeId int64, deviceId, appKey, tenantId string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendAddsharedMessage start")
	//ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
	//	HomeId: homeId,
	//})
	//if err != nil {
	//	iotlogger.LogHelper.WithTag("method", "SendAddsharedMessage").Error(err)
	//	return
	//}
	//if ret.Code != 200 {
	//	iotlogger.LogHelper.WithTag("method", "SendAddsharedMessage").Error(ret.Message)
	//	return
	//}
	var (
		childType int32 = 23
		//subject    string = ""
		tplCode    string = iotconst.APP_MESSAGE_ADD_SHARED
		opUserName string = ""
	)
	for _, user := range data.UserList {
		uId := iotutil.ToInt64(user.Uid)
		if uId == opUserId {
			opUserName = user.NickName
		}
	}

	deviceMap := map[string]string{}
	for _, device := range data.DeviceList {
		deviceMap[device.Data.Did] = device.Data.DeviceName
	}
	SendDeviceAppMessage(childType, appKey, tenantId, deviceId, opUserName+"用户邀请你一起操作"+deviceMap[deviceId]+"设备,请查看.", tplCode, homeId, userId, map[string]string{
		"userName":   opUserName,
		"deviceName": deviceMap[deviceId],
	})
}

// SendReceiveShareMessage 接受者已接受设备	24
// 接受者已接受设备	接受者“{{.userName}}”已接受设备共享邀请
func SendReceiveShareMessage(userId, belongUserId int64, homeId int64, deviceId, appKey, tenantId string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendReceiveShareMessage start")
	res, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: iotutil.ToInt64(userId),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendReceiveShareMessage").Error(err)
		return
	}
	if res.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendReceiveShareMessage").Error(res.Message)
		return
	}
	if len(res.Data) == 0 {
		iotlogger.LogHelper.WithTag("method", "SendReceiveShareMessage").Error("未找到用户信息")
		return
	}

	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendReceiveShareMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendReceiveShareMessage").Error(ret.Message)
		return
	}
	var (
		childType int32 = 24
		//subject    string = "接受共享"
		tplCode    string = iotconst.APP_MESSAGE_RECEIVE_SHARED
		opUserName string = res.Data[0].NickName
	)

	deviceMap := map[string]string{}
	for _, device := range ret.Data.DeviceList {
		deviceMap[device.Data.Did] = device.Data.DeviceName
	}
	SendDeviceAppMessage(childType, appKey, tenantId, deviceId, opUserName+"用户已接受你共享的"+deviceMap[deviceId]+"设备.", tplCode, homeId, belongUserId, map[string]string{
		"userName":   opUserName,
		"deviceName": deviceMap[deviceId],
	})
}

// SendCancelShareMessage 共享者取消共享给接受者	25
// 共享者取消共享给接受者	共享者“{{.userName}}”取消设备共享
func SendCancelShareMessage(userId, belongUserId int64, homeId int64, deviceId, appKey, tenantId string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendCancelShareMessage start")
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendCancelShareMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendCancelShareMessage").Error(ret.Message)
		return
	}
	var (
		childType int32 = 25
		//subject    string = "共享者取消共享给接受者"
		tplCode    string = iotconst.APP_MESSAGE_CANCEL_SHARED
		opUserName string = ""
	)
	for _, user := range ret.Data.UserList {
		uId := iotutil.ToInt64(user.Uid)
		if uId == belongUserId {
			opUserName = user.NickName
		}
	}

	deviceMap := map[string]string{}
	for _, device := range ret.Data.DeviceList {
		deviceMap[device.Data.Did] = device.Data.DeviceName
	}
	SendDeviceAppMessage(childType, appKey, tenantId, deviceId, opUserName+"用户已取消共享"+deviceMap[deviceId]+"设备给你.", tplCode, homeId, userId, map[string]string{
		"userName":   opUserName,
		"deviceName": deviceMap[deviceId],
	})
}

// SendCancelReceiveSharedMessage 接受者取消接受设备，拒绝接受设备	26
// 接受者取消接受设备，拒绝接受设备	接受者“{{.userName}}”取消接受设备共享
func SendCancelReceiveSharedMessage(userId, belongUserId int64, homeId int64, deviceId, appKey, tenantId string) {
	defer iotutil.PanicHandler()
	iotlogger.LogHelper.Infof("SendCancelReceiveSharedMessage start")
	res, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: iotutil.ToInt64(userId),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendCancelReceiveSharedMessage").Error(err)
		return
	}
	if res.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendCancelReceiveSharedMessage").Error(res.Message)
		return
	}
	if len(res.Data) == 0 {
		iotlogger.LogHelper.WithTag("method", "SendCancelReceiveSharedMessage").Error("未找到用户信息")
		return
	}

	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendCancelReceiveSharedMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendCancelReceiveSharedMessage").Error(ret.Message)
		return
	}
	var (
		childType int32 = 26
		//subject    string = "接受者取消接受设备"
		tplCode    string = iotconst.APP_MESSAGE_REFUSE_RECEIVE_SHARED
		opUserName string = res.Data[0].NickName
	)

	deviceMap := map[string]string{}
	for _, device := range ret.Data.DeviceList {
		deviceMap[device.Data.Did] = device.Data.DeviceName
	}
	SendDeviceAppMessage(childType, appKey, tenantId, deviceId, opUserName+"用户已拒绝接受"+deviceMap[deviceId]+"设备.", tplCode, homeId, belongUserId, map[string]string{
		"userName":   opUserName,
		"deviceName": deviceMap[deviceId],
	})
}

// SendAutoDisbandGroupMessage	19
// 解除群组(自动解除群组)	"群组名称"中的设备已经少于2个，已自动解散群组
func SendAutoDisbandGroupMessage(appInfo MessageAppInfo, data *protosService.UcHomeDetail, opUserId int64, homeId int64, groupName string) {
	defer iotutil.PanicHandler()
	data = loadHomeUserList(homeId)
	var (
		pushTo    string  = "home"
		childType int32   = 19
		subject   string  = "解除群组"
		tplCode   string  = iotconst.APP_MESSAGE_AUTO_DISBAND_GROUP_HOME
		userIds   []int64 = make([]int64, 0)
	)
	for _, user := range data.UserList {
		userId, err := iotutil.ToInt64AndErr(user.Uid)
		if err != nil {
			continue
		}
		userIds = append(userIds, userId)
	}
	if len(userIds) == 0 {
		iotlogger.LogHelper.WithTag("method", "SendAutoDisbandGroupMessage").Error("未找到推送用户，homeId:" + iotutil.ToString(homeId))
		return
	}
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"groupName": groupName,
	})
}

func loadHomeUserList(homeId int64) (data *protosService.UcHomeDetail) {
	if data == nil {
		ret, err := rpc.UcHomeUserService.HomeUserLists(context.Background(), &protosService.UcHomeUserFilter{HomeId: homeId})
		if err != nil {
			iotlogger.LogHelper.WithTag("method", "SendAutoDisbandGroupMessage").Error(err)
			return
		}
		if ret.Code != 200 {
			iotlogger.LogHelper.WithTag("method", "SendAutoDisbandGroupMessage").Error(ret.Message)
			return
		}
		userList := make([]*protosService.UserHome, 0)
		for _, u := range ret.Data {
			userList = append(userList, &protosService.UserHome{
				Uid:  iotutil.ToString(u.UserId),
				Role: u.RoleType,
			})
		}
		data = &protosService.UcHomeDetail{UserList: userList}
	}
	return data
}

// SendDisbandGroupMessage	18
// 解除群组	"用户名称"将"家庭名称"家庭中"群组名称"移除了
func SendDisbandGroupMessage(appInfo MessageAppInfo, opUserId, homeId int64, groupName string) {
	defer iotutil.PanicHandler()
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendDisbandGroupMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendDisbandGroupMessage").Error(ret.Message)
		return
	}

	var (
		pushTo    string  = "home"
		childType int32   = 18
		subject   string  = "解除群组"
		tplCode   string  = iotconst.APP_MESSAGE_DISBAND_GROUP_HOME
		userName  string  = ""
		userIds   []int64 = make([]int64, 0)
		homeName  string  = ret.Data.Data.Name
	)
	for _, user := range ret.Data.UserList {
		userId := iotutil.ToInt64(user.Uid)
		userIds = append(userIds, userId)
		//显示操作人员的名称
		if userId == opUserId {
			userName = user.NickName
			continue
		}
	}
	if len(userIds) == 0 {
		iotlogger.LogHelper.WithTag("method", "SendDisbandGroupMessage").Error("未找到推送用户，homeId:" + iotutil.ToString(homeId))
		return
	}
	SendHomeAppMessage(appInfo, pushTo, homeId, userIds, tplCode, subject, childType, map[string]string{
		"userName":  userName,
		"homeName":  HomeLanguage(appInfo.Lang, homeName),
		"groupName": groupName,
	})
}
