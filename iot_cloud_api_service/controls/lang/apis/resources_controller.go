package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/lang/services"
	services3 "cloud_platform/iot_cloud_api_service/controls/oem/services"
	dictEntitys "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	services2 "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
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
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/h2non/filetype"

	"github.com/tealeg/xlsx"

	"github.com/xuri/excelize/v2"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Resourcescontroller LangResourcesController

type LangResourcesController struct{} //部门操作控制器

var resourcesServices = apiservice.LangResourcesService{}

var columnFlags = []string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}

var tempPath = iotconst.GetWorkTempDir() + string(filepath.Separator)

// ResourceImport 资源导入
func (s *LangResourcesController) ResourceImport(c *gin.Context) {
	//参数解析和验证
	//lang := c.PostForm("lang")
	packageName := c.PostForm("packageName")
	belongType := c.PostForm("belongType")

	belongTypeInt, err := iotutil.ToInt32Err(belongType)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	multiFiles, _ := c.MultipartForm()
	file := multiFiles.File["file"][0]
	err = SaveLangByFile(belongTypeInt, 0, packageName, file, controls.WithUserContext(c))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SaveLangByFile 保持翻译从文件
func SaveLangByFile(belongTypeInt int32, belongId int64, packageName string, file *multipart.FileHeader, ctx context.Context) error {
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
					BelongType: belongTypeInt,
					BelongId:   belongId,
					Lang:       lang.LangType,
					Code:       colCode,
					Value:      colValue,
				})
			} else {
				saveMap[lang.LangType] = &proto.LangResourcesList{
					List: []*proto.LangResources{
						&proto.LangResources{
							BelongType: belongTypeInt,
							BelongId:   belongId,
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
		BelongType:   belongTypeInt,
		BelongId:     belongId,
		PackageName:  packageName,
		LangResource: saveMap,
	}
	res, err := rpc.ClientLangResourcesPackageService.Create(ctx, &saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// ResourceImportExcelTemplate 资源导出excel模板
func (s *LangResourcesController) ResourceImportExcelTemplate(c *gin.Context) {
	//导出生成excel附件
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("data")
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.Value = "编码"

	//获取支持的语言
	lang, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_language_type)
	for k, _ := range lang.GetData() {
		cell = headerRow.AddCell()
		cell.Value = k
	}

	//保存模板临时文件
	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err := file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		iotgin.ResErrCli(c, err)
		return
	}
	//发送完文件后删除对应文件
	//defer func() {
	//	os.Remove(tempPathFile)
	//}()
	fileName := "lang-template-" + time.Now().Format("20060102150400") + ".xlsx"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)

	//删除临时文件
	go func() {
		//延时3秒删除临时文件
		defer iotutil.PanicHandler()
		time.Sleep(3 * time.Second)
		os.Remove(tempPathFile)
	}()
}

// ResourcePackageDetail 资源详情
func (LangResourcesController) ResourcePackageDetail(c *gin.Context) {
	belongType := c.Query("belongType")
	if belongType == "" {
		iotgin.ResBadRequest(c, "belongType")
		return
	}
	belongTypeInt, _ := iotutil.ToInt32Err(belongType)
	ctx := controls.WithUserContext(c)
	res, err := rpc.ClientLangResourcesPackageService.Lists(ctx, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: belongTypeInt,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResBadRequest(c, res.Message)
		return
	}
	if len(res.Data) == 0 {
		iotgin.ResSuccess(c, nil)
		return
	}
	resObj := entitys.LangResourcePackage_pb2e(res.Data[0])
	//查询当前用户
	if res.Data[0].UpdatedBy != 0 {
		resUser, err := rpc.ClientSysUserService.Find(ctx, &proto.SysUserFilter{Id: res.Data[0].UpdatedBy})
		if err == nil {
			resObj.UploadUser = resUser.Data[0].UserNickname
		}
	}
	iotgin.ResSuccess(c, resObj)
}

// ResourcePackageDetailList 资源包详情列表
func (LangResourcesController) ResourcePackageDetailList(c *gin.Context) {
	belongType := c.Query("belongType")
	if belongType == "" {
		iotgin.ResBadRequest(c, "belongType")
		return
	}
	belongTypeInt, _ := iotutil.ToInt32Err(belongType)
	ctx := controls.WithUserContext(c)
	res, err := rpc.ClientLangResourcesPackageService.Lists(ctx, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: belongTypeInt,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResBadRequest(c, res.Message)
		return
	}
	resultList := make([]*entitys.LangResourcePackageEntitys, 0)
	for i, d := range res.Data {
		resObj := entitys.LangResourcePackage_pb2e(d)
		//查询当前用户
		if res.Data[i].UpdatedBy != 0 {
			resUser, err := rpc.ClientSysUserService.Find(ctx, &proto.SysUserFilter{Id: res.Data[i].UpdatedBy})
			if err == nil {
				resObj.UploadUser = resUser.Data[0].UserNickname
			}
		}
		resultList = append(resultList, resObj)
	}

	iotgin.ResSuccess(c, resultList)
}

// GetExport 导出的get方法
func (s *LangResourcesController) GetExport(c *gin.Context) {
	//默认是APP的资源下载
	belongType := c.Query("belongType")
	if belongType == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 belongType"))
		return
	}
	belongId := c.DefaultQuery("belongId", "0")
	belongTypeInt, _ := iotutil.ToInt32Err(belongType)
	packageId := c.DefaultQuery("packageId", "0")
	packageIdInt, _ := iotutil.ToInt64AndErr(packageId)
	belongIdInt, err := iotutil.ToInt64AndErr(belongId)
	if belongIdInt == 0 && packageIdInt == 0 {
		iotgin.ResErrCli(c, errors.New("参数错误 belongType / packageId"))
		return
	}
	thisContext := controls.WithUserContext(c)
	res, err := rpc.ClientLangResourcesPackageService.Lists(thisContext, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
			Id:         packageIdInt,
		},
	})

	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}

	if len(res.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("未找到任何可导出的资源"))
		return
	}

	//导出生成excel附件
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.Value = "编码"
	for _, lang := range res.Data[0].Langs {
		cell = headerRow.AddCell()
		cell.Value = lang
	}

	rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &proto.LangResourcesListRequest{
		Query: &proto.LangResources{BelongType: belongTypeInt, BelongId: belongIdInt, PackageId: packageIdInt},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rep.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rep.Message))
		return
	}
	//map[code]map[lang]value
	resultArr := map[string]map[string]string{}

	for _, d := range rep.Data {
		if _, ok := resultArr[d.Code]; !ok {
			resultArr[d.Code] = map[string]string{"code": d.Code, d.Lang: d.Value}
		} else {
			resultArr[d.Code][d.Lang] = d.Value
		}
	}

	for _, row := range resultArr {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.Value = row["code"]

		for _, lang := range res.Data[0].Langs {
			cell = headerRow.AddCell()
			cell.Value = row[lang]
		}
	}

	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err = file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		iotgin.ResErrCli(c, err)
		return
	}
	//发送完文件后删除对应文件
	//defer func() {
	//	os.Remove(tempPathFile)
	//}()
	fileName := "lang-" + time.Now().Format("20060102150400") + ".xlsx"
	//		return fileName, tempPathFile, nil

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// CustomResourceDetail 获取自定义资源的使用次数
func (LangResourcesController) CustomResourceDetail(c *gin.Context) {
	appKey := c.Query("appKey")
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	res, err := rpc.ClientLangCustomResourceService.ResourceUseRecord(controls.WithUserContext(c), &proto.ResourceOperationRecordRequest{
		LimitDays: 3,
		AppKey:    appKey,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResBadRequest(c, res.Message)
		return
	}
	iotgin.ResSuccess(c, map[string]int64{"import": res.ImportCount, "export": res.ImportCount})
}

// CustomResourceExport 导出自定义资源（如果不存在自定义资源，则导出基础app资源）
func (s *LangResourcesController) CustomResourceExport(c *gin.Context) {
	//默认是APP的资源下载
	appId := c.Query("appKey")
	if appId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 appKey"))
		return
	}
	thisContext := controls.WithUserContext(c)
	tenantId := c.GetString("tenantId")
	userId := c.GetInt64("userId")
	appIdInt := iotutil.ToInt64(appId)
	ctx := controls.WithUserContext(c)
	oemApp, err := rpc.ClientOemAppService.FindById(ctx, &proto.OemAppFilter{Id: appIdInt})
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if oemApp.Code != 200 {
		iotgin.ResBadRequest(c, oemApp.Message)
		return
	}
	if len(oemApp.Data) == 0 {
		iotgin.ResBadRequest(c, "APP不存在")
		return
	}
	appInfo := oemApp.Data[0]
	appKey := appInfo.AppKey

	if _, exportCount, err := s.checkImportLimit(c, appKey); err != nil || exportCount >= 3 {
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		iotgin.ResBadRequest(c, "24小时内，仅限导出3次")
		return
	}

	//获取资源部信息，读取支持的语言
	res, err := rpc.ClientLangResourcesPackageService.Lists(thisContext, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: 3,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	if len(res.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("未找到任何可导出的资源"))
		return
	}

	//导出生成excel附件
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.Value = "编码"

	for _, lang := range res.Data[0].Langs {
		cell = headerRow.AddCell()
		cell.Value = lang
	}

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			//AppId:  appIdInt,
			BelongType: 3,
			AppKey:     appKey,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if crs.Code != 200 {
		iotgin.ResErrCli(c, errors.New(crs.Message))
		return
	}

	resultArr := map[string]map[string]string{}
	//如果自定义为0，则读取基础的翻译
	if len(crs.Data) == 0 {
		rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &proto.LangResourcesListRequest{
			Query: &proto.LangResources{BelongType: 3, AppTemplateId: appInfo.AppTemplateId},
		})
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		if rep.Code != 200 {
			iotgin.ResErrCli(c, errors.New(rep.Message))
			return
		}
		for _, d := range rep.Data {
			if _, ok := resultArr[d.Code]; !ok {
				resultArr[d.Code] = map[string]string{"code": d.Code, d.Lang: d.Value}
			} else {
				resultArr[d.Code][d.Lang] = d.Value
			}
		}
	} else {
		for _, d := range crs.Data {
			if _, ok := resultArr[d.Code]; !ok {
				resultArr[d.Code] = map[string]string{"code": d.Code, d.Lang: d.Value}
			} else {
				resultArr[d.Code][d.Lang] = d.Value
			}
		}
	}
	for _, row := range resultArr {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.Value = row["code"]

		for _, lang := range res.Data[0].Langs {
			cell = headerRow.AddCell()
			cell.Value = row[lang]
		}
	}

	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err = file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		iotgin.ResErrCli(c, err)
		return
	}
	fileName := "custom-lang-" + time.Now().Format("20060102150400") + ".xlsx"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(tempPathFile)

	go s.SetExportCount(tenantId, appKey, userId, 2)
}

