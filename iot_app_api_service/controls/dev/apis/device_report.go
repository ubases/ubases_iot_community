package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
)

var DeviceReportsapis DeviceReportApis

type DeviceReportApis struct {
}

// GetDaysDetail
// @Summary 获取当日运行数据
// @Description
// @Tags 设备
// @Accept application/json
// @Param deviceId query string true "设备Id"
// @Param productKey query string true "产品Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/runRecord/daysGroupDetail [get]
func (a *DeviceReportApis) GetDaysDetail(c *gin.Context) {
	deviceId := c.Query("deviceId")
	if deviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	ctx := controls.WithUserContext(c)
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	tz := controls.GetTimezone(c)
	nowTime := time.Now()
	dayBefore7 := nowTime.AddDate(0, 0, -7).Unix()
	toDay := nowTime.Unix()

	proRes, err := rpc.ProductService.Find(ctx, &protosService.OpmProductFilter{ProductKey: productKey})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if proRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(proRes.Message))
		return
	}

	res, err := rpc.IotDeviceLogService.ProductEventLogReport(ctx, &protosService.ProductLogRequest{
		DeviceId:   deviceId,
		ProductKey: productKey,
		Timezone:   tz,
		TimeRange:  []int64{dayBefore7, toDay},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	//数据格式转换
	iotgin.ResSuccess(c, convertGroupData(tenantId, lang, productKey, proRes.Data[0].Name, res.List))
}

// 转换分组数据
func convertGroupData(tenantId string, lang string, productKey, productName string, datas []*protosService.ProductEventLogItem) []map[string]interface{} {
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Date > datas[i].Date
	})
	productName = getProductNameLang(tenantId, lang, productKey, productName)
	langMap := getThingsModelLang(tenantId)

	mapList := make(map[string][]map[string]string)
	for _, d := range datas {
		if _, ok := mapList[d.Date]; !ok {
			mapList[d.Date] = []map[string]string{}
		}
		//设备上线、下线事件日志逻辑
		if d.OnlineStatus != "" {
			funcDesc := d.OnlineStatus //不实现翻译
			if lang == "zh" {
				switch d.OnlineStatus {
				case "online":
					funcDesc = "上线"
				default:
					funcDesc = "下线"
				}
			}
			mapList[d.Date] = append(mapList[d.Date], map[string]string{
				"time":  d.Time,
				"from":  d.From,
				"value": fmt.Sprintf("%s %s", productName, funcDesc), //空气净化器关闭
			})
			continue
		}
		funcs := make([]string, 0)
		from := "app"
		if d.From != "" {
			from = d.From
		}
		//
		for _, m := range d.Properties {
			//名称翻译
			langKey := fmt.Sprintf("%s_%s_%s_name", lang, productKey, m.Identifier)
			nameDesc := iotutil.MapGetStringVal(langMap[langKey], m.Name)
			//属性值翻译
			langKey = fmt.Sprintf("%s_%s_%s_%v_name", lang, productKey, m.Identifier, m.Value)
			valueDesc := iotutil.MapGetStringVal(langMap[langKey], m.Value)
			funcs = append(funcs, fmt.Sprintf("%s %s", nameDesc, valueDesc))
		}
		mapList[d.Date] = append(mapList[d.Date], map[string]string{
			"time":  d.Time,
			"from":  from,
			"value": fmt.Sprintf("%s %s", productName, strings.Join(funcs, "、")), //空气净化器关闭
		})
	}
	res := make([]map[string]interface{}, 0)
	for key, list := range mapList {
		sort.Slice(list, func(i, j int) bool {
			return list[i]["fullTime"] > list[j]["fullTime"]
		})
		res = append(res, map[string]interface{}{
			"data": key,
			"list": list,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return iotutil.ToString(res[i]["data"]) > iotutil.ToString(res[j]["data"])
	})
	return res
}

// ClearDetail
// @Summary 清理运行记录
// @Description
// @Tags 设备
// @Accept application/json
// @Param deviceId query string true "设备Id"
// @Param productKey query string true "产品Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/runRecord/clearDetail [get]
func (a *DeviceReportApis) ClearDetail(c *gin.Context) {
	deviceId := c.Query("deviceId")
	if deviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	ctx := controls.WithUserContext(c)
	res, err := rpc.IotDeviceLogService.ClearDeviceLogs(ctx, &protosService.ProductLogRequest{
		DeviceId:   deviceId,
		ProductKey: productKey,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	//数据格式转换
	iotgin.ResSuccessMsg(c)
}

// 获取物理模型翻译数据
func getThingsModelLang(tenantId string) map[string]string {
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return make(map[string]string)
	}
	return langMap
}

// 获取产品名称的翻译
func getProductNameLang(tenantId, lang, productKey, productName string) string {
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_NAME)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return productName
	}
	productName = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_name", lang, productKey)], productName)
	return productName
}

