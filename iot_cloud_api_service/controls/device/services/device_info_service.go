package services

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/common"
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	"cloud_platform/iot_cloud_api_service/controls/device/services/extract"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	services2 "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IotDeviceInfoService struct {
	Ctx context.Context
}

func (s *IotDeviceInfoService) SetContext(ctx context.Context) *IotDeviceInfoService {
	s.Ctx = ctx
	return s
}

var tempPath = iotconst.GetWorkTempDir() + string(filepath.Separator)

// 设备信息详细
func (s *IotDeviceInfoService) GetIotDeviceInfoDetail(id string) (*entitys.IotDeviceInfoEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientIotDeviceInfoServer.Find(s.Ctx, &protosService.IotDeviceInfoFilter{Id: rid})
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
	return entitys.IotDeviceInfo_pb2e(data), err
}

// 设置查询条件
func deviceListQueryToRequest(filter entitys.IotDeviceInfoQuery) (*protosService.IotDeviceTriadListRequest, error) {
	var (
		isOnline          int32 = -1 //是否在线
		isActive          int32 = -1 //是否激活
		queryStartTime    int64 = 0  //激活或者创建查询开始时间
		queryEndTime      int64 = 0  //激活或者创建查询结束时间
		batchId           int32 = 0  //批次编号
		deviceNature            = "" //设备性质
		err               error
		enableZeroBatchId bool = false
	)
	if filter.Query.IsOnline != nil {
		isOnline = *filter.Query.IsOnline
	}
	if filter.Query.IsActive != nil {
		isActive = *filter.Query.IsActive
	}
	if filter.Query.DeviceNature != 0 {
		deviceNature = iotutil.ToString(filter.Query.DeviceNature)
	}
	if filter.Query.StartTime != nil && filter.Query.StartTime != "" {
		queryStartTime = iotutil.ToInt64(filter.Query.StartTime)
	}
	if filter.Query.EndTime != nil && filter.Query.EndTime != "" {
		queryEndTime = iotutil.ToInt64(filter.Query.EndTime)
	}
	if filter.Query.BatchId != nil && iotutil.ToString(filter.Query.BatchId) != "" {
		batchId, err = iotutil.ToInt32Err(filter.Query.BatchId)
		if err != nil {
			return nil, errors.New("请输入正确格式的批次参数")
		}
		//是否启动批次号0的查询（兼容处理）
		enableZeroBatchId = batchId == 0
	}

	return &protosService.IotDeviceTriadListRequest{
		Page:        int64(filter.Page),
		PageSize:    int64(filter.Limit),
		SearchKey:   filter.SearchKey,
		IsOnlyCount: filter.IsOnlyCount,
		Query: &protosService.IotDeviceTriad{
			DeviceNatureKey:   filter.Query.DeviceNature,
			BatchId:           batchId,
			Status:            isActive,
			ProductId:         filter.Query.ProductId,
			ProductKey:        filter.Query.ProductKey,
			Did:               filter.Query.Did,
			UseType:           0,  //只显示真实设备
			IsTest:            -1, //真实测试设备和真实设备都需要显示
			SerialNumber:      filter.Query.SerialNumber,
			ExportCount:       filter.Query.ExportCount,
			IsExport:          filter.Query.IsExport,
			IsQueryExport:     filter.Query.IsQueryExport,
			EnableZeroBatchId: enableZeroBatchId,
			PlatformCode:      filter.Query.PlatformCode,
			IsPlatform:        filter.IsPlatform,
			DeviceInfo: &protosService.IotDeviceInfo{
				OnlineStatus:   isOnline,
				DeviceNature:   deviceNature,
				DeviceName:     filter.Query.DeviceName,
				QueryStartTime: queryStartTime,
				QueryEndTime:   queryEndTime,
			},
			IsQueryTriadData: filter.Query.IsQueryTriadData,
		},
	}, nil
}

