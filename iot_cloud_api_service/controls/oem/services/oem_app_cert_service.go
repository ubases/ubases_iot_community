package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/crypto/pkcs12"

	"github.com/go-resty/resty/v2"
)

type OemAppCertService struct {
	Ctx context.Context
}

func (s OemAppCertService) SetContext(ctx context.Context) OemAppCertService {
	s.Ctx = ctx
	return s
}

// 获取iosCert
func (s OemAppCertService) GetIosCertByAppIdVersion(appId string, version string) (*protosService.OemAppIosCert, error) {
	res, err := rpc.ClientOemAppIosCertService.Find(s.Ctx, &protosService.OemAppIosCertFilter{
		AppId:   iotutil.ToInt64(appId),
		Version: version,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}

	return res.Data[0], nil
}

// 获取androidCert
func (s OemAppCertService) GetAndroidCertByAppIdVersion(appId string, version string) (*protosService.OemAppAndroidCert, error) {
	res, err := rpc.ClientOemAppAndroidCertService.Find(s.Ctx, &protosService.OemAppAndroidCertFilter{
		AppId:   iotutil.ToInt64(appId),
		Version: version,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}

	return res.Data[0], nil
}

// 获取push cert
func (s OemAppCertService) GetPushCertByAppIdVersion(appId string, version string) (*protosService.OemAppPushCert, int32, error) {
	res, err := rpc.ClientOemAppPushCertService.Find(s.Ctx, &protosService.OemAppPushCertFilter{
		AppId:   iotutil.ToInt64(appId),
		Version: version,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, 0, err
	}
	if res.Code != 200 {
		return nil, 0, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, 1, errors.New("参数错误,未找到记录")
	}

	return res.Data[0], 0, nil
}

// 获取APP信息
func (s OemAppCertService) GetAppInfo(appId string) (*protosService.OemApp, error) {
	res, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: iotutil.ToInt64(appId),
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data[0], nil
}

// 保存iosCert
func (s OemAppCertService) SaveIosCert(req entitys.OemAppIosCertReq) (string, error) {
	//验证iOS证书的有效性
	teamId, err := CheckDistCert(&req)
	if err != nil {
		return "", err
	}

	data, errFind := s.GetIosCertByAppIdVersion(req.AppId, req.Version)
	if errFind != nil {
		return "", errFind
	}
	res, err := rpc.ClientOemAppIosCertService.UpdateFields(s.Ctx, &protosService.OemAppIosCertUpdateFieldsRequest{
		Fields: []string{"dist_provision", "dist_cert", "dist_cert_secret", "dist_cert_official"},
		Data: &protosService.OemAppIosCert{
			DistProvision:    req.DistProvision,  //测试证书,mobileprovision文件
			DistCert:         req.DistCert,       //p12文件
			DistCertSecret:   req.DistCertSecret, //证书密码
			Id:               data.Id,
			DistCertOfficial: req.DistCertOfficial, //正式发布证书,mobileprovision文件
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新证书时间
	svc := OemAppService{Ctx: s.Ctx}
	svc.UpdateLastCertUpdateTimeAndTeamId(iotutil.ToInt64(req.AppId), 1, teamId)
	return "success", err
}

// 获取IosCert
func (s OemAppCertService) GetIosCert(req entitys.OemAppCommonReq) (*entitys.OemAppIosCertRes, error) {
	res, err := rpc.ClientOemAppIosCertService.Find(s.Ctx, &protosService.OemAppIosCertFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil {
		var nodata = entitys.OemAppIosCertRes{}
		nodata.AppId = req.AppId
		nodata.Version = req.Version
		return &nodata, nil
	}
	var data = entitys.OemAppIosCertRes{}
	data.AppId = iotutil.ToString(res.Data[0].AppId)
	data.DistCert = res.Data[0].DistCert
	data.DistProvision = res.Data[0].DistProvision
	data.DistCertSecret = res.Data[0].DistCertSecret
	data.Version = res.Data[0].Version
	data.DistCertOfficial = res.Data[0].DistCertOfficial
	return &data, err
}

// 保存AndroidCert
func (s OemAppCertService) SaveAndroidCert(req entitys.OemAppAndroidCertSaveReq) (string, error) {
	data, errFind := s.GetAndroidCertByAppIdVersion(req.AppId, req.Version)
	if errFind != nil {
		return "", errFind
	}
	data.Resign = req.Resign
	data.CertSha256 = req.CertSha256
	res, err := rpc.ClientOemAppAndroidCertService.UpdateAll(s.Ctx, data)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新证书时间
	svc := OemAppService{Ctx: s.Ctx}
	svc.UpdateLastCertUpdateTime(iotutil.ToInt64(req.AppId), 2)
	return "success", err
}

// 获取AndroidCert
func (s OemAppCertService) GetAndroidCert(req entitys.OemAppCommonReq) (*entitys.OemAppAndroidCertDetailRes, error) {
	res, err := rpc.ClientOemAppAndroidCertService.Find(s.Ctx, &protosService.OemAppAndroidCertFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil {
		var nodata = entitys.OemAppAndroidCertDetailRes{}
		nodata.AppId = req.AppId
		nodata.Version = req.Version
		return &nodata, nil
	}
	var data = entitys.OemAppAndroidCertDetailRes{}
	data.AppId = iotutil.ToString(res.Data[0].AppId)
	data.KsMd5 = res.Data[0].KsMd5
	data.KsSha1 = res.Data[0].KsSha1
	data.KsSha256 = res.Data[0].KsSha256
	data.Resign = res.Data[0].Resign
	data.CertSha256 = res.Data[0].CertSha256
	data.Version = res.Data[0].Version
	data.GoogleSignCert = res.Data[0].GoogleSignCert
	data.HwSignCert = res.Data[0].HwSignCert
	data.Id = iotutil.ToString(res.Data[0].Id)

	return &data, err
}

// 保存iosCert
func (s OemAppCertService) SaveIosPushCert(req entitys.OemAppIosPushCertSaveReq) (string, error) {
	app, err := s.GetAppInfo(req.AppId)
	if err != nil {
		return "", err
	}

	//APP开发类型(1 oem, 2 自定义, 3 sdk)
	if app.AppDevType == iotconst.APP_DEV_TYPE_OEM {
		appCert, errFind0 := s.GetIosCertByAppIdVersion(req.AppId, req.Version)
		if errFind0 != nil {
			return "", errFind0
		}
		//验证push证书的有效性
		if err := CheckPushCert(app.IosPkgName, appCert.GetDistCert(), appCert.GetDistCertSecret(), &req); err != nil {
			return "", err
		}
	}
	data, _, errFind := s.GetPushCertByAppIdVersion(req.AppId, req.Version)
	if errFind != nil {
		return "", errFind
	}
	apns := iotutil.ToStringByUrl(req)
	res, err := rpc.ClientOemAppPushCertService.UpdateFields(s.Ctx, &protosService.OemAppPushCertUpdateFieldsRequest{
		Fields: []string{"apns"},
		Data: &protosService.OemAppPushCert{
			Id:   data.Id,
			Apns: apns,
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

// 获取AndroidCert
func (s OemAppCertService) GetIosPushCert(req entitys.OemAppCommonReq) (*entitys.OemAppIosPushCertDetailRes, error) {
	res, err := rpc.ClientOemAppPushCertService.Find(s.Ctx, &protosService.OemAppPushCertFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil {
		var nodata = entitys.OemAppIosPushCertDetailRes{}
		nodata.AppId = req.AppId
		nodata.Version = req.Version
		return &nodata, nil
	}
	var data = entitys.OemAppIosPushCertDetailRes{}
	iotutil.JsonToStruct(res.Data[0].Apns, &data)
	data.AppId = req.AppId
	data.Version = req.Version
	return &data, err
}

// 保存 Android push cert
func (s OemAppCertService) SaveAndroidPushCert(req entitys.OemAppAndroidPushCertSaveReq) (string, error) {
	appData, err := s.GetAppInfo(req.AppId)
	if err != nil {
		return "", err
	}
	data, _, errFind := s.GetPushCertByAppIdVersion(req.AppId, req.Version)
	if errFind != nil {
		return "", errFind
	}

	var jpush = make(map[string]interface{}, 0)
	jpush["jpushKey"] = req.JpushKey
	jpush["jpushSecret"] = req.JpushSecret
	jpush["apnsProduction"] = true
	jpushJson := iotutil.ToStringByUrl(jpush)

	var fcm = make(map[string]interface{}, 0)
	fcm["fcmId"] = req.FcmId
	fcm["fcmKey"] = req.FcmKey
	fcm["fcmJson"] = req.FcmJson
	fcm["fcmServerJson"] = req.FcmServerJson

	//检查json格式
	if req.FcmJson != "" {
		err = checkJson(req.FcmJson, nil)
		if err != nil {
			return "", err
		}
	}
	if req.FcmServerJson != "" {
		err = checkJson(req.FcmServerJson, nil)
		if err != nil {
			return "", err
		}
	}

	fcmJson := iotutil.ToStringByUrl(fcm)

	var huawei = make(map[string]interface{}, 0)
	huawei["huaweiId"] = req.HuaweiId
	huawei["huaweiClientSecret"] = req.HuaweiClientSecret
	huawei["huaweiSecret"] = req.HuaweiSecret
	huawei["huaweiJson"] = req.HuaweiJson
	//检查华为推送证书的json格式
	if req.HuaweiJson != "" {
		err = checkJson(req.HuaweiJson, nil)
		if err != nil {
			return "", err
		}
	}
	huaweiJson := iotutil.ToStringByUrl(huawei)

	var mi = make(map[string]interface{}, 0)
	mi["miId"] = req.MiId
	mi["miKey"] = req.MiKey
	mi["miSecret"] = req.MiSecret
	mi["miDevChanelId"] = req.MiDevChanelId
	mi["miChannelId"] = req.MiChannelId
	miJson := iotutil.ToStringByUrl(mi)

	var vivo = make(map[string]interface{}, 0)
	vivo["vivoId"] = req.VivoId
	vivo["vivoKey"] = req.VivoKey
	vivo["vivoSecret"] = req.VivoSecret
	vivoJson := iotutil.ToStringByUrl(vivo)

	var oppo = make(map[string]interface{}, 0)
	oppo["oppoId"] = req.OppoId
	oppo["oppoKey"] = req.OppoKey
	oppo["oppoSecret"] = req.OppoSecret
	oppo["oppoMasterSecret"] = req.OppoMasterSecret
	oppo["oppoChanelId"] = req.OppoChanelId
	oppo["oppoChanelName"] = req.OppoChanelName
	oppoJson := iotutil.ToStringByUrl(oppo)

	var honor = make(map[string]interface{}, 0)
	honor["honorAppId"] = req.HonorAppId
	honor["honorAppSecret"] = req.HonorAppSecret
	honor["honorClientId"] = req.HonorClientId
	honor["honorClientSecret"] = req.HonorClientSecret
	honorJson := iotutil.ToStringByUrl(honor)

	res, err := rpc.ClientOemAppPushCertService.UpdateFields(s.Ctx, &protosService.OemAppPushCertUpdateFieldsRequest{
		Fields: []string{"jpush", "fcm", "huawei", "xiaomi", "vivo", "oppo", "honor"},
		Data: &protosService.OemAppPushCert{
			Id:     data.Id,
			Jpush:  jpushJson,
			Fcm:    fcmJson,
			Huawei: huaweiJson,
			Xiaomi: miJson,
			Vivo:   vivoJson,
			Oppo:   oppoJson,
			Honor:  honorJson,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//appCachedCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_APPPUSH_DATA_PREFIX+appData.AppKey, jpush)
	//if appCachedCmd.Err() != nil {
	//	iotlogger.LogHelper.Errorf("app推送数据缓存异常，%s", appCachedCmd.Err().Error())
	//}
	iotredis.GetClient().Del(context.Background(), iotconst.HKEY_APPPUSH_DATA_PREFIX+appData.AppKey)
	return "success", err
}

// checkJson 检查是否json文件，检查json中的格式内容
func checkJson(url string, checkFormat func(data map[string]interface{}) error, fields ...string) error {
	jsonByte, err := DownloadFile(url)
	if err != nil {
		return err
	}
	if jsonByte == nil {
		return err
	}
	jsonMap, jsonErr := iotutil.IsJSON(jsonByte)
	if jsonErr == nil {
		return nil
	}
	if checkFormat != nil {
		if err = checkFormat(jsonMap); err != nil {
			return err
		}
	}
	for _, field := range fields {
		if _, ok := jsonMap[field]; !ok {
			return errors.New(field + " not exists")
		}
	}
	return jsonErr
}

func (s OemAppCertService) ConvertPushCertDetailRes(data *protosService.OemAppPushCert) *entitys.OemAppAndroidPushCertDetailRes {
	var res = entitys.OemAppAndroidPushCertDetailRes{
		OppoChanelId: "DeviceAccount", //默认值
	}
	res.AppId = iotutil.ToString(data.AppId)
	res.Version = data.Version
	if data.Jpush != "" {
		jpush := iotutil.JsonToMap(data.Jpush)
		if jpush != nil {
			res.JpushKey = iotutil.MapGetStringVal(jpush["jpushKey"], "")
			res.JpushSecret = iotutil.MapGetStringVal(jpush["jpushSecret"], "")
		}
	}

	if data.Fcm != "" {
		fcm := iotutil.JsonToMap(data.Fcm)
		if fcm != nil {
			res.FcmId = iotutil.MapGetStringVal(fcm["fcmId"], "")
			res.FcmKey = iotutil.MapGetStringVal(fcm["fcmKey"], "")
			res.FcmJson = iotutil.MapGetStringVal(fcm["fcmJson"], "")
			res.FcmServerJson = iotutil.ToString(fcm["fcmServerJson"])
		}
	}

	if data.Huawei != "" {
		huawei := iotutil.JsonToMap(data.Huawei)
		if huawei != nil {
			res.HuaweiId = iotutil.MapGetStringVal(huawei["huaweiId"], "")
			res.HuaweiClientSecret = iotutil.MapGetStringVal(huawei["huaweiClientSecret"], "")
			res.HuaweiSecret = iotutil.MapGetStringVal(huawei["huaweiSecret"], "")
			res.HuaweiJson = iotutil.MapGetStringVal(huawei["huaweiJson"], "")
		}
	}

	if data.Xiaomi != "" {
		mi := iotutil.JsonToMap(data.Xiaomi)
		if mi != nil {
			res.MiId = iotutil.MapGetStringVal(mi["miId"], "")
			res.MiKey = iotutil.MapGetStringVal(mi["miKey"], "")
			res.MiSecret = iotutil.MapGetStringVal(mi["miSecret"], "")
			res.MiDevChanelId = iotutil.MapGetStringVal(mi["miDevChanelId"], "")
			res.MiChannelId = iotutil.MapGetStringVal(mi["miChannelId"], "")
		}
	}

	if data.Vivo != "" {
		vivo := iotutil.JsonToMap(data.Vivo)
		if vivo != nil {
			res.VivoId = iotutil.MapGetStringVal(vivo["vivoId"], "")
			res.VivoKey = iotutil.MapGetStringVal(vivo["vivoKey"], "")
			res.VivoSecret = iotutil.MapGetStringVal(vivo["vivoSecret"], "")
		}
	}

	if data.Oppo != "" {
		oppo := iotutil.JsonToMap(data.Oppo)
		if oppo != nil {
			res.OppoId = iotutil.MapGetStringVal(oppo["oppoId"], "")
			res.OppoKey = iotutil.MapGetStringVal(oppo["oppoKey"], "")
			res.OppoSecret = iotutil.MapGetStringVal(oppo["oppoSecret"], "")
			res.OppoMasterSecret = iotutil.MapGetStringVal(oppo["oppoMasterSecret"], "")
			res.OppoChanelId = iotutil.MapGetStringVal(oppo["oppoChanelId"], "DeviceAccount")
			res.OppoChanelName = iotutil.MapGetStringVal(oppo["oppoChanelName"], "")
		}
	}

	if data.Honor != "" {
		oppo := iotutil.JsonToMap(data.Honor)
		if oppo != nil {
			res.HonorAppId = iotutil.MapGetStringVal(oppo["honorAppId"], "")
			res.HonorAppSecret = iotutil.MapGetStringVal(oppo["honorAppSecret"], "")
			res.HonorClientId = iotutil.MapGetStringVal(oppo["honorClientId"], "")
			res.HonorClientSecret = iotutil.MapGetStringVal(oppo["honorClientSecret"], "")
		}
	}
	return &res
}

// 获取 android push cert
func (s OemAppCertService) GetAndroidPushCert(req entitys.OemAppCommonReq) (*entitys.OemAppAndroidPushCertDetailRes, error) {
	data, _, errFind := s.GetPushCertByAppIdVersion(req.AppId, req.Version)
	if errFind != nil && errFind.Error() != "record not found" {
		return nil, errFind
	}
	res := s.ConvertPushCertDetailRes(data)
	return res, nil
}

// 重新生成证书keystore
func (s OemAppCertService) Regenerate(id string) (string, error) {
	res, err := rpc.ClientOemAppAndroidCertService.Regenerate(s.Ctx, &protosService.OemAppAndroidCertFilter{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return "", errors.New(res.Message)
	}

	return "ok", nil
}

func CheckDistCert(req *entitys.OemAppIosCertReq) (string, error) {
	now := time.Now()
	if req.DistCertOfficial == "" {
		return "", nil
	}

	//ios测试证书mobileprovision
	DistProvision, err := DownloadFile(req.DistProvision)
	if err != nil {
		return "", err
	}
	DistProvisionMP, err := parseMobileprovision(DistProvision)
	if err != nil {
		return "", err
	}

	//ios生产证书mobileprovision
	DistCertOfficial, err := DownloadFile(req.DistCertOfficial)
	if err != nil {
		return "", err
	}
	DistCertOfficialMP, err := parseMobileprovision(DistCertOfficial)
	if err != nil {
		return "", err
	}

	//发布证书p12
	DistCert, err := DownloadFile(req.DistCert)
	if err != nil {
		return "", err
	}
	DistCertP12, err := parseP12(DistCert, req.DistCertSecret)
	if err != nil {
		return "", err
	}

	//以下依据规则进行有效性校验

	//1、测试证书是否过期
	if DistProvisionMP.IsExpired(now) {
		return "", fmt.Errorf("测试证书已过期.")
	}

	//2、生产证书是否过期
	if DistCertOfficialMP.IsExpired(now) {
		return "", fmt.Errorf("生产证书已过期.")
	}

	//3、验证发布证书是否有效
	if now.After(DistCertP12.NotAfter) || now.Before(DistCertP12.NotBefore) {
		return "", fmt.Errorf("发布证书已过期.")
	}

	//4、发布证书组织,必须是"Apple Inc."
	/*	isOrganizationOK := false
		for _, v := range DistCertP12.Subject.Organization {
			if v == "Apple Inc." {
				isOrganizationOK = true
				break
			}
		}
		if !isOrganizationOK {
			return fmt.Errorf("发布证书组织单位Apple公司")
		}*/

	//5 发布证书的常用名称必须是"Apple Distribution:"开头
	if !strings.HasPrefix(DistCertP12.Subject.CommonName, "Apple Distribution:") {
		return "", fmt.Errorf("发布证书的CommonName必须是Apple Distribution开头的字符串")
	}

	//6 生产证书，取出application-identifier内容，前缀要和发布证书的{teamID}一致。
	isTeamIDOK := false
	var teamId string
	for _, v := range DistCertP12.Subject.OrganizationalUnit {
		if v == DistCertOfficialMP.Entitlements.TeamIdentifier {
			isTeamIDOK = true
			teamId = v
			break
		}
		if strings.HasPrefix(DistCertOfficialMP.Entitlements.ApplicationIDentifier, v) {
			isTeamIDOK = true
			teamId = v
			break
		}
	}
	if !isTeamIDOK {
		return "", fmt.Errorf("生产证书的teamID要和发布证书teamID一致")
	}

	//6 测试证书，取出application-identifier内容，前缀要和发布证书的{teamID}一致。
	isTeamIDOKTest := false
	for _, v := range DistCertP12.Subject.OrganizationalUnit {
		if v == DistProvisionMP.Entitlements.TeamIdentifier {
			isTeamIDOKTest = true
			break
		}
		if strings.HasPrefix(DistProvisionMP.Entitlements.ApplicationIDentifier, v) {
			isTeamIDOKTest = true
			break
		}
	}
	if !isTeamIDOKTest {
		return "", fmt.Errorf("测试证书的teamID要和发布证书teamID一致")
	}

	//测试证书
	if len(DistProvisionMP.ProvisionedDevices) == 0 {
		return "", fmt.Errorf("测试证书要包含ProvisionedDevice")
	}

	//生产证书
	if len(DistCertOfficialMP.ProvisionedDevices) > 0 {
		return "", fmt.Errorf("生产证书不能包含ProvisionedDevice")
	}

	//OEM APP相关证书(iOS版)
	//验证p12证书，包括发布证书和推送证书
	//1,所有证书都验证有效期
	//2,组织单位teamID，发布证书、push证书都要一致
	//3,组织,必须是Apple
	//4,CommonName，常用名称前缀，发布证书和推送证书分别固定
	//5,推送证书CommonName(常用名称)后缀，跟APP的包名一致

	//验证mobileprovision，包括生产证书和iOS证书
	//1,取出application-identifier内容，其前缀(第1个.号前的内容)要和发布证书的{teamID}一致。
	//2,取出创建时间、截止时间，验证当前是否在有效期
	//3,生产证书不能包含ProvisionedDevice，测试证书必须包含ProvisionedDevice

	return teamId, nil
}

func CheckPushCert(packageName, distUrl, pass string, req *entitys.OemAppIosPushCertSaveReq) error {
	//下载发布证书p12
	DistCert, err := DownloadFile(distUrl)
	if err != nil {
		return err
	}
	DistCertP12, err := parseP12(DistCert, pass)
	if err != nil {
		return err
	}

	//下载push证书p12
	pushCert, err := DownloadFile(req.ApnsCert)
	if err != nil {
		return err
	}
	pushCertP12, err := parseP12(pushCert, req.ApnsSecret)
	if err != nil {
		return err
	}
	//1、验证push证书是否有效
	now := time.Now()
	if now.After(pushCertP12.NotAfter) || now.Before(pushCertP12.NotBefore) {
		return fmt.Errorf("push证书已过期.")
	}

	//2 push证书要和发布证书的teamID证书一致
	if len(DistCertP12.Subject.OrganizationalUnit) != len(pushCertP12.Subject.OrganizationalUnit) {
		return fmt.Errorf("push证书组织单位错误")
	}
	sort.Strings(DistCertP12.Subject.OrganizationalUnit)
	sort.Strings(pushCertP12.Subject.OrganizationalUnit)
	for i, v0 := range DistCertP12.Subject.OrganizationalUnit {
		if v0 != pushCertP12.Subject.OrganizationalUnit[i] {
			return fmt.Errorf("push证书和发布证书teamID不一致")
		}
	}

	//5 发布证书的常用名称必须是"Apple Distribution:"开头
	list := strings.Split(pushCertP12.Subject.CommonName, ":")
	if len(list) != 2 {
		return fmt.Errorf("push证书的CommonName错误")
	}
	if list[0] != "Apple Push Services" {
		return fmt.Errorf("push证书的CommonName不是Apple Push Services错误")
	}

	packageNamePush := strings.Trim(list[1], " ")
	if packageNamePush != packageName {
		return fmt.Errorf("push证书的包名和iOS APP包名不一致")
	}

	return nil
}

func DownloadFile(url string) ([]byte, error) {
	var retErr error
	for i := 0; i < 3; i++ {
		client := resty.New()
		resp, err := client.R().Get(url)
		if err != nil {
			retErr = err
		} else {
			if resp.StatusCode() == 200 {
				return resp.Body(), nil
			} else {
				retErr = fmt.Errorf("%d", resp.StatusCode())
			}
		}
	}
	return nil, retErr
}

func parseP12(p12Data []byte, password string) (*x509.Certificate, error) {
	priv, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		return nil, err
	}
	if err := priv.(*rsa.PrivateKey).Validate(); err != nil {
		return nil, err
	}
	return cert, nil
}

func parseMobileprovision(mpData []byte) (*MobileProvision, error) {
	info, err := NewContentInfo(mpData)
	if err != nil {
		return nil, err
	}
	return info.GetContent()
}
