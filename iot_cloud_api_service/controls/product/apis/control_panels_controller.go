package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	apis2 "cloud_platform/iot_cloud_api_service/controls/lang/apis"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"
	"strconv"

	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"

	"cloud_platform/iot_cloud_api_service/controls/common/apis"

	"github.com/gin-gonic/gin"
)

var ControlpanelsController ControlPanelsController

// ControlPanelController ctrl
type ControlPanelsController struct{}

func (ct *ControlPanelsController) Add(c *gin.Context) {
	//参数解析和验证
	name := c.PostForm("name")
	nameEn := c.PostForm("nameEn")
	productIdStr := c.PostForm("productId")
	productTypeIdStr := c.PostForm("productTypeId")
	desc := c.PostForm("desc")
	status := c.PostForm("status")
	productId, _ := strconv.ParseInt(productIdStr, 10, 64)
	productTypeId, err2 := iotutil.ToInt64AndErr(productTypeIdStr)
	if err2 != nil || name == "" {
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
	previewSize := int32(0)
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
			previewSize = int32(f.Size)
			hasPreview = true
			break
		}
	}
	if !hasPreview || !hasPanel {
		iotgin.ResErrCli(c, errors.New("请确保控制面板和预览图是否正确上传"))
		return
	}
	req := &protosService.PmControlPanels{
		Id:            controlPanelId,
		Name:          name,
		NameEn:        nameEn,
		Lang:          "zh_CN",
		Desc:          desc,
		Url:           panelFile,
		UrlName:       panelFileName,
		PanelSize:     panelFileSize,
		PanelKey:      panelFileKey,
		PreviewUrl:    panelPreview,
		PreviewName:   previewName,
		PreviewSize:   previewSize,
		ProductTypeId: productTypeId,
		ProductId:     productId,
		LangFileName:  langFileName,
		CreatedBy:     controls.GetUserId(c),
	}
	if status == "" {
		req.Status = 2
	} else {
		req.Status = iotutil.ToInt32(status)
	}
	rep, err := rpc.ClientControlPanelsService.Create(ctx, req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_control_panels", rep.Data, "name", name, nameEn)
	iotgin.ResSuccessMsg(c)
}

func (ct *ControlPanelsController) Update(c *gin.Context) {
	//参数解析和验证
	idStr := c.PostForm("id")
	name := c.PostForm("name")
	nameEn := c.PostForm("nameEn")
	desc := c.PostForm("desc")
	var productId int64
	var productTypeId int64
	productIdStr := c.PostForm("productId")
	if productIdStr != "" {
		id, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			iotgin.ResBadRequest(c, "")
			return
		}
		productId = id
	}
	productTypeIdStr := c.PostForm("productTypeId")
	if productTypeIdStr != "" {
		id, err := strconv.ParseInt(productTypeIdStr, 0, 64)
		if err != nil {
			iotgin.ResBadRequest(c, "")
			return
		}
		productTypeId = id
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || name == "" || id == 0 {
		iotgin.ResBadRequest(c, "")
		return
	}

	//两个文件，一个面板文件，一个缩略图
	panelFile := ""
	panelFileName := ""
	panelFileKey := ""
	panelFileSize := int32(0)
	panelPreview := ""
	previewName := ""
	previewSize := int32(0)
	langFileName := ""
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
				previewSize = int32(f.Size)
				break
			}
		}
	}

	req := &protosService.PmControlPanels{
		Id:            id,
		Name:          name,
		NameEn:        nameEn,
		Lang:          "zh_CN",
		Desc:          desc,
		Url:           panelFile,
		PanelKey:      panelFileKey,
		UrlName:       panelFileName,
		PanelSize:     panelFileSize,
		PreviewUrl:    panelPreview,
		PreviewName:   previewName,
		PreviewSize:   previewSize,
		ProductTypeId: productTypeId,
		ProductId:     productId,
		LangFileName:  langFileName,
		UpdatedBy:     controls.GetUserId(c),
	}
	rep, err := rpc.ClientControlPanelsService.Update(context.Background(), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_control_panels", rep.Data, "name", name, nameEn)
	// 更新产品控制面板时，删除缓存控制面板更新信息
	//if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.CONTROL_PANEL_IS_UPDATE, req.ProductId)); err != nil {
	//	iotgin.ResErrCli(c, err)
	//	return
	//}
	//清理所有使用该面板的产品面板缓存
	go ct.ClearPanelCached(id)
	iotgin.ResSuccessMsg(c)
}

// ClearPanelCached 清理面板缓存
func (ct *ControlPanelsController) ClearPanelCached(id int64) {
	iotlogger.LogHelper.Infof("开始清理面板缓存，id：%v", id)
	defer iotutil.PanicHandler("面板缓存清理失败", id)
	langKeysCached := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG_KEYS, id)
	keysCmd := iotredis.GetClient().HGetAll(context.Background(), langKeysCached)
	if keysCmd.Err() == nil {
		keys := []string{}
		//分次执行
		limit := 100
		for k, _ := range keysCmd.Val() {
			keys = append(keys, k)
			//每100条删除一次
			if len(keys) >= limit {
				iotredis.GetClient().Del(context.Background(), keys...)
				keys = make([]string, 0)
			}
		}
		//剩余不足一页缓存删除
		if len(keys) > 0 {
			iotredis.GetClient().Del(context.Background(), keys...)
		}
	}
}

func (ct *ControlPanelsController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqsvc := &protosService.PmControlPanelsFilter{Id: id}
	resp, err := rpc.ClientControlPanelsService.FindById(context.Background(), reqsvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.PmControlPanelsDetails_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

func (ct *ControlPanelsController) GetList(c *gin.Context) {
	var req entitys.PmControlPanelsQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqsvc := &protosService.PmControlPanelsListRequest{
		Page:      int64(req.Page),
		PageSize:  int64(req.Limit),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	if req.Query != nil {
		reqsvc.Query = &protosService.PmControlPanelsListFilter{
			Name: req.Query.Name,
		}
		if req.Query.ProductTypeId != "" {
			ProductTypeId, err := strconv.ParseInt(req.Query.ProductTypeId, 0, 64)
			if err == nil {
				reqsvc.Query.ProductTypeId = ProductTypeId
			}
		}
		if req.Query.ProductId != "" {
			ProductId, err := strconv.ParseInt(req.Query.ProductId, 0, 64)
			if err == nil {
				reqsvc.Query.ProductId = ProductId
			}
		}
		if req.Query.Status > 0 {
			reqsvc.Query.Status = int32(req.Query.Status)
		}
		if req.Query.Associate > 0 {
			reqsvc.Query.Associate = int32(req.Query.Associate)
		}
	}
	resp, err := rpc.ClientControlPanelsService.Lists(context.Background(), reqsvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}

	var ret []*entitys.PmControlPanelsDetailsEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.PmControlPanelsDetails_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

func (ct *ControlPanelsController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	req := &protosService.PmControlPanels{Id: id}
	resp, err := rpc.ClientControlPanelsService.DeleteById(context.Background(), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (ct *ControlPanelsController) SetStatus(c *gin.Context) {
	var req entitys.PmControlPanelsStatusEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.Status < 1 || req.Status > 2 {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	reqsvr := &protosService.PmControlPanelsUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.PmControlPanels{
			Id:     int64(id),
			Status: req.Status,
		},
	}
	resp, err := rpc.ClientControlPanelsService.UpdateFields(context.Background(), reqsvr)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccessMsg(c)
}
