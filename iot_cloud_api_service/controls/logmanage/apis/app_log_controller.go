package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/controls/logmanage/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

func GetAppLogUserList(c *gin.Context) {
	var req protosService.AppLogUserListReq
	err := c.BindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Errorf("bind app log user list req param error: %v", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppLogUserListReqParam, nil)
		return
	}
	resp, err := rpc.ClientAppLogService.GetAppLogUserList(controls.WithOpenUserContext(c), &req)
	if err != nil {
		iotlogger.LogHelper.Helper.Errorf("rpc call get app log user list error: %v", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	list := []*entitys.AppLogUser{}
	for i := range resp.Data.List {
		list = append(list, entitys.AppLogUserPb2Db(resp.Data.List[i]))
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, list, resp.Data.Total, int(req.Page))
}

func GetAppLogRecordsList(c *gin.Context) {
	var req entitys.AppLogRecordsListReq
	err := c.BindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Errorf("bind app log records list req param error: %v", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppLogRecordsListReqParam, nil)
		return
	}
	resp, err := rpc.ClientAppLogService.GetAppLogRecordsList(controls.WithOpenUserContext(c), entitys.AppLogRecordsListReqDB2Pb(&req))
	if err != nil {
		iotlogger.LogHelper.Helper.Errorf("rpc call get app log records list error: %v", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	list := []*entitys.AppLogRecords{}
	iosModuleModes, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_ios_mobile_mode)
	for i, r := range resp.Data.List {
		//匹配IOS手机型号
		if v, ok := r.Details["system"]; ok {
			sysInfo := GetSystemInfo(v)
			if strings.Index(sysInfo.Model, "iPhone") != -1 {
				if v := iosModuleModes.ValueStr(sysInfo.Model); v != "" {
					sysInfo.Model = v
					resp.Data.List[i].Details["system"] = fmt.Sprintf("%v,%v,%v", sysInfo.Os, sysInfo.Model, sysInfo.Version)
				}
			}
		}
		list = append(list, entitys.AppLogRecordsPb2Db(resp.Data.List[i]))
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, list, resp.Data.Total, int(req.Page))
}

// 系统信息
type SystemInfo struct {
	Os      string `json:"os"`
	Version string `json:"version"`
	Model   string `json:"model"`
}

func GetSystemInfo(xSysInfo string) SystemInfo {
	sysArrs := strings.Split(xSysInfo, ",")
	if len(sysArrs) == 4 {
		return SystemInfo{
			Os:      sysArrs[0],
			Model:   sysArrs[1] + "," + sysArrs[2],
			Version: sysArrs[3],
		}
	}
	//最多支取3个
	if len(sysArrs) > 3 {
		sysArrs = sysArrs[0:3]
	}
	sysNewArr := make([]string, 4)
	for i := 0; i < len(sysArrs); i++ {
		sysNewArr[i] = sysArrs[i]
	}
	return SystemInfo{
		Os:      sysNewArr[0],
		Model:   sysNewArr[1],
		Version: sysNewArr[2],
	}
}
