package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/intelligence/entitys"
	"cloud_platform/iot_app_api_service/controls/intelligence/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

var SceneIntelligencecontroller SceneIntelligenceController

type SceneIntelligenceController struct {
} //用户操作控制器

var SceneIntelligenceconService = services.SceneIntelligenceconService{}

// @Summary 新增/更新智能场景
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.OldSceneIntelligenceForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/save [post]
func (SceneIntelligenceController) SaveIntelligence(c *gin.Context) {
	var odlSceneIntelligenceForm entitys.OldSceneIntelligenceForm
	err := c.ShouldBind(&odlSceneIntelligenceForm)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//旧场景对象转换成新场景对象
	sceneIntelligenceForm := entitys.Intelligence_old2new(&odlSceneIntelligenceForm)
	errCode, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).SaveIntelligence(sceneIntelligenceForm)
	if err != nil {
		if errCode == -1 {
			iotgin.ResFailCode(c, "场景下的设备无效", errCode)
			return
		}
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)

}

// @Summary 查询智能场景列表
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceQueryForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/list [post]
func (SceneIntelligenceController) GetIntelligenceList(c *gin.Context) {
	var filter entitys.SceneIntelligenceQueryForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//if filter.Type == 0 {
	//	iotgin.ResBadRequest(c, "type不能为空")
	//	return
	//}
	if filter.HomeId == 0 {
		iotgin.ResBadRequest(c, "HomeId不能为空")
		return
	}
	//设置默认分页参数
	if filter.Limit == 0 {
		filter.Page = 1
		filter.Limit = 10
	}
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	res, total, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceList(lang, tenantId, filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(total), int(filter.Page))

}

