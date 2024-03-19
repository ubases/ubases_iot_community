package services

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go-micro.dev/v4/logger"
)

type OemAppUiConfigService struct {
	Ctx context.Context
}

func (s OemAppUiConfigService) SetContext(ctx context.Context) OemAppUiConfigService {
	s.Ctx = ctx
	return s
}

// 底部菜单字符串转结构
func (s OemAppUiConfigService) ButtomMenuConvertStruct(strJson string) *entitys.OemButtomMenuListRes {
	var bottomMenu = entitys.OemButtomMenuListRes{}
	iotutil.JsonToStruct(strJson, &bottomMenu)
	return &bottomMenu
}

// 底部菜单结构转字符串
func (s OemAppUiConfigService) ButtomMenuConvertString(req *entitys.OemButtomMenuListRes) string {
	strJson := iotutil.ToStringByUrl(req)
	return strJson
}

// 保存应用图标
func (s OemAppUiConfigService) SaveIcon(req entitys.OemAppUiConfigIcon) (string, error) {
	appId, err := iotutil.ToInt64AndErr(req.Id)
	if err != nil {
		return "", err
	}
	res, err := rpc.ClientOemAppService.SaveIcon(s.Ctx, &protosService.OemAppSaveIconReq{
		Id:      appId,
		IconUrl: req.IconUrl,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(appId)
	return "success", err
}

// 获取应用图标
func (s OemAppUiConfigService) GetIcon(req entitys.OemAppCommonReq) (*entitys.OemAppUiConfigIcon, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
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
		var nodata = entitys.OemAppUiConfigIcon{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}

	var data = entitys.OemAppUiConfigIcon{}
	data.AppId = req.AppId
	data.IconUrl = res.Data[0].IconUrl
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 保存ios应用图标
func (s OemAppUiConfigService) SaveIosLaunchScreen(req entitys.OemAppUiConfigIosLaunchScreen) (string, error) {
	////配合献敏, 只有用户上传了ios启动图. 才会把启动图打包. 如果用户未上传献敏会自己处理启动图

	//查询配置信息获取APPID
	appUI, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return "", err
	}
	req.IsUse = 1 //用户主动上传启动图
	reqStr := iotutil.ToString(req)
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"ios_launch_screen"},
		Data: &protosService.OemAppUiConfig{
			Id:              iotutil.ToInt64(req.Id),
			IosLaunchScreen: reqStr,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(appUI.Data[0].AppId)
	return "success", err
}

// 获取ios启动图
func (s OemAppUiConfigService) GetIosLaunchScreen(req entitys.OemAppCommonReq) (*entitys.OemAppUiConfigIosLaunchScreen, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppUiConfigIosLaunchScreen{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppUiConfigIosLaunchScreen{}
	if res.Data[0].IosLaunchScreen == "" {
		data.AppId = req.AppId
		data.Id = iotutil.ToString(res.Data[0].Id)
	}
	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].IosLaunchScreen, &data)
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 保存android应用图标
func (s OemAppUiConfigService) SaveAndroidLaunchScreen(req entitys.OemAppUiConfigAndroidLaunchScreen) (string, error) {
	//查询配置信息获取APPID
	appUI, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return "", err
	}

	//配合献敏, 只有用户上传了ios启动图. 才会把启动图打包. 如果用户未上传献敏会自己处理启动图
	req.IsUse = 1 //用户主动上传.
	reqStr := iotutil.ToString(req)
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"android_launch_screen"},
		Data: &protosService.OemAppUiConfig{
			Id:                  iotutil.ToInt64(req.Id),
			AndroidLaunchScreen: reqStr,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(appUI.Data[0].AppId)

	return "success", err
}

// 获取ios启动图
func (s OemAppUiConfigService) GetAndroidLaunchScreen(req entitys.OemAppCommonReq) (*entitys.OemAppUiConfigAndroidLaunchScreen, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppUiConfigAndroidLaunchScreen{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppUiConfigAndroidLaunchScreen{}
	if res.Data[0].AndroidLaunchScreen == "" {
		data.AppId = req.AppId
	}
	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].AndroidLaunchScreen, &data)
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 保存主题颜色
func (s OemAppUiConfigService) SaveThemeColors(req entitys.OemAppUiConfigThemeColors) (string, error) {
	//查询配置信息获取APPID
	appUI, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return "", err
	}

	reqStr := iotutil.ToString(req)
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"theme_colors"},
		Data: &protosService.OemAppUiConfig{
			Id:          iotutil.ToInt64(req.Id),
			ThemeColors: reqStr,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(appUI.Data[0].AppId)
	return "success", err
}

