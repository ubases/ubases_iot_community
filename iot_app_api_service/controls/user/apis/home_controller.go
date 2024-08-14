package apis

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var Homecontroller HomeController

type HomeController struct {
} //家庭操作控制器

var homeServices = services.AppHomeService{}

// @Summary 获取家庭详情
// @Description
// @Tags home
// @Accept application/json
// @Param id path string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/details/{id} [get]
func (HomeController) Details(c *gin.Context) {
	strId := c.Param("id")
	homeId, err := iotutil.ToInt64AndErr(strId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if homeId == 0 {
		iotgin.ResBusinessP(c, "homeId is empty")
		return
	}
	userId := controls.GetUserId(c)
	// 检查登录接口中的参数
	res, rspCode, err := homeServices.SetContext(controls.WithUserContext(c)).Details(c, homeId, userId)
	if rspCode != 0 {
		iotgin.ResErrCliCustomCode(c, err, rspCode)
		return
	}
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 添加家庭
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.UcHomeEntitys true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/add [post]
func (HomeController) Add(c *gin.Context) {
	req := entitys.UcHomeEntitys{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	req.Sid = iotutil.ToString(controls.GetRegionInt(c))
	name := req.Name
	if strings.TrimSpace(name) == "" {
		iotgin.ResBadRequest(c, "名称为空")
		return
	}
	//address := req.Address
	//if strings.TrimSpace(address) == "" {
	//	iotgin.ResBadRequest(c, "地址为空")
	//	return
	//}
	city := req.City
	//if len(req.CountryId) == 0 {
	//	if strings.TrimSpace(city) == "" {
	//		iotgin.ResBadRequest(c, "城市为空")
	//		return
	//	}
	//	coordType, _ := iotutil.ToInt64AndErr(req.CoordType)
	//	if coordType == 0 {
	//		iotgin.ResBadRequest(c, "坐标类型为空")
	//		return
	//	}
	//}
	roomList := req.RoomList
	if roomList != nil && len(roomList) > 0 {
		if len(roomList) > 99 {
			iotgin.ResBadRequest(c, "最大房间数量为100个")
			return
		}
	}
	// 获取区域数据，将区域ID转换为区域名称
	if len(req.CountryId) != 0 && city == "" {
		areaIds := make([]int64, 0)
		countryId := iotutil.MapGetInt64Val(req.CountryId, 0) //geo.Country,
		if countryId != 0 {
			areaIds = append(areaIds, countryId)
		}
		provinceId := iotutil.MapGetInt64Val(req.ProvinceId, 0) //geo.Province,
		if provinceId != 0 {
			areaIds = append(areaIds, provinceId)
		}
		cityId := iotutil.MapGetInt64Val(req.CityId, 0) //geo.City,
		if cityId != 0 {
			areaIds = append(areaIds, cityId)
		}
		if len(areaIds) > 0 {
			areaMap := map[int64]string{}
			areaList, err := rpc.ClientAreaService.Lists(context.Background(), &proto.SysAreaListRequest{Query: &proto.SysArea{
				AreaIds: []int64{countryId, cityId, provinceId},
				Pid:     -1,
			}})
			lang := controls.GetLang(c)
			if err == nil {
				for _, a := range areaList.Data {
					if lang == "zh" {
						areaMap[a.Id] = a.ChineseName
					} else {
						areaMap[a.Id] = a.EnglishName
					}
				}
			}
			req.Country = iotutil.MapGetStringVal(areaMap[countryId], "")   //geo.Country,
			req.Province = iotutil.MapGetStringVal(areaMap[provinceId], "") //geo.Province,
			req.City = iotutil.MapGetStringVal(areaMap[cityId], "")         //geo.City,
		}
	} else {
		geo, err := controls.Geoip(c.ClientIP())
		areaMap := map[string]string{}
		var (
			countryId  int64 = 0
			provinceId int64 = 0
			cityId     int64 = 0
		)
		if err != nil {
			iotlogger.LogHelper.Errorf("get address by ip[%s], error:%s", c.ClientIP(), err.Error())
		} else {
			//去除尾部 省、市
			if strings.HasSuffix(geo.Province, "省") {
				geo.Province = strings.TrimSuffix(geo.Province, "省")
			}
			if strings.HasSuffix(geo.City, "市") {
				geo.City = strings.TrimSuffix(geo.City, "市")
			}
			req.City = geo.City
			req.Country = geo.Country
			req.Province = geo.Province
			areaList, err := rpc.ClientAreaService.Lists(context.Background(), &proto.SysAreaListRequest{Query: &proto.SysArea{
				EnableGetCode: true,
				Country:       geo.Country,
				City:          geo.City,
				Province:      geo.Province,
				Pid:           -1,
			}})
			if err == nil {
				for _, a := range areaList.Data {
					areaMap[a.ChineseName] = iotutil.ToString(a.Id)
					areaMap[a.EnglishName] = iotutil.ToString(a.Id)
				}
			}
			countryId = iotutil.MapGetInt64Val(areaMap[geo.Country], 0)   //geo.Country,
			provinceId = iotutil.MapGetInt64Val(areaMap[geo.Province], 0) //geo.Province,
			cityId = iotutil.MapGetInt64Val(areaMap[geo.City], 0)         //geo.City,
			req.Address = fmt.Sprintf("%v%v%v", geo.Country, geo.Province, geo.City)
		}
		req.CountryId = iotutil.ToString(countryId)
		req.ProvinceId = iotutil.ToString(provinceId)
		req.CityId = iotutil.ToString(cityId)
	}
	userId := controls.GetUserId(c)
	err := homeServices.SetContext(controls.WithUserContext(c)).AddHome(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 修改家庭信息
// @Description
// @Tags home
// @Accept application/json
// @Param id query string true "家庭Id"
// @Param data body entitys.UpdateHomeEntitys true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/update/{id} [post]
func (HomeController) Update(c *gin.Context) {
	homeId := c.Param("id")
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}
	req := entitys.UpdateHomeEntitys{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	// 获取区域数据，将区域ID转换为区域名称
	if len(req.CountryId) != 0 && req.City == "" {
		areaIds := make([]int64, 0)
		countryId := iotutil.MapGetInt64Val(req.CountryId, 0) //geo.Country,
		if countryId != 0 {
			areaIds = append(areaIds, countryId)
		}
		provinceId := iotutil.MapGetInt64Val(req.ProvinceId, 0) //geo.Province,
		if provinceId != 0 {
			areaIds = append(areaIds, provinceId)
		}
		cityId := iotutil.MapGetInt64Val(req.CityId, 0) //geo.City,
		if cityId != 0 {
			areaIds = append(areaIds, cityId)
		}
		if len(areaIds) > 0 {
			areaMap := map[int64]string{}
			areaList, err := rpc.ClientAreaService.Lists(context.Background(), &proto.SysAreaListRequest{Query: &proto.SysArea{
				AreaIds: []int64{countryId, cityId, provinceId},
				Pid:     -1,
			}})
			lang := controls.GetLang(c)
			if err == nil {
				for _, a := range areaList.Data {
					if lang == "zh" {
						areaMap[a.Id] = a.ChineseName
					} else {
						areaMap[a.Id] = a.EnglishName
					}
				}
			}
			req.Country = iotutil.MapGetStringVal(areaMap[countryId], req.Country)    //geo.Country,
			req.Province = iotutil.MapGetStringVal(areaMap[provinceId], req.Province) //geo.Province,
			req.City = iotutil.MapGetStringVal(areaMap[cityId], req.City)             //geo.City,
		}
	}
	userId := controls.GetUserId(c)
	err := homeServices.SetContext(controls.WithUserContext(c)).UpdateHome(iotutil.ToString(userId), homeId, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 删除家庭
// @Description
// @Tags home
// @Accept application/json
// @Param id query string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/delete/{id} [post]
func (HomeController) Delete(c *gin.Context) {
	homeId := c.Param("id")
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}

	userId := controls.GetUserId(c)
	code, err := homeServices.SetContext(controls.WithUserContext(c)).Delete(iotutil.ToInt64(homeId), userId, c.ClientIP())
	if code != 0 {
		iotgin.ResErrCliCustomCode(c, err, int(code))
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (HomeController) SendMsgTest(c *gin.Context) {
	homeId := c.Param("id")
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}

	userId := controls.GetUserId(c)
	code, err := homeServices.SetContext(controls.WithUserContext(c)).SendMsgTest(iotutil.ToInt64(homeId), userId, c.ClientIP())
	if code != 0 {
		iotgin.ResErrCliCustomCode(c, err, int(code))
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 发送邀请码
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.SendInvitationCode true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/sendInvitationCode [post]
func (HomeController) SendInvitationCode(c *gin.Context) {
	req := entitys.SendInvitationCode{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	userId := controls.GetUserId(c)
	appKey := c.Request.Header.Get("appKey")
	tenantId := c.GetHeader("tenantId")
	if strings.TrimSpace(appKey) == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	if strings.TrimSpace(tenantId) == "" {
		iotgin.ResBadRequest(c, "tenantId")
		return
	}
	resp, msg := homeServices.SetContext(controls.WithUserContext(c)).SendInvitationCode(req.HomeId, userId, appKey, tenantId)
	if msg != "" {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, resp)
}

// @Summary 成员加入家庭
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.JoinHome true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/joinHome [post]
func (HomeController) JoinHome(c *gin.Context) {
	req := entitys.JoinHome{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	code := req.Code
	if strings.TrimSpace(code) == "" {
		iotgin.ResBadRequest(c, "验证码为空")
		return
	}
	code = strings.ToUpper(code)
	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotgin.ResBadRequest(c, "tenantId")
		return
	}
	rspCode, err := homeServices.SetContext(controls.WithUserContext(c)).JoinHome(code, iotutil.ToString(userId), iotutil.ToString(appKey), iotutil.ToString(tenantId))
	if rspCode != 0 {
		iotgin.ResErrCliCustomCode(c, err, int(rspCode))
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 家庭成员角色设置
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.SetRole true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/setRole [post]
func (HomeController) SetRole(c *gin.Context) {
	req := entitys.SetRole{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭id为空")
		return
	}
	thirdUserId := req.UserId
	if strings.TrimSpace(thirdUserId) == "" {
		iotgin.ResBadRequest(c, "第三方用户id为空")
		return
	}
	roleType, ierr := strconv.Atoi(iotutil.ToString(req.Role))
	if ierr != nil {
		roleType = 0
	}
	if !(roleType == 2 || roleType == 3) {
		iotgin.ResBadRequest(c, "角色类型参数有误")
		return
	}

	userId := controls.GetUserId(c)
	resp := homeServices.SetContext(controls.WithUserContext(c)).SetRole(iotutil.ToInt64(userId), iotutil.ToInt64(homeId), iotutil.ToInt64(thirdUserId), iotutil.ToInt32(roleType))
	if resp != "" {
		iotgin.ResBusinessP(c, resp)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 移除家庭成员
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.RemoveMembers true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/removeMembers [post]
func (HomeController) RemoveMembers(c *gin.Context) {
	req := entitys.RemoveMembers{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭id为空")
		return
	}
	thirdUserId := req.UserId
	if strings.TrimSpace(thirdUserId) == "" {
		iotgin.ResBadRequest(c, "第三方用户id为空")
		return
	}

	userId := controls.GetUserId(c)
	resp := homeServices.SetContext(controls.WithUserContext(c)).RemoveMembers(iotutil.ToInt64(userId), iotutil.ToInt64(homeId), iotutil.ToInt64(thirdUserId), c.ClientIP())
	if resp != "" {
		iotgin.ResBusinessP(c, resp)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 转移家庭所有权
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.TransferOwnership true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/transferOwnership [post]
func (HomeController) TransferOwnership(c *gin.Context) {
	req := entitys.TransferOwnership{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭id为空")
		return
	}
	thirdUserId := req.UserId
	if strings.TrimSpace(thirdUserId) == "" {
		iotgin.ResBadRequest(c, "第三方用户id为空")
		return
	}

	userId := controls.GetUserId(c)
	resp := homeServices.SetContext(controls.WithUserContext(c)).TransferOwnership(iotutil.ToInt64(userId), iotutil.ToInt64(homeId), iotutil.ToInt64(thirdUserId))
	if resp != "" {
		iotgin.ResBusinessP(c, resp)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 退出家庭
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.Quit true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/quit [post]
func (HomeController) Quit(c *gin.Context) {
	req := entitys.Quit{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭id为空")
		return
	}

	userId := controls.GetUserId(c)
	resp := homeServices.SetContext(controls.WithUserContext(c)).Quit(iotutil.ToInt64(userId), iotutil.ToInt64(homeId))
	if resp != "" {
		iotgin.ResBusinessP(c, resp)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 家庭房间列表
// @Description
// @Tags home
// @Accept application/json
// @Param homeId query string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/roomList/{homeId} [get]
func (HomeController) RoomList(c *gin.Context) {
	homeId := c.Param("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}
	homeIdInt, err := iotutil.ToInt64AndErr(homeId)
	if err != nil {
		iotgin.ResBadRequest(c, "家庭编号异常")
		return
	}
	resp, msg := homeServices.SetContext(controls.WithUserContext(c)).RoomList(homeIdInt)
	if msg != "" {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, resp)
}

// @Summary 家庭设备列表
// @Description
// @Tags home
// @Accept application/json
// @Param homeId query string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/deviceList/{homeId} [get]
func (HomeController) DeviceList(c *gin.Context) {
	homeId := c.Param("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}
	resp, msg := homeServices.SetContext(controls.WithUserContext(c)).DeviceList(iotutil.ToInt64(homeId))
	if msg != "" {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, resp)
}

// @Summary 用户设备列表
// @Description
// @Tags home
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /user/deviceList [get]
func (HomeController) UserDeviceList(c *gin.Context) {
	userId := controls.GetUserId(c)
	if userId == 0 {
		iotgin.ResBadRequest(c, "用户编号不存在")
		return
	}
	resp, msg := homeServices.SetContext(controls.WithUserContext(c)).UserDeviceList(iotutil.ToInt64(userId))
	if msg != "" {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, resp)
}

// @Summary 获取配网信息
// @Description
// @Tags home
// @Accept application/json
// @Param homeId path string true "家庭ID"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/serverAlloc/{homeId} [get]
func (HomeController) ServerAlloc(c *gin.Context) {
	homeId := c.Param("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "家庭编号为空")
		return
	}
	serverInfo := entitys.ServerEntity{Nationcode: "CN"}
	if config.Global.DeviceMQTT.Ip != "" && config.Global.DeviceMQTT.Port > 0 {
		serverInfo.MqttIp = config.Global.DeviceMQTT.Ip
		serverInfo.MqttPort = config.Global.DeviceMQTT.Port
	}

	//区域兼容处理regionId := controls.GetRegion(c)
	mqtt, err := services.GetRegionMqttByHomeId(iotutil.ToInt64(homeId))
	if err == nil {
		if mqtt.MqttServer != "" {
			serverInfo.MqttIp = mqtt.MqttServer
			serverInfo.MqttPort = mqtt.MqttPort
		}
	}

	//TODO 特殊处理，针对老香薰机设备开通1885端口
	//if productKey != "" && config.Global.DeviceMQTT.NoAclProducts != nil && iotutil.InArray(productKey, config.Global.DeviceMQTT.NoAclProducts) {
	//	serverInfo.MqttPort = 1885
	//}
	iotgin.ResSuccess(c, serverInfo)
}

// @Summary 家庭添加设备
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AddDevParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/addDev [post]
func (HomeController) AddDev(c *gin.Context) {
	req := entitys.AddDevParam{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "家庭id为空")
		return
	}
	name := req.Name
	if strings.TrimSpace(name) == "" {
		iotgin.ResBadRequest(c, "名称为空")
		return
	}
	devId := req.DevId
	if strings.TrimSpace(devId) == "" {
		iotgin.ResBadRequest(c, "设备编号为空")
		return
	}
	model := req.Model
	if strings.TrimSpace(model) == "" {
		iotgin.ResBadRequest(c, "model为空")
		return
	}
	state := req.State
	if strings.TrimSpace(state) == "" {
		iotgin.ResBadRequest(c, "状态为空")
		return
	}

	//userId := controls.GetUserId(c)
	//resp := services.AddDev(iotutil.ToInt64(userId),iotutil.ToInt64(homeId))
	//if resp != "" {
	//	iotgin.ResBusinessP(c, resp)
	//	return
	//}
	iotgin.ResSuccessMsg(c)
}

// @Summary 家庭设备排序
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.SetDevSort true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/home/setDevSort [post]
func (HomeController) SetDevSort(c *gin.Context) {
	req := entitys.SetDevSort{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "homeId  is empty")
		return
	}
	devParamlist := req.ParamList
	if strings.TrimSpace(iotutil.ToString(devParamlist)) == "" {
		iotgin.ResBadRequest(c, "paramList  is empty")
		return
	}
	if len(devParamlist) == 0 {
		iotgin.ResBadRequest(c, "paramList  is empty")
		return
	}

	err := homeServices.SetContext(controls.WithUserContext(c)).SetDevSort(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
