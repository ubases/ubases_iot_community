package apis

import (
	"bytes"
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/common/entitys"
	"cloud_platform/iot_app_api_service/controls/common/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

type CommonController struct {
}

var Commoncontroller CommonController

var commonService = services.CommonService{}

var imageSavePath = "D://upload/upload/temp/"

//文件上传
//oss文件上传

// @Summary 本地文件下载
// @Description 本地文件下载
// @Tags 文件
// @Accept application/json
// @Param filename path string true "文件名称"
// @Success 0 {object} object "{"code": 200, "data": {}, "msg": "error message"}"
// @Router /common/local/getfile [post]
func (CommonController) GetLocalFile(c *gin.Context) {
	fileName := c.Param("fileName")
	bucket := c.Param("bucket")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	var path = fmt.Sprintf("%s%s/%s", imageSavePath, bucket, fileName)
	c.File(path)
}

// @Summary 本地文件上传
// @Description 本地文件上传
// @Tags 文件
// @Accept application/json
// @Param filename path string true "文件名称"
// @Router /common/local/uploadfile [post]
func (CommonController) UploadLocalFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		iotgin.ResFailCode(c, "upload image failed ", -1)
		return
	}
	//network_guide
	bucket := c.PostForm("bucket")

	//获取文件后缀
	headerFileName := header.Filename
	ext := path.Ext(headerFileName)

	var buffer bytes.Buffer
	buffer.WriteString(iotutil.Uuid())
	buffer.WriteString(ext)
	filename := buffer.String()

	//创建目录
	var path = fmt.Sprintf("%s%s/", imageSavePath, bucket)
	iotutil.CheckAndCreateFolder(path)

	var filepath string = fmt.Sprintf("%s%s", path, filename)
	out, err := os.Create(filepath)
	if err != nil {
		iotgin.ResFailCode(c, "upload images failed，"+err.Error(), -1)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		iotgin.ResFailCode(c, "copy failed copy  file failed", -1)
		return
	}
	iotgin.ResSuccess(c, filename)
}

// RegionList 区域列表
// @Summary 区域信息表列表
// @Description 区域信息表列表
// @Tags 通用
// @Param lang header string true "语言"
// @Param tenantId query string true "租户Id，如果提供了租户Id则查询租户Id是否有效"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/regionList [get]
func (CommonController) RegionList(c *gin.Context) {
	lang := c.GetHeader("lang")
	ip := c.ClientIP()
	resp, msg := commonService.RegionList(lang, ip)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	tenantId := c.Query("tenantId")
	if tenantId != "" {
		oc, err := rpc.ClientOpenCompany.Find(context.Background(), &proto.OpenCompanyFilter{TenantId: tenantId})
		if err != nil {
			iotgin.ResFailCode(c, err.Error(), iotgin.FAIL_TENANTID_CODE)
			return
		}
		if oc.Code != 200 && len(oc.Data) == 0 {
			iotgin.ResFailCode(c, "租户Id不存在", iotgin.FAIL_TENANTID_CODE)
			return
		}
	}
	iotgin.ResSuccess(c, resp)
}

// DictList 字典列表
// @Summary 字典列表
// @Description 字典列表
// @Tags 通用
// @Param data body entitys.DictListParam true "查询的字典参数"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/dic/list [post]
func (CommonController) DictList(c *gin.Context) {
	req := entitys.DictListParam{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	dictTypeList := req.DictTypeList
	if strings.TrimSpace(iotutil.ToString(dictTypeList)) == "" {
		iotgin.ResBadRequest(c, "paramList  is empty")
		return
	}
	if len(dictTypeList) == 0 {
		iotgin.ResBadRequest(c, "paramList  is empty")
		return
	}
	resp, msg := commonService.SetContext(controls.WithUserContext(c)).DictList(dictTypeList)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

func (CommonController) CustomLangList(c *gin.Context) {
	appKey := c.GetHeader("appKey")
	resp, msg := commonService.SetContext(controls.WithUserContext(c)).CustomLangList(appKey)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

// CustomResourceExport 区域列导出自定义资源（如果不存在自定义资源，则导出基础app资源）表
// @Summary 客户APP自定定义资源数据
// @Description 客户APP自定定义资源数据
// @Tags 通用
// @Param appKey header string true "APP KEY"
// @Param lang header string true "语言"
// @Param tenantId header string true "租户Id"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/customLang/list [get]
func (s *CommonController) CustomResourceExport(c *gin.Context) {
	//默认是APP的资源下载
	appKey := c.GetHeader("appKey")
	if appKey == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 appKey"))
		return
	}
	key := c.Query("key")

	ctx := controls.WithUserContext(c)
	result := map[string]interface{}{}
	cachedKey := persist.GetRedisKey(iotconst.APP_CUSTOM_LANG, appKey)
	dataCachedCmd := iotredis.GetClient().Get(ctx, cachedKey)
	if dataCachedCmd.Err() == nil {
		err := json.Unmarshal([]byte(dataCachedCmd.Val()), &result)
		if err == nil {
			if key != "" && key == iotutil.ToString(result["key"]) {
				iotgin.ResSuccessMsg(c)
				return
			}
			iotgin.ResSuccess(c, result)
			return
		} else {
			fmt.Println(err.Error())
		}
	}

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(ctx, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongType: 3,
			AppKey:     appKey,
		},
	})
	if err == nil && crs.Code == 200 && crs.Data != nil && len(crs.Data) > 0 {
		resLang := commonService.LangCustomResultConvert(crs.Data)
		res := map[string]interface{}{}
		for k, v := range resLang {
			res[k] = v
		}
		res["key"] = iotutil.Md5(iotutil.ToString(resLang))
		if err := iotredis.GetClient().Set(ctx, cachedKey, iotutil.ToString(res), 600*time.Second).Err(); err != nil {
			iotgin.ResSuccess(c, res)
			return
		}
		iotgin.ResSuccess(c, res)
	} else {
		res := s.loadCommonResources(appKey, key, ctx, c)
		iotgin.ResSuccess(c, res)
	}
}

