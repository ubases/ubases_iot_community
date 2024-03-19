package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstrings"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

type OemAppService struct {
	Ctx context.Context
}

func (s OemAppService) SetContext(ctx context.Context) OemAppService {
	s.Ctx = ctx
	return s
}

// 创建oemapp
func (s OemAppService) AddOemApp(req entitys.OemAppEntitysAddReq) (map[string]string, error) {
	var resMap = make(map[string]string, 0)
	id := iotutil.GetNextSeqInt64()
	version := "1.0.0"
	req.AppType = 1
	req.AppTemplateId = "0"
	req.AppTemplateVersion = "1.0.0"
	res, err := rpc.ClientOemAppService.Create(s.Ctx, &protosService.OemApp{
		Id:                 id,
		Name:               req.Name,
		Version:            version,
		Status:             1,
		Channel:            req.Channel,
		IosPkgName:         req.IosPkgName,
		IosTeamId:          req.IosTeamId,
		AndroidPkgName:     req.AndroidPkgName,
		Region:             req.Region,
		AppType:            req.AppType,
		AppTemplateId:      iotutil.ToInt64(req.AppTemplateId),
		AppTemplateVersion: req.AppTemplateVersion,
		AppDevType:         req.AppDevType,
		AppIconUrl:         req.AppIconUrl,
	})
	if err != nil {
		return resMap, err
	}
	if res.Code != 200 {
		return resMap, errors.New(res.Message)
	}

	resMap["appId"] = iotutil.ToString(id)
	resMap["version"] = version

	return resMap, err
}

// 变更名称 TODO  具体详细逻辑未实现
func (s OemAppService) ChangeOemAppName(req entitys.OemAppChangeNameReq) (string, error) {
	reqApp := &protosService.OemAppUpdateFieldsRequest{}
	// Oem App和自定义app更新的时候字段不一致，需做兼容处理
	if req.AppIconUrl == "" {
		reqApp.Fields = []string{"name", "name_en"}
		reqApp.Data = &protosService.OemApp{
			Name:   req.Name,
			NameEn: req.NameEn,
			Id:     iotutil.ToInt64(req.Id),
		}
	} else {
		reqApp.Fields = []string{"name", "app_icon_url"}
		reqApp.Data = &protosService.OemApp{
			Name:       req.Name,
			AppIconUrl: req.AppIconUrl,
			Id:         iotutil.ToInt64(req.Id),
		}
	}
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, reqApp)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 变更名称 TODO  具体详细逻辑未实现
func (s OemAppService) UpdateOemAppTemplate(req entitys.OemAppUpdateTemplateReq) (string, error) {
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"app_type", "app_template_id", "app_template_version", "status"},
		Data: &protosService.OemApp{
			Id:                 iotutil.ToInt64(req.Id),
			AppType:            req.AppType,
			AppTemplateId:      iotutil.ToInt64(req.AppTemplateId),
			AppTemplateVersion: req.AppTemplateVersion,
			Status:             1,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// oemapp 列表
func (s OemAppService) QueryOemAppList(filter entitys.OemAppQuery, tenantId string) ([]*entitys.OemAppEntityListRes, int64, error) {
	rep, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.OemApp{
			TenantId:   tenantId,
			Status:     filter.Status,
			AppDevType: filter.AppDevType,
			IsDefault:  filter.IsDefault,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var resultList = []*entitys.OemAppEntityListRes{}
	for _, item := range rep.Data {

		tmp := entitys.OemApp_pb2eList(item)
		resIcon, _ := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
			AppId:   item.Id,
			Version: item.Version,
		})
		if resIcon != nil && len(resIcon.Data) > 0 {
			tmp.IocnUrl = resIcon.Data[0].IconUrl
		}
		resultList = append(resultList, tmp)
	}
	return resultList, rep.Total, err
}

