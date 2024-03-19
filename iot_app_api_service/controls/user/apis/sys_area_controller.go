package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	apiservice "cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Areacontroller SysAreaController

type SysAreaController struct{} //部门操作控制器

var sysAreaServices = apiservice.SysAreaService{}

// @Summary 获取区域级联列表
// @Description
// @Tags APP
// @Accept application/json
// @Param parentId path string true "parentId"
// @Param showChild path string true "showChild"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/treeData/{parentId}/{showChild} [get]
func (s SysAreaController) GetAreas(c *gin.Context) {
	pid := c.Param("parentId")
	if pid == "" {
		iotgin.ResBadRequest(c, "parentId")
		return
	}
	showChild := c.Param("showChild")
	showChildBool := false
	if showChild == "" {
		iotgin.ResBadRequest(c, "showChild")
		return
	}
	if showChild == "true" {
		showChildBool = true
	}
	pidInt64, _ := iotutil.ToInt64AndErr(pid)
	res, _, err := sysAreaServices.SetContext(controls.WithUserContext(c)).QuerySysAreaList(entitys.SysAreaQuery{
		Query: &entitys.SysAreaFilter{
			Pid:       &pidInt64,
			ShowChild: showChildBool,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, s.convertTreeData(pidInt64, res))
}

type AreaTreeData struct {
	entitys.SysAreaEntitys
	Children *[]*AreaTreeData `json:"children"`
}

func (s SysAreaController) convertTreeData(pid int64, areaList []*entitys.SysAreaEntitys) []*AreaTreeData {
	treeList := func() []*AreaTreeData {
		treeList := []*AreaTreeData{}
		flatPtr := []*AreaTreeData{}
		for _, src := range areaList {
			area := entitys.SysAreaEntitys{
				Id:              src.Id,
				Pid:             src.Pid,
				Level:           src.Level,
				Path:            src.Path,
				Code:            src.Code,
				AreaNumber:      src.AreaNumber,
				AreaPhoneNumber: src.AreaPhoneNumber,
				Abbreviation:    src.Abbreviation,
				Iso:             src.Iso,
				ChineseName:     src.ChineseName,
				EnglishName:     src.EnglishName,
				Pinyin:          src.Pinyin,
				Name:            src.Name,
			}
			t := AreaTreeData{
				SysAreaEntitys: area,
			}
			flatPtr = append(flatPtr, &t)
		}
		for m := range flatPtr {
			for n := range flatPtr {
				if flatPtr[m].Id == flatPtr[n].Pid {
					if flatPtr[m].Children == nil {
						flatPtr[m].Children = &[]*AreaTreeData{}
					}
					*(flatPtr[m].Children) = append(*(flatPtr[m].Children), flatPtr[n])
				}
			}
		}
		for _, j := range flatPtr {
			if j.Pid == pid {
				treeList = append(treeList, j)
			}
		}
		return treeList
	}()
	return treeList
}

// @Summary 获取区域列表
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/list [get]
func (SysAreaController) QueryList(c *gin.Context) {
	var filter entitys.SysAreaQuery
	res, total, err := sysAreaServices.SetContext(controls.WithUserContext(c)).QuerySysAreaTwoLevelList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, 0)
}

// @Summary 获取设备功能设置列表
// @Description
// @Tags 设备控制面板
// @Accept application/json
// @Param id query string true "Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/detail/{id} [post]
func (SysAreaController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := sysAreaServices.GetSysAreaDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 修改区域
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/edit [post]
func (SysAreaController) Edit(c *gin.Context) {
	var req entitys.SysAreaEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := sysAreaServices.UpdateSysArea(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// @Summary 新增区域
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/add [post]
func (SysAreaController) Add(c *gin.Context) {
	var req entitys.SysAreaEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := sysAreaServices.AddSysArea(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// @Summary 删除区域
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/area/delete [post]
func (SysAreaController) Delete(c *gin.Context) {
	var req entitys.SysAreaFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = sysAreaServices.DeleteSysArea(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
