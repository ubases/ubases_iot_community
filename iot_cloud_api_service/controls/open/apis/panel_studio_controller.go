package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	apis2 "cloud_platform/iot_cloud_api_service/controls/lang/apis"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var Panelcontroller OpmPanelController

type OpmPanelController struct{} //部门操作控制器

var panelServices = apiservice.OpmPanelService{}

func (OpmPanelController) QueryList(c *gin.Context) {
	var filter entitys.OpmPanelQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.OpmPanelFilter{}
	}
	filter.Query.TenantId = controls.GetTenantId(c)

	res, total, err := panelServices.SetContext(controls.WithUserContext(c)).QueryOpmPanelList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmPanelController) QueryDropDownList(c *gin.Context) {
	var filter entitys.OpmPanelQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := panelServices.SetContext(controls.WithUserContext(c)).QueryOpmPanelList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmPanelController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := panelServices.SetContext(controls.WithUserContext(c)).GetOpmPanelDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmPanelController) Edit(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).UpdateOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) EditStudio(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).UpdateOpmPanelStudio(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) Add(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//req.CreatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.CreatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).AddOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) Delete(c *gin.Context) {
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
	err = panelServices.SetContext(controls.WithUserContext(c)).DeleteOpmPanel(entitys.OpmPanelFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmPanelController) SetStatus(c *gin.Context) {
	var req entitys.OpmPanelFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = panelServices.SetContext(controls.WithUserContext(c)).SetStatusOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// AddCustom 面板新增，上传翻译文件、面板文件
func (OpmPanelController) AddCustom(c *gin.Context) {
	var (
		//参数解析和验证
		name                   = c.PostForm("name")
		productIdStr           = c.PostForm("productId")
		productTypeIdStr       = c.PostForm("baseProductId")
		status                 = c.PostForm("status")
		desc                   = c.PostForm("desc")
		code                   = c.PostForm("code")
		panelType        int32 = 3 //c.PostForm("panelType")
		productId, _           = strconv.ParseInt(productIdStr, 10, 64)
		productTypeId, _       = iotutil.ToInt64AndErr(productTypeIdStr)
		tenantId               = controls.GetTenantId(c)
	)
	if name == "" {
		iotgin.ResBadRequest(c, "")
		return
	}
	//两个文件，一个面板文件，一个缩略图
	form, err := c.MultipartForm()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	controlPanelId := iotutil.GetNextSeqInt64()
	ctx := controls.WithUserContext(c)
	//最先检查语言包是否有效
	langFiles := form.File["langFile"]
	langFileName := ""
	if langFiles != nil && len(langFiles) > 0 {
		langFileName = langFiles[0].Filename
		err = apis2.SaveLangByFile(4, controlPanelId, name, langFiles[0], ctx)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
	}
	//上传控制页面
	panelFile := ""
	panelFileName := ""
	panelFileKey := ""
	panelFileSize := int32(0)

	hasPanel := false
	files := form.File["panelFile"]
	for _, file := range files {
		f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "zip")
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		} else {
			panelFile = f.FullPath
			panelFileName = file.Filename
			panelFileKey = f.Key
			panelFileSize = int32(file.Size)
			hasPanel = true
			break
		}
	}
	//上传预览图
	panelPreview := ""
	previewName := ""
	hasPreview := false
	files2 := form.File["panelPreview"]
	for _, file := range files2 {
		f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "png", "jpeg", "jpg")
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		} else {
			panelPreview = f.FullPath
			previewName = f.Name
			hasPreview = true
			break
		}
	}
	if !hasPreview || !hasPanel {
		iotgin.ResErrCli(c, errors.New("请确保控制面板和预览图是否正确上传"))
		return
	}

	req := entitys.OpmPanelEntitys{
		Id:            controlPanelId,
		TenantId:      tenantId,
		PanelName:     name,
		PanelType:     panelType,
		PanelUrl:      panelFile,
		PanelUrlName:  panelFileName,
		PanelSize:     panelFileSize,
		PanelKey:      panelFileKey,
		PreviewName:   previewName,
		PreviewUrl:    panelPreview,
		ProductId:     productId,
		BaseProductId: productTypeId,
		LangFileName:  langFileName,
		Remark:        desc,
		PanelCode:     code,
	}
	if status == "" {
		req.Status = 1
	} else {
		req.Status = iotutil.ToInt32(status)
	}

	id, err := panelServices.SetContext(controls.WithUserContext(c)).AddOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// EditCustom 面板修改，上传翻译文件、面板文件
func (OpmPanelController) EditCustom(c *gin.Context) {
	var (
		//参数解析和验证
		idStr            = c.PostForm("id")
		id, _            = iotutil.ToInt64AndErr(idStr)
		name             = c.PostForm("name")
		productIdStr     = c.PostForm("productId")
		productTypeIdStr = c.PostForm("baseProductId")
		desc             = c.PostForm("desc")
		code             = c.PostForm("code")
		productId, _     = strconv.ParseInt(productIdStr, 10, 64)
		productTypeId, _ = iotutil.ToInt64AndErr(productTypeIdStr)
		tenantId         = controls.GetTenantId(c)
		panelFile        = ""
		panelFileName    = ""
		panelFileKey     = ""
		panelFileSize    = int32(0)
		panelPreview     = ""
		previewName      = ""
		langFileName     = ""
	)
	if name == "" {
		iotgin.ResBadRequest(c, "")
		return
	}
	//编辑时不一定上传
	form, err := c.MultipartForm()
	if err == nil {
		ctx := controls.WithUserContext(c)
		//最先检查语言包是否有效
		langFiles := form.File["langFile"]
		if langFiles != nil && len(langFiles) > 0 {
			langFileName = langFiles[0].Filename
			err = apis2.SaveLangByFile(4, id, name, langFiles[0], ctx)
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			}
		}

		files := form.File["panelFile"]
		for _, file := range files {
			f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "zip")
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			} else {
				panelFile = f.FullPath
				panelFileName = file.Filename
				panelFileKey = f.Key
				panelFileSize = int32(file.Size)
				break
			}
		}
		files2 := form.File["panelPreview"]
		for _, file := range files2 {
			f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "png", "jpeg", "jpg")
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			} else {
				panelPreview = f.FullPath
				previewName = f.Name
				break
			}
		}
	}

	req := entitys.OpmPanelEntitys{
		Id:            id,
		TenantId:      tenantId,
		PanelName:     name,
		PanelType:     3,
		PanelUrl:      panelFile,
		PanelUrlName:  panelFileName,
		PanelSize:     panelFileSize,
		PanelKey:      panelFileKey,
		PreviewName:   previewName,
		PreviewUrl:    panelPreview,
		ProductId:     productId,
		BaseProductId: productTypeId,
		LangFileName:  langFileName,
		Remark:        desc,
		PanelCode:     code,
	}
	_, err = panelServices.SetContext(controls.WithUserContext(c)).UpdateOpmPanelV2(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	go apiservice.ClearProductPanelLangByPanelId(req.Id)
	iotgin.ResSuccessMsg(c)
}