func (s LangResourcesController) SetExportCount(tenantId, appKey string, userId int64, opType int32) error {
	defer iotutil.PanicHandler()
	crs, err := rpc.ClientLangCustomResourceService.CreateOpRecord(context.Background(), &proto.CustomerResourceRecord{
		Id:       iotutil.GetNextSeqInt64(),
		TenantId: tenantId,
		UserId:   userId,
		AppKey:   appKey,
		OpType:   opType,
	})
	if err != nil {
		return err
	}
	if crs.Code != 200 {
		return errors.New(crs.Message)
	}
	return nil
}

// CustomResourceJsonData 导出自定义资源（如果不存在自定义资源，则导出基础app资源）
func (s *LangResourcesController) CustomResourceJsonData(c *gin.Context) {
	//默认是APP的资源下载
	appKey := c.Query("appKey")
	lang := c.Query("lang")
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	if lang == "" {
		iotgin.ResBadRequest(c, "lang")
		return
	}
	thisContext := controls.WithUserContext(c)
	res, err := rpc.ClientLangResourcesPackageService.Lists(thisContext, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: 3,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	if len(res.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("未找到任何可导出的资源"))
		return
	}

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			AppKey: appKey,
			Lang:   lang,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if crs.Code != 200 {
		iotgin.ResErrCli(c, errors.New(crs.Message))
		return
	}

	resultArr := map[string]map[string]string{}
	//如果自定义为0，则读取基础的翻译
	if len(crs.Data) == 0 {
		rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &proto.LangResourcesListRequest{
			Query: &proto.LangResources{BelongType: 3},
		})
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		if rep.Code != 200 {
			iotgin.ResErrCli(c, errors.New(rep.Message))
			return
		}
		for _, d := range rep.Data {
			if _, ok := resultArr[d.Lang]; !ok {
				resultArr[d.Lang] = map[string]string{d.Code: d.Value}
			} else {
				resultArr[d.Lang][d.Code] = d.Value
			}
		}
	} else {
		for _, d := range crs.Data {
			if _, ok := resultArr[d.Lang]; !ok {
				resultArr[d.Lang] = map[string]string{d.Code: d.Value}
			} else {
				resultArr[d.Lang][d.Code] = d.Value
			}
		}
	}
	iotgin.ResSuccess(c, resultArr)
}

