package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls/intelligence/entitys"
	services2 "cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SceneIntelligenceconService struct {
	Ctx context.Context
}

func (s SceneIntelligenceconService) SetContext(ctx context.Context) SceneIntelligenceconService {
	s.Ctx = ctx
	return s
}

// 创建/编辑智能场景
func (s SceneIntelligenceconService) SaveIntelligence(req entitys.SceneIntelligenceForm) (int, error) {
	saveObj, err := entitys.Intelligence_e2pb(&req)
	if err != nil {
		return 0, err
	}
	//查询家庭信息
	homeRet, err := rpc.UcHomeService.FindById(s.Ctx, &protosService.UcHomeFilter{
		Id: req.HomeId,
	})
	if err != nil {
		return 0, err
	}
	if homeRet.Code != 200 {
		return 0, errors.New(homeRet.Message)
	}
	//如果是天气，将用户的当前城市报错给任务
	conEffStatus := make([]bool, 0)
	taskEffStatus := make([]bool, 0)

	conHasDev := true
	taskHasDev := true
	for i, condition := range saveObj.SceneIntelligenceCondition {
		if iotconst.ConditionType(condition.ConditionType) == iotconst.CONDITION_TYPE_WEATHER {
			saveObj.SceneIntelligenceCondition[i].WeatherCity = homeRet.Data[0].City
			saveObj.SceneIntelligenceCondition[i].WeatherArea = homeRet.Data[0].District
			saveObj.SceneIntelligenceCondition[i].WeatherCountry = homeRet.Data[0].Country
		} else if iotconst.ConditionType(condition.ConditionType) == iotconst.CONDITION_TYPE_SATACHANGE {
			//检查是否有效
			conHasDev = true
			deviceInfos := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+condition.DeviceDid, iotconst.FIELD_ONLINE)
			if deviceInfos.Err() == nil {
				list := deviceInfos.Val()
				onlineStatus := ""
				if len(list) > 0 {
					onlineStatus = iotutil.ToString(list[0])
				}
				if onlineStatus == "" {
					//设备已移除
					conEffStatus = append(conEffStatus, false)
					continue
				}
				conEffStatus = append(conEffStatus, true)
			}
		}
	}
	for _, task := range saveObj.SceneIntelligenceTask {
		if iotconst.TaskType(task.TaskType) == iotconst.TASK_TYPE_DEVICE {
			//检查是否有效
			taskHasDev = true
			deviceInfos := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+task.ObjectId, iotconst.FIELD_ONLINE)
			if deviceInfos.Err() == nil {
				list := deviceInfos.Val()
				onlineStatus := ""
				if len(list) > 0 {
					onlineStatus = iotutil.ToString(list[0])
				}
				if onlineStatus == "" {
					//设备已移除
					taskEffStatus = append(taskEffStatus, false)
					continue
				}
				taskEffStatus = append(taskEffStatus, true)
			}
		}
	}

	if taskHasDev || conHasDev {
		taskEff := taskHasDev && iotutil.BoolArrayIsVal(taskEffStatus, true)
		connEff := conHasDev && iotutil.BoolArrayIsVal(conEffStatus, true)
		if !taskEff && !connEff {
			return -2, errors.New("无有效设备")
		}
	}

	res, err := rpc.SceneIntelligenceService.Create(s.Ctx, saveObj)
	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Message)
	}
	s.clearIntelligenceListCachedByHomdId(req.HomeId)
	return 0, err
}