// 获取app详细
func (s OemAppService) QueryOemAppDetail(id string, tenantId string) (*entitys.OemAppEntityDetailRes, error) {
	res, err := rpc.ClientOemAppService.Find(s.Ctx, &protosService.OemAppFilter{
		Id:       iotutil.ToInt64(id),
		TenantId: tenantId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		if res.Message == "record not found" {
			return nil, errors.New("app记录不存在")
		}
		return nil, errors.New(res.Message)
	}
	result := entitys.OemApp_pb2eDetail(res.Data[0])
	//TODO 需要从构建表的逻辑来决定此数据
	//此字段暂时弃用
	result.BuildStatus = 1

	resIcon, _ := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   res.Data[0].Id,
		Version: res.Data[0].Version,
	})
	if resIcon != nil && len(resIcon.Data) > 0 {
		result.IconUrl = resIcon.Data[0].IconUrl
	}

	return result, err
}

// 删除oem app
func (s OemAppService) DeleteOemApp(id string) (string, error) {
	res, err := rpc.ClientOemAppService.DeleteById(s.Ctx, &protosService.OemApp{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	return "success", err
}

// 更新oemapp当前操作步骤 TODO  具体详细逻辑未实现
func (s OemAppService) ChangeOemAppCurrentStep(req entitys.OemAppChangeCurrentStepReq) (string, error) {
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"current_step"},
		Data: &protosService.OemApp{
			Id:          iotutil.ToInt64(req.Id),
			CurrentStep: req.CurrentStep,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 更新最后构建时间
func (s OemAppService) UpdateLastBuildTime(id int64) (string, error) {
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"last_build_time"},
		Data: &protosService.OemApp{
			Id:            id,
			LastBuildTime: timestamppb.New(time.Now()),
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 更新最后证书更新时间
func (s OemAppService) UpdateLastCertUpdateTime(id int64, os int) (string, error) {
	fields := make([]string, 0)
	data := &protosService.OemApp{Id: id}
	switch os {
	case 1:
		fields = append(fields, "last_ios_cert_update_time")
		data.LastIosCertUpdateTime = timestamppb.New(time.Now())
	default:
		fields = append(fields, "last_android_cert_update_time")
		data.LastAndroidCertUpdateTime = timestamppb.New(time.Now())
	}
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: fields,
		Data:   data,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

func (s OemAppService) UpdateLastCertUpdateTimeAndTeamId(id int64, os int, teamId string) (string, error) {
	fields := make([]string, 0)
	data := &protosService.OemApp{Id: id}
	switch os {
	case 1:
		fields = append(fields, "last_ios_cert_update_time")
		data.LastIosCertUpdateTime = timestamppb.New(time.Now())
	default:
		fields = append(fields, "last_android_cert_update_time")
		data.LastAndroidCertUpdateTime = timestamppb.New(time.Now())
	}
	if teamId == "" {
		fields = append(fields, "ios_team_id")
		data.IosTeamId = teamId
	}
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: fields,
		Data:   data,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 更新最后证书更新时间
func (s OemAppService) UpdateUIConfigUpdateTime(id int64) (string, error) {
	fields := make([]string, 0)
	data := &protosService.OemApp{Id: id}
	fields = append(fields, "last_ui_update_time")
	data.LastUiUpdateTime = timestamppb.New(time.Now())
	res, err := rpc.ClientOemAppService.UpdateFields(context.Background(), &protosService.OemAppUpdateFieldsRequest{
		Fields: fields,
		Data:   data,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 保存地图
func (s OemAppService) SaveMap(req entitys.OemAppMap) (string, error) {
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"amap_key", "googlemap_key"},
		Data: &protosService.OemApp{
			Id:           iotutil.ToInt64(req.AppId),
			GooglemapKey: req.GooglemapKey,
			AmapKey:      req.AmapKey,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// 获取地图
func (s OemAppService) GetMap(req entitys.OemAppCommonReq) (*entitys.OemAppMap, error) {
	res, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id:      iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppMap{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppMap{}
	data.AmapKey = res.Data[0].AmapKey
	data.GooglemapKey = res.Data[0].GooglemapKey
	data.AppId = req.AppId

	return &data, err
}

// CheckMap 检查是否配置的地图参数
func (s OemAppService) CheckMap(req entitys.OemAppCommonReq) (bool, error) {
	res, err := s.GetMap(req)
	if err != nil {
		return false, err
	}
	if res.AmapKey == "" && res.GooglemapKey == "" {
		return false, nil
	}
	return true, nil
}

// 根据数据字典类型.获取数据值
func GetBaseDataValue(dictType string, ctydiy context.Context) map[string]interface{} {
	res, err := rpc.TConfigDictDataServerService.Lists(ctydiy, &protosService.ConfigDictDataListRequest{
		Query: &protosService.ConfigDictData{
			DictType: dictType,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	if res == nil || res.Code != 200 {
		if res != nil {
			iotlogger.LogHelper.Errorf(res.Message)
		} else {
			iotlogger.LogHelper.Errorf("未获取到参数")
		}
		return nil
	}
	var dicMap = make(map[string]interface{})
	for _, v := range res.Data {
		dicMap[v.DictLabel] = v.DictValue
	}
	return dicMap
}

var OemAppEnv string

func GetOemAppEnv() string {
	if OemAppEnv == "" {
		OemAppEnv = os.Getenv("ENVIRONMENT")
	}
	return OemAppEnv
}

// 获取所有构建跑的二维码链接
func (s OemAppService) OemAppBuildPackageQrCodeUrl(req entitys.OemAppCommonReq) (string, error) {
	mp := GetBaseDataValue("oem_app_package_domain", s.Ctx)
	//http://127.0.0.1:8080/v1/platform/web/open/oem
	url := iotutil.ToString(mp[GetOemAppEnv()])
	url += "/app/build/package/qrcode?appId=" + req.AppId + "&version=" + req.Version
	return url, nil
}

// 获取自定义app的二维码链接
func (s OemAppService) OemAppCustomPackageQrCodeUrl(req entitys.OemAppCommonReq) (string, error) {
	mp := GetBaseDataValue("oem_app_package_domain", s.Ctx)
	//http://127.0.0.1:8080/v1/platform/web/open/oem
	url := iotutil.ToString(mp[GetOemAppEnv()])
	url += "/app/custom/package/qrcode?appId=" + req.AppId + "&version=" + req.Version + "&os=" + iotutil.ToString(req.Os)
	return url, nil
}

func (s OemAppService) GetAppInfo(appId int64) (*protosService.OemApp, error) {
	res, err := rpc.ClientOemAppService.Find(s.Ctx, &protosService.OemAppFilter{
		Id: appId,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res.Data[0], nil
}

func (s OemAppService) GetAppMap(tenantId string) (map[int64]*protosService.OemApp, error) {
	defaultMap := make(map[int64]*protosService.OemApp)
	res, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Query: &protosService.OemApp{TenantId: tenantId}})
	//
	if err != nil && err.Error() != "record not found" {
		return defaultMap, err
	}
	if res.Code != 200 {
		return defaultMap, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return defaultMap, errors.New("参数错误,未找到记录")
	}

	resMap := make(map[int64]*protosService.OemApp)
	for _, d := range res.Data {
		resMap[d.Id] = d
	}
	return resMap, nil
}

func (s OemAppService) GetAppKeyMap(tenantId string) (map[string]*protosService.OemApp, error) {
	defaultMap := make(map[string]*protosService.OemApp)
	res, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Query: &protosService.OemApp{TenantId: tenantId}})
	//
	if err != nil && err.Error() != "record not found" {
		return defaultMap, err
	}
	if res.Code != 200 {
		return defaultMap, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return defaultMap, errors.New("参数错误,未找到记录")
	}

	resMap := make(map[string]*protosService.OemApp)
	for _, d := range res.Data {
		resMap[d.AppKey] = d
	}
	return resMap, nil
}

func (s OemAppService) OemAppBuildPackage(req entitys.OemAppCommonReq) ([]*entitys.OemAppBuildPackage, error) {
	appId := iotutil.ToInt64(req.AppId)
	//获取APP信息
	app, err := s.GetAppInfo(appId)
	if err != nil {
		return nil, err
	}
	var hasUpdate bool
	//最后IOS证书更新时间
	if app.LastIosCertUpdateTime != nil && app.LastBuildTime != nil {
		if app.LastBuildTime.AsTime().Unix() < app.LastIosCertUpdateTime.AsTime().Unix() {
			hasUpdate = true
		}
	}
	//最后安卓证书更新时间
	if app.LastAndroidCertUpdateTime != nil && app.LastBuildTime != nil {
		if app.LastBuildTime.AsTime().Unix() < app.LastAndroidCertUpdateTime.AsTime().Unix() {
			hasUpdate = true
		}
	}
	//UI修改时间
	if app.LastUiUpdateTime != nil && app.LastBuildTime != nil {
		if app.LastBuildTime.AsTime().Unix() < app.LastUiUpdateTime.AsTime().Unix() {
			hasUpdate = true
		}
	}
	showBuildBtn := hasUpdate

	resCert, err := rpc.ClientOemAppAndroidCertService.Find(s.Ctx, &protosService.OemAppAndroidCertFilter{
		AppId:   appId,
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if resCert.Code != 200 && resCert.Message != gorm.ErrRecordNotFound.Error() {
		return nil, errors.New(resCert.Message)
	}
	if resCert.Data == nil || len(resCert.Data) == 0 {
		return nil, errors.New("证书未配置")
	}
	googleSignCert := resCert.Data[0].GoogleSignCert
	hwSignCert := resCert.Data[0].HwSignCert

	res, err := rpc.ClientOemAppBuildRecordService.Lists(s.Ctx, &protosService.OemAppBuildRecordListRequest{
		OrderKey:  "start_time",
		OrderDesc: "desc",
		Query: &protosService.OemAppBuildRecord{
			AppId:   iotutil.ToInt64(req.AppId),
			Version: req.Version,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != gorm.ErrRecordNotFound.Error() {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = make([]*entitys.OemAppBuildPackage, 0)
		return nodata, nil
	}
	//构建记录 每个平台只取一条记录[后续改用sql方式查询,先按os分组.然后获取最新的start_time第一条记录]
	//只获取构建每个平台最新的那条构建记录
	var mp = s.initBuildPackage(res.Data)
	var list = make([]*entitys.OemAppBuildPackage, 0)
	//是否全部超时
	buildingAndTimeoutRecordMap := map[int64]int32{}
	for _, v := range mp {
		list = append(list, s.setBuildPackage(showBuildBtn, hasUpdate, googleSignCert, hwSignCert, v, buildingAndTimeoutRecordMap))
	}
	//检查并更新app状态
	if app.Status == 2 {
		s.SetAppTimeoutStatus(appId, buildingAndTimeoutRecordMap)
	}
	//排序
	if len(list) > 1 {
		s.SortBuildPackage(list)
	}
	return list, nil
}

// 初始化生成报列表
func (s OemAppService) initBuildPackage(req []*protosService.OemAppBuildRecord) map[int32]*protosService.OemAppBuildRecord {
	res := make(map[int32]*protosService.OemAppBuildRecord)
	osArr := []int32{1, 2, 3}
	for _, os := range osArr {
		/**
		 	"cause": "用户主动取消构建",
			"os": 1,
			"status": 4,
			"message": "",
			"timeSurplus": 120,
			"url": "",
			"urlMd5": "",
			"urlAab": "",
			"urlExt": "",
			"urlAabExt": "",
			"urlValidity": 0,
			"qrCodeUrl": "",
			"urlIos": "",
			"signCerts": [],
			"buildId": "8825075850152935424",
			"hasUpdate": false,
			"showBuildBtn": true
		*/
		res[os] = &protosService.OemAppBuildRecord{
			BuildId: "0",
			Os:      os,
			Status:  1,
		}
	}
	for _, vv := range req {
		v := res[vv.Os]
		if v.BuildId == "0" {
			res[vv.Os] = vv
		}
	}
	return res
}

func (s OemAppService) setBuildPackage(showBuildBtn, hasUpdate bool, googleSignCert, hwSignCert string,
	v *protosService.OemAppBuildRecord, buildingAndTimeoutRecordMap map[int64]int32) *entitys.OemAppBuildPackage {
	var (
		ext    = ""
		extAab = ""
		url    = ""
		urlAab = ""
		urlIos = "" //ios正式下载包
	)
	if v.PkgUrl != "" {
		arrTmp := make([]string, 0)
		iotutil.JsonToStruct(v.PkgUrl, &arrTmp)
		url = arrTmp[0]
		if v.Os == iotconst.OS_IOS {
			ext = ".ipa"
			//如果长度等于三, 则第三个是正式包下载地址.
			if len(arrTmp) == 3 {
				//url = arrTmp[2]
				urlIos = arrTmp[2]
			}
		} else if v.Os == iotconst.OS_ANDROID_CHINA {
			ext = ".apk"
		}
		//如果aab安装包不为空.
		if arrTmp[1] != "" {
			if v.Os == 1 {
				extAab = ".plist"
				ext = ".ipa"
			} else {
				extAab = ".aab"
				ext = ".apk"
			}
			urlAab = arrTmp[1]
		}
	}
	//构建按钮显示要求
	if v.Status == 3 || v.Status == 4 {
		showBuildBtn = true
	}

	tmp := entitys.OemAppBuildPackage{
		Os:           v.Os,
		Status:       int64(v.Status),
		Cause:        v.BuildResultMsg, //失败原因
		TimeSurplus:  120,
		Url:          url,
		UrlAab:       urlAab,
		UrlExt:       ext,
		UrlAabExt:    extAab,
		UrlIos:       urlIos,
		BuildId:      iotutil.ToString(v.Id),
		HasUpdate:    hasUpdate,
		ShowBuildBtn: showBuildBtn,
	}
	if tmp.Url != "" {
		dd, _ := time.ParseDuration("168h")
		tmp.UrlValidity = v.EndTime.AsTime().Add(dd).Unix() //time.Now().Add(dd).Unix()
	}

	tmp.SignCerts = make([]*entitys.SignCertItems, 0)
	switch tmp.Os {
	case iotconst.OS_IOS:
	case iotconst.OS_ANDROID:
		if googleSignCert != "" {
			tmp.SignCerts = append(tmp.SignCerts, &entitys.SignCertItems{Type: iotconst.APP_GOOGLE, SignCertUrl: googleSignCert})
		}
	case iotconst.OS_ANDROID_CHINA:
		if hwSignCert != "" {
			tmp.SignCerts = append(tmp.SignCerts, &entitys.SignCertItems{Type: iotconst.APP_HAIWEI, SignCertUrl: hwSignCert})
		}
	}
	//检查，构建状态为构建中，结束时间小于，如果APP构建已经超时，则将状态修复为失败
	if v.Status == 2 {
		if v.EndTime.AsTime().Before(time.Now()) {
			s.SetRecordTimeout(v, &tmp)
			//构建中，并且超时
			buildingAndTimeoutRecordMap[v.Id] = 1 //构建中已超时
		} else {
			//构建中，并且未超时
			buildingAndTimeoutRecordMap[v.Id] = 2 //构建未超时
		}
	} else {
		if v.Status == 3 {
			buildingAndTimeoutRecordMap[v.Id] = 3 //成功
		} else if v.Status == 1 {
			buildingAndTimeoutRecordMap[v.Id] = 0 //未选择的构建
		} else {
			buildingAndTimeoutRecordMap[v.Id] = 4 //失败
		}
	}
	return &tmp
}

// 设置APP构建状态
func (s OemAppService) SetAppTimeoutStatus(appId int64, recordMap map[int64]int32) {
	//恢复APP的状态（有2不能改，已改情况不修改
	//全部为0，app状态为构建中，修改为成功
	//部分为0，部分为1，app状态为构建中，修改为失败
	//部分为0，部分为1，部分为2，app状态为构建中，不修改
	has0 := false
	has1 := false
	has2 := false
	for _, v := range recordMap {
		switch v {
		case 0:
			has0 = true
		case 1:
			has1 = true
		case 2:
			has2 = true
		}
	}
	if has0 && !has1 && !has2 {
		//全部为0，app状态为构建中，修改为成功
		s.OemAppUpdateStatus(appId, 1)
	} else if has0 && has1 && has2 {
		//不修改
	} else if has0 && has1 && !has2 {
		//部分为0，部分为1，app状态为构建中，修改为失败
		s.OemAppUpdateStatus(appId, 3)
	}
}

// 设置超时状态
func (s OemAppService) SetRecordTimeout(v *protosService.OemAppBuildRecord, tmp *entitys.OemAppBuildPackage) {
	var req entitys.OemAppBuildFinishNotifyReq
	req.BuildId = iotutil.ToString(v.Id)
	req.BuildProgress = 100
	req.BuildResult = 2
	//req.EndTime = time.Now().Unix()
	req.CommitID = v.CommitId
	req.Status = 4
	req.BuildResultMsg = "超时自动取消"
	pkgUrl := ""
	//存放aab或是plist
	pkgAabOrPlistUrl := ""
	//ipa正式包
	pkgIpaUrl := ""
	//arrPkgUrl 说明
	//android 国内 apk安装包和空字符串和空字符串
	//Android 海外 apk安装包和aab安装包和空字符串
	//ios  ipa安装包和plist文件和ios正式安装包
	var arrPkgUrl = []string{pkgUrl, pkgAabOrPlistUrl, pkgIpaUrl}
	req.PkgURL = iotutil.ToStringByUrl(arrPkgUrl)
	serviceBuild := OemAppBuildRecordService{}
	_, err := serviceBuild.SetContext(s.Ctx).BuildFinishNotify(req)
	if err == nil {
		tmp.Status = req.Status
		tmp.Cause = req.BuildResultMsg
	}
}

// SortBuildPackage 按系统排查
func (s OemAppService) SortBuildPackage(list []*entitys.OemAppBuildPackage) {
	sort.Slice(list, func(i, j int) bool { // asc
		return list[i].Os > list[j].Os
	})
}

// OemAppUpdateStatus 上架APP [1.配置中  2.构建中  3.构建完成  4.上架中  5.已上架]
func (s OemAppService) OemAppUpdateStatus(appId int64, status int32) (string, error) {
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OemApp{
			Id:     appId,
			Status: status,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// OemAppPublishing 上架APP
func (s OemAppService) OemAppPublishing(req entitys.OemAppCommonReq) (string, error) {

	var bu = OemAppBuildRecordService{}
	bu.Ctx = s.Ctx
	isSuccess, errBu := bu.IsBuildSuccess(iotutil.ToInt64(req.AppId), req.Version)
	if errBu != nil {
		return "", errBu
	}
	if !isSuccess {
		return "", errors.New("未找到构建成功的安装包")
	}

	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OemApp{
			Id:     iotutil.ToInt64(req.AppId),
			Status: 4,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}

// OemAppPublish 上架APP
func (s OemAppService) OemAppPublish(req entitys.OemAppCommonReq) (string, error) {
	appId, _ := iotutil.ToInt64AndErr(req.AppId)
	// 添加版本号必须要高于最新版本号
	//respApp, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
	//	Id: appId,
	//})
	//if err != nil {
	//	return "", err
	//}
	//appInfo := respApp.Data[0]
	res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OemApp{
			Id:     appId,
			Status: 5,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//
	appSrc := OemAppService{Ctx: s.Ctx}
	buildPackages, err := appSrc.OemAppBuildPackage(req)
	if err != nil {
		return "", err
	}
	for _, pkg := range buildPackages {
		//只有构建成功的包，才需要添加到上架记录中
		if pkg.Status == 3 {
			//批量推送上架列表
			versionSvc := OemAppCustomRecordService{Ctx: s.Ctx}
			pkgUrl := pkg.Url
			switch pkg.Os {
			case iotconst.OS_IOS:
				pkgUrl = pkg.UrlIos
			case iotconst.OS_ANDROID:
				pkgUrl = pkg.UrlAab
			case iotconst.OS_ANDROID_CHINA:
				pkgUrl = pkg.Url
			}
			versionSvc.CreateOemAppCustomRecord(&entitys.OemAppCustomRecordEntitys{
				AppId:       appId,
				Version:     req.Version,
				PkgUrl:      pkgUrl,
				PkgMd5:      pkg.UrlMd5,
				Os:          pkg.Os,
				Status:      1,
				Description: "",
			})
		}
	}
	return "success", err
}

// SetAppVersionRecord 设置APP版本记录
func (s OemAppService) SetAppVersionRecord(appInfo *protosService.OemApp, pkg *entitys.OemAppBuildPackage) error {
	// 通过模板生成plist文件
	var plistUrl string
	if pkg.Os == 1 {
		respUi, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
			AppId:   appInfo.Id,
			Version: "1.0.0",
		})
		if err != nil {
			return err
		}
		m := map[string]string{}
		err = json.Unmarshal([]byte(respUi.Data[0].IosLaunchScreen), &m)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrShouldBindJSON)
		}
		plist := entitys.TemplatePlistEntitys{
			Title:            fmt.Sprintf("%v上架%v", appInfo.Name, appInfo.Version),
			Version:          appInfo.Version,
			BundleIdentifier: appInfo.IosPkgName,
			DisplayImage:     m["displayImage"],
			FullSizeImage:    m["fullSizeImage"],
			SoftwarePackage:  pkg.Url,
		}
		plistUrl, err = genPlistFile(plist)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrCloudGenOrUploadPlist)
		}
	}
	req := &protosService.OemAppCustomRecord{}
	req.Id = iotutil.GetNextSeqInt64()
	req.PlistUrl = plistUrl
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err := rpc.ClientOemAppCustomRecordService.Create(s.Ctx, req)
	if err != nil {
		return err
	}
	return err
}

// OemAppUpdateVersion 更新app版本
func (s OemAppService) OemAppUpdateVersion(req entitys.OemAppVersionUpdateReq) (string, error) {
	appId := iotutil.ToInt64(req.AppId)
	resFind, errFind := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: appId,
	})
	if errFind != nil {
		return "", errFind
	}
	oldVersion := resFind.Data[0].Version
	if oldVersion != req.Version {
		return "", errors.New("原版本号参数不正确")
	}

	if !iotstrings.VersionCompared(req.NewVersion, oldVersion) {
		return "", errors.New("更新版本号必须大于原版号")
	}

	res, err := rpc.ClientOemAppService.UpdateVersion(s.Ctx, &protosService.OemAppUpdateVersionReq{
		AppId:      appId,
		OldVersion: oldVersion,
		NewVersion: req.NewVersion,
	})

	// res, err := rpc.ClientOemAppService.UpdateFields(s.Ctx, &protosService.OemAppUpdateFieldsRequest{
	// 	Fields: []string{"status","version","current_step"},
	// 	Data: &protosService.OemApp{
	// 		Id: iotutil.ToInt64(req.AppId),
	// 		Version: req.NewVersion,
	// 		Status: 1,
	// 		CurrentStep: 1,
	// 	},
	// })

	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", err
}
