package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/h2non/filetype"
	"github.com/xuri/excelize/v2"

	"github.com/gin-gonic/gin"
)

var ResourcesPackagecontroller LangResourcesPackageController

type LangResourcesPackageController struct{} //部门操作控制器

// Get 查询详情
func (s *LangResourcesPackageController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.LangResourcePackageFilter{Id: id}
	resp, err := rpc.ClientLangResourcesPackageService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ctx := controls.WithUserContext(c)
	ret := entitys.LangResourcePackage_pb2e(resp.Data[0])
	//查询当前用户
	if resp.Data[0].UpdatedBy != 0 {
		resUser, err := rpc.ClientSysUserService.Find(ctx, &proto.SysUserFilter{Id: resp.Data[0].UpdatedBy})
		if err == nil {
			ret.UploadUser = resUser.Data[0].UserNickname
		}
	}
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *LangResourcesPackageController) List(c *gin.Context) {
	var req entitys.LangResourcePackageQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.LangResourcePackageListRequest{}
	if req.Query != nil {
		reqSvc.Query = &protosService.LangResourcePackage{
			AppTemplateId:   req.Query.AppTemplateId,
			AppTemplateType: req.Query.AppTemplateType,
			PackageName:     req.Query.PackageName,
			BelongType:      3,
		}
	}
	resp, err := rpc.ClientLangResourcesPackageService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	appTemplates, err := rpc.ClientOemAppTemplateService.Lists(context.Background(), &proto.OemAppTemplateListRequest{})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	appTemplateMap := map[int64]string{}
	for _, r := range appTemplates.Data {
		appTemplateMap[r.Id] = r.Version
	}
	resultList := make([]*entitys.LangResourcePackageEntitys, 0)
	for _, d := range resp.Data {
		resObj := entitys.LangResourcePackage_pb2e(d)
		if d.AppTemplateId != 0 {
			resObj.AppTemplateVersion = appTemplateMap[d.AppTemplateId]
		}
		resultList = append(resultList, resObj)
	}
	iotgin.ResPageSuccess(c, resultList, resp.Total, int(req.Page))
}

// Add 新增
func (s *LangResourcesPackageController) Add(c *gin.Context) {
	//参数解析和验证
	packageName := c.PostForm("packageName")
	belongType := c.PostForm("belongType")
	appTemplateType := c.PostForm("appTemplateType") //APP模板类型
	appTemplateId := c.PostForm("appTemplateId")     //APP模板Id

	belongTypeInt, err := iotutil.ToInt32Err(belongType)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appTemplateIdInt, err := iotutil.ToInt64AndErr(appTemplateId)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appTemplateTypeInt, err := iotutil.ToInt32Err(appTemplateType)
	req := entitys.LangResourcePackageEntitys{
		PackageName:     packageName,
		BelongType:      belongTypeInt,
		AppTemplateId:   appTemplateIdInt,
		AppTemplateType: appTemplateTypeInt,
	}

	multiFiles, _ := c.MultipartForm()
	file := multiFiles.File["file"][0]
	req.CreatedBy = controls.GetUserId(c)
	err = s.SaveLangByFile(req, file, controls.WithUserContext(c))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//清理缓存
	cachedKey := persist.GetRedisKey(iotconst.APP_COMMON_LANG, appTemplateId)
	if err := iotredis.GetClient().Del(context.Background(), cachedKey).Err(); err != nil {
		iotlogger.LogHelper.Errorf("缓存删除失败，LangResourcesPackageController.Add，err: %v", err.Error())
	}
	iotgin.ResSuccessMsg(c)
}

// Update 修改
func (s *LangResourcesPackageController) Update(c *gin.Context) {
	//参数解析和验证
	packageId := c.PostForm("id")
	packageName := c.PostForm("packageName")
	belongType := c.PostForm("belongType")
	appTemplateType := c.PostForm("appTemplateType") //APP模板类型
	appTemplateId := c.PostForm("appTemplateId")     //APP模板Id

	belongTypeInt, err := iotutil.ToInt32Err(belongType)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appTemplateIdInt, err := iotutil.ToInt64AndErr(appTemplateId)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	packageIdInt, err := iotutil.ToInt64AndErr(packageId)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appTemplateTypeInt, err := iotutil.ToInt32Err(appTemplateType)
	req := entitys.LangResourcePackageEntitys{
		Id:              packageIdInt,
		PackageName:     packageName,
		BelongType:      belongTypeInt,
		AppTemplateId:   appTemplateIdInt,
		AppTemplateType: appTemplateTypeInt,
	}

	multiFiles, err := c.MultipartForm()
	if err == nil {
		files := multiFiles.File["file"]
		if files != nil && len(files) > 0 {
			file := multiFiles.File["file"][0]
			req.UpdatedBy = controls.GetUserId(c)
			err = s.SaveLangByFile(req, file, controls.WithUserContext(c))
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			}
		} else {
			s.UpdateResourcesPackage(req, controls.WithUserContext(c))
		}
	} else {
		s.UpdateResourcesPackage(req, controls.WithUserContext(c))
	}
	//清理缓存
	cachedKey := persist.GetRedisKey(iotconst.APP_COMMON_LANG, appTemplateId)
	if err := iotredis.GetClient().Del(context.Background(), cachedKey).Err(); err != nil {
		iotlogger.LogHelper.Errorf("缓存删除失败，LangResourcesPackageController.Add，key: %v, err: %v", cachedKey, err.Error())
	}
	iotgin.ResSuccessMsg(c)
}