// 开关智能场景
func (s SceneIntelligenceconService) UpdateIntelligenceStatus(req entitys.SceneIntelligenceForm) error {
	userId, _ := metadata.Get(s.Ctx, "userId")
	userIdI := iotutil.ToInt64(userId)
	if userIdI == 0 {
		return errors.New("用户Id获取失败")
	}
	res, err := rpc.SceneIntelligenceService.UpdateFields(s.Ctx, &protosService.SceneIntelligenceUpdateFieldsRequest{
		Fields: []string{"status", "updated_at", "updated_by"},
		Data: &protosService.SceneIntelligence{
			Id:        req.Id,
			Status:    iotutil.IfInt32(req.Status, 1, 2),
			UpdatedAt: timestamppb.Now(),
			UpdatedBy: userIdI,
		},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	s.clearIntelligenceListCachedById(req.Id)
	return err
}

// 一键执行和自动执行顺序调整
func (s SceneIntelligenceconService) UpdateIntelligenceSortNo(req entitys.SceneIntelligenceForm) error {
	userId, _ := metadata.Get(s.Ctx, "userId")
	userIdI := iotutil.ToInt64(userId)
	if userIdI == 0 {
		return errors.New("用户Id获取失败")
	}
	res, err := rpc.SceneIntelligenceService.UpdateFields(s.Ctx, &protosService.SceneIntelligenceUpdateFieldsRequest{
		Fields: []string{"sort_no", "updated_at", "updated_by"},
		Data: &protosService.SceneIntelligence{
			Id:        req.Id,
			SortNo:    req.SortNo,
			UpdatedAt: timestamppb.Now(),
			UpdatedBy: userIdI,
		},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	s.clearIntelligenceListCachedById(req.Id)
	return err
}

// 删除智能场景
func (s SceneIntelligenceconService) DeleteIntelligence(id int64) error {

	res, err := rpc.SceneIntelligenceService.DeleteById(s.Ctx, &protosService.SceneIntelligence{
		Id: id,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	s.clearIntelligenceListCachedById(id)
	return err
}

// 查询智能场景列表
func (s SceneIntelligenceconService) GetIntelligenceList(lang, tenantId string, req entitys.SceneIntelligenceQueryForm) (rets []*entitys.SceneIntelligenceVo, total int, err error) {
	//处理缓存
	respData := make([]*entitys.SceneIntelligenceVo, 0)
	err = cached.RedisStore.Get(persist.GetRedisKey(iotconst.APP_HOME_INTELLIGENCE_DATA, iotutil.ToString(req.HomeId)), &respData)
	if err == nil {
		return respData, 0, err
	}

	ret, err := rpc.SceneIntelligenceService.Lists(s.Ctx, &protosService.SceneIntelligenceListRequest{
		Page:     int64(req.Page),
		PageSize: int64(req.Limit),
		Query: &protosService.SceneIntelligence{UserId: req.UserId, HomeId: req.HomeId, Type: req.Type,
			IntelligenceId: req.IntelligenceId, ProductKey: req.ProductKey, DeviceId: req.DeviceId, EnableDisplay: req.EnableDisplay},
	})
	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Message)
	}
	var (
		dataList = make([]*entitys.SceneIntelligenceVo, 0)
	)

	//TODO 如果需要对物模型进行翻译,启用下列代码
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	for _, datum := range ret.Data {
		entity := entitys.Intelligence_pb2eVo(lang, tenantId, langMap, datum)
		dataList = append(dataList, entity)
	}
	//mapstructure.WeakDecode(ret.Data, &dataList)

	//返回结果t
	var (
		limit    = req.Limit
		offset   = limit * (req.Page - 1)
		indexEnd = limit + offset
	)
	total = len(dataList)
	//组装返回结果
	hasPage := req.Page != 0 || req.Limit != 0
	if hasPage {
		if offset > total {
			return nil, total, err
		} else {
			if (offset + limit) > total {
				indexEnd = total
			}
			dataList, total, err = dataList[offset:indexEnd], total, err
		}
	}
	if err != nil {
		cerr := cached.RedisStore.Set(persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(req.HomeId)), &respData, 600*time.Second)
		if cerr != nil {
			iotlogger.LogHelper.Error("缓存场景数据失败，error：%s", cerr.Error())
		}
	}
	return dataList, total, err
}

// 查询智能场景详情
func (s SceneIntelligenceconService) GetIntelligenceDetail(req entitys.SceneIntelligenceQueryForm) (res *entitys.OldSceneIntelligenceForm, err error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
	)
	//查询智能场景详情
	ret, err := rpc.SceneIntelligenceService.FindById(s.Ctx, &protosService.SceneIntelligenceFilter{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	//查询智能条件列表
	conditionRet, err := rpc.SceneIntelligenceConditionService.Lists(s.Ctx, &protosService.SceneIntelligenceConditionListRequest{
		Page:     1,
		PageSize: 20,
		Query: &protosService.SceneIntelligenceCondition{
			IntelligenceId: req.Id,
		},
	})

	if err != nil {
		return nil, err
	}
	if conditionRet != nil && conditionRet.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	//查询智能任务列表
	taskRet, err := rpc.SceneIntelligenceTaskService.Lists(s.Ctx, &protosService.SceneIntelligenceTaskListRequest{
		Page:     1,
		PageSize: 20,
		Query: &protosService.SceneIntelligenceTask{
			IntelligenceId: req.Id,
		},
	})

	if err != nil {
		return nil, err
	}
	if taskRet != nil && taskRet.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	//绑定设备所在房间名称
	homeId := ret.Data[0].HomeId
	userServies := services.AppHomeService{Ctx: s.Ctx}
	roomMap, _ := userServies.GetRoomList(iotutil.ToInt64(homeId))

	//用户设备列表
	deviceHomeList, err := rpc.IotDeviceHomeService.UserDevList(s.Ctx, &protosService.IotDeviceHomeHomeId{HomeId: homeId})
	if err != nil {
		return nil, err
	}
	if deviceHomeList.Code != 200 {
		return nil, errors.New(deviceHomeList.Message)
	}

	var deviceRoomMap map[string]string = make(map[string]string)
	for _, dev := range deviceHomeList.DevList {
		if dev.RoomId != "" {
			deviceRoomMap[dev.Did] = roomMap[iotutil.ToInt64(dev.RoomId)]
		}
	}

	tlsOp := entitys.TlsInfo{}
	if ret.Data != nil && len(ret.Data) > 0 {
		//TODO 如果需要对物模型进行翻译,启用下列代码
		cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
		langMap, _ := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

		sceneIntelligenceForm := entitys.Intelligence_pb2e(ret.Data[0])
		//条件proto转entity
		if conditionRet.Data != nil && len(conditionRet.Data) > 0 {
			sceneIntelligenceForm.Condition = s.setConditionList(lang, langMap, conditionRet.Data, &tlsOp)
		}
		//任务proto转entity
		if taskRet.Data != nil && len(taskRet.Data) > 0 {
			sceneIntelligenceForm.Task = s.setTaskList(lang, langMap, taskRet.Data, &tlsOp)
		}
		//智能场景form转oldform
		oldSceneIntelligenceForm := entitys.Intelligence_new2old(sceneIntelligenceForm, tenantId, lang, deviceRoomMap)
		return &oldSceneIntelligenceForm, nil
	} else {
		return nil, errors.New("查询智能场景失败")
	}
}

// 设置条件返回列表
func (s SceneIntelligenceconService) setConditionList(lang string, langMap map[string]string, conditionList []*protosService.SceneIntelligenceCondition, tlsOp *entitys.TlsInfo) []entitys.SceneIntelligenceConditionForm {
	condition := entitys.Condition_pb2e(conditionList)
	//条件翻译处理
	for i, c := range condition {
		if iotconst.ConditionType(c.ConditionType) == iotconst.CONDITION_TYPE_SATACHANGE {
			multiple, dataType := tlsOp.GetMultipleAndDataType(c.ProductKey, c.DevicePropKey)
			langKey := fmt.Sprintf("%s_%s_%s_name", lang, c.ProductKey, c.DevicePropIdentifier)
			funcDesc := iotutil.MapGetStringVal(langMap[langKey], c.DevicePropDesc)
			//功能值翻译
			langKey = fmt.Sprintf("%s_%s_%s_%v_name", lang, c.ProductKey, c.DevicePropIdentifier, c.DevicePropValue)
			funcValueDesc := iotutil.MapGetStringVal(langMap[langKey], c.DevicePropValue)
			condition[i].Desc = fmt.Sprintf("%v: %v", funcDesc, funcValueDesc)
			//增加倍数和数据类型返回
			condition[i].DevicePropMultiple = multiple
			condition[i].DevicePropDataType = dataType

			//获取条件中设备的状态
			onlineStatus, _ := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+c.DeviceDid, "onlineStatus").Result()
			if onlineStatus == "" {
				//设备已移除
				condition[i].FailureFlag = 2
			} else if onlineStatus == "offline" {
				//设备已离线
				condition[i].FailureFlag = 1
			}
		}
	}
	return condition
}

// 创建任务返回列表
func (s SceneIntelligenceconService) setTaskList(lang string, langMap map[string]string, taskList []*protosService.SceneIntelligenceTask, tlsOp *entitys.TlsInfo) []entitys.SceneIntelligenceTaskForm {
	task := entitys.Task_pb2e(taskList)
	for i, t := range task {
		//任务条件处理
		if iotconst.TaskType(t.TaskType) == iotconst.TASK_TYPE_DEVICE && t.Functions != nil {
			//获取属性的倍数和数据类型
			//返回倍数、数据类型
			for j, function := range t.Functions {
				multiple, dataType := tlsOp.GetMultipleAndDataType(t.ProductKey, function.FuncKey)
				//功能键翻译
				langKey := fmt.Sprintf("%s_%s_%s_name", lang, t.ProductKey, function.FuncIdentifier)
				funcDesc := iotutil.MapGetStringVal(langMap[langKey], function.FuncDesc)
				//功能值翻译
				langKey = fmt.Sprintf("%s_%s_%s_%v_name", lang, t.ProductKey, function.FuncIdentifier, function.FuncValue)

				//funcValueDesc := iotutil.MapGetStringVal(langMap[langKey], function.FuncValueDesc)
				funcValueDesc := function.FuncValueDesc
				if !iotutil.IsNumeric(function.FuncValueDesc) {
					funcValueDesc = iotutil.MapGetStringVal(langMap[langKey], function.FuncValue)
				}

				//t.Functions[j].FuncDesc = fmt.Sprintf("%v: %v", funcDesc, funcValueDesc)
				t.Functions[j].ShowDesc = fmt.Sprintf("%v: %v", funcDesc, funcValueDesc)
				//增加倍数和数据类型返回
				t.Functions[j].Multiple = multiple
				t.Functions[j].DataType = dataType
			}
			//获取条件中设备的状态
			onlineStatus, _ := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+t.ObjectId, "onlineStatus").Result()
			if onlineStatus == "" {
				//设备已移除
				t.FailureFlag = 2
			} else if onlineStatus == "offline" {
				//设备已离线
				t.FailureFlag = 1
			}
			task[i] = t
		}
	}
	return task
}

// 查询智能场景执行结果日志
func (s SceneIntelligenceconService) GetIntelligenceResultLogList(req entitys.SceneIntelligenceQueryForm) (rets []*entitys.SceneIntelligenceLogVo, total int, err error) {
	userId, _ := metadata.Get(s.Ctx, "userId")
	userIdI := iotutil.ToInt64(userId)
	if userIdI == 0 {
		return nil, 0, errors.New("用户Id获取失败")
	}

	ret, err := rpc.SceneIntelligenceLogService.Lists(s.Ctx, &protosService.SceneIntelligenceLogListRequest{
		//Page:      int64(req.Page),
		//PageSize:  int64(req.Limit),
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.SceneIntelligenceLog{UserId: userIdI,
			CreatedAt: timestamppb.New(time.Unix(timestamppb.Now().Seconds-7*24*3600, 0))},
	})
	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Message)
	}
	var (
		dataList = make([]*entitys.SceneIntelligenceLogVo, len(ret.Data))
	)

	mapstructure.WeakDecode(ret.Data, &dataList)

	//返回结果t
	var (
		limit    = req.Limit
		offset   = limit * (req.Page - 1)
		indexEnd = limit + offset
	)
	total = len(dataList)
	//组装返回结果
	hasPage := req.Page != 0 || req.Limit != 0
	if hasPage {
		if offset > total {
			return nil, total, err
		} else {
			if (offset + limit) > total {
				indexEnd = total
			}
			return dataList[offset:indexEnd], total, err
		}
	} else {
		return dataList, total, err
	}
}

