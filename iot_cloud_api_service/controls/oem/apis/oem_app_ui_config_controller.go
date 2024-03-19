package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var OemAppUiConfigcontroller OemAppUiConfigController

var serviceUiConfig apiservice.OemAppUiConfigService

type OemAppUiConfigController struct {
} //用户操作控制器

// 保存应用图标
func (OemAppUiConfigController) SaveIcon(c *gin.Context) {
	var req entitys.OemAppUiConfigIcon
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveIcon(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取应用图标
func (OemAppUiConfigController) GetIcon(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetIcon(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存应用图标
func (OemAppUiConfigController) SaveIosLaunchScreen(c *gin.Context) {
	var req entitys.OemAppUiConfigIosLaunchScreen
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveIosLaunchScreen(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取应用图标
func (OemAppUiConfigController) GetIosLaunchScreen(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetIosLaunchScreen(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存应用图标
func (OemAppUiConfigController) SaveAndroidLaunchScreen(c *gin.Context) {
	var req entitys.OemAppUiConfigAndroidLaunchScreen
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveAndroidLaunchScreen(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取应用图标
func (OemAppUiConfigController) GetAndroidLaunchScreen(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetAndroidLaunchScreen(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存主题色
func (OemAppUiConfigController) SaveThemeColors(c *gin.Context) {
	var req entitys.OemAppUiConfigThemeColors
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveThemeColors(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取应用图标
func (OemAppUiConfigController) GetThemeColors(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetThemeColors(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存个性化
func (OemAppUiConfigController) SavePersonalize(c *gin.Context) {
	var req entitys.OemAppUiConfigPersonalize
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SavePersonalize(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取个性化
func (OemAppUiConfigController) GetPersonalize(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetPersonalize(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存功能配置
func (OemAppUiConfigController) SaveFunctionConfigThird(c *gin.Context) {
	var req entitys.OemAppThirdServiceReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveFunctionConfigThird(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取功能配置
func (OemAppUiConfigController) GetFunctionConfigThird(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetFunctionConfigThird(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存功能配置
func (OemAppUiConfigController) SaveFunctionConfigVoice(c *gin.Context) {
	var req entitys.OemAppFunctionVoiceReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveFunctionConfigVoice(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取功能配置
func (OemAppUiConfigController) GetFunctionConfigVoice(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetFunctionConfigVoice(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存APP自动升级功能配置
func (OemAppUiConfigController) SaveFunctionConfigAutoUpgrade(c *gin.Context) {
	var req entitys.OemAppAutoUpgradeServiceReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveFunctionConfigAutoUpgrade(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取APP自动升级功能配置
func (OemAppUiConfigController) GetFunctionConfigAutoUpgrade(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetFunctionConfigAutoUpgrade(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存功能配置
func (OemAppUiConfigController) SaveFunctionConfig(c *gin.Context) {
	var req entitys.OemAppFunctionConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveFunctionConfig(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取功能配置
func (OemAppUiConfigController) GetFunctionConfig(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetFunctionConfig(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取默认房间列表
func (OemAppUiConfigController) GetRoomList(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetRoomList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存默认房间
func (OemAppUiConfigController) SaveRoom(c *gin.Context) {
	var req entitys.OemAppRoomEntityReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveRoom(req.Id, req.Room)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	iotgin.ResSuccess(c, res)
}

// 删除默认房间
func (OemAppUiConfigController) DeleteRoom(c *gin.Context) {
	id := c.Query("id")
	roomId := c.Query("roomId")
	if id == "" || roomId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).DeleteRoom(id, roomId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 恢复默认房间
func (OemAppUiConfigController) RecoverDefaultRoom(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).RecoverDefaultRoom(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取房间自定义图标
func (OemAppUiConfigController) GetRoomIconList(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetRoomIconList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存房间自定义图标.
func (OemAppUiConfigController) SaveRoomIconsList(c *gin.Context) {
	var req entitys.OemAppRoomIconsRes
	errReq := c.ShouldBindJSON(&req)
	if errReq != nil {
		iotgin.ResErrCli(c, errReq)
		return
	}
	err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveRoomIconsList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, "ok")
}

// 获取底部菜单列表
func (OemAppUiConfigController) GetButtonMenu(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetButtonMenu(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存底部菜单文字颜色
func (OemAppUiConfigController) SaveButoomMenuFontColor(c *gin.Context) {
	var req entitys.OemButtomMenuColorReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).SaveButoomMenuFontColor(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 保存底部菜单文字颜色
func (OemAppUiConfigController) AddButoomMenu(c *gin.Context) {
	uiId := c.Query("id")
	var req entitys.OemButtomMenuEntity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).AddButoomMenu(uiId, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改自定义底部菜单.
func (OemAppUiConfigController) UpdateButoomMenu(c *gin.Context) {
	uiId := c.Query("id")
	var req entitys.OemButtomMenuEntity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).UpdateButoomMenu(uiId, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppUiConfigController) DeleteButoomMenu(c *gin.Context) {
	uiId := c.Query("id")
	menuId := c.Query("menuId")
	if uiId == "" || menuId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).DeleteButoomMenu(uiId, menuId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppUiConfigController) GetButoomMenuDetail(c *gin.Context) {
	uiId := c.Query("id")
	menuId := c.Query("menuId")
	if uiId == "" || menuId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).GetButoomMenuDetail(uiId, menuId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 恢复默认值.
func (OemAppUiConfigController) RecoverDefault(c *gin.Context) {
	var req entitys.OemAppRecoverDefaultReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceUiConfig.SetContext(controls.WithOpenUserContext(c)).RecoverDefault(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