func (s *IotDeviceInfoService) getQueryProductIds(isPlatform bool, deviceName string) (map[int64]*protosService.OpmProduct, []int64, error) {
	//if deviceName == "" {
	//	return map[int64]*protosService.OpmProduct{}, []int64{}, nil
	//}
	//查询产品信息，Map数据结构，Key=productId
	pSvc := services2.OpmProductService{Ctx: s.Ctx}
	proMap, err := pSvc.QueryProductMap(isPlatform, deviceName)
	if err != nil {
		return nil, nil, err
	}

	//输入设备名称查询
	queryProductIds := make([]int64, 0)
	queryProductIdMap := make(map[int64]int64)
	if deviceName == "" {
		return proMap, queryProductIds, nil
	}
	for _, product := range proMap {
		if strings.Index(product.Name, deviceName) != -1 {
			if _, ok := queryProductIdMap[product.Id]; !ok {
				queryProductIdMap[product.Id] = product.Id
			}
		}
	}
	for _, v := range queryProductIdMap {
		queryProductIds = append(queryProductIds, v)
	}
	return proMap, queryProductIds, nil
}

// QueryIotDeviceInfoList 设备信息列表
func (s *IotDeviceInfoService) QueryIotDeviceInfoList(filter entitys.IotDeviceInfoQuery, setPlatform func(string) string) ([]*entitys.IotDeviceInfoEntitys, int64, *protosService.IotDeviceTriadListRequest, error) {
	resultList := make([]*entitys.IotDeviceInfoEntitys, 0) //返回数据结构
	//获取租户编号集合
	var tenantIds []string
	if filter.Query.Developer != "" {
		companyUserSvc := services2.OpenUserService{Ctx: s.Ctx}
		tenantIds, _ = companyUserSvc.GetUserCompanyTenantIds(filter.Query.Developer)
		if tenantIds == nil || len(tenantIds) == 0 {
			return resultList, 0, nil, nil
		}
	}

	var err error
	proMap := make(map[int64]*protosService.OpmProduct)
	queryProductIds := make([]int64, 0)
	//获取模糊查询的产品名称
	if filter.SearchKey != "" || filter.Query.DeviceName != "" {
		queryProName := filter.SearchKey
		if filter.SearchKey != "" {
			queryProName = filter.SearchKey
		}
		if filter.Query.DeviceName != "" {
			queryProName = filter.Query.DeviceName
		}
		//输入设备名称对应的产品Id
		proMap, queryProductIds, err = s.getQueryProductIds(filter.IsPlatform, queryProName)
		if err != nil {
			return nil, 0, nil, err
		}
	} else {
		proMap, queryProductIds, err = s.getQueryProductIds(filter.IsPlatform, "")
		if err != nil {
			return nil, 0, nil, err
		}
	}
	//转换请求参数
	requestParams, err := deviceListQueryToRequest(filter)
	if err != nil {
		return nil, 0, nil, err
	}
	requestParams.Query.DeveloperTenantIds = tenantIds
	requestParams.Query.QueryProductIds = queryProductIds
	iotlogger.LogHelper.Infof("准备调用ClientIotDeviceServer.Lists")
	rep, err := rpc.ClientIotDeviceServer.Lists(s.Ctx, requestParams)
	if err != nil {
		iotlogger.LogHelper.Errorf("调用ClientIotDeviceServer.Lists异常, %s", err.Error())
		return nil, 0, nil, err
	}
	if rep.Code != 200 {
		iotlogger.LogHelper.Errorf("调用ClientIotDeviceServer.Lists异常, %s", rep.Message)
		return nil, 0, nil, errors.New(rep.Message)
	}
	if len(rep.Data) == 0 {
		return resultList, rep.Total, nil, nil
	}
	//租户信息查询（获取租户对应的账号和公司名称）
	tenantIdMap := make(map[string]string)
	for _, item := range rep.Data {
		tenantIdMap[item.TenantId] = item.TenantId
	}
	//租户Map数据读取
	tenantMap, _ := services2.OpenCompanyService{Ctx: s.Ctx}.GetCompanyList(iotutil.Keys(tenantIdMap))
	//设备性质
	deviceNature, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_device_nature)

	for _, item := range rep.Data {
		var info *entitys.IotDeviceInfoEntitys
		//已经激活的设备数据，存在deviceInfo数据，激活状态、激活时间等信息
		if item.DeviceInfo != nil {
			info = entitys.IotDeviceInfo_pb2e(item.DeviceInfo)
			info.Did = item.Did
			info.Sn = item.SerialNumber
			//TODO 在线状态的转换，此处需要统一（待优化）
			if item.DeviceInfo.OnlineStatus != 1 {
				info.OnlineStatus = 2
			}
		} else {
			//未激活数据，指定默认值
			info = entitys.IotDeviceInfo_pb2e(&protosService.IotDeviceInfo{
				Did:          item.Did,
				OnlineStatus: 2,
				ActiveStatus: "2",
				Sn:           item.SerialNumber,
			})
		}
		info.DeviceNatureKey = item.DeviceNatureKey
		info.ActiveStatus = iotutil.ToString(item.Status)
		//激活状态处理
		if info.ActiveStatus != "1" {
			info.ActiveStatus = "2"
		}
		//租户信息赋值，显示公司名称和归属开发者账号
		if v, ok := tenantMap[item.TenantId]; ok && v != nil {
			info.CompanyName = v.Name
			info.Account = v.UserName
		}
		info.Id = item.Id
		info.ProductName = item.ProductKey
		info.ProductId = item.ProductId
		info.ProductKey = item.ProductKey
		info.UserName = item.UserName
		info.Passward = item.Passward
		info.BatchId = iotutil.ToInt64(item.BatchId)
		//赋值产品信息（产品名称、设备性质和Wifi标识）
		if item.ProductId != 0 {
			if v, ok := proMap[item.ProductId]; ok {
				info.ProductName = v.Name
				info.DeviceNatureKey = v.AttributeType
				info.WifiFlag = v.WifiFlag
			}
			//采用默认名称
			if info.DeviceName == "" {
				info.DeviceName = info.ProductName
			}
		}
		//设备性质转换为中文返回
		info.DeviceNature = deviceNature.Value(info.DeviceNatureKey)
		//导出次数
		info.ExportCount = item.ExportCount
		if item.CreatedAt != nil {
			info.CreatedAt = item.CreatedAt.AsTime().Unix()
			if item.CreatedAt.AsTime().Unix() < 0 {
				info.CreatedAt = 0
			}
		}
		if item.UpdatedAt != nil {
			info.UpdatedAt = item.UpdatedAt.AsTime().Unix()
			if item.UpdatedAt.AsTime().Unix() < 0 {
				info.UpdatedAt = 0
			}
		}
		info.PlatformCode = item.PlatformCode
		if setPlatform != nil {
			info.PlatformName = setPlatform(item.PlatformCode)
		}
		if item.ExportTimeList != "" {
			//通过逗号分割ExportList
			var timeList = strings.Split(item.ExportTimeList, ",")
			//排除ExportList中的空值字符串
			info.ExportList = iotutil.RemoveEmptyString(timeList)
		}
		resultList = append(resultList, info)
	}
	return resultList, rep.Total, requestParams, err
}

