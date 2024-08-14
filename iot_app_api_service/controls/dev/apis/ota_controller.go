package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"sort"

	"github.com/gin-gonic/gin"
)

var Otacontroller OtaController

type OtaController struct {
} //固件信息

// CheckOtaVersion 检查设备固件是否需要升级
func (s *OtaController) CheckOtaVersion(c *gin.Context) {
	var (
		productKey = c.DefaultQuery("productKey", "") //产品Key
		preVersion = c.DefaultQuery("version", "")    //设备当前版本
		deviceId   = c.DefaultQuery("deviceId", "")   //设备Id
	)
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	if deviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}
	//调用服务获取升级信息
	ret, err := rpc.ClientOtaPublishService.CheckOtaVersion(controls.WithUserContext(c), &protosService.CheckOtaVersionRequest{
		ProductKey: productKey,
		Version:    preVersion,
		DeviceId:   deviceId,
	})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	//IsAuto作用于是否显示启动自动升级，这里值为2固定出现选择框
	simpleRet := entitys.CheckOtaVersionSimpleResponse{IsAuto: 2, IsAutoUpgrade: false}
	//获取是否自动升级
	isAutoUpgrade := false
	rdCmd := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId, iotconst.FIELD_IS_AUTH_UPGRADE)
	if rdCmd.Err() == nil {
		if len(rdCmd.Val()) > 0 {
			isAutoUpgrade = rdCmd.Val()[0] == "true"
			simpleRet.IsAutoUpgrade = isAutoUpgrade
		}
	}

	if ret.Code == 101 || ret.OtaPkg == nil {
		iotgin.ResSuccessWithCode(c, 101, simpleRet)
		return
	}
	if ret.Code != 200 {
		iotgin.ResErrCli(c, errors.New(ret.Message))
		return
	}
	lang := controls.GetLang(c)
	otaInfo := entitys.CheckOtaVersion_Pd2E(ret, lang)
	otaInfo.IsAutoUpgrade = isAutoUpgrade
	iotgin.ResSuccess(c, otaInfo)
}

// CheckOtaUpgradeList 检查设备所有固件升级记录，并返回给前端
func (s *OtaController) CheckOtaUpgradeList(c *gin.Context) {
	var (
		productKey = c.DefaultQuery("productKey", "") //产品Key
		deviceId   = c.DefaultQuery("deviceId", "")   //设备Id
	)
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	if deviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}

	//调用服务获取升级信息
	ret, err := rpc.ClientOtaPublishService.CheckOtaUpgradeList(controls.WithUserContext(c), &protosService.CheckOtaVersionRequest{
		ProductKey: productKey,
		DeviceId:   deviceId,
	})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}

	//获取是否自动升级
	var (
		lang                 = controls.GetLang(c)
		isAutoUpgrade        = false
		otaState      string = ""
		progress      int32  = 0
	)
	devStatus := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId)
	if devStatus.Err() == nil {
		isAutoUpgrade = devStatus.Val()[iotconst.FIELD_IS_AUTH_UPGRADE] == "true"
		if val, ok := devStatus.Val()[iotconst.FIELD_UPGRADE_STATE]; ok && val != "" {
			otaState = iotutil.ToString(val)
		}
		if val, ok := devStatus.Val()[iotconst.FIELD_UPGRADE_PROGRESS]; ok && val != "" {
			progress, _ = iotutil.ToInt32Err(val)
		}
	}
	//根据固件类型分组固件列表
	firmwareGroup := make(map[int32][]*protosService.CheckOtaVersionResponse)
	for _, response := range ret.UpgradeList {
		if _, ok := firmwareGroup[response.OtaPkg.FirmwareType]; !ok {
			firmwareGroup[response.OtaPkg.FirmwareType] = make([]*protosService.CheckOtaVersionResponse, 0)
		}
		firmwareGroup[response.OtaPkg.FirmwareType] = append(firmwareGroup[response.OtaPkg.FirmwareType], response)
	}
	list := make([]*entitys.CheckOtaFirmwares, 0)
	hasUpdate := false
	for firmwareType, firmwares := range firmwareGroup {
		f := new(entitys.CheckOtaFirmwares)
		f.FirmwareType = firmwareType
		f.FirmwareList = make([]*entitys.CheckOtaVersionResponse, 0)
		for _, firmware := range firmwares {
			otaInfo := entitys.CheckOtaVersion_Pd2E(firmware, lang)
			if otaInfo.HasUpdate {
				f.HasUpdate = true
				hasUpdate = true
			}
			f.FirmwareList = append(f.FirmwareList, otaInfo)
		}
		list = append(list, f)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].FirmwareType < list[j].FirmwareType
	})
	iotgin.ResSuccess(c, map[string]interface{}{
		"isAuto":        true,
		"hasUpdate":     hasUpdate,
		"isAutoUpgrade": isAutoUpgrade,
		"otaState":      otaState,
		"progress":      progress,
		"list":          list,
	})
}

// SetAutoUpgradeSwitch 设置自动升级授权
// @Summary 设置自动升级授权
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.SetAutoUpgradeRequest true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/ota/setAutoUpgrade [post]
func (s *OtaController) SetAutoUpgradeSwitch(c *gin.Context) {
	req := entitys.SetAutoUpgradeRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.DeviceId == "" {
		iotgin.ResBadRequest(c, "deviceId")
		return
	}
	deviceInfo := map[string]interface{}{
		iotconst.FIELD_IS_AUTH_UPGRADE: req.IsAutoUpgrade,
	}
	//设置自动升级开关
	rdCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+req.DeviceId, deviceInfo)
	if rdCmd.Err() != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