func (s LangResourcesController) checkImportLimit(c *gin.Context, appKey string) (importCount, exportCount int64, err error) {
	res, err := rpc.ClientLangCustomResourceService.ResourceUseRecord(controls.WithUserContext(c), &proto.ResourceOperationRecordRequest{
		LimitDays: 3,
		AppKey:    appKey,
	})
	if err != nil {
		return 0, 0, err
	}
	if res.Code != 200 {
		return 0, 0, errors.New(res.Message)
	}
	return res.ImportCount, res.ExportCount, nil
}

// CustomResourceImport 自定义资源导入  TODO 翻译的key需要进行验证！！！
func (s *LangResourcesController) CustomResourceImport(c *gin.Context) {
	//参数解析和验证
	appId := c.PostForm("appKey")
	if appId == "" {
		iotgin.ResBadRequest(c, "参数错误 appKey")
		return
	}
	appIdInt := iotutil.ToInt64(appId)
	tenantId := c.GetString("tenantId")
	userId := c.GetInt64("userId")

	ctx := controls.WithUserContext(c)
	appInfo, err := services3.OemAppService{Ctx: ctx}.GetAppInfo(appIdInt)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appKey := appInfo.AppKey

	//检查导入限制
	if importCount, _, err := s.checkImportLimit(c, appKey); err != nil || importCount >= 3 {
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		iotgin.ResBadRequest(c, "24小时内，仅限导入3次")
		return
	}
	saveMap, err := s.loadResourcesListByExcel(c, appInfo)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	saveObj := proto.ImportLangCustomResource{
		AppKey:       appKey,
		LangResource: saveMap,
	}
	res, err := rpc.ClientLangCustomResourceService.ImportCreate(ctx, &saveObj)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResBadRequest(c, res.Message)
		return
	}
	go s.SetExportCount(tenantId, appKey, userId, 1)

	//缓存清理
	cachedKey := persist.GetRedisKey(iotconst.APP_CUSTOM_LANG, appInfo.AppKey)
	if err := iotredis.GetClient().Del(context.Background(), cachedKey).Err(); err != nil {
		iotlogger.LogHelper.Errorf("缓存删除失败，LangResourcesController.CustomResourceImport，key: %v, err: %v", cachedKey, err.Error())
	}
	iotgin.ResSuccessMsg(c)
}

