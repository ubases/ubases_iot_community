package apis

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var OemAppcontroller OemAppController

var serviceApp apiservice.OemAppService

type OemAppController struct {
} //用户操作控制器

// 创建oem app
func (OemAppController) Add(c *gin.Context) {
	var req entitys.OemAppEntitysAddReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).AddOemApp(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OemAppController) QueryList(c *gin.Context) {
	var filter entitys.OemAppQuery
	err := c.BindQuery(&filter)
	tenantId, _ := c.Get("tenantId")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).QueryOemAppList(filter, iotutil.ToString(tenantId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OemAppController) QueryListByTenantId(c *gin.Context) {
	var filter entitys.OemAppQuery
	err := c.BindQuery(&filter)
	if filter.TenantId == "" {
		iotgin.ResBadRequest(c, "tenantId")
		return
	}
	res, total, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).QueryOemAppList(filter, filter.TenantId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 修改app名称
func (OemAppController) ChangeName(c *gin.Context) {
	var req entitys.OemAppChangeNameReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).ChangeOemAppName(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 修改操作步骤
func (OemAppController) UpdateCurrentStep(c *gin.Context) {
	var req entitys.OemAppChangeCurrentStepReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).ChangeOemAppCurrentStep(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 修改app名称
func (OemAppController) UpdateTemplate(c *gin.Context) {
	var req entitys.OemAppUpdateTemplateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).UpdateOemAppTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 操作详细
func (OemAppController) GetOemAppDetail(c *gin.Context) {
	id := c.Query("appId")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("appId参数错误"))
		return
	}
	tenantId := c.GetString("tenantId")
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).QueryOemAppDetail(id, tenantId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 删除oemapp
func (OemAppController) DeleteOemApp(c *gin.Context) {
	id := c.Query("appId")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("appId参数错误"))
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).DeleteOemApp(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存功能配置
func (OemAppController) SaveMap(c *gin.Context) {
	var req entitys.OemAppMap
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).SaveMap(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取功能配置
func (OemAppController) GetMap(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).GetMap(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取构建包
func (OemAppController) OemAppBuildPackage(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppBuildPackage(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 上架中
func (OemAppController) OemAppPublishing(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppPublishing(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 上架
func (OemAppController) OemAppPublish(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppPublish(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 更新版本
func (OemAppController) OemAppUpdateVersion(c *gin.Context) {
	var req entitys.OemAppVersionUpdateReq
	err := c.ShouldBindQuery(&req)

	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppUpdateVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取功能配置
func (OemAppController) OemAppBuildPackageQrCodeUrl(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppBuildPackageQrCodeUrl(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取自定义app二维码链接
func (OemAppController) OemAppCustomPackageQrCodeUrl(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).OemAppCustomPackageQrCodeUrl(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (s *OemAppController) GetThirdLoginJson(c *gin.Context) {
	appKey := controls.GetAppKeyByHost(c)
	iosPkgName := ""
	if appKey == "" {
		appKey = config.Global.DefaultApp.AppKey
		iosPkgName = config.Global.DefaultApp.IosPkgName
	}
	iotgin.ResJSON(c, http.StatusOK, s.GetDeveloperApp(appKey, iosPkgName))
}

// GetBaseApp 获取公版的APP第三方登录信息apple-app-site-association
func (s *OemAppController) GetBaseApp() map[string]interface{} {
	//公版APP
	return map[string]interface{}{
		"applinks": map[string]interface{}{
			"apps": []map[string]interface{}{},
			"details": []map[string]interface{}{
				{
					"appID": "X86K9K38VS.com.aithinker.iot.aihome",
					"paths": []string{
						"/aithingsios/wechat/*",
						"/aithingsios/qq/*",
						"/aithingsios/alipay/*",
						"/aithingsios/weibo/*",
						"/aithingsios/facebook/*",
						"/aithingsios/linkedin/*",
						"/aithingsios/twitter/*",
						"/aithingsios/line/*",
						"/aithingsios/google/*",
						"/aithingsios/douyin/*",
					},
				},
			},
		},
	}
}

// GetDeveloperApp 获取公版的APP第三方登录信息apple-app-site-association
func (s *OemAppController) GetDeveloperApp(appKey, iosPkgName string) map[string]interface{} {
	//TODO　获取APP信息，TeamId和Ios包名
	appInfo, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{AppKey: appKey})
	if err != nil {
		return nil
	}
	if appInfo.Code != 200 {
		return nil
	}
	appData := appInfo.Data[0]
	//查询开发者APP信息
	return map[string]interface{}{
		"applinks": map[string]interface{}{
			"apps": []map[string]interface{}{},
			"details": []map[string]interface{}{
				{
					"appID": fmt.Sprintf("%s.%s", appData.IosTeamId, appData.IosPkgName),
					"paths": []string{
						fmt.Sprintf("/ai%s/wechat/*", appKey),
						fmt.Sprintf("/ai%s/qq/*", appKey),
						fmt.Sprintf("/ai%s/alipay/*", appKey),
						fmt.Sprintf("/ai%s/weibo/*", appKey),
						fmt.Sprintf("/ai%s/facebook/*", appKey),
						fmt.Sprintf("/ai%s/linkedin/*", appKey),
						fmt.Sprintf("/ai%s/twitter/*", appKey),
						fmt.Sprintf("/ai%s/line/*", appKey),
						fmt.Sprintf("/ai%s/google/*", appKey),
						fmt.Sprintf("/ai%s/douyin/*", appKey),
					},
				},
			},
		},
	}
}