// 查询智能场景执行结果日志
func (s SceneIntelligenceconService) GetIntelligenceResultLogGroupList(req entitys.SceneIntelligenceQueryForm) (rets []map[string]interface{}, err error) {
	homeIdInt := req.HomeId
	if homeIdInt == 0 {
		return nil, errors.New("HomeId获取失败")
	}
	ret, err := rpc.SceneIntelligenceLogService.Lists(s.Ctx, &protosService.SceneIntelligenceLogListRequest{
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Page:      1,
		PageSize:  500,
		Query: &protosService.SceneIntelligenceLog{
			//UserId:    userIdI,
			HomeId:    homeIdInt,
			CreatedAt: timestamppb.New(time.Unix(timestamppb.Now().Seconds-7*24*3600, 0)),
		},
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	return converGroupData(ret.Data), nil

}

func converGroupData(result []*protosService.SceneIntelligenceLog) []map[string]interface{} {
	dateList := map[string][]*entitys.SceneIntelligenceLogEntitys{}
	currentTime := time.Now()
	zeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	ServenDayZeroTime := zeroTime.AddDate(0, 0, -6) //获取七天前的时间
	for _, row := range result {
		createAt := row.CreatedAt.AsTime()
		if !createAt.After(ServenDayZeroTime) {
			continue
		}
		date := iotutil.TimeFormat(createAt)
		list, ok := dateList[date]
		if !ok {
			list = []*entitys.SceneIntelligenceLogEntitys{}
		}
		list = append(list, entitys.SceneIntelligenceLog_pb2e(row))
		dateList[date] = list
	}
	resultArr := []map[string]interface{}{}
	num := 1
	for key, value := range dateList {
		valueTime, _ := iotutil.GetTimeByStr(key)
		resultArr = append(resultArr, map[string]interface{}{
			"groupTime": key,
			"item":      value,
			"month":     time.Date(valueTime.Year(), valueTime.Month(), valueTime.Day(), 0, 0, 0, 0, currentTime.Location()).Month(),
			"num":       num,
			"weeks":     time.Date(valueTime.Year(), valueTime.Month(), valueTime.Day(), 0, 0, 0, 0, currentTime.Location()).Weekday(),
		})
		num = num + 1
	}
	//根据sort进行排序
	sort.Slice(resultArr, func(i, j int) bool {
		return resultArr[i]["groupTime"].(string) > resultArr[j]["groupTime"].(string)
	})
	return resultArr
}

// 查询智能场景任务执行结果列表
func (s SceneIntelligenceconService) GetIntelligenceTaskResultList(resultId int64) (rets []*entitys.SceneIntelligenceResultTaskVo, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	ret, err := rpc.SceneIntelligenceResultTaskService.Lists(s.Ctx, &protosService.SceneIntelligenceResultTaskListRequest{
		Page:     1,
		PageSize: 999,
		Query:    &protosService.SceneIntelligenceResultTask{ResultId: resultId},
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	var dataList []*entitys.SceneIntelligenceResultTaskVo

	//TODO 如果需要对物模型进行翻译,启用下列代码
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	for _, src := range ret.Data {
		entity := entitys.ResultTask_pb2e(src)
		if src.TaskType == int32(iotconst.TASK_TYPE_DEVICE) {
			if src.Functions != "" {
				var functions []entitys.TaskFunction
				json.Unmarshal([]byte(src.Functions), &functions)
				if functions != nil {
					var funcKey, funcDesc, funcValue, funcResult []string
					for _, function := range functions {
						funcKey = append(funcKey, function.FuncIdentifier)

						//功能键翻译
						langKey := fmt.Sprintf("%s_%s_%s_name", lang, entity.ProductKey, function.FuncIdentifier)
						function.FuncDesc = iotutil.MapGetStringVal(langMap[langKey], function.FuncDesc)
						funcDesc = append(funcDesc, function.FuncDesc)

						//功能值翻译
						langKey = fmt.Sprintf("%s_%s_%s_%v_name", lang, entity.ProductKey, function.FuncIdentifier, function.FuncValue)
						//数值转换，更具倍数对功能值进行转换
						//if function.DataType != "" && function.Multiple != nil {
						//	multiple, err := iotutil.ToFloat64Err(function.Multiple)
						//	if err != nil && multiple > 0 {
						//		fv, err := iotutil.ToFloat64Err(function.FuncValue)
						//		if err != nil && fv > 0 {
						//			function.FuncValue = fmt.Sprintf("%.1f", fv/multiple)
						//		}
						//	}
						//}
						//如果是数值，则使用FuncValueDesc进行显示
						if iotutil.IsNumeric(function.FuncValueDesc) {
							function.FuncValue = function.FuncValueDesc
						} else {
							function.FuncValue = iotutil.MapGetStringVal(langMap[langKey], function.FuncValue)
						}
						funcValue = append(funcValue, function.FuncValue)

						funcResult = append(funcResult, fmt.Sprintf("%s：%s", function.FuncDesc, function.FuncValue))
					}
					entity.FuncKey = strings.Join(funcKey, "、")
					entity.FuncValue = strings.Join(funcValue, "、")
					entity.FuncDesc = strings.Join(funcDesc, "、")
					entity.FuncResult = strings.Join(funcResult, "、")
				}
			} else {
				//历史兼容
				if entity.FuncKey != "" {
					langKey := fmt.Sprintf("%s_%s_%s_%s_name", lang, entity.ProductKey, entity.FuncKey, entity.FuncValue)
					entity.FuncDesc = iotutil.MapGetStringVal(langMap[langKey], entity.FuncDesc)
					entity.FuncResult = entity.FuncDesc
				}
			}
		} else if src.TaskType == int32(iotconst.TASK_TYPE_DELAYED) {
			entity.FuncResult = entity.FuncDesc
		}
		dataList = append(dataList, entity)
	}
	//mapstructure.WeakDecode(ret.Data, &dataList)
	//返回结果t
	return dataList, err
}

// 删除智能场景日志
func (s SceneIntelligenceconService) DeleteIntelligenceLog() error {
	userId, _ := metadata.Get(s.Ctx, "userId")
	userIdI := iotutil.ToInt64(userId)
	if userIdI == 0 {
		return errors.New("用户Id获取失败")
	}
	res, err := rpc.SceneIntelligenceLogService.Delete(s.Ctx, &protosService.SceneIntelligenceLog{
		UserId: userIdI,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

// GetTaskOrWhereByProductKey 获取产品对应的智能条件和智能任务
func (s SceneIntelligenceconService) GetTaskOrWhereByProductKey(productId int64, condType string) (rets *entitys.ProductThingsModel, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	ret, err := rpc.ProductThingsModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productId,
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	retPro, err := rpc.ProductService.FindById(context.Background(), &protosService.OpmProductFilter{Id: productId})
	if err != nil {
		return nil, err
	}
	if retPro != nil && retPro.Code != 200 {
		return nil, errors.New(retPro.Message)
	}

	var (
		thingsModel        = new(entitys.ProductThingsModel)
		productKey  string = retPro.Data[0].ProductKey
	)

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	for _, property := range ret.Data.Properties {
		switch condType {
		case "triggerCond":
			//TODO 这里的productKey可以修改为property.productKey
			name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, property.Identifier)], property.Name)
			//TODO DataSpecsList 也需要进行翻译
			//TODO DataSpecs 也需要进行翻译
			//是否启用作为条件
			if property.TriggerCond == 1 {
				thingsModel.Properties = append(thingsModel.Properties, &entitys.ThingModelProperties{
					Name:          name,
					Identifier:    iotutil.ToString(property.Dpid), //property.Identifier,
					DataType:      property.DataType,
					DataSpecs:     CheckDataSpec(property.DataSpecs),
					DefaultVal:    property.DefaultVal,
					DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, property.DataSpecsList, langMap, nil, iotutil.ToString(property.Dpid)),
				})
			}
		case "execCond":
			//TODO 这里的productKey可以修改为property.productKey
			name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, property.Identifier)], property.Name)
			//TODO DataSpecsList 也需要进行翻译
			//TODO DataSpecs 也需要进行翻译
			//是否启用作为执行动作
			if property.ExecCond == 1 {
				thingsModel.Properties = append(thingsModel.Properties, &entitys.ThingModelProperties{
					Name:          name,
					Identifier:    iotutil.ToString(property.Dpid), //property.Identifier,
					DataType:      property.DataType,
					DataSpecs:     CheckDataSpec(property.DataSpecs),
					DefaultVal:    property.DefaultVal,
					DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, property.DataSpecsList, langMap, nil, iotutil.ToString(property.Dpid)),
				})
			}
		}
	}
	return thingsModel, nil
}

