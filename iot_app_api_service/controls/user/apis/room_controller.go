package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"strings"

	"github.com/gin-gonic/gin"
)

var Roomcontroller RoomController

type RoomController struct {
} //家庭房间操作控制器

var roomServices = services.AppRoomService{}

// @Summary 添加家庭房间
// @Description
// @Tags room
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/add [post]
func (RoomController) Add(c *gin.Context) {
	req := entitys.UcHomeRoomEntitys{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	//BindQuery
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "homeId  is empty")
		return
	}
	name := req.RoomName
	if strings.TrimSpace(name) == "" {
		iotgin.ResBadRequest(c, "roomName  is empty")
		return
	}
	icon := req.IconUrl
	if strings.TrimSpace(icon) == "" {
		iotgin.ResBadRequest(c, "iconUrl  is empty")
		return
	}
	userId := controls.GetUserId(c)
	err := roomServices.SetContext(controls.WithUserContext(c)).AddRoom(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 家庭房间排序
// @Description
// @Tags room
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/setSort [post]
func (RoomController) SetSort(c *gin.Context) {
	req := entitys.SetRoomSort{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "homeId  is empty")
		return
	}
	roomParamlist := req.RoomParamlist
	if strings.TrimSpace(iotutil.ToString(roomParamlist)) == "" {
		iotgin.ResBadRequest(c, "paramlist  is empty")
		return
	}
	if len(roomParamlist) == 0 {
		iotgin.ResBadRequest(c, "paramlist  is empty")
		return
	}

	err := roomServices.SetContext(controls.WithUserContext(c)).SetSort(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 家庭房间设备排序
// @Description
// @Tags room
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/setSort [post]
func (RoomController) SetDevSort(c *gin.Context) {
	req := entitys.SetRoomSort{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId := req.HomeId
	if strings.TrimSpace(homeId) == "" {
		iotgin.ResBadRequest(c, "homeId  is empty")
		return
	}
	roomParamlist := req.RoomParamlist
	if strings.TrimSpace(iotutil.ToString(roomParamlist)) == "" {
		iotgin.ResBadRequest(c, "paramlist  is empty")
		return
	}
	if len(roomParamlist) == 0 {
		iotgin.ResBadRequest(c, "paramlist  is empty")
		return
	}

	err := roomServices.SetContext(controls.WithUserContext(c)).SetDevSort(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 房间详情
// @Description
// @Tags room
// @Accept application/json
// @Param homeId query string true "家庭Id"
// @Param roomId query string true "房间Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/details/{homeId}/{roomId} [get]
func (RoomController) Details(c *gin.Context) {
	homeId := c.Param("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "homeId")
		return
	}
	roomId := c.Param("roomId")
	if roomId == "" {
		iotgin.ResBadRequest(c, "roomId")
		return
	}
	res, err := roomServices.SetContext(controls.WithUserContext(c)).Details(homeId, roomId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 删除房间
// @Description
// @Tags room
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/delete [post]
func (RoomController) Delete(c *gin.Context) {
	var req entitys.UcHomeRoomFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	userId := controls.GetUserId(c)
	err = roomServices.SetContext(controls.WithUserContext(c)).Delete(req, userId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 修改家庭房间
// @Description
// @Tags room
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/room/update [post]
func (RoomController) Update(c *gin.Context) {
	req := entitys.UcHomeRoomEntitys{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId := controls.GetUserId(c)
	err = roomServices.SetContext(controls.WithUserContext(c)).Update(req, userId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
