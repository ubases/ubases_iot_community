package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/intelligence/apis"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotnats/jetstream"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	r := e.Group(webApiPrefix)
	r.Use(controls.AuthCheck)

	//一键执行和自动执行场景删除
	r.POST("/intelligence/del/:id", iotgin.AppLogger(jetstream.GetJsPublisherMgr()), apis.SceneIntelligencecontroller.DeleteIntelligence)

	//一键执行和自动执行详情信息
	r.POST("/intelligence/info", apis.SceneIntelligencecontroller.GetIntelligenceDetail)

	//智能列表数据
	r.POST("/intelligence/list", apis.SceneIntelligencecontroller.GetIntelligenceList)

	//一键执行和自动场景新增和修改
	r.POST("/intelligence/save", iotgin.AppLogger(jetstream.GetJsPublisherMgr()), apis.SceneIntelligencecontroller.SaveIntelligence)

	//一键执行和自动执行顺序调整  （no)
	r.POST("/intelligence/setSort", apis.SceneIntelligencecontroller.UpdateIntelligenceSortNo)

	//自动场景任务的开关接口
	r.POST("/intelligence/setSwitch", apis.SceneIntelligencecontroller.UpdateIntelligenceStatus)

	//场景日志
	//日志按天分组查询（默认查询最近7天）
	r.POST("/intelligence/logList", apis.SceneIntelligencecontroller.GetIntelligenceResultLogList)

	//场景日志详情
	r.POST("/intelligence/log/result/:id", apis.SceneIntelligencecontroller.GetIntelligenceTaskResultList)

	//清空场景日志
	r.POST("/intelligence/logDel", apis.SceneIntelligencecontroller.DeleteIntelligenceLog)

	//执行任务自动化智能列表
	r.POST("/intelligence/execList", apis.SceneIntelligencecontroller.GetExecList)

	//获取选中产品能进行智能条件和智能任务的物模型列表
	// /uapi/product/propsinfo/{model}
	r.GET("/product/propsinfo/:condType/:productId", apis.SceneIntelligencecontroller.GetTaskOrWhereByProductKey)
	r.GET("/product/propsInfoV2/:condType/:productId", apis.SceneIntelligencecontroller.GetTaskOrWhereByProductKeyV2)

	//一键执行
	r.POST("/intelligence/excute/submit/:id", apis.SceneIntelligencecontroller.OneKeyExec)

	//场景模板 /v1/platform/web/open/sceneTemplate/list
	r.POST("/sceneTemplate/list", apis.SceneTemplatecontroller.List)
	r.GET("/sceneTemplate/detail", apis.SceneTemplatecontroller.Get)

	r.POST("/intelligence/clear", apis.SceneIntelligencecontroller.ClearIntelligence)
}