// GetTaskOrWhereByProductKeyV2 与GetTaskOrWhereByProductKey的区别为dpid和identifier的返回值
func (s SceneIntelligenceconService) GetTaskOrWhereByProductKeyV2(productId int64, condType, deviceId string) (rets *entitys.ProductThingsModel, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	ret, err := rpc.ProductThingsModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productId,
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	retPro, err := rpc.ProductService.FindById(context.Background(), &protosService.OpmProductFilter{Id: productId})
	if err != nil {
		return nil, err
	}
	if retPro != nil && retPro.Code != 200 {
		return nil, errors.New(retPro.Message)
	}

	var (
		thingsModel        = new(entitys.ProductThingsModel)
		productKey  string = retPro.Data[0].ProductKey
	)

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	funDescMap, _ := getFunctionSetMap(deviceId)
	for _, property := range ret.Data.Properties {
		name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, property.Identifier)], property.Name)
		//客户自定义功能描述获取
		if funDescMap != nil {
			if v, ok := funDescMap[fmt.Sprintf("%s_%s", productKey, property.Dpid)]; ok && v != "" {
				name = v
			}
		}
		switch condType {
		case "triggerCond":
			//是否启用作为条件
			if property.TriggerCond == 1 {
				//自定义触发条件参数
				if property.TriggerCondArgs != "" {
					switch property.DataType {
					case "ENUM", "BOOL":
						property.DataSpecsList = property.TriggerCondArgs
					case "DOUBLE", "FLOAT", "INT", "TEXT":
						property.DataSpecs = property.TriggerCondArgs
					}
				}
				thingsModel.Properties = append(thingsModel.Properties, &entitys.ThingModelProperties{
					Name:          name,
					Identifier:    property.Identifier,
					DpId:          iotutil.ToString(property.Dpid),
					DataType:      property.DataType,
					DataSpecs:     CheckDataSpec(property.DataSpecs),
					DefaultVal:    property.DefaultVal,
					DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, property.DataSpecsList, langMap, funDescMap, iotutil.ToString(property.Dpid)),
				})
			}
		case "execCond":
			//是否启用作为执行动作
			if property.ExecCond == 1 {
				//自定义执行任务参数
				if property.ExecCondArgs != "" {
					switch property.DataType {
					case "ENUM", "BOOL":
						property.DataSpecsList = property.ExecCondArgs
					case "DOUBLE", "FLOAT", "INT", "TEXT":
						property.DataSpecs = property.ExecCondArgs
					}
				}
				thingsModel.Properties = append(thingsModel.Properties, &entitys.ThingModelProperties{
					Name:          name,
					Identifier:    property.Identifier,
					DpId:          iotutil.ToString(property.Dpid),
					DataType:      property.DataType,
					DataSpecs:     CheckDataSpec(property.DataSpecs),
					DefaultVal:    property.DefaultVal,
					DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, property.DataSpecsList, langMap, funDescMap, iotutil.ToString(property.Dpid)),
				})
			}
		}
	}
	return thingsModel, nil
}

