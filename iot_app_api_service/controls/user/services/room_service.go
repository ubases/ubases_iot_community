package services

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"

	"go-micro.dev/v4/metadata"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppRoomService struct {
	Ctx context.Context
}

func (s AppRoomService) SetContext(ctx context.Context) AppRoomService {
	s.Ctx = ctx
	return s
}

func (s AppRoomService) AddRoom(req entitys.UcHomeRoomEntitys, userId int64) error {
	homeId, err := iotutil.ToInt64AndErr(req.HomeId)
	if err != nil {
		return errors.New("家庭Id不能为空")
	}
	req.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	saveObj := entitys.UcHomeRoom_e2pb(&req)
	saveObj.CreatedAt = timestamppb.Now()
	ctx := context.Background()
	res, err := rpc.UcHomeRoomService.Create(ctx, saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	// 删除家庭详情缓存
	homeSvc := AppHomeService{Ctx: context.Background()}
	homeSvc.clearHomeDetailsCached(homeId, ReadHomeRoomListsCachedKey(homeId))
	return err
}

func (s AppRoomService) Details(homeId, roomId string) (map[string]interface{}, error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	roomObj := make(map[string]interface{})

	deviceHomeList, err := rpc.IotDeviceHomeService.UserDevList(context.Background(), &protosService.IotDeviceHomeHomeId{HomeId: iotutil.ToInt64(homeId)})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间详情失败，原因:%s", err.Error())
		return nil, err
	}
	if deviceHomeList.Code != 200 {
		return nil, errors.New(deviceHomeList.Message)
	}

	roomRes, err := rpc.UcHomeRoomService.Lists(context.Background(), &protosService.UcHomeRoomListRequest{
		Query: &protosService.UcHomeRoom{
			HomeId: iotutil.ToInt64(homeId),
		},
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取房间列表失败，原因:%s", err.Error())
		return nil, err
	}
	if roomRes.Code != 200 {
		return nil, errors.New(roomRes.Message)
	}
	roomMap := map[int64]*protosService.UcHomeRoom{}
	defaultRooms := GetDefaultRooms(lang, tenantId, appKey)

	//对选择的房间名称进行翻译
	for _, room := range roomRes.Data {
		if room.RoomTemplateId != 0 {
			if dfVal, ok := defaultRooms[iotutil.ToString(room.RoomTemplateId)]; ok {
				room.RoomName = dfVal
			}
		}
		roomMap[room.Id] = room
	}

	deviceSlice := []entitys.HomeDevice{}
	otherRoomDeviceSlice := []entitys.HomeDevice{}
	productCached := controls.ProductCachedData{}
	for _, v := range deviceHomeList.DevList {
		devInfo := entitys.IotDeviceInfo_pb2db(v)
		productInfo, err := productCached.GetProduct(iotutil.ToString(devInfo.ProductKey))
		if err == nil && productInfo != nil {
			devInfo.ProductKey = productInfo.ProductKey
			devInfo.ProductPic = productInfo.ImageUrl
		}

		devEntitys := entitys.HomeDevice{
			Did:         devInfo.Did,
			ProductId:   devInfo.ProductId,
			DeviceName:  devInfo.DeviceName,
			ProductPic:  devInfo.ProductPic,
			RoomName:    devInfo.RoomName,
			RoomId:      devInfo.RoomId,
			ProductKey:  devInfo.ProductKey,
			SecretKey:   devInfo.SecretKey,
			MqttServer:  devInfo.MqttServer,
			OnlineState: devInfo.OnlineStatus,
			DevSwitch:   0,
			DevType:     1,
			Sort:        iotutil.ToInt32(v.DevSort),
		}

		if roomId == devInfo.RoomId {
			deviceSlice = append(deviceSlice, devEntitys)
		} else {
			otherRoomDeviceSlice = append(otherRoomDeviceSlice, devEntitys)
		}
	}
	if val, ok := roomMap[iotutil.ToInt64(roomId)]; ok {
		roomObj["name"] = val.RoomName
		roomObj["icon"] = val.IconUrl
		roomObj["sort"] = val.Sort
	}
	roomObj["roomId"] = roomId
	roomObj["deviceList"] = deviceSlice
	roomObj["otherRoomDeviceSlice"] = otherRoomDeviceSlice
	return roomObj, nil
}

func (s AppRoomService) Delete(req entitys.UcHomeRoomFilter, userId int64) error {
	//先查询是为了获取HomeId来清除家庭详情缓存
	findResp, err := rpc.UcHomeRoomService.FindById(s.Ctx, &protosService.UcHomeRoomFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if findResp.Code != 200 {
		return errors.New(findResp.Message)
	}
	rep, err := rpc.UcHomeRoomService.Delete(s.Ctx, &protosService.UcHomeRoom{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}

	homeId := findResp.Data[0].HomeId
	homeSvc := AppHomeService{Ctx: context.Background()}
	homeSvc.clearHomeDetailsCached(homeId, ReadHomeRoomListsCachedKey(homeId))
	return nil
}

func (s AppRoomService) Update(req entitys.UcHomeRoomEntitys, userId int64) error {
	findResp, err := rpc.UcHomeRoomService.FindById(s.Ctx, &protosService.UcHomeRoomFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if findResp.Code != 200 {
		return errors.New(findResp.Message)
	}
	req.UpdatedAt = time.Now()
	res, err := rpc.UcHomeRoomService.Update(s.Ctx, entitys.UcHomeRoom_e2pb(&req))
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	homeId := findResp.Data[0].HomeId
	homeSvc := AppHomeService{Ctx: s.Ctx}
	homeSvc.clearHomeDetailsCached(homeId, ReadHomeRoomListsCachedKey(homeId))
	return nil
}

func (s AppRoomService) SetSort(req entitys.SetRoomSort) error {
	homeId := iotutil.ToInt64(req.HomeId)
	paramList := []*protosService.UcHomeSortParam{}
	for _, v := range req.RoomParamlist {
		paramList = append(paramList, &protosService.UcHomeSortParam{
			Sort:   iotutil.ToInt32(v.Sort),
			RoomId: iotutil.ToInt64(v.RoomId),
		})
	}
	_, err := rpc.UcHomeRoomService.SetSort(s.Ctx, &protosService.UcHomeRoomSortRequest{
		HomeId: homeId,
		Data:   paramList,
	})
	if err != nil {
		return err
	}

	homeSvc := AppHomeService{Ctx: s.Ctx}
	homeSvc.clearHomeDetailsCached(homeId, ReadHomeRoomListsCachedKey(homeId))
	return nil
}

func (s AppRoomService) SetDevSort(req entitys.SetRoomSort) error {
	homeId := iotutil.ToInt64(req.HomeId)
	paramList := []*protosService.DevSortParam{}
	for _, v := range req.RoomParamlist {
		paramList = append(paramList, &protosService.DevSortParam{
			Sort:  iotutil.ToInt32(v.Sort),
			DevId: v.DevId,
		})
	}
	_, err := rpc.IotDeviceHomeService.SetDevSort(s.Ctx, &protosService.SetDevSortRequest{
		HomeId: homeId,
		Data:   paramList,
	})
	if err != nil {
		return err
	}

	homeSvc := AppHomeService{Ctx: context.Background()}
	homeSvc.clearHomeDetailsCached(homeId, ReadHomeRoomListsCachedKey(homeId))
	return nil
}