// 获取主题颜色
func (s OemAppUiConfigService) GetThemeColors(req entitys.OemAppCommonReq) (*entitys.OemAppUiConfigThemeColors, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppUiConfigThemeColors{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppUiConfigThemeColors{}
	if res.Data[0].ThemeColors == "" {
		data.AppId = req.AppId
	}
	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].ThemeColors, &data)
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 保存个性化
func (s OemAppUiConfigService) SavePersonalize(req entitys.OemAppUiConfigPersonalize) (string, error) {
	//查询配置信息获取APPID
	appUI, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return "", err
	}
	reqStr := iotutil.ToString(req)
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"personalize"},
		Data: &protosService.OemAppUiConfig{
			Id:          iotutil.ToInt64(req.Id),
			Personalize: reqStr,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(appUI.Data[0].AppId)
	return "success", err
}

// 获取个性化
func (s OemAppUiConfigService) GetPersonalize(req entitys.OemAppCommonReq) (*entitys.OemAppUiConfigPersonalize, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppUiConfigPersonalize{}
		nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppUiConfigPersonalize{}
	if res.Data[0].Personalize == "" {
		data.AppId = req.AppId
	}

	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].Personalize, &data)
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 获取功能配置三方服务
func (s OemAppUiConfigService) GetFunctionConfigThird(req entitys.OemAppCommonReq) (*entitys.OemAppThirdServiceRes, error) {
	resApp, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: iotutil.ToInt64(req.AppId),
	})
	if err != nil {
		return nil, err
	}
	if resApp.Code != 200 {
		return nil, errors.New(resApp.Message)
	}
	var (
		appKey      = resApp.Data[0].AppKey
		thirdDomain = config.Global.Service.ThirdDomain
	)
	//oem_app_third_login
	thirdLogins := GetBaseDataValue("oem_app_third_login", s.Ctx)
	thirds := make([]entitys.OemAppThird, 0)
	for k, v := range thirdLogins {
		thirds = append(thirds, entitys.OemAppThird{
			ThirdCode: iotutil.ToString(v),
			ThirdName: k,
			IsCheck:   0,
		})
	}
	res, err := rpc.ClientOemAppFunctionConfigService.Find(s.Ctx, &protosService.OemAppFunctionConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}

	var data = entitys.OemAppThirdServiceRes{}
	if res.Data[0].Thirds != "" {
		iotutil.JsonToStruct(res.Data[0].Thirds, &data.ThirdList)
	} else {
		data.ThirdList = make([]entitys.OemAppThird, 0)
	}
	data.Id = iotutil.ToString(res.Data[0].Id)

	//拼接UniversalLink,从配置中获取，支持非泛域名方式
	if strings.Index(thirdDomain, "%s") != -1 {
		thirdDomain = fmt.Sprintf(thirdDomain, appKey)
	}
	for i, t := range thirds {
		for _, t2 := range data.ThirdList {
			if t.ThirdCode == t2.ThirdCode {
				thirds[i].ThirdAppId = t2.ThirdAppId
				thirds[i].ThirdAppKey = t2.ThirdAppKey
				thirds[i].UniversalLink = t2.UniversalLink
				thirds[i].IsCheck = t2.IsCheck
			}
		}
		//拼接UniversalLink
		thirds[i].UniversalLink = fmt.Sprintf(thirdDomain+"/ai%s/%s/", appKey, t.ThirdCode)
	}
	data.ThirdList = thirds
	return &data, err
}