// 加载APP的基础翻译数据
func (s CommonController) loadCommonResources(appKey string, verifyKey string, ctx context.Context, c *gin.Context) interface{} {
	//查询APP信息
	appRes, err := rpc.ClientOemAppService.Find(ctx, &proto.OemAppFilter{AppKey: appKey})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return nil
	}
	if appRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(appRes.Message))
		return nil
	}
	if len(appRes.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("参数异常"))
		return nil
	}

	appTemplateId := appRes.Data[0].AppTemplateId

	result := map[string]map[string]interface{}{}
	cachedKey := persist.GetRedisKey(iotconst.APP_COMMON_LANG, appTemplateId)
	dataCachedCmd := iotredis.GetClient().Get(ctx, cachedKey)
	if dataCachedCmd.Err() == nil {
		err := json.Unmarshal([]byte(dataCachedCmd.Val()), &result)
		if err == nil {
			if verifyKey != "" && verifyKey == iotutil.ToString(result["verifyKey"]) {
				return nil
			}
			return result
		}
	}

	rep, err := rpc.ClientLangResourcesService.Lists(ctx, &proto.LangResourcesListRequest{
		Query: &proto.LangResources{BelongType: 3, AppTemplateId: appTemplateId},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return nil
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return nil
	}
	resLang := commonService.LangResultConvert(rep.Data)
	res := map[string]interface{}{}
	for k, v := range resLang {
		res[k] = v
	}
	res["key"] = iotutil.Md5(iotutil.ToString(resLang))

	//不需要设置redis的过期时间，当用户重新上传了之后删除缓存即可；
	if err := iotredis.GetClient().Set(ctx, cachedKey, iotutil.ToString(res), 0).Err(); err != nil {
		return resLang
	}
	return res
}

