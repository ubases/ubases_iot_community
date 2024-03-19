package services

import (
	"cloud_platform/iot_cloud_api_service/config"
	commonApi "cloud_platform/iot_cloud_api_service/controls/common/apis"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppCustomRecordService struct {
	Ctx context.Context
}

func (s OemAppCustomRecordService) SetContext(ctx context.Context) OemAppCustomRecordService {
	s.Ctx = ctx
	return s
}

func (s OemAppCustomRecordService) CreateOemAppCustomRecord(obj *entitys.OemAppCustomRecordEntitys) error {
	// 添加版本号必须要高于最新版本号
	respApp, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: obj.AppId,
	})
	if err != nil {
		return err
	}
	var checkVersion string
	switch obj.Os {
	case 1:
		checkVersion = respApp.Data[0].IosVersion
	case 2:
		checkVersion = respApp.Data[0].AndroidInterVersion
	case 3:
		checkVersion = respApp.Data[0].AndroidOuterVersion
	}
	if checkVersion != "" {
		comInt, err := iotutil.VerCompare(obj.Version, checkVersion)
		if err != nil {
			return goerrors.New("", "版本号比较错误, 请确认版本号格式", ioterrs.ErrCloudRequestParam)
		}
		if checkVersion != "" && comInt != 1 {
			return goerrors.New("", "当前版本号必须大于最新版本号", ioterrs.ErrCloudVersionTooLow)
		}
	}
	// 通过模板生成plist文件
	var plistUrl string
	if obj.Os == 1 {
		respUi, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
			AppId:   obj.AppId,
			Version: "1.0.0",
		})
		if err != nil {
			return err
		}
		m := map[string]string{}
		err = json.Unmarshal([]byte(respUi.Data[0].IosLaunchScreen), &m)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrShouldBindJSON)
		}
		plist := entitys.TemplatePlistEntitys{
			Title:            respApp.Data[0].Name,
			Version:          obj.Version,
			BundleIdentifier: respApp.Data[0].IosPkgName,
			DisplayImage:     m["displayImage"],
			FullSizeImage:    m["fullSizeImage"],
			SoftwarePackage:  obj.PkgUrl,
		}
		plistUrl, err = genPlistFile(plist)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrCloudGenOrUploadPlist)
		}
	}
	req := entitys.OemAppCustomRecord_e2pb(obj)
	req.Id = iotutil.GetNextSeqInt64()
	req.PlistUrl = plistUrl
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err = rpc.ClientOemAppCustomRecordService.Create(s.Ctx, req)
	if err != nil {
		return err
	}
	reqApp := &protosService.OemAppUpdateFieldsRequest{}
	if obj.Os == 1 {
		reqApp.Fields = []string{"ios_version"}
		reqApp.Data = &protosService.OemApp{
			Id:         obj.AppId,
			IosVersion: obj.Version,
		}
	} else if obj.Os == 2 {
		reqApp.Fields = []string{"android_inter_version"}
		reqApp.Data = &protosService.OemApp{
			Id:                  obj.AppId,
			AndroidInterVersion: obj.Version,
		}
	} else if obj.Os == 3 {
		reqApp.Fields = []string{"android_outer_version"}
		reqApp.Data = &protosService.OemApp{
			Id:                  obj.AppId,
			AndroidOuterVersion: obj.Version,
		}
	}
	_, err = rpc.ClientOemAppService.UpdateFields(s.Ctx, reqApp)
	if err != nil {
		return err
	}
	return nil
}