// 获取产品信息
func (s SceneIntelligenceconService) GetProductInfo(productKey string) (*protosService.OpmProduct, error) {
	ret, err := rpc.ProductService.Find(s.Ctx, &protosService.OpmProductFilter{
		ProductKey: productKey,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret.Data[0], nil
}

// GetProductFunctions 获取设备的功能列表 dataMode数据模式
func (s SceneIntelligenceconService) GetProductFunctions(deviceId string, dataMode int32) (rets *entitys.ProductThingsModel, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")

	//设备换成信息
	deviceInfo := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId)
	if deviceInfo.Err() != nil {
		return nil, deviceInfo.Err()
	}
	productId, err := iotutil.ToInt64AndErr(deviceInfo.Val()["productId"])
	if err != nil {
		return nil, errors.New("产品Id异常")
	}
	productKey := iotutil.ToString(deviceInfo.Val()["productKey"])

	//查询产品信息
	proInfo, err := s.GetProductInfo(productKey)
	if err != nil {
		return nil, err
	}

	ret, err := rpc.ProductThingsModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productId,
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	var (
		thingsModel = new(entitys.ProductThingsModel)
	)

	//面板样式规则
	if proInfo.StyleLinkage != "" {
		thingsModel.StyleLinkage, _ = iotutil.JsonToMapErr(proInfo.StyleLinkage)
	}

	//翻译数据获取
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	funDescMap, _ := getFunctionSetMap(deviceId)

	dpidMaps := make(map[int32]*protosService.OpmThingModelProperties)
	for _, property := range ret.Data.Properties {
		if property.AllowAppointment == 1 {
			dpidMaps[property.Dpid] = property
		}
	}

	//tree结构
	thingsModel.Properties = setChildrens(0, lang, productKey, langMap, funDescMap, dataMode, ret.Data.Properties)

	//功能规则
	proSvc := services2.ProductService{Ctx: context.Background()}
	thingsModel.Rules, err = proSvc.GetFunctionRules(productKey, 2, dpidMaps)
	if err != nil {
		iotlogger.LogHelper.Error("功能列表接口-功能规则获取失败，error：%s", err.Error())
	}
	return thingsModel, nil
}

func setChildrens(parentId int64, lang string, productKey string,
	langMap map[string]string, funDescMap map[string]string, dataMode int32,
	list []*protosService.OpmThingModelProperties) []*entitys.ThingModelProperties {

	var resList []*entitys.ThingModelProperties
	for _, property := range list {
		if property.ParentId == parentId {
			//预约数据获取
			dataSpecs := property.DataSpecs
			dataSpecsList := property.DataSpecsList
			if dataMode == 2 {
				if property.AllowAppointment != 1 {
					continue
				}
				if property.AppointmentArgs != "" {
					switch property.DataType {
					case "ENUM", "BOOL":
						dataSpecsList = property.AppointmentArgs
					case "DOUBLE", "FLOAT", "INT", "TEXT":
						dataSpecs = property.AppointmentArgs
					}
				}
			}
			name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, property.Identifier)], property.Name)
			//客户自定义功能描述获取
			if funDescMap != nil {
				if v, ok := funDescMap[fmt.Sprintf("%s_%s", productKey, property.Dpid)]; ok && v != "" {
					name = v
				}
			}
			childrenList := setChildrens(property.Id, lang, productKey, langMap, funDescMap, dataMode, list)

			//是否启用作为执行动作
			// if property.ExecCond == 1 {
			resList = append(resList, &entitys.ThingModelProperties{
				Name:          name,
				Identifier:    property.Identifier,
				DpId:          iotutil.ToString(property.Dpid),
				DataType:      property.DataType,
				RwFlag:        property.RwFlag,
				DataSpecs:     CheckDataSpec(dataSpecs),
				DefaultVal:    property.DefaultVal,
				DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, dataSpecsList, langMap, funDescMap, iotutil.ToString(property.Dpid)),
				Childrens:     childrenList,
				Sort:          property.Sort,
			})
		}
	}

	//排序问题
	sort.Slice(resList, func(i, j int) bool {
		return resList[i].Sort < resList[j].Sort
	})

	return resList
}