// Delete 删除
func (s *LangResourcesPackageController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.LangResourcePackage{Id: id}
	resp, err := rpc.ClientLangResourcesPackageService.DeleteById(context.Background(), reqSvc)
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

// SaveLangByFile 保持翻译从文件
func (s *LangResourcesPackageController) SaveLangByFile(req entitys.LangResourcePackageEntitys, file *multipart.FileHeader, ctx context.Context) error {
	fileOpen, err := file.Open()
	realType, err := filetype.MatchReader(fileOpen)
	if err != nil {
		return err
	}
	realExt := realType.Extension
	wantType := []string{"xls", "xlsx"}
	//如果没有设置想要的后缀，则不限制，如果设置则以设备为准  realtype.MIME.Subtype
	if len(wantType) != 0 && !iotutil.ArraysExistsString(wantType, realExt) {
		return errors.New(fmt.Sprintf("文件类型错误,要求%s,实际%s", strings.Join(wantType, "、"), realExt))
	}
	fileOpen.Seek(0, 0)
	f, err := excelize.OpenReader(fileOpen)
	if err != nil {
		return err
	}

	sheetName := "Sheet1"
	rows, _ := f.GetRows(sheetName)
	if len(rows) <= 2 {
		return errors.New("导入的Excel无任何数据或者工作表名称不为“Sheet1”")
	}

	//获取语言类型
	var langTypes []*struct {
		Flag     string
		LangDesc string //描述
		LangType string //语言表示（en、cn、jp）
	}
	for _, flag := range columnFlags {
		langType, _ := f.GetCellValue(sheetName, fmt.Sprintf("%s1", flag))
		langDesc, _ := f.GetCellValue(sheetName, fmt.Sprintf("%s2", flag))
		if langType == "" {
			break
		}
		langTypes = append(langTypes, &struct {
			Flag     string
			LangDesc string
			LangType string
		}{flag, langDesc, langType})
	}
	if len(langTypes) == 0 {
		return errors.New("未获取到任何语言翻译")
	}

	saveMap := map[string]*proto.LangResourcesList{}
	codeCheck := map[string]int{}
	repeatCode := ""
	for i := 2; i <= len(rows); i++ {
		colCode, _ := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", "A", i))
		if _, ok := codeCheck[colCode]; ok {
			repeatCode = colCode
			break
		}
		codeCheck[colCode] = 1
		for _, lang := range langTypes {
			colValue, _ := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", lang.Flag, i))
			//TODO colValue is null 是否需要插入到数据库呢？
			if colValue == "" {
				continue
			}
			if _, ok := saveMap[lang.LangType]; ok {
				saveMap[lang.LangType].List = append(saveMap[lang.LangType].List, &proto.LangResources{
					BelongType: req.BelongType,
					BelongId:   0,
					Lang:       lang.LangType,
					Code:       colCode,
					Value:      colValue,
				})
			} else {
				saveMap[lang.LangType] = &proto.LangResourcesList{
					List: []*proto.LangResources{
						{
							BelongType: req.BelongType,
							BelongId:   0,
							Lang:       lang.LangType,
							Code:       colCode,
							Value:      colValue,
						},
					},
				}
			}
		}
	}
	if repeatCode != "" {
		return errors.New(fmt.Sprintf("重复编码【%s】", repeatCode))
	}

	saveObj := proto.LangResourcePackage{
		Id:              req.Id,
		BelongType:      req.BelongType,
		BelongId:        0,
		PackageName:     req.PackageName,
		LangResource:    saveMap,
		AppTemplateId:   req.AppTemplateId,
		AppTemplateType: req.AppTemplateType,
		FileSize:        file.Size,
		FileName:        file.Filename,
		CreatedBy:       req.CreatedBy,
		UpdatedBy:       req.UpdatedBy,
	}
	res, err := rpc.ClientLangResourcesPackageService.CreateV2(ctx, &saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// SaveLangByFile 保持翻译从文件
func (s *LangResourcesPackageController) UpdateResourcesPackage(req entitys.LangResourcePackageEntitys, ctx context.Context) error {
	saveObj := proto.LangResourcePackage{
		Id:              req.Id,
		PackageName:     req.PackageName,
		AppTemplateId:   req.AppTemplateId,
		AppTemplateType: req.AppTemplateType,
		UpdatedBy:       req.UpdatedBy,
	}
	res, err := rpc.ClientLangResourcesPackageService.Update(ctx, &saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}
