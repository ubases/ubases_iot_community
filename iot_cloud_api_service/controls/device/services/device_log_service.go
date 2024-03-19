package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common"
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IotDeviceLogService struct {
}

// 设备日志详细
func (s IotDeviceLogService) GetIotDeviceLogDetail(id string) (*entitys.IotDeviceLogEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientIotDeviceLogServer.FindById(context.Background(), &protosService.IotDeviceLogFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]
	return entitys.IotDeviceLog_pb2e(data), err
}

// QueryIotDeviceLogList 设备日志列表
func (s IotDeviceLogService) QueryIotDeviceLogList(filter entitys.IotDeviceLogQuery) ([]*entitys.IotDeviceLogEntitys, int64, error) {
	productMap := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+filter.Query.Did)
	if productMap.Err() != nil {
		return nil, 0, nil
	}
	productInfo := productMap.Val()
	productKey := productInfo["productKey"]
	if productKey == "" {
		return nil, 0, nil
	}
	var ns string
	if filter.Query.EventType != 0 {
		ns = entitys.EventTypeToNs(productKey, filter.Query.Did, filter.Query.EventType)
	}

	//默认时间处理
	if filter.Query.StartTime == 0 || filter.Query.EndTime == 0 {
		filter.Query.StartTime = iotutil.BeginningOfPointDay(-7).Unix()
		filter.Query.EndTime = time.Now().Unix()
	} else {
		//设置开始00:00:00，结束23:59:59
		startTime := time.Unix(filter.Query.StartTime, 0)
		filter.Query.StartTime = iotutil.GetTodaySartTime(startTime).Unix()
		endTime := time.Unix(filter.Query.EndTime, 0)
		filter.Query.EndTime = iotutil.GetTodayLastTime(endTime).Unix()
	}
	rep, err := rpc.ClientIotDeviceLogServer.ProductReportLogRecord(context.Background(), &protosService.ProductLogRequest{
		DeviceId:   filter.Query.Did,
		ProductKey: productKey,
		Ns:         ns,
		Origin:     filter.Query.Origin,
		Identifier: filter.Query.EventKey,
		EventType:  filter.Query.EventType,
		TimeRange:  []int64{filter.Query.StartTime, filter.Query.EndTime},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	tenantId := iotutil.ToString(productInfo["tenantId"])
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	var resultList = []*entitys.IotDeviceLogEntitys{}
	//如果是IsOnlyCount
	if rep.Records == nil {
		return resultList, 0, nil
	}
	//数据格式转换
	for _, src := range rep.Records {
		eventType, eventTypeName := entitys.ToEventTypeNameExt(src.Ns)

		row := &entitys.IotDeviceLogEntitys{
			Id:             src.MsgId,
			Did:            src.Did,
			ReportTime:     src.CreatedAt,
			EventType:      eventType, //src.EventType,
			EventTypeName:  eventTypeName,
			OriginType:     3,
			OriginTypeName: "客户端",
			//OriginDetail:   src.Ns,
			CreatedAt: src.CreatedAt,
		}
		//来源转换
		switch src.From {
		case "app":
			row.OriginType = 3
			row.OriginTypeName = "客户端"
		case "device":
			row.OriginType = 1
			row.OriginTypeName = "设备本身"
		case "cloud":
			row.OriginType = 2
			row.OriginTypeName = "云端"
		default:
			row.OriginType = 4
			row.OriginTypeName = "其它"
		}
		funcDescs, funcKeys, funcValues := []string{}, []string{}, []string{}
		if src.OnlineStatus != "" {
			funcDesc := src.OnlineStatus //不实现翻译
			switch src.OnlineStatus {
			case "online":
				funcDesc = "设备上线"
			default:
				funcDesc = "设备下线"
			}
			row.FuncKey = "在线状态" //funcDesc //iotconst.FIELD_ONLINE
			row.FuncValue = src.OnlineStatus
			//row.FuncDesc = fmt.Sprintf("%s:%s", row.FuncKey, funcDesc) //不实现翻译
			row.FuncKeyName = funcDesc
			resultList = append(resultList, row)
			continue
		}
		if src.Properties != nil {
			for _, property := range src.Properties {
				//翻译事件名称
				name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("zh_%s_%s_name", productKey, property.Identifier)], property.Identifier)
				//事件名称
				funcKeys = append(funcKeys, name)
				funcValues = append(funcValues, fmt.Sprintf("%v:%v", property.Identifier, property.Value))
				//翻译值
				val := iotutil.MapGetStringVal(langMap[fmt.Sprintf("zh_%s_%s_%s_name",
					productKey, property.Identifier, property.Value)], property.Value)
				//事件详情（翻译）
				funcDescs = append(funcDescs, fmt.Sprintf("%v:%v", name, val)) //append(funcDescs, fmt.Sprintf("%s", val))
			}
		}
		row.FuncKey = strings.Join(funcKeys, "、")
		row.FuncValue = strings.Join(funcValues, "、")
		row.FuncDesc = strings.Join(funcDescs, "<br/>")
		row.FuncKeyName = strings.Join(funcDescs, "、")
		resultList = append(resultList, row)
	}
	return resultList, int64(len(resultList)), nil
}