// {"dataType":"INT","max":"10000000","min":"0","step":"1","unit":"mg"}
func CheckDataSpec(dataSpecs string) string {
	if dataSpecs == "" {
		return dataSpecs
	}
	mapData, err := iotutil.JsonToMapErr(dataSpecs)
	if err == nil {
		return dataSpecs
	}
	if _, ok := mapData["multiple"]; !ok {
		mapData["multiple"] = 1
		dataSpecs = iotutil.ToString(mapData)
	}
	return dataSpecs
}

// 获取功能名称自定义设置Map数据
func getFunctionSetMap(deviceId string) (map[string]string, error) {
	if deviceId == "" {
		return map[string]string{}, nil
	}
	setRes, err := rpc.IotDeviceInfoService.GetDeviceFunctionSetList(context.Background(), &protosService.IotDeviceFunctionSet{
		DeviceId: deviceId,
	})
	if err != nil {
		return nil, err
	}
	resMap := make(map[string]string)
	for _, l := range setRes.Data {
		if l.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_SET {
			resMap[fmt.Sprintf("%s_%s", l.ProductKey, l.FuncKey)] = l.CustomDesc
		} else if l.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_VALUE_SET {
			resMap[fmt.Sprintf("%s_%s_%v", l.ProductKey, l.FuncKey, l.FuncValue)] = l.CustomDesc
		}
	}
	return resMap, nil
}