// 保存功能配置三方服务
func (s OemAppUiConfigService) SaveFunctionConfigThird(req entitys.OemAppThirdServiceReq) (string, error) {
	thirds := iotutil.ToString(req.ThirdList)
	res, err := rpc.ClientOemAppFunctionConfigService.UpdateFields(s.Ctx, &protosService.OemAppFunctionConfigUpdateFieldsRequest{
		Fields: []string{"thirds"},
		Data: &protosService.OemAppFunctionConfig{
			Id:     iotutil.ToInt64(req.Id),
			Thirds: thirds,
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

// 获取APP自动更新功能配置
func (s OemAppUiConfigService) GetFunctionConfigAutoUpgrade(req entitys.OemAppCommonReq) (*entitys.OemAppAutoUpgradeServiceRes, error) {
	res, err := rpc.ClientOemAppFunctionConfigService.Find(s.Ctx, &protosService.OemAppFunctionConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}

	var data = entitys.OemAppAutoUpgradeServiceRes{}
	if res.Data[0].AutoUpgrade != "" {
		iotutil.JsonToStruct(res.Data[0].AutoUpgrade, &data.AutoUpgrade)
	}
	data.Id = iotutil.ToString(res.Data[0].Id)
	return &data, err
}

// 保存APP自动更新功能配置
func (s OemAppUiConfigService) SaveFunctionConfigAutoUpgrade(req entitys.OemAppAutoUpgradeServiceReq) (string, error) {
	autoUpgrade := iotutil.ToString(req.AutoUpgrade)
	res, err := rpc.ClientOemAppFunctionConfigService.UpdateFields(s.Ctx, &protosService.OemAppFunctionConfigUpdateFieldsRequest{
		Fields: []string{"auto_upgrade"},
		Data: &protosService.OemAppFunctionConfig{
			Id:          iotutil.ToInt64(req.Id),
			AutoUpgrade: autoUpgrade,
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

// 获取功能配置语音服务
func (s OemAppUiConfigService) GetFunctionConfigVoice(req entitys.OemAppCommonReq) (*entitys.OemAppFunctionVoiceRes, error) {
	res, err := rpc.ClientOemAppFunctionConfigService.Find(s.Ctx, &protosService.OemAppFunctionConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var data = entitys.OemAppFunctionVoiceRes{}
	if res.Data[0].Voices != "" {
		iotutil.JsonToStruct(res.Data[0].Voices, &data.VoiceList)
	} else {
		data.VoiceList = make([]entitys.OemAppVoice, 0)
	}
	data.Id = iotutil.ToString(res.Data[0].Id)
	//查询文档列表
	introduceRes, err := rpc.ClientOemAppIntroduceService.Lists(s.Ctx, &protosService.OemAppIntroduceListRequest{
		Query: &protosService.OemAppIntroduce{
			AppId:       iotutil.ToInt64(req.AppId),
			ContentType: 4,
			// Version:     req.Version,
			//VoiceCode:   vcode,
		},
	})
	if err != nil {
		return nil, err
	}
	if introduceRes.Code != 200 {
		return nil, errors.New(introduceRes.Message)
	}
	//文档列表转换为Map
	introduceMap := make(map[string][]*protosService.OemAppIntroduce)
	for _, d := range introduceRes.Data {
		// if _, ok := introduceMap[d.VoiceCode]; ok {
		// 	continue
		// }
		introduceMap[d.VoiceCode] = append(introduceMap[d.VoiceCode], d)
	}
	//动态复制翻译
	for i, v := range data.VoiceList {
		langs := make([]string, 0)
		if introduces, ok := introduceMap[v.VoiceCode]; ok {
			for _, vv := range introduces {
				if !iotutil.ArraysExistsString(langs, vv.Lang) {
					langs = append(langs, vv.Lang)
				}
			}
		}
		data.VoiceList[i].Langs = langs
	}
	return &data, err
}

// 获取语控授权url
func (s OemAppUiConfigService) GetVoiceAuthUrl(appKey string, voiceCode string) (string, string, string) {
	domain := config.Global.Service.ThirdDomain
	//说明是https协议, 否则是http协议
	//tenantId, _ := metadata.Get(s.Ctx, "tenantid")
	if strings.Index(domain, "%s") != -1 {
		domain = fmt.Sprintf(domain, appKey)
	}
	gateway := ""
	//语控网关接口方法名
	switch voiceCode {
	case "xiaomi":
		gateway = "xiaomiIoTGateWay"
	case "xiaoai":
		gateway = "xiaoaiIoTGateWay"
	case "tianmao":
		gateway = "tianMaoIotGateWay"
	case "google":
		gateway = "googleFulfillment"
	case "alexa":
		gateway = "alexaIoTGateWay"
	}
	gate := fmt.Sprintf("%s/api/%s", domain, gateway)
	auth := fmt.Sprintf("%s/oauth/authorize", domain)
	token := fmt.Sprintf("%s/oauth/token", domain)
	return gate, auth, token
}

// 保存功能配置语音服务
func (s OemAppUiConfigService) SaveFunctionConfigVoice(req entitys.OemAppFunctionVoiceReq) (string, error) {
	if len(req.VoiceList) > 0 {
		for i := 0; i < len(req.VoiceList); i++ {
			if req.VoiceList[i].ClientId == "" {
				//生成密钥
				req.VoiceList[i].ClientId = iotutil.GetRandomString(16)
				req.VoiceList[i].ClientScrect = iotutil.GetRandomString(32)
			}
			if req.VoiceList[i].GatewayUrl == "" {
				gate, auth, token := s.GetVoiceAuthUrl(req.AppKey, req.VoiceList[i].VoiceCode)
				req.VoiceList[i].AuthUrl = auth
				req.VoiceList[i].GatewayUrl = gate
				req.VoiceList[i].TokenUrl = token
			}
		}
	}
	voices := iotutil.ToString(req.VoiceList)
	res, err := rpc.ClientOemAppFunctionConfigService.UpdateFields(s.Ctx, &protosService.OemAppFunctionConfigUpdateFieldsRequest{
		Fields: []string{"voices"},
		Data: &protosService.OemAppFunctionConfig{
			Id:     iotutil.ToInt64(req.Id),
			Voices: voices,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	// 将clientId, clientSecret, domain同步到iot_smart_speaker_service
	reqInfos := &protosService.ClientInfoReq{}
	for i := range req.VoiceList {
		reqInfo := &protosService.ClientInfo{
			ClientId:     req.VoiceList[i].ClientId,
			ClientSecret: req.VoiceList[i].ClientScrect,
			Domain:       "https://localhost",
		}
		reqInfos.ClientInfo = append(reqInfos.ClientInfo, reqInfo)
	}
	respInfo, err := rpc.ClientVoiceService.CreateClientInfo(s.Ctx, reqInfos)
	if err != nil {
		return "", err
	}
	if respInfo.Code != 200 {
		return "", errors.New(respInfo.Msg)
	}
	return "success", err
}

// 保存功能配置
func (s OemAppUiConfigService) SaveFunctionConfig(req entitys.OemAppFunctionConfig) (string, error) {
	res, err := rpc.ClientOemAppFunctionConfigService.UpdateFields(s.Ctx, &protosService.OemAppFunctionConfigUpdateFieldsRequest{
		Fields: []string{"about_us", "weather", "geo"},
		Data: &protosService.OemAppFunctionConfig{
			Id: iotutil.ToInt64(req.Id),
			//Eula: req.Eula,
			//PrivacyPolicy: req.Privacypolicy,
			AboutUs: req.Aboutus,
			Weather: req.Weather,
			Geo:     req.Geo,
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

// 获取功能配置
func (s OemAppUiConfigService) GetFunctionConfig(req entitys.OemAppCommonReq) (*entitys.OemAppFunctionConfig, error) {
	appId := iotutil.ToInt64(req.AppId)
	res, err := rpc.ClientOemAppFunctionConfigService.Find(s.Ctx, &protosService.OemAppFunctionConfigFilter{
		AppId:   appId,
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppFunctionConfig{}
		//nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppFunctionConfig{}

	//data.Eula = res.Data[0].Eula
	data.Aboutus = res.Data[0].AboutUs
	data.Weather = res.Data[0].Weather
	data.Geo = res.Data[0].Geo
	data.Id = iotutil.ToString(res.Data[0].Id)
	//data.AppId = req.AppId

	//获取是否配置了关于我们
	doc := OemAppIntroduceService{Ctx: s.Ctx}
	data.HasAboutus, _ = doc.CheckAppIntroduce(appId, 3)

	//获取是否配置了地图参数
	appSvc := OemAppService{Ctx: s.Ctx}
	data.HasGeo, _ = appSvc.CheckMap(req)

	return &data, err
}

// 默认房间进行升序
func (s OemAppUiConfigService) SortRoomList(list []*entitys.OemAppRoomEntity) {
	sort.Slice(list, func(i, j int) bool { // asc
		return list[i].RoomSort < list[j].RoomSort
	})
}

// 获取房间自定义图标
func (s OemAppUiConfigService) GetRoomIconList(req entitys.OemAppCommonReq) (*entitys.OemAppRoomIconsRes, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var nodata entitys.OemAppRoomIconsRes
	nodata.RoomIcons = make([]string, 0)
	if len(res.Data) == 0 || res.Data == nil {
		return &nodata, nil
	}
	var data entitys.OemAppRoomIconsRes
	if res.Data[0].Room == "" {
		return &nodata, nil
	}
	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].RoomIcons, &data.RoomIcons)
	data.Id = iotutil.ToString(res.Data[0].Id)

	return &data, nil
}

// 设置房间自定义图标.
func (s OemAppUiConfigService) SaveRoomIconsList(req entitys.OemAppRoomIconsRes) error {
	strIcons := ""
	if len(req.RoomIcons) > 0 {
		strIcons = iotutil.ToStringByUrl(req.RoomIcons)
	}

	resConfig, err := s.getUiConfig(req.Id)
	if err != nil {
		return err
	}

	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"room_icons"},
		Data: &protosService.OemAppUiConfig{
			Id:        iotutil.ToInt64(req.Id),
			RoomIcons: strIcons,
		},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	//清空缓存
	s.ClearLangCached(resConfig.Data[0].AppId)

	return nil
}

// 获取默认房间列表
func (s OemAppUiConfigService) GetRoomList(req entitys.OemAppCommonReq) (*entitys.OemAppRoomEntityRes, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var nodata entitys.OemAppRoomEntityRes
	nodata.Id = ""
	nodata.RoomList = make([]*entitys.OemAppRoomEntity, 0)
	if len(res.Data) == 0 || res.Data == nil {
		return &nodata, nil
	}
	var data entitys.OemAppRoomEntityRes
	if res.Data[0].Room == "" {
		nodata.Id = iotutil.ToString(res.Data[0].Id)
		return &nodata, nil
	}
	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].Room, &data.RoomList)
	//排序
	s.SortRoomList(data.RoomList)

	//赋值
	data.Id = iotutil.ToString(res.Data[0].Id)

	return &data, err
}

// 保存默认房间
func (s OemAppUiConfigService) SaveRoom(id string, req entitys.OemAppRoomEntity) (string, error) {
	res, err := s.getUiConfig(id)
	if err != nil {
		return "", err
	}

	if len(res.Data) > 0 && res.Data[0] != nil {
		appSvc := OemAppService{Ctx: s.Ctx}
		appInfo, err := appSvc.GetAppInfo(res.Data[0].AppId)
		if err != nil {
			return "", err
		}

		var list []*entitys.OemAppRoomEntity
		iotutil.JsonToStruct(res.Data[0].Room, &list)
		list, errRoom := s.SaveRoomEntity(list, req)
		if errRoom != nil {
			return "", errRoom
		}
		//保存入库
		errDb := s.SaveRoomToDb(id, list)
		if errDb != nil {
			return "", errDb
		}
		//写入翻译
		s.SetRoomTranslate(req.RoomId, fmt.Sprintf("oem_app_rooms_%v", appInfo.AppKey), req.RoomName, s.Ctx)
		//清空缓存
		s.ClearLangCached(res.Data[0].AppId)

		return "ok", nil
	} else {
		return "", errors.New("参数错误.未找到数据")
	}
}

func (s OemAppUiConfigService) getUiConfig(id string) (*protosService.OemAppUiConfigResponse, error) {
	resConfig, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if resConfig.Code != 200 && resConfig.Message != "record not found" {
		return nil, errors.New(resConfig.Message)
	}
	return resConfig, nil
}

func (s OemAppUiConfigService) ClearLangCached(appId int64) error {
	//删除缓存
	appInfo, err := rpc.ClientOemAppService.Find(s.Ctx, &protosService.OemAppFilter{
		Id: appId,
	})
	if err != nil {
		return err
	}
	if appInfo.Code != 200 && appInfo.Message != "record not found" {
		return errors.New(appInfo.Message)
	}
	//清空缓存
	key := fmt.Sprintf("%s_%s_%s_%s", appInfo.Data[0].TenantId, "zh", iotconst.LANG_OEM_APP_ROOMS, appInfo.Data[0].AppKey)
	r := iotredis.GetClient().Del(context.Background(), key)
	if r.Err() != nil {
		iotlogger.LogHelper.Error(r.Err())
	}
	key = fmt.Sprintf("%s_%s_%s_%s", appInfo.Data[0].TenantId, "en", iotconst.LANG_OEM_APP_ROOMS, appInfo.Data[0].AppKey)
	r = iotredis.GetClient().Del(context.Background(), key)
	if r.Err() != nil {
		iotlogger.LogHelper.Error(r.Err())
	}
	return nil
}

func (s OemAppUiConfigService) ClearLangCachedByAppKey(appKey string) error {
	//删除缓存
	appInfo, err := rpc.ClientOemAppService.Find(s.Ctx, &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		return err
	}
	if appInfo.Code != 200 && appInfo.Message != "record not found" {
		return errors.New(appInfo.Message)
	}
	//清空缓存
	key := fmt.Sprintf("%s_%s_%s_%s", appInfo.Data[0].TenantId, "zh", iotconst.LANG_OEM_APP_ROOMS, appInfo.Data[0].AppKey)
	r := iotredis.GetClient().Del(context.Background(), key)
	if r.Err() != nil {
		iotlogger.LogHelper.Error(r.Err())
	}
	key = fmt.Sprintf("%s_%s_%s_%s", appInfo.Data[0].TenantId, "en", iotconst.LANG_OEM_APP_ROOMS, appInfo.Data[0].AppKey)
	r = iotredis.GetClient().Del(context.Background(), key)
	if r.Err() != nil {
		iotlogger.LogHelper.Error(r.Err())
	}
	return nil
}

// 删除默认房间
func (s OemAppUiConfigService) DeleteRoom(id string, roomId string) (string, error) {
	res, err := s.getUiConfig(id)
	if err != nil {
		return "", err
	}

	if len(res.Data) > 0 && res.Data[0] != nil {
		var list []*entitys.OemAppRoomEntity
		iotutil.JsonToStruct(res.Data[0].Room, &list)
		resList := s.DeleteRoomEntity(list, roomId)
		if len(resList) <= 0 {
			return "", errors.New("请至少保留一个房间")
		}
		//保存入库
		errDb := s.SaveRoomToDb(id, resList)
		if errDb != nil {
			return "", errDb
		}

		//清空缓存
		s.ClearLangCached(res.Data[0].AppId)

		return "ok", nil
	} else {
		return "", errors.New("参数错误.未找到数据")
	}
}

// 恢复默认房间
func (s OemAppUiConfigService) RecoverDefaultRoom(id string) (string, error) {
	resConfig, err := s.getUiConfig(id)
	if err != nil {
		return "", err
	}

	//默认房间
	strRoomList := ""
	roomlist := s.GetBaseDataValueList("default_rooms")
	if len(roomlist) > 0 {
		var mapRoomList = make([]map[string]interface{}, 0)
		for _, v := range roomlist {
			var mapRoom = make(map[string]interface{})
			mapRoom["roomId"] = iotutil.ToString(iotutil.GetNextSeqInt64())
			mapRoom["roomName"] = v.DictLabel
			mapRoom["roomImage"] = v.Listimg
			mapRoom["roomSort"] = v.DictSort
			mapRoomList = append(mapRoomList, mapRoom)
		}
		strRoomList = iotutil.ToStringByUrl(mapRoomList)
	}

	//默认房间自定义图标
	strRoomIconList := ""
	roomIconList := s.GetBaseDataValueList("room_icons")
	if len(roomIconList) > 0 {
		var arrRoomIconList = make([]string, 0)
		for _, v := range roomIconList {
			if v.Listimg != "" {
				arrRoomIconList = append(arrRoomIconList, v.Listimg)
			}
		}
		strRoomIconList = iotutil.ToStringByUrl(arrRoomIconList)
	}

	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"room", "room_icons"},
		Data: &protosService.OemAppUiConfig{
			Id:        iotutil.ToInt64(id),
			Room:      strRoomList,
			RoomIcons: strRoomIconList,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//清空缓存
	s.ClearLangCached(resConfig.Data[0].AppId)

	return "ok", nil
}

// 根据数据字典类型获取字典数据列表
func (s OemAppUiConfigService) GetBaseDataValueList(dictType string) []*protosService.ConfigDictData {
	res, err := rpc.TConfigDictDataServerService.Lists(s.Ctx, &protosService.ConfigDictDataListRequest{
		Page:     1,
		PageSize: 100000,
		Query: &protosService.ConfigDictData{
			DictType: dictType,
		},
	})
	if err != nil {
		logger.Error(err.Error())
	}
	//排序
	if len(res.Data) > 1 {
		s.SortDictList(res.Data)
	}
	return res.Data
}

// 字典排序
func (s OemAppUiConfigService) SortDictList(list []*protosService.ConfigDictData) {
	sort.Slice(list, func(i, j int) bool { // asc
		return list[i].DictSort < list[j].DictSort
	})
}

// 房间结构保存到集合
func (s OemAppUiConfigService) SaveRoomEntity(list []*entitys.OemAppRoomEntity, req entitys.OemAppRoomEntity) ([]*entitys.OemAppRoomEntity, error) {
	isExists := false
	for _, v := range list {
		if req.RoomId == "" {
			//判断是否存在
			if v.RoomName == req.RoomName || v.RoomSort == req.RoomSort {
				isExists = true
				break
			}
		} else {
			//判断是否存在
			if (v.RoomName == req.RoomName || v.RoomSort == req.RoomSort) && v.RoomId != req.RoomId {
				isExists = true
				break
			}
		}
	}
	if isExists {
		return nil, errors.New("房间名称或序号已存在")
	}
	if req.RoomId == "" {
		req.RoomId = iotutil.ToString(iotutil.GetNextSeqInt64())
		list = append(list, &req)
	} else {
		for i := 0; i < len(list); i++ {
			if list[i].RoomId == req.RoomId {
				list[i].RoomImage = req.RoomImage
				list[i].RoomName = req.RoomName
				list[i].RoomSort = req.RoomSort
			}
		}
	}
	return list, nil
}

// 房间结构保存到集合
func (s OemAppUiConfigService) DeleteRoomEntity(list []*entitys.OemAppRoomEntity, roomId string) []*entitys.OemAppRoomEntity {
	var rslist []*entitys.OemAppRoomEntity
	for _, v := range list {
		if v.RoomId != roomId {
			rslist = append(rslist, v)
		}
	}
	return rslist
}

// 设置房间自定义图标.
func (s OemAppUiConfigService) SaveRoomToDb(id string, list []*entitys.OemAppRoomEntity) error {

	strRoom := ""
	if len(list) > 0 {
		strRoom = iotutil.ToString(list)
	}
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"room"},
		Data: &protosService.OemAppUiConfig{
			Id:   iotutil.ToInt64(id),
			Room: strRoom,
		},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return nil
}

func (s OemAppUiConfigService) SetRoomTranslate(sourceRowId string, sourceTable string, roomName string, ctx context.Context) {
	//设置翻译
	translateList := make([]*protosService.BatchSaveTranslateItem, 0)
	translateItem := protosService.BatchSaveTranslateItem{
		Lang:       "zh",
		FieldName:  "name",
		FieldType:  0,
		FieldValue: roomName,
	}
	translateList = append(translateList, &translateItem)
	_, err := rpc.ClientLangTranslateService.BatchCreate(ctx, &protosService.BatchSaveTranslate{
		SourceRowId:  sourceRowId,
		SourceTable:  sourceTable,
		PlatformType: 2,
		List:         translateList,
	})
	if err != nil {
		return
	}
}

// 获取底部菜单列表.
func (s OemAppUiConfigService) GetButtonMenu(req entitys.OemAppCommonReq) (*entitys.OemButtomMenuListRes, error) {
	res, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}

	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemButtomMenuListRes{}
		return &nodata, nil
	}

	var data = entitys.OemButtomMenuListRes{}
	if res.Data[0].BottomMenu == "" {
		return nil, errors.New("底部菜单初始化错误")
	}

	//josn字符串转结构.
	iotutil.JsonToStruct(res.Data[0].BottomMenu, &data)
	data.Id = iotutil.ToString(res.Data[0].Id)

	return &data, err
}

// 保存底部菜单文字颜色
func (s OemAppUiConfigService) SaveButoomMenuFontColor(req entitys.OemButtomMenuColorReq) (string, error) {
	//获取界面配置对象.
	resUi, errUi := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if errUi != nil {
		return "", errUi
	}
	if resUi.Code != 200 {
		return "", errors.New(resUi.Message)
	}
	//获取底部菜单的json
	strJson := resUi.Data[0].BottomMenu
	//把底部菜单的json字符串转化为结构对象
	bottonMenu := s.ButtomMenuConvertStruct(strJson)

	//赋值文字颜色文字
	bottonMenu.SelColor = req.SelColor
	bottonMenu.DefColor = req.DefColor
	//把结构重新转为字符串存储入库
	resButtonMenuJson := s.ButtomMenuConvertString(bottonMenu)

	//底部菜单入库
	err := s.SaveButtomMenuToDb(req.Id, resButtonMenuJson)
	if err != nil {
		return "", nil
	}
	return "success", nil
}

// 保存底部菜单json字符串到数据库.
func (s OemAppUiConfigService) SaveButtomMenuToDb(id string, strButtomMenu string) error {
	res, err := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
		Fields: []string{"bottom_menu"},
		Data: &protosService.OemAppUiConfig{
			Id:         iotutil.ToInt64(id),
			BottomMenu: strButtomMenu,
		},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// 新增自定义菜单.
func (s OemAppUiConfigService) AddButoomMenu(uiId string, req entitys.OemButtomMenuEntity) (string, error) {
	res, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(uiId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return "", errors.New(res.Message)
	}

	// if req.Position != 3  || req.Position != 4{
	// 	return "",errors.New("菜单位置不正确")
	// }

	bottomMenu := s.ButtomMenuConvertStruct(res.Data[0].BottomMenu)
	req.MenuId = iotutil.ToString(iotutil.GetNextSeqUint64())
	//把菜单新增到json里面
	req.Required = 2

	bottomMenu.MenuList = append(bottomMenu.MenuList, req)
	//菜单结构转字符串
	strJson := s.ButtomMenuConvertString(bottomMenu)
	//菜单字符串入库
	errMenu := s.SaveButtomMenuToDb(uiId, strJson)
	if errMenu != nil {
		return "", errMenu
	}
	//返回菜单id

	//更新APP UI更新时间
	appSvc := OemAppService{}
	appSvc.UpdateUIConfigUpdateTime(res.Data[0].AppId)

	return req.MenuId, nil
}

// 修改自定义菜单.
func (s OemAppUiConfigService) UpdateButoomMenu(uiId string, req entitys.OemButtomMenuEntity) (string, error) {
	res, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(uiId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return "", errors.New(res.Message)
	}
	bottomMenu := s.ButtomMenuConvertStruct(res.Data[0].BottomMenu)
	//把菜单结构修改一下.
	for i, _ := range bottomMenu.MenuList {
		posi := bottomMenu.MenuList[i].Position
		//自定义菜单
		if bottomMenu.MenuList[i].MenuId == req.MenuId && (posi == 3 || posi == 4) {
			bottomMenu.MenuList[i].Name = req.Name
			bottomMenu.MenuList[i].WebUrl = req.WebUrl
			bottomMenu.MenuList[i].DefImage = req.DefImage
			bottomMenu.MenuList[i].SelImage = req.SelImage
			bottomMenu.MenuList[i].Position = req.Position
			break
		}
		//默认菜单
		if bottomMenu.MenuList[i].MenuId == req.MenuId && (posi == 1 || posi == 2 || posi == 5) {
			bottomMenu.MenuList[i].DefImage = req.DefImage
			bottomMenu.MenuList[i].SelImage = req.SelImage
			break
		}

	}
	//菜单结构转字符串
	strJson := s.ButtomMenuConvertString(bottomMenu)
	//菜单字符串入库
	errMenu := s.SaveButtomMenuToDb(uiId, strJson)
	if errMenu != nil {
		return "", errMenu
	}
	//返回菜单id
	return req.MenuId, nil
}

// 修改自定义菜单.
func (s OemAppUiConfigService) DeleteButoomMenu(uiId string, menuId string) (string, error) {
	res, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(uiId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return "", errors.New(res.Message)
	}
	//获取到数据库里面的底部菜单结构
	bottomMenu := s.ButtomMenuConvertStruct(res.Data[0].BottomMenu)
	//定义一个新的底部菜单结构集合
	var newbuttomMenuList = make([]entitys.OemButtomMenuEntity, 0)
	//过滤到需要删除的结构.新增到新的菜单结构集合
	for _, v := range bottomMenu.MenuList {
		if v.MenuId != menuId {
			newbuttomMenuList = append(newbuttomMenuList, v)
		}
	}
	bottomMenu.MenuList = newbuttomMenuList
	//菜单结构转字符串
	strJson := s.ButtomMenuConvertString(bottomMenu)
	//菜单字符串入库
	errMenu := s.SaveButtomMenuToDb(uiId, strJson)
	if errMenu != nil {
		return "", errMenu
	}
	//返回菜单id
	return "success", nil
}

// 获取自定义菜单详细
func (s OemAppUiConfigService) GetButoomMenuDetail(uiId string, menuId string) (*entitys.OemButtomMenuEntity, error) {
	res, err := rpc.ClientOemAppUiConfigService.FindById(s.Ctx, &protosService.OemAppUiConfigFilter{
		Id: iotutil.ToInt64(uiId),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	//获取到数据库里面的底部菜单结构
	bottomMenu := s.ButtomMenuConvertStruct(res.Data[0].BottomMenu)
	//定义一个新的底部菜单结构集合
	var resMenu = entitys.OemButtomMenuEntity{}
	//过滤到需要删除的结构.新增到新的菜单结构集合
	for _, v := range bottomMenu.MenuList {
		if v.MenuId == menuId {
			//newbuttomMenuList = append(newbuttomMenuList, v)
			resMenu = v
			break
		}
	}
	//返回菜单id
	return &resMenu, nil
}

// 恢复默认值
func (s OemAppUiConfigService) RecoverDefault(req entitys.OemAppRecoverDefaultReq) (int32, error) {

	//底部菜单恢复默认值
	if req.DefType == 1 {
		resMenu, errMenu := rpc.ClientOemAppDefMenuService.Lists(s.Ctx, &protosService.OemAppDefMenuListRequest{
			Query: &protosService.OemAppDefMenu{
				Required: 1,
			},
		})
		if errMenu != nil {
			return 0, errMenu
		}
		if resMenu.Code != 200 && resMenu.Message != "record not found" {
			return 0, errors.New(resMenu.Message)
		}

		var buttomMenu = make(map[string]interface{})
		//TODO  需要改为默认值[建议后续把默认值配置到字典]
		buttomMenu["selColor"] = "#3B7CFF"
		buttomMenu["defColor"] = "#343A40"
		var menuList []map[string]interface{}
		for _, v := range resMenu.Data {
			var tmp = make(map[string]interface{})
			tmp["menuId"] = iotutil.ToString(v.Id)
			tmp["menuKey"] = v.MenuKey
			tmp["required"] = v.Required
			tmp["selImage"] = v.SelImage
			tmp["defImage"] = v.DefImage
			tmp["name"] = v.Name
			tmp["position"] = v.Position
			menuList = append(menuList, tmp)
		}
		buttomMenu["menuList"] = menuList
		strJsonButtomMenu := iotutil.ToString(buttomMenu)
		//fmt.Println(strJsonButtomMenu)

		resUp, errUp := rpc.ClientOemAppUiConfigService.UpdateFields(s.Ctx, &protosService.OemAppUiConfigUpdateFieldsRequest{
			Fields: []string{"bottom_menu"},
			Data: &protosService.OemAppUiConfig{
				Id:         iotutil.ToInt64(req.Id),
				BottomMenu: strJsonButtomMenu,
			},
		})
		if errUp != nil {
			return 0, errUp
		}
		if resUp.Code != 200 && resUp.Message != "record not found" {
			return 0, errors.New(resUp.Message)
		}

	} else if req.DefType == 2 {
		personalize := GetBaseDataValue("oem_app_default_personalize", s.Ctx)
		_, errPer := s.SavePersonalize(entitys.OemAppUiConfigPersonalize{
			AppId:                      req.AppId,
			Id:                         req.Id,
			Loginregisterbackgroundurl: iotutil.ToString(personalize["loginRegisterBackgroundUrl"]),
			Loginregisterlogourl:       iotutil.ToString(personalize["loginRegisterLogoUrl"]),
			Nodataurl:                  iotutil.ToString(personalize["noDataUrl"]),
			Defaultavatarurl:           iotutil.ToString(personalize["defaultAvatarUrl"]),
		})
		if errPer != nil {
			return 0, errPer
		}
	} else {
		return 0, errors.New("defType参数错误")
	}

	return 1, nil
}
