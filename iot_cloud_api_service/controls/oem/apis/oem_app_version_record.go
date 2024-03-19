package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

var OemAppVersionRecordcontroller OemAppVersionRecordController

var serviceAppVersionRecord apiservice.OemAppVersionRecordService

type OemAppVersionRecordController struct {
} //用户操作控制器

// 更新版本
func (OemAppVersionRecordController) GetOemAppVersionRecordList(c *gin.Context) {
	req := entitys.OemAppVersionRecordQuery{}
	err := c.ShouldBind(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind oem app version list request error: ", err.Error())
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	iotlogger.LogHelper.Helper.Debugf("info: %v", req)
	res, err := serviceAppVersionRecord.SetContext(controls.WithOpenUserContext(c)).GetOemAppVersionRecordList(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get oem app version list error: ", err.Error())
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, res, int64(len(res)), int(req.Page))
}
