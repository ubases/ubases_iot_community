package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	services "cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"sort"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppShareDeviceService struct {
	Ctx context.Context
}

func (s AppShareDeviceService) SetContext(ctx context.Context) AppShareDeviceService {
	s.Ctx = ctx
	return s
}

// ShareDeviceList 共享设备列表
func (s AppShareDeviceService) ShareDeviceList(homeId string) ([]entitys.SharedDeviceListEntityDto, error) {
	ret, err := rpc.IotDeviceHomeService.HomeDevListExcludeVirtualDevices(context.Background(), &protosService.IotDeviceHomeHomeId{
		HomeId: iotutil.ToInt64(homeId),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
		return nil, errors.New("record not found")
	}

	devIdList := []string{}
	roomIdList := []int64{}

	productIds := []int64{}
	for _, v := range ret.DevList {
		productIds = append(productIds, v.ProductId)
		roomIdList = append(roomIdList, iotutil.ToInt64(v.RoomId))
	}
	productInfoList := &protosService.OpmProductResponse{}
	if len(productIds) > 0 {
		productInfoList, err = rpc.ProductService.ListsByProductIds(context.Background(), &protosService.ListsByProductIdsRequest{
			ProductIds: productIds,
		})
	}

	roomMap := map[string]string{}
	ucHomeRoomResponse := &protosService.UcHomeRoomResponse{}
	if len(roomIdList) > 0 {
		ucHomeRoomResponse, err = rpc.UcHomeRoomService.FindByIds(context.Background(), &protosService.UcHomeRoomFilter{
			RoomIds: roomIdList,
		})
		for _, room := range ucHomeRoomResponse.Data {
			roomMap[iotutil.ToString(room.Id)] = room.RoomName
		}
	}

	deviceMap := map[string]interface{}{}
	for _, device := range ret.DevList {
		deviceInfo := device
		var productPic, productKey string
		if productInfoList != nil && len(productInfoList.Data) > 0 {
			for _, productInfo := range productInfoList.Data {
				if device.ProductId == productInfo.Id {
					productPic = productInfo.ImageUrl
					productKey = productInfo.ProductKey
					break
				}
			}
		}

		sharedDeviceObj := entitys.SharedDeviceListEntityDto{
			DevId:    deviceInfo.Did,
			Model:    productKey,
			RoomId:   deviceInfo.RoomId,
			Name:     deviceInfo.DeviceName,
			Time:     deviceInfo.BindTime.AsTime().Unix(),
			Pic:      productPic,
			RoomName: roomMap[deviceInfo.RoomId],
			Type:     0,
		}

		deviceMap[sharedDeviceObj.DevId] = sharedDeviceObj

		devIdList = append(devIdList, deviceInfo.Did)
	}

	deviceSharedResult, err := rpc.IotDeviceSharedService.Lists(context.Background(), &protosService.IotDeviceSharedListRequest{
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.IotDeviceShared{
			HomeId:    iotutil.ToInt64(homeId),
			DeviceIds: devIdList,
		},
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
		return nil, errors.New("record not found")
	}

	sharedDeviceResult := make([]entitys.SharedDeviceListEntityDto, 0)
	unSharedDeviceResult := make([]entitys.SharedDeviceListEntityDto, 0)
	deviceResult := make([]entitys.SharedDeviceListEntityDto, 0)

	sharedDeviceData := deviceSharedResult.Data
	sharedDeviceDataCount := len(sharedDeviceData)
	sharedDeviceKeyValue := make(map[string]interface{})
	if sharedDeviceDataCount > 0 {
		for _, sharedDeviceObj := range sharedDeviceData {
			sharedDeviceKeyValue[sharedDeviceObj.DeviceId] = sharedDeviceObj
		}
	}
	for _, deviceItem := range deviceMap {
		var mqttServer string
		if len(config.Global.AppMQTT.Addrs) > 0 {
			mqttServer = config.Global.AppMQTT.Addrs[0]
		} else {
			mqttServer = "ws://120.77.64.252:8883/mqtt"
		}
		deviceObj := deviceItem.(entitys.SharedDeviceListEntityDto)
		deviceObj.Type = 2
		deviceObj.MqttServer = mqttServer
		if sharedDeviceDataCount > 0 {
			if sharedDeviceValue, ok := sharedDeviceKeyValue[iotutil.ToString(deviceObj.DevId)]; ok {
				deviceObj.Type = 1
				deviceObj.Time = sharedDeviceValue.(*protosService.IotDeviceShared).SharedTime.AsTime().Unix()
				sharedDeviceResult = append(sharedDeviceResult, deviceObj)
				continue
			}
		}
		unSharedDeviceResult = append(unSharedDeviceResult, deviceObj)
	}

	for _, deviceItem := range sharedDeviceResult {
		deviceResult = append(deviceResult, deviceItem)
	}
	for _, deviceItem := range unSharedDeviceResult {
		deviceResult = append(deviceResult, deviceItem)
	}

	//根据sort进行排序
	sort.Slice(deviceResult, func(i, j int) bool {
		return deviceResult[i].Time > deviceResult[j].Time
	})
	return deviceResult, nil
}

// ShareUserList 共享用户列表
func (s AppShareDeviceService) ShareUserList(devId, homeId string) ([]map[string]interface{}, error) {
	deviceSharedResult, err := rpc.IotDeviceSharedService.Lists(context.Background(), &protosService.IotDeviceSharedListRequest{
		Query: &protosService.IotDeviceShared{
			HomeId:   iotutil.ToInt64(homeId),
			DeviceId: devId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if deviceSharedResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(deviceSharedResult.Message)
		return nil, errors.New("record not found")
	}

	deviceShareReceiveResult, err := rpc.IotDeviceShareReceiveService.Lists(context.Background(), &protosService.IotDeviceShareReceiveListRequest{
		Query: &protosService.IotDeviceShareReceive{
			HomeId:   iotutil.ToInt64(homeId),
			DeviceId: devId,
			//IsAgree:  1,   //未同意
		},
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if deviceShareReceiveResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(deviceShareReceiveResult.Message)
		return nil, errors.New("record not found")
	}

	sharedUserData := []entitys.SharedUserDto{}
	for _, deviceShared := range deviceSharedResult.Data {
		sharedUserData = append(sharedUserData, entitys.SharedUserDto{Time: deviceShared.SharedTime.AsTime().Unix(), UserId: iotutil.ToString(deviceShared.UserId),
			UserNickname: deviceShared.UserName, Type: 1, Photo: deviceShared.Photo, Id: iotutil.ToString(deviceShared.Id)})
	}

	for _, deviceShareReceive := range deviceShareReceiveResult.Data {
		var isAgree int32
		if deviceShareReceive.IsAgree == 1 {
			isAgree = 2
		} else if deviceShareReceive.IsAgree == 3 {
			isAgree = 3
		} else {
			continue //过滤2-已同意的数据
		}
		sharedUserData = append(sharedUserData, entitys.SharedUserDto{Time: deviceShareReceive.SharedTime.AsTime().Unix(), UserId: iotutil.ToString(deviceShareReceive.UserId),
			UserNickname: deviceShareReceive.UserName, Type: isAgree, Photo: deviceShareReceive.Photo, Id: iotutil.ToString(deviceShareReceive.Id)})
	}

	dateList := map[string][]entitys.SharedUserDto{}
	for _, row := range sharedUserData {
		createAt := time.Unix(row.Time, 0)

		date := iotutil.TimeFormatNew(createAt)
		list, ok := dateList[date]
		if !ok {
			list = []entitys.SharedUserDto{}
		}
		list = append(list, row)
		dateList[date] = list
	}

	resultArr := []map[string]interface{}{}
	for key, value := range dateList {
		resultArr = append(resultArr, map[string]interface{}{
			"date": key,
			"list": value,
		})
	}
	//根据sort进行排序
	sort.Slice(resultArr, func(i, j int) bool {
		return resultArr[i]["date"].(string) > resultArr[j]["date"].(string)
	})

	return resultArr, nil
}

// CancelShare 取消共享
func (s AppShareDeviceService) CancelShare(req entitys.CancelShare, userId, appKey, tenantId string) error {
	var belongUserId, homeId int64
	var deviceId string

	cancelType := req.Type
	if cancelType != 1 && cancelType != 2 {
		return errors.New("cancelType is wrong")
	}
	if cancelType == 1 {
		result, err := rpc.IotDeviceSharedService.FindById(context.Background(), &protosService.IotDeviceSharedFilter{
			Id: iotutil.ToInt64(req.Id),
		})
		if err != nil {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return err
		}
		if result.Code != 200 {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(result.Message)
			return errors.New("record not found")
		}
		deviceShare := result.Data[0]
		belongUserId = deviceShare.UserId
		homeId = deviceShare.HomeId
		deviceId = deviceShare.DeviceId

		_, err = rpc.IotDeviceSharedService.Delete(context.Background(), &protosService.IotDeviceShared{
			Id: iotutil.ToInt64(req.Id),
		})
		if err != nil {
			//iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return err
		}

		_, err = rpc.IotDeviceShareReceiveService.Delete(context.Background(), &protosService.IotDeviceShareReceive{
			UserId:   belongUserId,
			DeviceId: deviceId,
		})
		if err != nil {
			//iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return err
		}

		// 删除家庭详情缓存
		services.ClearHomeCached(belongUserId, false)
	} else {
		result, err := rpc.IotDeviceShareReceiveService.FindById(context.Background(), &protosService.IotDeviceShareReceiveFilter{
			Id: iotutil.ToInt64(req.Id),
		})
		if err != nil {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return err
		}
		if result.Code != 200 {
			iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(result.Message)
			return errors.New("record not found")
		}
		deviceShareReceive := result.Data[0]
		belongUserId = deviceShareReceive.UserId
		homeId = deviceShareReceive.HomeId
		deviceId = deviceShareReceive.DeviceId

		_, err = rpc.IotDeviceShareReceiveService.Delete(context.Background(), &protosService.IotDeviceShareReceive{
			Id: iotutil.ToInt64(req.Id),
		})
		if err != nil {
			//iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
			return err
		}
	}

	go services.SendCancelShareMessage(belongUserId, iotutil.ToInt64(userId), homeId, deviceId, appKey, tenantId)
	return nil
}

// AddShared 添加共享设备
func (s AppShareDeviceService) AddShared(belongUserId, appKey, tenantId string, regionServerId int64, req entitys.Addshared) (int, string) {
	var phone, email string

	if req.Type == 1 {
		if iotutil.CheckAllPhone("", req.Account) == false {
			return -1, "手机号码不合法"
		}
		phone = req.Account
	} else if req.Type == 2 {
		if iotutil.VerifyEmailFormat(req.Account) == false {
			return -1, "电子邮箱不合法"
		}
		email = req.Account
	} else {
		return -1, "用户类型有误"
	}

	//查询设备信息（如前端不传productKey，则需要从设备信息缓存中获取）

	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: iotutil.ToInt64(req.HomeId),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(err)
		return -1, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(ret.Message)
		return -1, ret.Message
	}
	data := ret.Data
	var currentUserRole int32
	for _, v := range data.UserList {
		if v.Uid == belongUserId {
			currentUserRole = v.Role
			break
		}
	}
	if currentUserRole == 3 {
		return -1, "家庭成员不能共享设备"
	}

	userInfo, err := rpc.TUcUserService.Lists(s.Ctx, &protosService.UcUserListRequest{
		Query: &protosService.UcUser{
			Id:       iotutil.ToInt64(belongUserId),
			Phone:    phone,
			Status:   1,
			Email:    email,
			AppKey:   appKey,
			TenantId: tenantId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("QueryQuestionTypeList error,%s", err.Error())
		return -1, err.Error()
	}
	if userInfo.Code == 200 && userInfo.Data != nil && len(userInfo.Data) > 0 {
		return ioterrs.ERROR_NOT_SHARE_YOURSELF.Code, ioterrs.ERROR_NOT_SHARE_YOURSELF.Msg
	}

	belongUserResult, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: iotutil.ToInt64(belongUserId),
		//AppKey:   req.AppKey,
		//TenantId: req.TenantId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return -1, err.Error()
	}
	if belongUserResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(belongUserResult.Message)
		return ioterrs.ERROR_USER_IS_NOT_EXIST.Code, ioterrs.ERROR_USER_IS_NOT_EXIST.Msg
	}
	userResult, err := rpc.TUcUserService.Find(context.Background(), &protosService.UcUserFilter{
		Phone:          phone,
		Password:       "",
		Email:          email,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return -1, err.Error()
	}
	if userResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(userResult.Message)
		return ioterrs.ERROR_USER_IS_NOT_EXIST.Code, ioterrs.ERROR_USER_IS_NOT_EXIST.Msg
	}
	userId := userResult.Data[0].Id

	deviceShareReceiveData, err := rpc.IotDeviceShareReceiveService.Find(context.Background(), &protosService.IotDeviceShareReceiveFilter{
		UserId:   userId,
		DeviceId: req.DevId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return -1, err.Error()
	}
	//if deviceShareReceiveData.Code != 200 {
	//	iotlogger.LogHelper.WithTag("method", "Addshared").Error(deviceShareReceiveData.Message)
	//	return -1,"record not found"
	//}
	if deviceShareReceiveData != nil && len(deviceShareReceiveData.Data) > 0 && deviceShareReceiveData.Data[0].IsAgree == 1 {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return ioterrs.ERROR_INVITATION_SENT.Code, ioterrs.ERROR_INVITATION_SENT.Msg
	}

	if deviceShareReceiveData != nil && len(deviceShareReceiveData.Data) > 0 && deviceShareReceiveData.Data[0].IsAgree == 2 {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return ioterrs.ERROR_DEVICE_ALREADY_SHARED.Code, ioterrs.ERROR_DEVICE_ALREADY_SHARED.Msg
	}

	productResult, err := rpc.ProductService.Find(context.Background(), &protosService.OpmProductFilter{
		ProductKey: req.ProductKey,
		TenantId:   tenantId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(err)
		return -1, err.Error()
	}
	if productResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "Addshared").Error(productResult.Message)
		return -1, "record not found"
	}

	smsCode := iotutil.Getcode()
	//600  10分钟
	//设置过期时间,有效期7天，60*60*24*7 = 604800
	res := iotredis.GetClient().Set(context.Background(), cached.APP+"_"+appKey+"_"+belongUserId+"_"+req.DevId, smsCode, 604800*time.Second) //有效期7天
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendSms,缓存smsCodeInt失败:%s", res.Err().Error())
		return -1, res.Err().Error()
	}
	homeId, _ := iotutil.ToInt64AndErr(req.HomeId)
	roomId, _ := iotutil.ToInt64AndErr(req.RoomId)
	_, err = rpc.IotDeviceShareReceiveService.Create(context.Background(), &protosService.IotDeviceShareReceive{
		Id:             iotutil.GetNextSeqInt64(),
		CustomName:     req.DevName,
		UserId:         iotutil.ToInt64(userId),
		UserName:       userResult.Data[0].NickName,
		Phone:          phone,
		Email:          email,
		DeviceId:       req.DevId,
		BelongUserId:   iotutil.ToInt64(belongUserId),
		BelongUserName: belongUserResult.Data[0].NickName,
		HomeId:         homeId,
		RoomId:         roomId,
		ProductKey:     req.ProductKey,
		ProductId:      productResult.Data[0].Id,
		ProductPic:     productResult.Data[0].ImageUrl,
		Photo:          userResult.Data[0].Photo,
		IsAgree:        1, //未同意
		SharedTime:     timestamppb.Now(),
		CreatedAt:      timestamppb.Now(),
		UpdatedAt:      timestamppb.Now(),
	})
	if err != nil {
		return -1, err.Error()
	}

	go services.SendAddSharedMessage(data, userId, iotutil.ToInt64(belongUserId), iotutil.ToInt64(req.HomeId), req.DevId, appKey, tenantId)
	return 0, ""
}

// ReceiveSharedList 接收共享列表
func (s AppShareDeviceService) ReceiveSharedList(userId int64) ([]entitys.ReceiveSharedDto, error) {
	deviceShareReceiveResult, err := rpc.IotDeviceShareReceiveService.Lists(context.Background(), &protosService.IotDeviceShareReceiveListRequest{
		Query: &protosService.IotDeviceShareReceive{
			UserId: userId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if deviceShareReceiveResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(deviceShareReceiveResult.Message)
		return nil, errors.New("record not found")
	}

	receiveSharedDtoList := []entitys.ReceiveSharedDto{}
	for _, info := range deviceShareReceiveResult.Data {
		receiveSharedDtoList = append(receiveSharedDtoList, entitys.ReceiveSharedDto{
			Id:             iotutil.ToString(info.Id),
			DevName:        info.CustomName,
			ProductKey:     info.ProductKey,
			DeviceImg:      info.ProductPic,
			BelongUserName: info.BelongUserName,
			Status:         info.IsAgree,
		})
	}
	return receiveSharedDtoList, nil
}

// ReceiveShare 接收共享
func (s AppShareDeviceService) ReceiveShare(userId, id int64, appKey, tenantId string) error {
	result, err := rpc.IotDeviceShareReceiveService.FindById(context.Background(), &protosService.IotDeviceShareReceiveFilter{
		Id: id,
		//IsAgree: 2, //已同意
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return err
	}
	if result.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(result.Message)
		return errors.New("record not found")
	}

	deviceShareReceive := result.Data[0]

	smsCodeValue := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+iotutil.ToString(deviceShareReceive.BelongUserId)+"_"+deviceShareReceive.DeviceId)
	if deviceShareReceive.IsAgree == 1 && smsCodeValue.Val() == "" {
		s.expirationSetting(id) //设置共享邀请失效
		iotlogger.LogHelper.Error("设备共享邀请已失效")
		return errors.New("设备共享邀请已失效")
	}

	if deviceShareReceive.IsAgree == 2 {
		return errors.New("用户已同意共享,无需再同意")
	}

	if deviceShareReceive.IsAgree == 3 {
		return errors.New("设备共享邀请已失效")
	}

	deviceShareReceiveResult, err := rpc.IotDeviceShareReceiveService.Update(context.Background(), &protosService.IotDeviceShareReceive{
		Id:      id,
		IsAgree: 2, //已同意
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return err
	}
	if deviceShareReceiveResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(deviceShareReceiveResult.Message)
		return errors.New("record not found")
	}

	_, err = rpc.IotDeviceSharedService.Create(context.Background(), &protosService.IotDeviceShared{
		Id:             iotutil.GetNextSeqInt64(),
		CustomName:     deviceShareReceive.CustomName,
		UserId:         deviceShareReceive.UserId,
		UserName:       deviceShareReceive.UserName,
		Phone:          deviceShareReceive.Phone,
		Email:          deviceShareReceive.Email,
		BelongUserId:   deviceShareReceive.BelongUserId,
		BelongUserName: deviceShareReceive.BelongUserName,
		DeviceId:       deviceShareReceive.DeviceId,
		HomeId:         deviceShareReceive.HomeId,
		Photo:          deviceShareReceive.Photo,
		ProductKey:     deviceShareReceive.ProductKey,
		ProductId:      deviceShareReceive.ProductId,
		ProductPic:     deviceShareReceive.ProductPic,
		Sid:            0,
		SharedTime:     timestamppb.Now(),
		CreatedAt:      timestamppb.Now(),
		UpdatedAt:      timestamppb.Now(),
	})
	if err != nil {
		//iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return err
	}

	// 删除家庭详情缓存
	services.ClearHomeCached(userId, false)

	go services.SendReceiveShareMessage(userId, deviceShareReceive.BelongUserId, iotutil.ToInt64(deviceShareReceive.HomeId), deviceShareReceive.DeviceId, appKey, tenantId)
	return nil
}

func (s AppShareDeviceService) expirationSetting(id int64) error {
	//todo  还要把设备共享表中对应的数据删掉
	deviceShareReceiveResult, err := rpc.IotDeviceShareReceiveService.Update(context.Background(), &protosService.IotDeviceShareReceive{
		Id:      id,
		IsAgree: 3, //已过期
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "expirationSetting").Error(err)
		return err
	}
	if deviceShareReceiveResult.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "expirationSetting").Error(deviceShareReceiveResult.Message)
		return errors.New("record not found")
	}
	return nil
}

// CancelReceiveShared 取消接受共享
func (s AppShareDeviceService) CancelReceiveShared(req entitys.CancelReceiveShared, userId, appKey, tenantId string) error {
	result, err := rpc.IotDeviceShareReceiveService.FindById(context.Background(), &protosService.IotDeviceShareReceiveFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return err
	}
	if result.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(result.Message)
		return errors.New("record not found")
	}

	isAgree := result.Data[0].IsAgree
	_, err = rpc.IotDeviceShareReceiveService.Delete(context.Background(), &protosService.IotDeviceShareReceive{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if isAgree == 2 {
		_, err := rpc.IotDeviceSharedService.Delete(context.Background(), &protosService.IotDeviceShared{
			UserId:   result.Data[0].UserId,
			DeviceId: result.Data[0].DeviceId,
		})
		if err != nil {
			return err
		}

		// 删除家庭详情缓存
		services.ClearHomeCached(result.Data[0].UserId, false)
	}

	go services.SendCancelReceiveSharedMessage(result.Data[0].UserId, iotutil.ToInt64(result.Data[0].BelongUserId),
		iotutil.ToInt64(result.Data[0].HomeId), result.Data[0].DeviceId, appKey, tenantId)
	return nil
}