func (s OemAppCustomRecordService) UpdateOemAppCustomRecord(obj *entitys.OemAppCustomRecordEntitys) error {
	// 添加版本号必须要高于最新版本号
	respApp, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: obj.AppId,
	})
	if err != nil {
		return err
	}
	var checkVersion string
	switch obj.Os {
	case 1:
		checkVersion = respApp.Data[0].IosVersion
	case 2:
		checkVersion = respApp.Data[0].AndroidInterVersion
	case 3:
		checkVersion = respApp.Data[0].AndroidOuterVersion
	}
	comInt, err := iotutil.VerCompare(obj.Version, checkVersion)
	if err != nil {
		return goerrors.New("", "版本号比较错误, 请确认版本号格式", ioterrs.ErrCloudRequestParam)
	}
	if checkVersion != obj.Version && checkVersion != "" && comInt != 1 {
		return goerrors.New("", "当前版本号必须大于最新版本号", ioterrs.ErrCloudVersionTooLow)
	}
	// 通过模板生成plist文件
	var plistUrl string
	if obj.Os == 1 {
		respUi, err := rpc.ClientOemAppUiConfigService.Find(s.Ctx, &protosService.OemAppUiConfigFilter{
			AppId:   obj.AppId,
			Version: "1.0.0",
		})
		if err != nil {
			return err
		}
		m := map[string]string{}
		err = json.Unmarshal([]byte(respUi.Data[0].IosLaunchScreen), &m)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrShouldBindJSON)
		}
		plist := entitys.TemplatePlistEntitys{
			Title:            respApp.Data[0].Name,
			Version:          obj.Version,
			BundleIdentifier: respApp.Data[0].IosPkgName,
			DisplayImage:     m["displayImage"],
			FullSizeImage:    m["fullSizeImage"],
			SoftwarePackage:  obj.PkgUrl,
		}
		plistUrl, err = genPlistFile(plist)
		if err != nil {
			return goerrors.New("", err.Error(), ioterrs.ErrCloudGenOrUploadPlist)
		}
	}
	req := entitys.OemAppCustomRecord_e2pb(obj)
	req.PlistUrl = plistUrl
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err = rpc.ClientOemAppCustomRecordService.Update(s.Ctx, req)
	if err != nil {
		return err
	}
	// 如果更新的版本号比当前最新版本号要高，需要更新自定义app相关的最新版本号
	if obj.Version != "" && checkVersion != obj.Version {
		reqApp := &protosService.OemAppUpdateFieldsRequest{}
		if obj.Os == 1 {
			reqApp.Fields = []string{"ios_version"}
			reqApp.Data = &protosService.OemApp{
				Id:         obj.AppId,
				IosVersion: obj.Version,
			}
		} else if obj.Os == 2 {
			reqApp.Fields = []string{"android_inter_version"}
			reqApp.Data = &protosService.OemApp{
				Id:                  obj.AppId,
				AndroidInterVersion: obj.Version,
			}
		} else if obj.Os == 3 {
			reqApp.Fields = []string{"android_outer_version"}
			reqApp.Data = &protosService.OemApp{
				Id:                  obj.AppId,
				AndroidOuterVersion: obj.Version,
			}
		}
		_, err = rpc.ClientOemAppService.UpdateFields(s.Ctx, reqApp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s OemAppCustomRecordService) SetOemAppCustomRecord(obj *entitys.OemAppCustomRecordEntitys) error {
	req := &protosService.OemAppCustomRecordUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OemAppCustomRecord{
			Id:     iotutil.ToInt64(obj.Id),
			Status: obj.Status,
		},
	}
	_, err := rpc.ClientOemAppCustomRecordService.UpdateFields(s.Ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// 更新上架记录
func (s OemAppCustomRecordService) SetLaunchMarkets(obj *entitys.OemAppCustomRecordEntitys) error {
	req := &protosService.OemAppCustomRecordUpdateFieldsRequest{
		Fields: []string{"launch_markets"},
		Data: &protosService.OemAppCustomRecord{
			Id:            iotutil.ToInt64(obj.Id),
			LaunchMarkets: iotutil.ToString(obj.LaunchMarkets),
		},
	}
	_, err := rpc.ClientOemAppCustomRecordService.UpdateFields(s.Ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// 更新描述
func (s OemAppCustomRecordService) SetRemark(obj *entitys.OemAppCustomRecordEntitys) error {
	req := &protosService.OemAppCustomRecordUpdateFieldsRequest{
		Fields: []string{"description"},
		Data: &protosService.OemAppCustomRecord{
			Id:          iotutil.ToInt64(obj.Id),
			Description: obj.Description,
		},
	}
	_, err := rpc.ClientOemAppCustomRecordService.UpdateFields(s.Ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s OemAppCustomRecordService) DeleteOemAppCustomRecord(obj *entitys.OemAppCustomRecordEntitys) error {
	_, err := rpc.ClientOemAppCustomRecordService.DeleteById(s.Ctx, &protosService.OemAppCustomRecord{
		Id: obj.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s OemAppCustomRecordService) GetOemAppCustomRecord(id string) (*entitys.OemAppCustomRecordEntitys, error) {
	req := &protosService.OemAppCustomRecordFilter{
		Id: iotutil.ToInt64(id),
	}
	resp, err := rpc.ClientOemAppCustomRecordService.Find(s.Ctx, req)
	if err != nil && goerrors.FromError(err).GetDetail() == ioterrs.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	data := entitys.OemAppCustomRecord_pb2e(resp.Data[0])
	return data, nil
}

// 获取app自定义版本列表
func (s OemAppCustomRecordService) GetOemAppCustomRecordList(req *entitys.OemAppCustomRecordQuery) ([]*entitys.OemAppCustomRecordEntitys, int64, error) {
	reqV := &protosService.OemAppCustomRecordListRequest{
		Query: &protosService.OemAppCustomRecord{
			AppId:  iotutil.ToInt64(req.Query.AppId),
			Os:     req.Query.Os,
			Status: req.Query.Status,
		},
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	resp, err := rpc.ClientOemAppCustomRecordService.Lists(s.Ctx, reqV)
	if err != nil {
		return nil, 0, err
	}
	data := []*entitys.OemAppCustomRecordEntitys{}
	for i := range resp.Data {
		data = append(data, entitys.OemAppCustomRecord_pb2e(resp.Data[i]))
	}
	return data, resp.Total, nil
}

func (s OemAppCustomRecordService) CustomPackageQrCode(req entitys.OemAppCommonReq) (string, error) {
	resApp, errApp := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
		Id: iotutil.ToInt64(req.AppId),
	})
	if errApp != nil {
		return "", errApp
	}
	if resApp.Code != 200 {
		return "", errors.New(resApp.Message)
	}

	reqCustom := &protosService.OemAppCustomRecordListRequest{
		Query: &protosService.OemAppCustomRecord{
			AppId:   iotutil.ToInt64(req.AppId),
			Version: req.Version,
			Os:      req.Os,
		},
	}
	respCustom, err := rpc.ClientOemAppCustomRecordService.Lists(s.Ctx, reqCustom)
	if err != nil {
		return "", err
	}
	if respCustom.Data == nil || len(respCustom.Data) == 0 {
		return "", errors.New("查询不到自定义app版本记录数据")
	}

	var res = entitys.OemAppBuildPackageQrCodeRes{}
	res.AppName = resApp.Data[0].Name
	//图标从t_oem_app里面获取，需新增字段
	res.IconUrl = resApp.Data[0].AppIconUrl
	res.Version = req.Version
	res.IosUrl = "javascript:void(0);"
	res.AndroidUrl = "javascript:void(0);"
	res.AndroidAabUrl = "javascript:void(0);"

	res.IosClass = "download disabled"
	res.AndroidClass = "download disabled"
	res.AndroidAabClass = "download disabled"

	//ios
	for _, c := range respCustom.Data {
		if c.Os == 1 && c.PlistUrl != "" {
			//urlaab 如果是ios的时候. 存放的 plist文件
			//如果加入了itms-services://  模板引擎就会识别不了.导致 iosUrl连接错误.
			//res.IosUrl = "itms-services://?action=download-manifest&url="+v.UrlAab
			res.IosUrl = "itms-services://?action=download-manifest&url=" + c.PlistUrl
			res.IosClass = "download"
		}
		//android
		if c.Os == 2 && c.PkgUrl != "" {
			res.AndroidUrl = c.PkgUrl
			res.AndroidClass = "download"
		}
		//android aab
		if c.Os == 3 && c.PkgUrl != "" {
			res.AndroidAabUrl = c.PkgUrl
			res.AndroidAabClass = "download"
		}
	}

	templatePath := strings.Join([]string{iotconst.GetTemplatesDir(), "buildPackageQrCode.tmpl"}, string(filepath.Separator))
	cont, _ := iotutil.RenderHtmpTemplateByPath(templatePath, res)
	//此字段特殊处理. 因为会包含特殊字符itms-services:// 导致模板引擎失效.
	cont = strings.ReplaceAll(cont, "___IosUrl___", res.IosUrl)
	return cont, nil
}

func genPlistFile(plist entitys.TemplatePlistEntitys) (string, error) {
	plistPathDir := iotconst.GetIosPlistFileDir()
	//创建目录
	err := iotutil.MkDir(plistPathDir)
	if err != nil {
		logger.Error("创建目录" + plistPathDir + "错误:" + err.Error())
		return "", err
	}

	templatePath := strings.Join([]string{iotconst.GetTemplatesDir(), "plist.tmpl"}, string(filepath.Separator))
	content, _ := iotutil.RenderTextTemplateByPath(templatePath, plist)

	plistName := iotutil.GetRandomPureString(20) + "_ios.plist"
	plistPath := GetDirPath(plistPathDir, plistName)
	if err := iotutil.WriteFile(plistPath, content); err != nil {
		return "", err
	}
	url, err := commonApi.Upload(config.Global.Oss.UseOss, plistName, plistPath)
	if err != nil {
		logger.Error("上传oss错误:" + err.Error())
		return "", err
	}
	return url, nil
}

// 拼凑文件或目录的路径
func GetDirPath(dirName string, name string) string {
	return strings.Join([]string{dirName, name}, string(filepath.Separator))
}
