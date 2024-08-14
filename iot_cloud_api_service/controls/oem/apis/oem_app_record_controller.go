package apis

import (
	"bytes"
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
)

var OSS_OEMAPP_PLIST_DIR = "plist"

var OemAppBuildRecordcontroller OemAppBuildRecordController

var serviceBuild apiservice.OemAppBuildRecordService

type OemAppBuildRecordController struct {
} //用户操作控制器

// //修改app名称
// func (OemAppBuildRecordController) BuildFinishNotify(c *gin.Context) {
// 	var req entitys.OemAppBuildFinishNotifyReq
// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		iotgin.ResErrCli(c, err)
// 		return
// 	}
// 	id, err := serviceBuild.SetContext(controls.WithOpenUserContext(c)).BuildFinishNotify(req)
// 	if err != nil {
// 		iotgin.ResErrCli(c, err)
// 		return
// 	}
// 	iotgin.ResSuccess(c, id)
// }

// 修改app名称
func (OemAppBuildRecordController) BuildFinishNotify(c *gin.Context) {
	iotlogger.LogHelper.Info("BuildFinishNotify 开始..........>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	status := c.PostForm("status")
	iotlogger.LogHelper.Info("BuildFinishNotify status=", status)
	statusInt, err := iotutil.ToInt64AndErr(status)
	if err != nil {
		iotgin.ResBadRequest(c, "status")
		return
	}
	buildId := c.PostForm("buildId")
	iotlogger.LogHelper.Info("BuildFinishNotify buildId=", buildId)
	_, err = iotutil.ToInt64AndErr(buildId)
	if err != nil {
		iotgin.ResBadRequest(c, "buildId")
		return
	}
	iotlogger.LogHelper.Info("接收到" + buildId + "构建通知..........>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	commitId := c.PostForm("commitId")
	iotlogger.LogHelper.Info("BuildFinishNotify commitId=", commitId)
	buildResultMsg := c.PostForm("buildResultMsg")
	iotlogger.LogHelper.Info("BuildFinishNotify buildResultMsg=", buildResultMsg)
	//endTime := c.PostForm("endTime")
	buildResult := c.PostForm("buildResult")
	iotlogger.LogHelper.Info("BuildFinishNotify buildResult=", buildResult)
	buildResultInt, err := iotutil.ToInt32Err(buildResult)
	if err != nil {
		iotgin.ResBadRequest(c, "buildResult")
		return
	}
	pkgUrl := ""
	//存放aab或是plist
	pkgAabOrPlistUrl := ""
	//ipa正式包
	pkgIpaUrl := ""

	form, errForm := c.MultipartForm()
	if form != nil {
		if errForm != nil {
			iotgin.ResErrCli(c, errForm)
			return
		}
		files := form.File["pkgFile"]
		if len(files) > 0 {
			f, err := apis.SaveFileToQiniuStaticOSS(c, files[0], "oemapp_package", "apk", "zip", "ipa", "aab", "plist")
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			} else {
				pkgUrl = f.FullPath
			}
		}
		//ios的plist文件或是安卓海外的aab安装包
		files2 := form.File["pkgFile2"]
		if len(files2) > 0 {
			//获取上传文件的后缀名
			ext := path.Ext(files2[0].Filename)
			if ext == ".aab" {
				f2, err := apis.SaveFileToQiniuStaticOSS(c, files2[0], "oemapp_package", "apk", "zip", "ipa", "aab", "plist")
				if err != nil {
					iotgin.ResErrCli(c, err)
					return
				} else {
					pkgAabOrPlistUrl = f2.FullPath
				}
			} else if ext == ".ipa" {
				f2, err := apis.SaveFileToQiniuStaticOSS(c, files2[0], "oemapp_package", "apk", "zip", "ipa", "aab", "plist")
				if err != nil {
					iotgin.ResErrCli(c, err)
					return
				} else {
					pkgUrl = f2.FullPath
				}
			} else {
				iotgin.ResErrCli(c, errors.New("收到未知的文件"))
				return
			}
		}
		plist := form.File["manifestFile"]
		if len(plist) > 0 {
			//获取上传文件的后缀名
			ext := path.Ext(plist[0].Filename)
			if ext == ".plist" {
				//创建文件夹
				errMkDir := iotutil.MkDir(apiservice.DirTempBuildPlistRecord)
				if errMkDir != nil {
					iotgin.ResErrCli(c, errMkDir)
					return
				}

				//组合文件路径[上传的plist文件保存到服务器的路径(带文件名)]
				fileId := buildId + "_" + iotutil.GetRandomString(6)
				plistPathName := apiservice.DirTempBuildPlistRecord + "/" + fileId + ext
				//生成新的plist文件名
				plistNewName := fileId + "_new" + ext
				//生成新的plist文件路径(带文件名)
				plistNewPathName := apiservice.DirTempBuildPlistRecord + "/" + plistNewName

				//保存原始上传文件
				errUpload := c.SaveUploadedFile(plist[0], plistPathName)
				if errUpload != nil {
					iotgin.ResErrCli(c, errUpload)
					return
				}
				//读取文件内容(考虑后期直接从内存读取bytes转string 进行内容替换.  避免存储硬盘提高性能)
				b, errRead := ioutil.ReadFile(plistPathName)
				if errRead != nil {
					iotgin.ResErrCli(c, errRead)
					return
				}
				plistConent := string(b)
				//plist文件内容替换
				strTemp, errIcon := serviceBuild.SetContext(controls.WithOpenUserContext(c)).HandlerIosPlistFile(buildId, pkgUrl, plistConent)
				if errIcon != nil {
					iotgin.ResErrCli(c, errIcon)
					return
				}
				plistConent = strTemp

				//保存替换后的内容
				bby := bytes.NewBufferString(plistConent)
				iotutil.FileCreate(*bby, plistNewPathName)

				//把替换后的内容上传到oss
				urlPlist, errOss := apis.UploadStatic(config.Global.Oss.UseOss, plistNewName, plistNewPathName)
				if errOss != nil {
					iotgin.ResErrCli(c, errOss)
					return
				}
				//赋值pkgAabOrPlistUrl 变量
				pkgAabOrPlistUrl = urlPlist
			} else {
				iotgin.ResErrCli(c, errors.New("收到未知的文件"))
				return
			}
		}
		//ios正式包
		files3 := form.File["pkgFile3"]
		if len(files3) > 0 {
			f3, errF3 := apis.SaveFileToQiniuStaticOSS(c, files3[0], "oemapp_package", "apk", "zip", "ipa", "aab", "plist")
			if errF3 != nil {
				iotgin.ResErrCli(c, errF3)
				return
			} else {
				pkgIpaUrl = f3.FullPath
			}
		}
	}
	iotlogger.LogHelper.Info(buildId + "构建文件上传成功........" + pkgUrl + "..>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	var req entitys.OemAppBuildFinishNotifyReq
	req.BuildId = buildId
	req.BuildProgress = 100
	req.BuildResult = buildResultInt
	//req.EndTime = time.Now().Unix()
	req.CommitID = commitId
	req.Status = statusInt
	req.BuildResultMsg = buildResultMsg
	//arrPkgUrl 说明
	//android 国内 apk安装包和空字符串和空字符串
	//Android 海外 apk安装包和aab安装包和空字符串
	//ios  ipa安装包和plist文件和ios正式安装包
	var arrPkgUrl = []string{pkgUrl, pkgAabOrPlistUrl, pkgIpaUrl}
	req.PkgURL = iotutil.ToStringByUrl(arrPkgUrl)
	id, err := serviceBuild.SetContext(controls.WithOpenUserContext(c)).BuildFinishNotify(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func GetiOSAppIconUrl(buildId string) string {
	ret, err := serviceBuild.SetContext(context.Background()).GetIconUrl(buildId)
	if err != nil {
		return ""
	}
	return ret
}

// 打开二维码链接内容
func (OemAppBuildRecordController) BuildPackageQrCode(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}
	ret, errBr := serviceBuild.SetContext(controls.WithUserContext(c)).BuildPackageQrCode(req)
	if errBr != nil {
		c.Writer.WriteString(errBr.Error())
		return
	}
	c.Writer.WriteString(ret)
}

func (OemAppBuildRecordController) BuildFinishNotifyEx(c *gin.Context) {
	var req entitys.OemAppBuildFinishNotifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	id, err := serviceBuild.SetContext(controls.WithOpenUserContext(c)).BuildFinishNotify(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OemAppBuildRecordController) GetBuildAppIconUrl(c *gin.Context) {
	id := c.Query("buildId")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("buildId参数错误"))
		return
	}
	ret, err := serviceBuild.SetContext(controls.WithUserContext(c)).GetIconUrl(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, ret)
}
