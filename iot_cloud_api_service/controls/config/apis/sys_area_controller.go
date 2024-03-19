package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/config/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/config/services"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Areacontroller SysAreaController

type SysAreaController struct{} //部门操作控制器

var sysAreaServices = apiservice.SysAreaService{}

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
	pidInt64 := iotutil.ToInt64(pid)
	res, _, err := sysAreaServices.QuerySysAreaList(entitys.SysAreaQuery{
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

func (SysAreaController) QueryList(c *gin.Context) {
	var filter entitys.SysAreaQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := sysAreaServices.QuerySysAreaList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

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