// 从excel加载翻译数据
func (s LangResourcesController) loadResourcesListByExcel(c *gin.Context, appInfo *protosService.OemApp) (map[string]*proto.LangCustomResourcesList, error) {
	multiFiles, _ := c.MultipartForm()
	file := multiFiles.File["file"][0]
	fileOpen, err := file.Open()
	f, _ := excelize.OpenReader(fileOpen)
	if err != nil {
		return nil, err
	}

	sheetName := "Sheet1"
	rows, _ := f.GetRows(sheetName)
	if len(rows) <= 2 {
		return nil, errors.New("导入的Excel无任何数据或者工作表名称不为“Sheet1”")
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
		return nil, errors.New("未获取到任何语言翻译")
	}

	saveMap := map[string]*proto.LangCustomResourcesList{}
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
				saveMap[lang.LangType].List = append(saveMap[lang.LangType].List, &proto.LangCustomResources{
					AppKey: appInfo.AppKey,
					AppId:  appInfo.Id,
					Lang:   lang.LangType,
					Code:   colCode,
					Value:  colValue,
				})
			} else {
				saveMap[lang.LangType] = &proto.LangCustomResourcesList{
					List: []*proto.LangCustomResources{
						{
							AppKey: appInfo.AppKey,
							AppId:  appInfo.Id,
							Lang:   lang.LangType,
							Code:   colCode,
							Value:  colValue,
						},
					},
				}
			}
		}
	}
	if repeatCode != "" {
		return nil, errors.New(fmt.Sprintf("重复编码【%s】", repeatCode))
	}
	return saveMap, nil
}