// PanelCustomResourceExport 导出自定义资源（如果不存在自定义资源，则导出基础app资源）
func (s *CommonController) PanelCustomResourceExport(c *gin.Context) {
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 productKey"))
		return
	}
	thisContext := controls.WithUserContext(c)
	pro, err := rpc.ProductService.Find(thisContext, &proto.OpmProductFilter{ProductKey: productKey})
	if err != nil || len(pro.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("获取产品信息失败"))
		return
	}
	controlPanelId := pro.Data[0].ControlPanelId
	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongId:   controlPanelId,
			BelongType: 4,
			ProductKey: productKey,
		},
	})
	if err == nil && crs.Code == 200 && crs.Data != nil && len(crs.Data) > 0 {
		resLang := commonService.LangCustomResultConvert(crs.Data)
		iotgin.ResSuccess(c, resLang)
		return
	}
	rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &proto.LangResourcesListRequest{
		Query: &proto.LangResources{
			BelongId:   controlPanelId,
			BelongType: 4,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	resLang := commonService.LangResultConvert(rep.Data)
	iotgin.ResSuccess(c, resLang)
}

// GetWeather 获取天气
// @Summary 获取天气
// @Description 获取天气
// @Tags 通用
// @Param city query string true "城市"
// @Param province query string true "省"
// @Param lang query string true "语言"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/weather [get]
func (CommonController) GetWeather(c *gin.Context) {
	city := c.Query("city")
	province := c.Query("province")
	language := c.GetHeader("lang")
	resp, msg := commonService.GetWeather(city, province, language)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

// RegionList 区域列表
// @Summary 区域信息表列表
// @Description 区域信息表列表
// @Tags 通用
// @Param lang header string true "语言"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/region/{id} [get]
func (CommonController) RegionInfo(c *gin.Context) {
	lang := c.GetHeader("lang")
	regionId, _ := iotutil.ToInt64AndErr(c.Param("id"))
	if regionId == 0 {
		iotgin.ResBusinessP(c, "regionId is empty")
		return
	}
	resp, msg := commonService.RegionInfo(lang, regionId)
	if msg != nil {
		iotgin.ResBusinessP(c, msg.Error())
		return
	}
	iotgin.ResSuccess(c, resp)
}

// DictList 房间配置（默认房间、房间图标）
// @Summary 房间配置（默认房间、房间图标）
// @Description 房间配置（默认房间、房间图标）
// @Tags 通用
// @Param code path string true "编码"
// @Param appKey header string true "APP KEY"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/room/config/{code} [post]
func (CommonController) RoomConfigList(c *gin.Context) {
	code := c.Param("code") //room 和 icons
	if code == "" {
		iotgin.ResBadRequest(c, "code")
		return
	}
	appKey := c.GetHeader("appKey")
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	_, _, result := commonService.SetContext(controls.WithUserContext(c)).RoomConfigList(lang, tenantId, appKey, code)
	//if msgcode != 0 {
	//	iotgin.ResSuccess(c,result)
	//	//iotgin.ResBusinessP(c, msg)
	//	return
	//}
	iotgin.ResSuccess(c, result)
}

// VoiceService 第三方语音
// @Summary 第三方语音
// @Description 第三方语音
// @Tags 通用
// @Param voiceCode path string true "语音类型"
// @Param appKey header string true "APP Key"
// @Param lang header string true "lang"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/voice/config/{voiceCode} [get]
func (CommonController) VoiceService(c *gin.Context) {
	voiceCode := c.Param("voiceCode")
	if voiceCode == "" {
		iotgin.ResBadRequest(c, "voiceCode")
		return
	}

	appKey := c.GetHeader("appKey")
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	lang := c.GetHeader("lang")
	if lang == "" {
		iotgin.ResBadRequest(c, "lang")
		return
	}
	msgcode, _, result := commonService.SetContext(controls.WithUserContext(c)).VoiceService(appKey, voiceCode, lang)
	if msgcode != 0 {
		iotgin.ResSuccess(c, result)
		return
	}
	iotgin.ResSuccess(c, result)
}

// GetWeather 获取天气
// @Summary 获取天气
// @Description 获取天气
// @Tags 通用
// @Param appVersion query string true "APP版本"
// @Param account query string true "账号"
// @Param sizeType query string true "尺寸类型"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/flashscreen [get]
func (CommonController) GetFlashScreen(c *gin.Context) {
	appVersion := c.Query("appVersion")
	account := c.Query("account")
	sizeType := c.Query("sizeType")
	if appVersion == "" || sizeType == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppQueryParamIsNil, nil)
		return
	}
	data, err := commonService.SetContext(controls.WithUserContext(c)).GetFlashScreen(appVersion, account, iotutil.ToInt(sizeType))
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, data, int64(len(data)), 0)
}

type NowTime struct {
	Time    string `json:"time"`
	Weekday int    `json:"weekday"`
}

// GetNowTime 获取当前时间
// @Summary GetNowTime
// @Description GetNowTime
// @Tags 通用
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/nowtime [get]
func (CommonController) GetNowTime(c *gin.Context) {
	now := time.Now()
	nowTime := NowTime{
		Time:    iotutil.GetLocalTimeStr(now),
		Weekday: int(now.Local().Weekday()),
	}
	if nowTime.Weekday == 0 {
		nowTime.Weekday = 7
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nowTime)
}

