package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Documentscontroller OpmDocumentsController

type OpmDocumentsController struct{} //部门操作控制器

var documentsServices = apiservice.OpmDocumentsService{}

// 查询物模型文档设置列表
func (OpmDocumentsController) QueryList(c *gin.Context) {
	var filter entitys.OpmDocumentsQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.OpmDocumentsFilter{}
	}
	res, total, err := documentsServices.SetContext(controls.WithUserContext(c)).QueryOpmDocumentsList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 查询物模型文档设置列表
func (OpmDocumentsController) QueryListByProductId(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	res, _, err := documentsServices.SetContext(controls.WithUserContext(c)).QueryOpmDocumentsListByProductId(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 查询文档设置详情
func (OpmDocumentsController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := documentsServices.SetContext(controls.WithUserContext(c)).GetOpmDocumentsDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 新增文档设置
func (s *OpmDocumentsController) Add(c *gin.Context) {
	//参数解析和验证
	docCode := c.PostForm("docCode")
	docName := c.PostForm("docName")
	productIdStr := c.PostForm("productId")
	dataOriginStr := c.PostForm("dataOrigin")
	remark := c.PostForm("remark")
	productId, _ := strconv.ParseInt(productIdStr, 10, 64)

	if docCode == "" || docName == "" || productId == 0 {
		iotgin.ResBadRequest(c, "")
		return
	}

	ctx := controls.WithUserContext(c)
	proSvc := apiservice.OpmProductService{Ctx: ctx}
	proInfo, err := proSvc.GetOpmProductDetail(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}

	var dataOrigin int32 = 1
	if dataOriginStr == "" {
		dataOrigin, _ = iotutil.ToInt32Err(dataOriginStr)
	}
	//两个文件，一个面板文件，一个缩略图
	form, err := c.MultipartForm()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	filePath := ""
	fileName := ""
	fileKey := ""
	fileSize := int32(0)

	hasPanel := false
	files := form.File["file"]
	for _, file := range files {
		f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "pdf")
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		} else {
			filePath = f.FullPath
			fileName = file.Filename
			fileKey = f.Key
			fileSize = int32(file.Size)
			hasPanel = true
			break
		}
	}
	if !hasPanel {
		iotgin.ResErrCli(c, errors.New("请确保控制面板和预览图是否正确上传"))
		return
	}
	req := &protosService.OpmDocuments{
		DataOrigin: dataOrigin,
		OriginId:   productId,
		OriginKey:  proInfo.ProductKey,
		DocCode:    docCode,
		DocName:    docName,
		Remark:     remark,
		FilePath:   filePath,
		FileName:   fileName,
		FileSize:   fileSize,
		FileKey:    fileKey,
		CreatedBy:  controls.GetUserId(c),
	}
	rep, err := rpc.ClientOpmDocumentsService.Create(ctx, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	if filePath != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmDocuments, iotutil.ToString(req.Id), filePath)
	}
	iotgin.ResSuccessMsg(c)
}

// 编辑文档设置
func (s *OpmDocumentsController) Edit(c *gin.Context) {
	//参数解析和验证
	idStr := c.PostForm("id")
	docCode := c.PostForm("docCode")
	docName := c.PostForm("docName")
	productIdStr := c.PostForm("productId")
	dataOriginStr := c.PostForm("dataOrigin")
	remark := c.PostForm("remark")
	productId, _ := strconv.ParseInt(productIdStr, 10, 64)

	if idStr != "" || docCode == "" || docName == "" || productId == 0 {
		iotgin.ResBadRequest(c, "")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}

	ctx := controls.WithUserContext(c)
	proSvc := apiservice.OpmProductService{Ctx: ctx}
	proInfo, err := proSvc.GetOpmProductDetail(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}

	var dataOrigin int32 = 1
	if dataOriginStr == "" {
		dataOrigin, _ = iotutil.ToInt32Err(dataOriginStr)
	}
	filePath := ""
	fileName := ""
	fileKey := ""
	fileSize := int32(0)

	form, err := c.MultipartForm()
	if err == nil {
		hasPanel := false
		files := form.File["file"]
		for _, file := range files {
			f, err := apis.SaveFileToOSS(c, file, apis.ControlPanel, "pdf")
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			} else {
				filePath = f.FullPath
				fileName = file.Filename
				fileKey = f.Key
				fileSize = int32(file.Size)
				hasPanel = true
				break
			}
		}
		if !hasPanel {
			iotgin.ResErrCli(c, errors.New("请确保控制面板和预览图是否正确上传"))
			return
		}
	}
	req := &protosService.OpmDocuments{
		Id:         id,
		DataOrigin: dataOrigin,
		OriginId:   productId,
		OriginKey:  proInfo.ProductKey,
		DocCode:    docCode,
		DocName:    docName,
		Remark:     remark,
		FilePath:   filePath,
		FileName:   fileName,
		FileSize:   fileSize,
		FileKey:    fileKey,
		UpdatedBy:  controls.GetUserId(c),
	}
	rep, err := rpc.ClientOpmDocumentsService.Update(ctx, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 删除文档设置
func (OpmDocumentsController) Delete(c *gin.Context) {
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
	err = documentsServices.SetContext(controls.WithUserContext(c)).DeleteOpmDocuments(entitys.OpmDocumentsFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