func (s *IotDeviceInfoService) SetExportCount(requestParams *protosService.IotDeviceTriadListRequest) error {
	defer iotutil.PanicHandler()
	if requestParams == nil {
		return nil
	}
	iotlogger.LogHelper.Infof("准备调用SetExportCount")
	rep, err := rpc.ClientIotDeviceServer.SetExportCount(s.Ctx, requestParams)
	if err != nil {
		iotlogger.LogHelper.Errorf("SetExportCount异常, %s", err.Error())
		return err
	}
	if rep.Code != 200 {
		iotlogger.LogHelper.Errorf("调用SetExportCount异常, %s", rep.Message)
		return errors.New(rep.Message)
	}
	return err
}

// AddIotDeviceInfo 新增设备信息
func (s *IotDeviceInfoService) AddIotDeviceInfo(req entitys.IotDeviceInfoEntitys) (string, error) {
	saveObj := entitys.IotDeviceInfo_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.CreatedAt = timestamppb.Now()
	res, err := rpc.ClientIotDeviceInfoServer.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateIotDeviceInfo 修改设备信息
func (s *IotDeviceInfoService) UpdateIotDeviceInfo(req entitys.IotDeviceInfoEntitys) (string, error) {
	req.UpdatedAt = time.Now().Unix()
	res, err := rpc.ClientIotDeviceInfoServer.Update(s.Ctx, entitys.IotDeviceInfo_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteIotDeviceInfo 删除设备信息
func (s *IotDeviceInfoService) DeleteIotDeviceInfo(req entitys.IotDeviceInfoFilter) error {
	rep, err := rpc.ClientIotDeviceInfoServer.Delete(s.Ctx, &protosService.IotDeviceInfo{
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

// QueryIotDeviceCount 设备信息统计
func (s *IotDeviceInfoService) QueryIotDeviceCount(filter entitys.IotDeviceInfoQuery) (*entitys.IotDeviceInfoTotalCount, error) {
	var res = &entitys.IotDeviceInfoTotalCount{
		ActiveTotal: 0,
		DeviceTotal: 0,
		OnlineTotal: 0,
	}
	//获取租户编号集合
	var tenantIds []string
	if filter.Query.Developer != "" {
		companyUserSvc := services2.OpenUserService{Ctx: s.Ctx}
		tenantIds, _ = companyUserSvc.GetUserCompanyTenantIds(filter.Query.Developer)
		if tenantIds == nil || len(tenantIds) == 0 {
			return res, nil
		}
	}
	rep, err := rpc.ClientIotDeviceInfoServer.QueryCount(s.Ctx, &protosService.IotDeviceInfoListRequest{
		Query: &protosService.IotDeviceInfo{DeveloperTenantIds: tenantIds},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	res.ActiveTotal = rep.Data.ActiveTotal
	res.DeviceTotal = rep.Data.DeviceTotal
	res.OnlineTotal = rep.Data.OnlineTotal
	return res, nil
}

// QueryIotDeviceDetails 设备信息列表
func (s *IotDeviceInfoService) QueryIotDeviceDetails(lang, did string) (*entitys.IotDeviceInfoDetails, error) {
	rep, err := rpc.ClientIotDeviceInfoServer.QueryDetails(s.Ctx, &protosService.IotDeviceInfoFilter{
		Did: did,
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	var tls *protosService.OpmThingModelAllList
	if rep.Data.DeviceInfo.ProductId != 0 {
		//获取产品的物模型
		//TODO 需要另增加一个接口， ALLDetail内容较多，资源浪费
		proRes, err := rpc.ClientOpmProductService.FindByAllDetails(s.Ctx, &protosService.OpmProductPrimarykey{
			Id: rep.Data.DeviceInfo.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if rep.Code != 200 {
			return nil, errors.New(rep.Message)
		}
		rep.Data.DeviceInfo.ProductName = proRes.Product.Name
		if rep.Data.DeviceInfo.DeviceName == "" {
			rep.Data.DeviceInfo.DeviceName = proRes.Product.Name
		}
		rep.Data.DeviceInfo.ProductKey = proRes.Product.ProductKey
		if proRes.Module != nil {
			rep.Data.DeviceInfo.FirmwallKey = proRes.Module.FirmwareKey
		}
		//mcu固件Key获取
		if proRes.CustomFirmwares != nil && rep.Data.DeviceInfo.McuFirmwallKey == "" {
			for _, cf := range proRes.CustomFirmwares {
				if cf.FirmwareType == iotconst.FIRMWARE_TYPE_MCU {
					rep.Data.DeviceInfo.McuFirmwallKey = cf.FirmwareKey
				}
			}
		}
		tls = proRes.ThingModes
	}

	firmwareList := make([]*entitys.DeviceFirmwares, 0)
	req, err := rpc.ClientOpmProductService.FindByAllDetails(s.Ctx, &protosService.OpmProductPrimarykey{Id: rep.Data.DeviceInfo.ProductId})
	if err != nil {
		return nil, err
	}
	if req.Code == 200 {
		for _, v := range req.CustomFirmwares {
			firmwareList = append(firmwareList, &entitys.DeviceFirmwares{
				Name:    v.FirmwareName,
				Type:    v.FirmwareType,
				Key:     v.FirmwareKey,
				Version: v.Version,
			})
		}
	}

	var res = &entitys.IotDeviceInfoDetails{
		ActiveInfo:   entitys.IotDeviceInfo_activeInfo_pb2e(rep.Data.ActiveInfo),
		DeviceInfo:   entitys.IotDeviceInfo_deviceInfo_pb2e(rep.Data.DeviceInfo),
		DeviceStatus: convertDeviceStatus(rep.Data.DeviceInfo.TenantId, lang, rep.Data.DeviceInfo.ProductKey, rep.Data.DeviceStatus, tls),
		FirmwareList: firmwareList,
	}
	if rep.Data.DeviceInfo != nil {
		developer, _ := new(services.DeveloperCachedData).GetByTenantId(rep.Data.DeviceInfo.TenantId)
		if developer != nil {
			res.ActiveInfo.BelogOpenUser = developer.Name
			//res.Account = developer.UserName
		}
	}

	if res.ActiveInfo.ActiveApp != "" {
		app, err := extract.GetAppInfo(s.Ctx, res.ActiveInfo.ActiveApp)
		if err == nil && app != nil {
			res.ActiveInfo.ActiveApp = app.Name
		}
	}
	return res, nil
}

func convertDeviceStatus(tenantId, lang, productKey string, deviceStatus map[string]string, tls *protosService.OpmThingModelAllList) []*entitys.IotDeviceInfoStatus {
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, _ := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	if lang == "" {
		lang = "zh"
	}

	tslMap := map[string]*protosService.OpmThingModelProperties{}
	if tls != nil && tls.Properties != nil {
		for _, tslObj := range tls.Properties {
			tslMap[iotutil.ToString(tslObj.Dpid)] = tslObj
		}
	}

	list := make([]*entitys.IotDeviceInfoStatus, 0)
	for k, v := range deviceStatus {
		obj := &entitys.IotDeviceInfoStatus{
			AttrKey:       k,
			AttrKeyName:   k,
			AttrValue:     v,
			AttrValueName: v,
		}
		if val, ok := tslMap[k]; ok {
			obj.AttrKeyName = val.Name
			dataType := val.DataType
			//数值转换（BOOL类型特殊处理）
			if dataType == "BOOL" {
				if obj.AttrValue == "1" || obj.AttrValue == "true" {
					obj.AttrValue = "true"
				} else if obj.AttrValue == "0" {
					obj.AttrValue = "false"
				}
			}
			langKey := fmt.Sprintf("%s_%s_%s_%v_name", lang, productKey, val.Identifier, obj.AttrValue)
			obj.AttrValueName = iotutil.MapGetStringVal(langMap[langKey], obj.AttrValue)
			list = append(list, obj)
		}
	}

	//排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].AttrKey < list[j].AttrKey
	})
	return list
}

// Export 导出
func (s *IotDeviceInfoService) Export(mode int, filter entitys.IotDeviceInfoQuery) (string, string, error) {
	if filter.Query.StartTime == "" || filter.Query.EndTime == "" {
		return "", "", errors.New("时间参数不能为空")
	}
	//excel 样式文件
	headerStyle := common.ExcelHeaderStyle()
	contentStyle := common.ExcelContentStyle()

	//导出生成excel附件
	//timeTemplate := "2006-01-02 15:04:00"
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")

	sheet.SetColWidth(0, 0, 20)
	sheet.SetColWidth(1, 3, 10)
	sheet.SetColWidth(4, 5, 18)
	sheet.SetColWidth(6, 7, 10)
	sheet.SetColWidth(8, 8, 16)
	sheet.SetColWidth(10, 10, 14)
	sheet.SetColWidth(11, 11, 20)

	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备Id"

	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "产品Key"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "WIFI标识"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "用户名"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "密码"

	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备名称"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "是否绑定"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "是否在线"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "所属产品"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备性质"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备SN"
	if mode == 1 {
		cell = headerRow.AddCell()
		cell.SetStyle(headerStyle)
		cell.Value = "开发者"
		filter.IsPlatform = true
	}
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "首次激活时间"

	res, _, queryParams, err := s.QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		return "", "", err
	}
	deviceNature, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_device_nature)
	activeStatus, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_active_status)
	onlineStatus, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_online_status)
	for _, row := range res {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Did

		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.ProductKey
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.WifiFlag
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.UserName
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Passward

		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.DeviceName
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = activeStatus.ValueStr(row.ActiveStatus)
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = onlineStatus.Value(row.OnlineStatus)
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.ProductName
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = deviceNature.Value(row.DeviceNatureKey)
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Sn
		if mode == 1 {
			cell = headerRow.AddCell()
			cell.SetStyle(contentStyle)
			cell.Value = row.CompanyName + "\r\n" + row.Account //row.UserName
		}
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		if row.ActivatedTime != 0 {
			cell.Value = iotutil.TimeFullFormat(time.Unix(row.ActivatedTime, 0)) //row.UserName
		}
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
	fileName := "deviceInfo-" + time.Now().Format("20060102150400") + ".xlsx"

	//调用记录到处统计数据
	if filter.Query.IsQueryTriadData {
		go s.SetExportCount(queryParams)
	}

	return fileName, tempPathFile, nil
}

// Export 导出
func (s *IotDeviceInfoService) ExportTriad(userId int64, filter entitys.IotDeviceInfoQuery) (string, string, error) {
	ocs := services2.OpenCompanyService{Ctx: context.Background()}
	companyInfo, err := ocs.GetBaseInfo(userId)
	if err != nil {
		return "", "", errors.New("公司信息获取失败")
	}

	if filter.Query.StartTime == "" || filter.Query.EndTime == "" {
		return "", "", errors.New("时间参数不能为空")
	}
	//excel 样式文件
	headerStyle := common.ExcelHeaderStyle()
	contentStyle := common.ExcelContentStyle()
	//导出生成excel附件
	//timeTemplate := "2006-01-02 15:04:00"
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备Id"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "产品Key"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "WIFI标识"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "用户名"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "密码"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "设备SN"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "所属私有云平台"
	cell.SetStyle(headerStyle)
	cell.Value = "公司编码"
	cell.SetStyle(headerStyle)
	cell.Value = "公司名称"

	res, _, _, err := s.QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		return "", "", err
	}

	for _, row := range res {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Did
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.ProductKey
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.WifiFlag
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.UserName
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Passward
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Sn
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = config.Global.Service.PlatformCode
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = companyInfo.TenantId
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = companyInfo.CompanyName
	}

	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err = file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		return "", "", err
	}
	//发送完文件后删除对应文件（延时删除）
	//defer func() {
	//	os.Remove(tempPathFile)
	//}()
	fileName := "deviceTriad-" + time.Now().Format("20060102150400") + ".xlsx"
	return fileName, tempPathFile, nil
}

func (s *IotDeviceInfoService) ExportCsvTriad(userId int64, filter entitys.IotDeviceInfoQuery, setPlatform func(string) string) (string, string, error) {
	ocs := services2.OpenCompanyService{Ctx: context.Background()}
	var (
		tenantId     string
		companyName  string
		platformCode string
	)
	companyInfo, err := ocs.GetBaseInfo(userId)
	if err != nil {
		return "", "", errors.New("公司信息获取失败")
	}
	tenantId = companyInfo.TenantId
	companyName = companyInfo.CompanyName
	platformCode = config.Global.Service.PlatformCode

	if filter.Query.StartTime == "" || filter.Query.EndTime == "" {
		return "", "", errors.New("时间参数不能为空")
	}

	fileName := "deviceTriad-" + time.Now().Format("20060102150400") + ".csv"
	tempPathFile := tempPath + fileName
	file, err := os.Create(tempPathFile)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	//"产品Key", "WIFI标识",
	w.Write([]string{"设备Id", "用户名", "密码", "设备SN", "所属私有云平台编码", "公司编码", "公司名称"})

	res, _, _, err := s.QueryIotDeviceInfoList(filter, setPlatform)
	if err != nil {
		return "", "", err
	}

	for _, row := range res {
		//row.ProductKey, row.WifiFlag,
		w.Write([]string{row.Did, row.UserName,
			row.Passward, row.Sn, platformCode, tenantId, companyName,
		})
	}
	w.Flush()
	return fileName, tempPathFile, nil
}