// @Summary 查询智能场景详情
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceQueryForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/info [post]
func (SceneIntelligenceController) GetIntelligenceDetail(c *gin.Context) {
	var filter entitys.SceneIntelligenceQueryForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Id == 0 {
		iotgin.ResBadRequest(c, "id不能为空")
		return
	}
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceDetail(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 开关自动执行的智能场景
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/setSwitch [post]
func (SceneIntelligenceController) UpdateIntelligenceStatus(c *gin.Context) {
	var filter entitys.SceneIntelligenceForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//if filter.Status == 0 {
	//	iotgin.ResBadRequest(c, "status不能为空")
	//	return
	//}

	err = SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).UpdateIntelligenceStatus(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// @Summary 删除智能场景
// @Description
// @Tags intelligence
// @Accept application/json
// @Param id path string true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/del/:id [post]
func (SceneIntelligenceController) DeleteIntelligence(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	idi, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	err = SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).DeleteIntelligence(idi)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// @Summary 查询智能场景日志列表
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceQueryForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/logList [post]
func (SceneIntelligenceController) GetIntelligenceResultLogList(c *gin.Context) {
	var filter entitys.SceneIntelligenceQueryForm
	err := c.ShouldBind(&filter)

	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceResultLogGroupList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 查询智能场景详情
// @Description
// @Tags intelligence
// @Accept application/json
// @Param id path string true "场景Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/log/result/:id [post]
func (SceneIntelligenceController) GetIntelligenceTaskResultList(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	idi, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceTaskResultList(idi)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 清空智能场景日志
// @Description
// @Tags intelligence
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /intelligence/logDel [post]
func (SceneIntelligenceController) DeleteIntelligenceLog(c *gin.Context) {
	err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).DeleteIntelligenceLog()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// @Summary 查询执行任务的自动化智能场景列表
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceQueryForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/execList [post]
func (SceneIntelligenceController) GetExecList(c *gin.Context) {
	var filter entitys.SceneIntelligenceQueryForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.HomeId == 0 {
		iotgin.ResBadRequest(c, "HomeId不能为空")
		return
	}
	//filter.Type = 2
	userId := controls.GetUserId(c)
	filter.UserId = iotutil.ToInt64(userId)
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	res, _, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceList(lang, tenantId, filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//只获取有有效的场景
	newRes := make([]*entitys.SceneIntelligenceVo, 0)
	for _, v := range res {
		if v.FailureFlag != 4 {
			newRes = append(newRes, v)
		}
	}
	iotgin.ResPageSuccess(c, newRes, int64(len(newRes)), int(filter.Page))

}

// @Summary 一键执行和自动执行顺序调整
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/setSort [post]
func (SceneIntelligenceController) UpdateIntelligenceSortNo(c *gin.Context) {
	var filter entitys.SceneIntelligenceForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.SortNo == 0 {
		iotgin.ResBadRequest(c, "sortNo不能为空")
		return
	}

	err = SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).UpdateIntelligenceSortNo(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// @Summary 获取产品对应的智能条件和智能任务
// @Description
// @Tags intelligence
// @Accept application/json
// @Param productId path string true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/propsinfo/{condType}/{productId} [post]
func (SceneIntelligenceController) GetTaskOrWhereByProductKey(c *gin.Context) {
	productIdStr := c.Param("productId")
	if productIdStr == "" {
		iotgin.ResBadRequest(c, "产品编号不能为空")
		return
	}
	productId, err := iotutil.ToInt64AndErr(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "产品编号不合法")
		return
	}
	condType := c.Param("condType")
	if condType == "" {
		iotgin.ResBadRequest(c, "条件类型不能为空")
		return
	}
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetTaskOrWhereByProductKey(productId, condType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取产品对应的智能条件和智能任务
func (SceneIntelligenceController) GetTaskOrWhereByProductKeyV2(c *gin.Context) {
	productIdStr := c.Param("productId")
	if productIdStr == "" {
		iotgin.ResBadRequest(c, "产品编号不能为空")
		return
	}
	productId, err := iotutil.ToInt64AndErr(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "产品编号不合法")
		return
	}
	condType := c.Param("condType")
	if condType == "" {
		iotgin.ResBadRequest(c, "条件类型不能为空")
		return
	}
	deviceId := c.Query("deviceId")
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetTaskOrWhereByProductKeyV2(productId, condType, deviceId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GetProductFunctions 获取设备的功能列表
// @Summary 获取设备的功能列表
// @Description
// @Tags room
// @Accept application/json
// @Param devId path string true "设备Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/functions/{devId} [get]
func (SceneIntelligenceController) GetProductFunctions(c *gin.Context) {
	deviceId := c.Param("devId")
	if deviceId == "" {
		iotgin.ResBadRequest(c, "设备编号不能为空")
		return
	}

	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetProductFunctions(deviceId, 0)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GetAppointmentFunctions 获取设备的功能列表
// @Summary 获取设备的功能列表
// @Description
// @Tags room
// @Accept application/json
// @Param devId path string true "设备Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/functionsV2/{devId} [get]
func (SceneIntelligenceController) GetAppointmentFunctions(c *gin.Context) {
	deviceId := c.Param("devId")
	if deviceId == "" {
		iotgin.ResBadRequest(c, "设备编号不能为空")
		return
	}
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetProductFunctions(deviceId, 2)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// OneKeyExec 一键执行
func (SceneIntelligenceController) OneKeyExec(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		iotgin.ResBadRequest(c, "一键执行编号不能为空")
		return
	}
	id, err := iotutil.ToInt64AndErr(idStr)
	if err != nil {
		iotgin.ResBadRequest(c, "一键执行编号不合法")
		return
	}
	res, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).OneKeyExec(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, iotutil.ToString(res))

}

// @Summary 清理失效的场景
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneIntelligenceQueryForm true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Failure 400 {object} iotgin.ResponseModel 失败返回
// @Router /intelligence/clear [post]
func (SceneIntelligenceController) ClearIntelligence(c *gin.Context) {
	var filter entitys.SceneIntelligenceQueryForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.HomeId == 0 {
		iotgin.ResBadRequest(c, "HomeId不能为空")
		return
	}
	//设置默认分页参数
	filter.Page = 1
	filter.Limit = 10
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	res, _, err := SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).GetIntelligenceList(lang, tenantId, filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//删除场景
	for _, r := range res {
		if r.FailureFlag == 4 {
			SceneIntelligenceconService.SetContext(controls.WithUserContext(c)).DeleteIntelligence(r.Id)
		}
	}
	iotgin.ResSuccess(c, nil)
}
