package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var MsNoticeRecordcontroller MsNoticeRecordController

type MsNoticeRecordController struct{} //部门操作控制器

var msNoticeRecordServices = apiservice.MsNoticeRecordService{}

// @Summary 查询同时信息发送记录信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"code": "", "name":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/data/noticeInfo/sendRecord [get]
func (MsNoticeRecordController) QueryList(c *gin.Context) {
	var filter entitys.MsNoticeRecordQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.MsNoticeRecordFilter{}
	}
	//Method =1 邮件 =2 短信
	if filter.Query.Method == 2 {
		filter.Query.Method = 2
	} else if filter.Query.Method == 1 {
		if filter.Query.NoticeType == 1 {
			filter.Query.Method = 1
		} else if filter.Query.NoticeType == 2 {
			filter.Query.Method = 3
		} else {
			filter.Query.Method = 0 //使用methods多
			filter.Query.Methods = []int32{1, 3}
		}
	}
	res, total, err := msNoticeRecordServices.SetContext(controls.WithUserContext(c)).QueryMsNoticeRecordList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// @Summary 查询同时信息发送记录详细信息
// @Description
// @Tags APP
// @Accept application/json
// @Param id path string true "记录Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/data/noticeInfo/detail/{id} [get]
func (MsNoticeRecordController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := msNoticeRecordServices.SetContext(controls.WithUserContext(c)).GetMsNoticeRecordDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 删除同时信息发送记录信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"id": "" }"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/data/noticeInfo/delete/{id} [post]
func (MsNoticeRecordController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	idInt, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		iotgin.ResBadRequest(c, "id format")
		return
	}
	err = msNoticeRecordServices.SetContext(controls.WithUserContext(c)).DeleteMsNoticeRecord(entitys.MsNoticeRecordFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