func ConvertJsonByLang(lang string, productKey string, identifier string, jsonStr string, langMap map[string]string, funcSetMap map[string]string, dpid string) string {
	if jsonStr == "" {
		return jsonStr
	}
	var resObj []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &resObj)
	if err != nil {
		//panic(err)
		return jsonStr
	}
	if resObj == nil {
		return jsonStr
	}
	for i, item := range resObj {
		val := iotutil.ToString(item["value"])
		desc := item["desc"]
		name := item["name"]
		if desc == "" || desc == nil {
			desc = name
		}
		dataType := iotutil.ToString(item["dataType"])
		//数值转换（BOOL类型特殊处理）
		if dataType == "BOOL" {
			if val == "1" || val == "true" {
				val = "true"
			} else if val == "0" {
				val = "false"
			}
		}
		//如果有设置自定名称，则已自定义名称显示，否则使用翻译显示；
		if funcSetMap != nil {
			if v, ok := funcSetMap[fmt.Sprintf("%s_%v_%v", productKey, dpid, val)]; ok && v != "" {
				name = v
				resObj[i]["name"] = name
				resObj[i]["desc"] = name
				continue
			}
		}
		langKey := fmt.Sprintf("%s_%s_%s_%v_name", lang, productKey, identifier, val)
		desc = iotutil.MapGetStringVal(langMap[langKey], desc)
		resObj[i]["desc"] = desc
		resObj[i]["value"] = val
	}
	return iotutil.ToString(resObj)
}

// OneKeyExec 一键执行
func (s SceneIntelligenceconService) OneKeyExec(id int64) (resultId int64, err error) {
	res, err := rpc.SceneIntelligenceService.OneKeyExec(s.Ctx, &protosService.SceneIntelligencePrimarykey{
		Id: id,
	})
	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Message)
	}
	return res.Data, nil
}

// 清理场景列表缓存
func (s SceneIntelligenceconService) clearIntelligenceListCachedByHomdId(homeId int64) error {
	keys := []string{
		fmt.Sprintf(iotconst.APP_HOME_INTELLIGENCE_DATA, homeId),
	}
	pipe := cached.RedisStore.Pipeline()
	pipe.Del(context.Background(), keys...)
	if _, err := pipe.Exec(context.Background()); err != nil {
		return err
	}
	return nil
}

// 通过场景Id清理缓存
func (s SceneIntelligenceconService) clearIntelligenceListCachedById(intelligenceId int64) error {
	res, err := rpc.SceneIntelligenceService.FindById(s.Ctx, &protosService.SceneIntelligenceFilter{
		Id: intelligenceId,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return s.clearIntelligenceListCachedByHomdId(res.Data[0].HomeId)
}