// QueryPanelResourceV2 面板发翻译接口
// @Summary 面板发翻译接口
// @Description // @Summary 面板发翻译接口
// @Tags 通用
// @Param productKey query string true "产品Key"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/customLang/panel [get]
func (s *CommonController) QueryPanelResourceV2(c *gin.Context) {
	var (
		ctx                 = controls.WithUserContext(c)
		productKey          = c.DefaultQuery("productKey", "")
		belongIdInt   int64 = 0
		belongTypeInt int32 = 4
	)

	if productKey == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 productKey"))
		return
	}
	key := c.Query("key")

	result := map[string]interface{}{}
	cachedKey := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, productKey)
	dataCachedCmd := iotredis.GetClient().Get(ctx, cachedKey)

	if dataCachedCmd.Err() == nil {
		err := json.Unmarshal([]byte(dataCachedCmd.Val()), &result)
		if err == nil {
			if key != "" && key == iotutil.ToString(result["key"]) {
				iotgin.ResSuccessMsg(c)
				return
			}
			iotgin.ResSuccess(c, result)
			return
		}
	}

	pro, err := rpc.ProductService.Find(ctx, &proto.OpmProductFilter{ProductKey: productKey})
	if err != nil || len(pro.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("获取产品信息失败"))
		return
	}
	belongIdInt = pro.Data[0].ControlPanelId

	//获取面板Id对应的默认翻译信息
	langTypes, _ := s.getLangTypes(ctx, belongIdInt, belongTypeInt)

	//获取公共面板资源包
	public, err := s.getPublicResource(ctx, belongIdInt, belongTypeInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//获取面板自定义资源包
	custom, err := s.getCustomResource(ctx, belongIdInt, belongTypeInt, productKey)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(custom) > 0 {
		for key, _ := range public {
			if v, ok := custom[key]; ok {
				for ckey, cval := range v {
					public[key][ckey] = cval
				}
			}
		}
	}
	resLang := s.convertResouceResult(public, langTypes)
	res := map[string]interface{}{}
	for k, v := range resLang {
		res[k] = v
	}
	res["key"] = iotutil.Md5(iotutil.ToString(resLang))
	//设置缓存
	if err := iotredis.GetClient().Set(ctx, cachedKey, iotutil.ToString(res), 0).Err(); err != nil {
		iotlogger.LogHelper.WithTag("file", "CommonController").WithTag("method", "QueryPanelResourceV2").
			Errorf("缓存设置失败, key: %s, err: %s", cachedKey, err.Error())
	}

	//存储key值（将Keys存储起来，用于上传面板的时候清理；
	langKeysCached := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG_KEYS, belongIdInt)
	iotredis.GetClient().HMSet(ctx, langKeysCached, map[string]interface{}{cachedKey: "1"})

	iotgin.ResSuccess(c, res)
}

// 展示默认 zh、en
func (s CommonController) getLangTypes(ctx context.Context, belongIdInt int64, belongTypeInt int32) ([]string, error) {
	langTypes := []string{"zh", "en"}
	//res, err := rpc.ClientLangResourcesPackageService.Lists(ctx, &proto.LangResourcePackageListRequest{
	//	Query: &proto.LangResourcePackage{
	//		BelongType: belongTypeInt,
	//		BelongId:   belongIdInt,
	//	},
	//})
	//if err != nil {
	//	return langTypes, err
	//}
	//if res.Code != 200 {
	//	return langTypes, errors.New(res.Message)
	//}
	//if len(res.Data) == 0 {
	//	return langTypes, nil
	//}
	//var tempLangTyps []string = make([]string, 0)
	//for _, d := range res.Data[0].Langs {
	//	tempLangTyps = append(tempLangTyps, d)
	//}
	//
	//if err == nil {
	//	for _, lang := range tempLangTyps {
	//		noLang := ""
	//		for _, l := range langTypes {
	//			noLang = l
	//			if lang == l {
	//				noLang = ""
	//				break
	//			}
	//		}
	//		if noLang != "" {
	//			langTypes = append(langTypes, noLang)
	//		}
	//	}
	//}
	return langTypes, nil
}

// 转换返回格式
func (s CommonController) convertResouceResult(resultArr map[string]map[string]interface{}, langType []string) map[string]map[string]interface{} {
	langMap := map[string]map[string]interface{}{}
	for _, row := range resultArr {
		langKey := iotutil.ToString(row["code"])
		for _, lang := range langType {
			_, ok := langMap[lang]
			if !ok {
				langMap[lang] = map[string]interface{}{}
			}
			if v, ok := row[lang]; ok {
				langMap[lang][langKey] = v
			}
		}
	}
	return langMap
}

// 获取公共资源
func (s CommonController) getPublicResource(ctx context.Context, belongIdInt int64, belongTypeInt int32) (map[string]map[string]interface{}, error) {
	resultArr := map[string]map[string]interface{}{}
	rep, err := rpc.ClientLangResourcesService.Lists(ctx, &proto.LangResourcesListRequest{
		Query: &proto.LangResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	for i, d := range rep.Data {
		if _, ok := resultArr[d.Code]; !ok {
			resultArr[d.Code] = map[string]interface{}{"sort": i, "code": d.Code, d.Lang: d.Value}
		} else {
			resultArr[d.Code][d.Lang] = d.Value
		}
	}
	return resultArr, nil
}

// 获取自定义资源
func (s CommonController) getCustomResource(ctx context.Context, belongIdInt int64, belongTypeInt int32, productKey string) (map[string]map[string]interface{}, error) {
	resultArr := map[string]map[string]interface{}{}
	crs, err := rpc.ClientLangCustomResourceService.Lists(ctx, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
			ProductKey: productKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if crs.Code != 200 {
		return nil, errors.New(crs.Message)
	}
	for i, d := range crs.Data {
		if _, ok := resultArr[d.Code]; !ok {
			resultArr[d.Code] = map[string]interface{}{"sort": i, "code": d.Code, d.Lang: d.Value}
		} else {
			resultArr[d.Code][d.Lang] = d.Value
		}
	}
	return resultArr, nil
}
