package services

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"path/filepath"
	"strings"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppBuildRecordService struct {
	Ctx context.Context
}

var DirTempBuildPlistRecord = strings.Join([]string{iotconst.GetBuildRecordDir(), "plist"}, string(filepath.Separator))

func (s OemAppBuildRecordService) SetContext(ctx context.Context) OemAppBuildRecordService {
	s.Ctx = ctx
	return s
}

func (s OemAppBuildRecordService) CheckUpdateOemAppStatus(appId int64, version string) error {

	res, err := rpc.ClientOemAppBuildRecordService.Lists(s.Ctx, &protosService.OemAppBuildRecordListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.OemAppBuildRecord{
			AppId:   appId,
			Version: version,
		},
	})

	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	//app下的构建是否已经都完成.
	hasSuccess := false
	hasBuilding := false
	for _, v := range res.Data {
		//只要有成功的就显示成功
		if v.Status == 3 {
			hasSuccess = true
		} else if v.Status == 2 {
			hasBuilding = true
		}
	}
	var app = OemAppService{}
	app.Ctx = s.Ctx
	if !hasBuilding {
		if hasSuccess {
			app.OemAppUpdateStatus(iotutil.ToInt64(appId), 3)
		} else {
			app.OemAppUpdateStatus(iotutil.ToInt64(appId), 4)
		}
	}
	return nil
}

// 必须要有一个构建包成功才返回true
func (s OemAppBuildRecordService) IsBuildSuccess(appId int64, version string) (bool, error) {

	res, err := rpc.ClientOemAppBuildRecordService.Lists(s.Ctx, &protosService.OemAppBuildRecordListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.OemAppBuildRecord{
			AppId:   appId,
			Version: version,
		},
	})

	if err != nil && err.Error() != "record not found" {
		return false, err
	}
	if res.Code != 200 {
		return false, errors.New(res.Message)
	}

	//app下的构建是否已经都完成.
	isSuccess := false
	for _, v := range res.Data {
		if v.Status == 3 {
			isSuccess = true
			break
		}
	}
	return isSuccess, nil
}

func (s OemAppBuildRecordService) BuildFinishNotify(req entitys.OemAppBuildFinishNotifyReq) (string, error) {
	buildId, _ := iotutil.ToInt64AndErr(req.BuildId)
	endtime := timestamppb.Now()
	res, err := rpc.ClientOemAppBuildRecordService.UpdateFields(s.Ctx, &protosService.OemAppBuildRecordUpdateFieldsRequest{
		Fields: []string{"pkg_url", "status", "commit_id", "build_progress", "build_result", "build_result_msg", "end_time"},
		Data: &protosService.OemAppBuildRecord{
			Id:             buildId,
			PkgUrl:         req.PkgURL,
			Status:         int32(req.Status),
			CommitId:       req.CommitID,
			BuildProgress:  int32(req.BuildProgress),
			BuildResultMsg: req.BuildResultMsg,
			EndTime:        endtime,
			BuildResult:    req.BuildResult,
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	resFind, errFind := rpc.ClientOemAppBuildRecordService.FindById(s.Ctx, &protosService.OemAppBuildRecordFilter{
		Id: buildId,
	})
	if errFind != nil {
		return "", err
	}
	if resFind.Code != 200 {
		return "", errors.New(resFind.Message)
	}
	errCk := s.CheckUpdateOemAppStatus(resFind.Data[0].AppId, resFind.Data[0].Version)
	if errCk != nil {
		return "", errCk
	}

	return "success", nil
}

func (s OemAppBuildRecordService) BuildPackageQrCode(req entitys.OemAppCommonReq) (string, error) {
	appId, err := iotutil.ToInt64AndErr(req.AppId)
	if err != nil {
		return "", errors.New(err.Error())
	}

	resApp, errApp := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: appId,
	})
	if errApp != nil {
		return "", errApp
	}
	if resApp.Code != 200 {
		return "", errors.New(resApp.Message)
	}

	resUi, errUi := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   iotutil.ToInt64(req.AppId),
		Version: req.Version,
	})
	if errUi != nil {
		return "", errUi
	}
	if resUi.Code != 200 {
		return "", errors.New(resUi.Message)
	}

	var br = OemAppService{}
	br.Ctx = s.Ctx
	resBr, errBr := br.OemAppBuildPackage(req)
	if errBr != nil {
		return "", errBr
	}

	var res = entitys.OemAppBuildPackageQrCodeRes{}
	res.AppName = resApp.Data[0].Name
	res.IconUrl = resUi.Data[0].IconUrl
	res.Version = req.Version
	res.IosUrl = "javascript:void(0);"
	res.AndroidUrl = "javascript:void(0);"
	res.AndroidAabUrl = "javascript:void(0);"

	res.IosClass = "download disabled"
	res.AndroidClass = "download disabled"
	res.AndroidAabClass = "download disabled"

	for _, v := range resBr {
		//ios
		if v.Os == 1 && v.Url != "" {
			//urlaab 如果是ios的时候. 存放的 plist文件
			//如果加入了itms-services://  模板引擎就会识别不了.导致 iosUrl连接错误.
			//res.IosUrl = "itms-services://?action=download-manifest&url="+v.UrlAab
			res.IosUrl = "itms-services://?action=download-manifest&url=" + v.UrlAab
			res.IosClass = "download"
		}
		//android
		if v.Os == 2 && v.Url != "" {
			res.AndroidUrl = v.Url
			res.AndroidClass = "download"
		}
		//android aab
		if v.Os == 3 && v.UrlAab != "" {
			res.AndroidAabUrl = v.Url
			res.AndroidAabClass = "download"
		}
	}
	templatePath := strings.Join([]string{iotconst.GetTemplatesDir(), "buildPackageQrCode.tmpl"}, string(filepath.Separator))
	cont, _ := iotutil.RenderHtmpTemplateByPath(templatePath, res)
	//此字段特殊处理. 因为会包含特殊字符itms-services:// 导致模板引擎失效.
	cont = strings.ReplaceAll(cont, "___IosUrl___", res.IosUrl)
	return cont, nil
}