// AddIotDeviceLog 新增设备日志
func (s IotDeviceLogService) AddIotDeviceLog(req entitys.IotDeviceLogEntitys) (string, error) {
	saveObj := entitys.IotDeviceLog_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.CreatedAt = timestamppb.Now()
	res, err := rpc.ClientIotDeviceLogServer.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateIotDeviceLog 修改设备日志
func (s IotDeviceLogService) UpdateIotDeviceLog(req entitys.IotDeviceLogEntitys) (string, error) {
	res, err := rpc.ClientIotDeviceLogServer.Update(context.Background(), entitys.IotDeviceLog_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteIotDeviceLog 删除设备日志
func (s IotDeviceLogService) DeleteIotDeviceLog(req entitys.IotDeviceLogFilter) error {
	rep, err := rpc.ClientIotDeviceLogServer.Delete(context.Background(), &protosService.IotDeviceLog{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// Export 导出
func (this *IotDeviceLogService) Export(filter entitys.IotDeviceLogQuery) (string, string, error) {
	//导出生成excel附件
	timeTemplate := "2006-01-02 15:04:00"
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("data")
	headerRow := sheet.AddRow()

	// 标题样式
	headerStyle := common.ExcelHeaderStyle()
	contentStyle := common.ExcelContentStyle()
	cell := headerRow.AddCell()
	cell.Value = "序号"
	cell.SetStyle(headerStyle)
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "时区（GMT+8）"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备事件"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "事件名称"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "事件详情"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "来源"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "来源详情"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)

	res, _, err := this.QueryIotDeviceLogList(filter)
	if err != nil {
		return "", "", err
	}

	for i, row := range res {
		contentRow := sheet.AddRow()
		cell := contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = iotutil.ToString(i)
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = time.Unix(row.ReportTime, 0).Format(timeTemplate)
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = convertEventType(row.EventType)
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = convertEvent(row.FuncKey)
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.FuncDesc
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = convertOriginType(row.OriginType)
		cell = contentRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.OriginDetail
	}

	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err = file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		return "", "", err
	}
	//发送完文件后删除对应文件
	//defer func() {
	//	os.Remove(tempPathFile)
	//}()
	fileName := "deviceLog-" + time.Now().Format("20060102150400") + ".xlsx"
	return fileName, tempPathFile, nil
}

func convertEventType(eventType int32) string {
	eventTypeName := "其它"
	switch eventType {
	case 1: //数据上报
		eventTypeName = "数据上报"
	case 2: //指令下发
		eventTypeName = "指令下发"
	case 3: //设备信号量
		eventTypeName = "设备信号量"
	}
	return eventTypeName
}

func convertOriginType(originType int32) string {
	originTypeName := "其它设备"
	switch originType {
	case 1: //数据上报
		originTypeName = "设备本身"
	case 2: //指令下发
		originTypeName = "客户端"
	case 3: //设备信号量
		originTypeName = "其它设备"
	}
	return originTypeName
}

// 从物模型读取
var eventMap = map[string]string{
	"switch": "开关",
	"light":  "亮度",
}

func convertEvent(eventKey string) string {
	eventName, ok := eventMap[eventKey]
	if !ok {
		eventName = "未知"
	}
	return eventName
}