// GetDaysHourCount
// @Summary 获取日运行数据按照小时份分组
// @Description
// @Tags 设备
// @Accept application/json
// @Param deviceId query string true "设备Id"
// @Param productKey query string true "产品Id"
// @Param identifier query string true "功能标识"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/runRecord/GetDaysHourCount [get]
func (a *DeviceReportApis) GetDaysHourCount(c *gin.Context) {
	deviceId := c.Query("deviceId")
	if deviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	identifier := c.Query("identifier")
	if identifier == "" {
		//iotgin.ResBadRequest(c, "identifier")
		//return
		identifier = "methanal"
	}
	ctx := controls.WithUserContext(c)
	nowTime := time.Now()
	dayBefore7 := nowTime.AddDate(0, 0, -7).Unix()
	toDay := nowTime.Unix()

	res, err := rpc.IotDeviceLogService.ProductEventLogReport(ctx, &protosService.ProductLogRequest{
		DeviceId:   deviceId,
		ProductKey: productKey,
		TimeRange:  []int64{dayBefore7, toDay},
		Identifier: identifier,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	//数据格式转换
	iotgin.ResSuccess(c, convertGroupCount(identifier, res.List))
}

// 转换分组数据
func convertGroupCount(identifier string, datas []*protosService.ProductEventLogItem) []map[string]interface{} {
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Date > datas[i].Date
	})
	mapList := make(map[string][]map[string]interface{})
	//将查询数据转换为hour和value
	for _, d := range datas {
		if _, ok := mapList[d.Date]; !ok {
			mapList[d.Date] = []map[string]interface{}{}
		}
		currentVal := getPointIdentifierVal(identifier, d.Properties)
		mapList[d.Date] = append(mapList[d.Date], map[string]interface{}{
			"hour":  iotutil.ToInt32(strings.Split(d.Time, ":")[0]),
			"value": currentVal,
		})
	}

	//将指定日期下的时间按照消息小时合并
	res := make([]map[string]interface{}, 0)
	for key, list := range mapList {
		var maxVal, minVal float64 = 0, 0
		hourMap := map[int32]float64{}
		for _, l := range list {
			currVal := l["value"].(float64)
			if currVal > maxVal {
				maxVal = currVal
			} else if currVal < minVal {
				minVal = currVal
			}
			//TODO 需要合并1,3,5,7,9.......
			currHour := l["hour"].(int32)
			if _, ok := hourMap[currHour]; ok {
				hourMap[currHour] = hourMap[currHour] + currVal
			} else {
				hourMap[currHour] = currVal
			}
		}
		hourList := make([]map[string]interface{}, 0)
		var i int32 = 1
		for i = 1; i < 24; i += 2 {
			var tempVal float64 = 0
			for j := i - 1; j <= i+1; j++ {
				if j < 0 || j >= 24 {
					continue
				}
				if v, ok := hourMap[j]; ok {
					tempVal += v
				}
			}
			hourList = append(hourList, map[string]interface{}{
				"key":   i,
				"value": tempVal,
				"type":  "0",
			})
		}

		res = append(res, map[string]interface{}{
			"data": key,
			"max":  maxVal,
			"min":  minVal,
			"list": hourList,
		})
	}
	return res
}

// 获取指定物模型的值
func getPointIdentifierVal(identifier string, properties []*protosService.ProductLogEventProperties) float64 {
	var currentVal float64 = 0
	for _, m := range properties {
		if m.Identifier != identifier {
			continue
		}
		if m.Value == "" {
			continue
		}
		val, _ := iotutil.ToFloat64Err(m.Value)
		currentVal = val
	}
	return currentVal
}

// CreateTestData 生成测试数据
func CreateTestData(dateStr string) map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	var (
		max = 500
		min = 0
	)
	for i := 1; i < 24; i += 2 {
		rdNum := rand.Intn(201)
		//result, _ := rand.Int(rand.Reader, big.NewInt(200))
		if rdNum > max {
			max = rdNum
		}
		if rdNum < min {
			min = rdNum
		}
		res = append(res, map[string]interface{}{
			"key":   i,
			"value": rdNum,
			"type":  "0",
		})
	}
	for i := 1; i < 24; i += 2 {
		rdNum := rand.Intn(301) + 200
		if rdNum > max {
			max = rdNum
		}
		if rdNum < min {
			min = rdNum
		}
		res = append(res, map[string]interface{}{
			"key":   i,
			"value": rdNum,
			"type":  "1",
		})
	}
	return map[string]interface{}{
		"data": dateStr,
		"max":  max,
		"min":  min,
		"list": res,
	}
}

// GetDaysHourCountTest 获取每日每小时的测试数据
func (a *DeviceReportApis) GetDaysHourCountTest(c *gin.Context) {
	nowDate := time.Now()
	list := make([]map[string]interface{}, 0)
	for i := 0; i < 7; i++ {
		list = append(list, CreateTestData(nowDate.Add(-time.Duration(i)*time.Hour*24).Format("2006-01-02")))
	}
	iotgin.ResSuccess(c, list)
}