// =========================================
// 标准添删改查
func (LangResourcesController) QueryList(c *gin.Context) {
	var filter entitys.LangResourcesQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.LangResourcesFilter)
	}
	res, total, err := resourcesServices.SetContext(controls.WithOpenUserContext(c)).QueryLangResourcesList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (LangResourcesController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := resourcesServices.SetContext(controls.WithOpenUserContext(c)).GetLangResourcesDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (LangResourcesController) Edit(c *gin.Context) {
	var req entitys.LangResourcesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := resourcesServices.SetContext(controls.WithOpenUserContext(c)).UpdateLangResources(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangResourcesController) Add(c *gin.Context) {
	var req entitys.LangResourcesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := resourcesServices.SetContext(controls.WithOpenUserContext(c)).AddLangResources(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangResourcesController) Delete(c *gin.Context) {
	var req entitys.LangResourcesFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = resourcesServices.SetContext(controls.WithOpenUserContext(c)).DeleteLangResources(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// CustomResourcesBatchSave 资源翻译保存
func (s LangResourcesController) CustomResourcesSave(c *gin.Context) {
	var req entitys.BatchCustomResourcesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	thisContext := controls.WithUserContext(c)
	resourcesList := []*protosService.BatchCustomResourcesItem{}
	langKey := ""
	for _, item := range req.ResourcesList {
		translateItem := protosService.BatchCustomResourcesItem{
			Lang:    item.Lang,
			Value:   item.Value,
			LangKey: item.LangKey,
			Id:      item.Id,
		}
		langKey = item.LangKey
		resourcesList = append(resourcesList, &translateItem)
	}
	ret, err := rpc.ClientLangCustomResourceService.BatchSaveCustomResources(thisContext, &protosService.BatchCustomResourcesRequest{
		BelongId:      req.BelongId,
		BelongType:    req.BelongType,
		ProductKey:    req.ProductKey,
		LangKey:       langKey,
		ResourcesList: resourcesList,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if ret.Code != 200 {
		iotgin.ResErrCli(c, errors.New(ret.Message))
		return
	}

	//删除缓存
	if req.ProductKey != "" {
		if err := iotredis.GetClient().Del(thisContext, persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, req.ProductKey)).Err(); err != nil {
			iotlogger.LogHelper.Errorf("面板缓存删除失败，productKey:%v", req.ProductKey)
			//return
		}
	}
	clearCached(req.ProductKey, req.BelongType)
	iotgin.ResSuccessMsg(c)
}

// 清理面板缓存
func clearCached(productKey string, belongType int32) {
	if productKey == "" || belongType != 4 {
		return
	}
	cachedKey := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, productKey)
	iotredis.GetClient().Del(context.Background(), cachedKey)
}

// CustomResourcesGet 查询面板自定义资源
func (s *LangResourcesController) CustomResourcesGet(c *gin.Context) {
	langKey := c.Query("langKey")
	if langKey == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 langKey"))
		return
	}
	thisContext := controls.WithUserContext(c)

	belongId := c.Query("belongId")
	var belongIdInt int64 = 0
	if belongId != "" {
		belongIdInt = iotutil.ToInt64(belongId)
	}

	belongType := c.Query("belongType")
	var belongTypeInt int32 = 0
	if belongType != "" {
		belongTypeInt = iotutil.ToInt32(belongType)
	}
	productKey := c.DefaultQuery("productKey", "")

	////获取资源部信息，读取支持的语言
	//res, err := rpc.ClientLangResourcesPackageService.Lists(thisContext, &protosService.LangResourcePackageListRequest{
	//	Query: &protosService.LangResourcePackage{
	//		BelongType: belongTypeInt,
	//		BelongId:   belongIdInt,
	//	},
	//})
	//if err != nil {
	//	iotgin.ResErrCli(c, err)
	//	return
	//}
	//if res.Code != 200 {
	//	iotgin.ResErrCli(c, errors.New(res.Message))
	//	return
	//}
	//if len(res.Data) == 0 {
	//	iotgin.ResErrCli(c, errors.New("未找到任何可导出的资源"))
	//	return
	//}

	dicService := services2.BaseDataService{}
	langTypes := dicService.GetLangType()

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &protosService.LangCustomResourcesListRequest{
		Query: &protosService.LangCustomResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
			ProductKey: productKey,
			Code:       langKey,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if crs.Code != 200 {
		iotgin.ResErrCli(c, errors.New(crs.Message))
		return
	}

	resourcesRes := new(entitys.BatchCustomResourcesEntitys)
	resourcesRes.BelongId = belongIdInt
	resourcesRes.BelongType = belongTypeInt
	resourcesRes.ResourcesList = make([]entitys.BatchSaveCustomResourcesItem, 0)
	resMap := map[string]*entitys.BatchSaveCustomResourcesItem{}

	for i, langType := range langTypes {
		item := &entitys.BatchSaveCustomResourcesItem{
			Id:      0,
			Lang:    langType.Code,
			LangKey: langKey,
			Value:   "",
			Sort:    i,
		}
		//解决排序问题。
		switch item.Lang {
		case "zh":
			item.Sort = 1
		case "en":
			item.Sort = 2
		default:
			item.Sort = 2 + i
		}
		resMap[langType.Code] = item
	}

	//如果自定义为0，则读取基础的翻译
	if len(crs.Data) == 0 {
		rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &protosService.LangResourcesListRequest{
			Query: &protosService.LangResources{
				BelongType: belongTypeInt,
				BelongId:   belongIdInt,
				Code:       langKey,
			},
		})
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		if rep.Code != 200 {
			iotgin.ResErrCli(c, errors.New(rep.Message))
			return
		}
		for _, d := range rep.Data {
			if t, ok := resMap[d.Lang]; ok {
				resMap[d.Lang] = &entitys.BatchSaveCustomResourcesItem{
					Id:      d.Id,
					Lang:    d.Lang,
					LangKey: d.Code,
					Value:   d.Value,
					Sort:    t.Sort,
				}
			}
		}
	} else {
		for _, d := range crs.Data {
			if t, ok := resMap[d.Lang]; ok {
				resMap[d.Lang] = &entitys.BatchSaveCustomResourcesItem{
					Id:      d.Id,
					Lang:    d.Lang,
					LangKey: d.Code,
					Value:   d.Value,
					Sort:    t.Sort,
				}
			}
		}
	}

	for _, item := range resMap {
		resourcesRes.ResourcesList = append(resourcesRes.ResourcesList, *item)
	}

	sort.Slice(resourcesRes.ResourcesList, func(i, j int) bool {
		return resourcesRes.ResourcesList[i].Sort < resourcesRes.ResourcesList[j].Sort
	})

	iotgin.ResSuccess(c, resourcesRes)
}

// QueryCustomResource 查询面板自定义资源
func (s *LangResourcesController) QueryCustomResource(c *gin.Context) {
	var (
		thisContext         = controls.WithUserContext(c)
		belongId            = c.Query("belongId")
		belongType          = c.Query("belongType")
		productKey          = c.DefaultQuery("productKey", "")
		belongIdInt   int64 = 0
		belongTypeInt int32 = 0
	)
	//默认是APP的资源下载
	if belongId != "" {
		belongIdInt = iotutil.ToInt64(belongId)
	}
	if belongType != "" {
		belongTypeInt = iotutil.ToInt32(belongType)
	}
	//获取面板Id对应的默认翻译信息
	langTypes, _ := s.getLangTypes(thisContext, belongIdInt, belongTypeInt)

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
			ProductKey: productKey,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if crs.Code != 200 {
		iotgin.ResErrCli(c, errors.New(crs.Message))
		return
	}

	resultArr := map[string]map[string]interface{}{}
	//如果自定义为0，则读取基础的翻译
	if len(crs.Data) == 0 {
		resultArr, err = s.getPublicResource(thisContext, belongIdInt, belongTypeInt)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
	} else {
		for i, d := range crs.Data {
			if _, ok := resultArr[d.Code]; !ok {
				resultArr[d.Code] = map[string]interface{}{"sort": i, "code": d.Code, d.Lang: d.Value}
			} else {
				resultArr[d.Code][d.Lang] = d.Value
			}
		}
	}
	iotgin.ResSuccess(c, s.convertResouceResult(resultArr, langTypes))
}

func (s *LangResourcesController) QueryCustomResourceV2(c *gin.Context) {
	var (
		thisContext         = controls.WithUserContext(c)
		belongId            = c.Query("belongId")
		belongType          = c.Query("belongType")
		productKey          = c.DefaultQuery("productKey", "")
		belongIdInt   int64 = 0
		belongTypeInt int32 = 0
	)
	//默认是APP的资源下载
	if belongId != "" {
		belongIdInt = iotutil.ToInt64(belongId)
	}
	if belongType != "" {
		belongTypeInt = iotutil.ToInt32(belongType)
	}
	//获取面板Id对应的默认翻译信息
	langTypes, _ := s.getLangTypes(thisContext, belongIdInt, belongTypeInt)

	//resultArr := map[string]map[string]interface{}{}
	//获取公共面板资源包
	public, err := s.getPublicResource(thisContext, belongIdInt, belongTypeInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//获取面板自定义资源包
	custom, err := s.getCustomResource(thisContext, belongIdInt, belongTypeInt, productKey)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	for key, _ := range public {
		if v, ok := custom[key]; ok {
			for ckey, cval := range v {
				public[key][ckey] = cval
			}
		}
	}
	iotgin.ResSuccess(c, s.convertResouceResult(public, langTypes))
}

func (s LangResourcesController) getLangTypes(ctx context.Context, belongIdInt int64, belongTypeInt int32) ([]string, error) {
	langTypes := []string{"zh", "en"}
	res, err := rpc.ClientLangResourcesPackageService.Lists(ctx, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
		},
	})
	if err != nil {
		return langTypes, err
	}
	if res.Code != 200 {
		return langTypes, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return langTypes, nil
	}
	var tempLangTyps []string = make([]string, 0)
	for _, d := range res.Data[0].Langs {
		tempLangTyps = append(tempLangTyps, d)
	}

	if err == nil {
		for _, lang := range tempLangTyps {
			noLang := ""
			for _, l := range langTypes {
				noLang = l
				if lang == l {
					noLang = ""
					break
				}
			}
			if noLang != "" {
				langTypes = append(langTypes, noLang)
			}
		}
	}
	return langTypes, nil
}

// 转换返回格式
func (s LangResourcesController) convertResouceResult(resultArr map[string]map[string]interface{}, langType []string) []map[string]interface{} {
	newResultArr := make([]map[string]interface{}, 0)
	for _, row := range resultArr {
		rowMap := map[string]interface{}{}
		rowMap["id"] = row["code"]
		rowMap["langKey"] = row["code"]
		rowMap["sort"] = row["sort"]
		if rowMap["sort"] == nil {
			rowMap["sort"] = 1
		}
		rowMap["showLangKey"] = row["code"]
		for _, lang := range langType {
			switch lang {
			case "zh":
				rowMap["name"] = row[lang]
			case "en":
				rowMap["nameEn"] = row[lang]
			default:
				rowMap[lang] = row[lang]
			}
		}
		newResultArr = append(newResultArr, rowMap)
	}
	sort.Slice(newResultArr, func(i, j int) bool {
		return iotutil.ToInt32(newResultArr[i]["sort"]) < iotutil.ToInt32(newResultArr[j]["sort"])
		//return iotutil.ToString(newResultArr[i]["langKey"]) > iotutil.ToString(newResultArr[j]["langKey"])
	})
	return newResultArr
}

// 获取公共资源
func (s LangResourcesController) getPublicResource(ctx context.Context, belongIdInt int64, belongTypeInt int32) (map[string]map[string]interface{}, error) {
	resultArr := map[string]map[string]interface{}{}
	rep, err := rpc.ClientLangResourcesService.Lists(ctx, &proto.LangResourcesListRequest{
		Query: &proto.LangResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	for i, d := range rep.Data {
		if _, ok := resultArr[d.Code]; !ok {
			resultArr[d.Code] = map[string]interface{}{"sort": i, "code": d.Code, d.Lang: d.Value}
		} else {
			resultArr[d.Code][d.Lang] = d.Value
		}
	}
	return resultArr, nil
}

// 获取自定义资源
func (s LangResourcesController) getCustomResource(ctx context.Context, belongIdInt int64, belongTypeInt int32, productKey string) (map[string]map[string]interface{}, error) {
	resultArr := map[string]map[string]interface{}{}
	crs, err := rpc.ClientLangCustomResourceService.Lists(ctx, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongType: belongTypeInt,
			BelongId:   belongIdInt,
			ProductKey: productKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if crs.Code != 200 {
		return nil, errors.New(crs.Message)
	}
	for i, d := range crs.Data {
		if _, ok := resultArr[d.Code]; !ok {
			resultArr[d.Code] = map[string]interface{}{"sort": i, "code": d.Code, d.Lang: d.Value}
		} else {
			resultArr[d.Code][d.Lang] = d.Value
		}
	}
	return resultArr, nil
}

func (s LangResourcesController) SetLangs(lanyKey, value string, langTypes []*dictEntitys.DictKeyVal, data []*entitys.BatchSaveCustomResourcesItem) []*entitys.BatchSaveCustomResourcesItem {
	defaultData := []*entitys.BatchSaveCustomResourcesItem{}
	for _, langType := range langTypes {
		var currData *entitys.BatchSaveCustomResourcesItem
		for _, d := range data {
			if langType.Code == d.Lang {
				currData = d
			}
		}
		defaultItem := &entitys.BatchSaveCustomResourcesItem{
			//SourceTable: sourceTable,
			//SourceRowId: sourceRowId,
			Lang:    langType.Code,
			LangKey: lanyKey,
			Value:   value,
		}
		if currData != nil {
			defaultItem = currData
		}
		defaultData = append(defaultData, defaultItem)
	}
	return defaultData
}