// 处理plist文件.生成57和512尺寸的图标并且替换占位符
func (s OemAppBuildRecordService) HandlerIosPlistFile(buildId string, ipaOssUrl string, plistConent string) (string, error) {
	res, err := rpc.ClientOemAppBuildRecordService.FindById(s.Ctx, &protosService.OemAppBuildRecordFilter{
		Id: iotutil.ToInt64(buildId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	appId := res.Data[0].AppId
	//文件中间名避免重复构建的时候. oss上传文件的时候. 出现文件存在的错误.
	fileMiddleName := buildId + "_" + iotutil.ToString(res.Data[0].Os) + res.Data[0].Version
	resUi, errUi := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   appId,
		Version: res.Data[0].Version,
	})
	if errUi != nil {
		return "", errUi
	}
	if resUi.Code != 200 {
		return "", errors.New(resUi.Message)
	}
	iconOssUrl := resUi.Data[0].IconUrl

	iconOssDownloadPath := DirTempBuildPlistRecord + "/icon_" + fileMiddleName + ".png"
	//保存1024x1024的icon图
	iotutil.DownloadFile(iconOssUrl, iconOssDownloadPath)
	//生成指定大小icon图标
	iconLocalName57 := "icon_57x57_" + fileMiddleName + ".png"
	iconLocalPath57 := DirTempBuildPlistRecord + "/" + iconLocalName57

	iconLocalName512 := "icon_512x512_" + fileMiddleName + ".png"
	iconLocalPath512 := DirTempBuildPlistRecord + "/" + iconLocalName512
	err57 := iotutil.ImageResize(iconOssDownloadPath, iconLocalPath57, 57, 57)
	if err57 != nil {
		logger.Error("iotutil.ImageResize 57 error:" + err57.Error())
		return "", err57
	}
	err512 := iotutil.ImageResize(iconOssDownloadPath, iconLocalPath512, 512, 512)
	if err512 != nil {
		logger.Error("iotutil.ImageResize 512 error:" + err512.Error())
		return "", err512
	}

	//上传57尺寸的icon图标到oss
	iconUrl57, errOss57 := apis.UploadStatic(config.Global.Oss.UseOss, iconLocalName57, iconLocalPath57)
	if errOss57 != nil {
		logger.Error("apis.Upload 57 error:" + errOss57.Error())
		return "", errOss57
	}
	//上传512尺寸的icon图标到oss
	iconUrl512, errOss512 := apis.UploadStatic(config.Global.Oss.UseOss, iconLocalName512, iconLocalPath512)
	if errOss512 != nil {
		logger.Error("apis.Upload 512 error:" + errOss512.Error())
		return "", errOss512
	}

	plistConent = strings.ReplaceAll(plistConent, "https://www.xxx.com/__ipa__.ipa", ipaOssUrl)
	plistConent = strings.ReplaceAll(plistConent, "https://www.xxx.com/__57__.png", iconUrl57)
	plistConent = strings.ReplaceAll(plistConent, "https://www.xxx.com/__512__.png", iconUrl512)

	return plistConent, nil

}

func (s OemAppBuildRecordService) GetIconUrl(buildId string) (string, error) {
	res, err := rpc.ClientOemAppBuildRecordService.FindById(s.Ctx, &protosService.OemAppBuildRecordFilter{
		Id: iotutil.ToInt64(buildId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	appId := res.Data[0].AppId
	//文件中间名避免重复构建的时候. oss上传文件的时候. 出现文件存在的错误.
	resUi, errUi := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
		AppId:   appId,
		Version: res.Data[0].Version,
	})
	if errUi != nil {
		return "", errUi
	}
	if resUi.Code != 200 {
		return "", errors.New(resUi.Message)
	}
	if len(resUi.Data) == 0 {
		return "", errors.New("no data")
	}
	return resUi.Data[0].IconUrl, nil
}
