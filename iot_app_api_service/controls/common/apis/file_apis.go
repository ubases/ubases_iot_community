package apis

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotoss/file_store"
	"cloud_platform/iot_common/iotutil"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/h2non/filetype"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"fullPath"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Key      string `json:"key"`
}

var tempPath = iotconst.GetWorkTempDir() + string(filepath.Separator)

var Filecontroller File

type File struct {
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单文件，2：多文件, 3：base64图片)
// @Param file formData file true "file"
// @Param source formData file true "source （3 七牛）"
// @Param dir formData file true "dir 目录名"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e File) UploadFile(c *gin.Context) {
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1": // 单个文件
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		iotgin.ResSuccess(c, fileResponse)
		return
	case "2": // 多个文件
		multipartFile := e.multipleFile(c, urlPrefix)
		iotgin.ResSuccess(c, multipartFile)
		return
	case "3": // base64图片文件
		fileResponse = e.baseImg(c, fileResponse, urlPrefix)
		iotgin.ResSuccess(c, fileResponse)
	default:
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		iotgin.ResSuccess(c, fileResponse)
		return
	}

}

func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()

	dir, _ := c.GetPostForm("dir")
	fileName := fmt.Sprintf("%s/%s", dir, guid+".jpg")

	err := iotutil.IsNotExistMkDir(tempPath)
	err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
	if err != nil {
		iotgin.ResFailCode(c, "初始化文件路径失败", 500)
		return FileResponse{}
	}
	base64File := tempPath + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size:     iotutil.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	source, _ := c.GetPostForm("source")
	url, err := thirdUpload(config.Global.Oss.UseOss, fileName, base64File)
	if err != nil {
		iotgin.ResFailCode(c, "上传第三方失败", 500)
		return fileResponse
	}
	if source == "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	} else {
		fileResponse.Path = fileName
		fileResponse.FullPath = url
	}
	return fileResponse
}

func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	source, _ := c.GetPostForm("source")
	dir, _ := c.GetPostForm("dir")
	var multipartFile []FileResponse
	for _, f := range files {
		guid := uuid.New().String()
		fileName := fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(f.Filename))

		err := iotutil.IsNotExistMkDir(tempPath)
		err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
		if err != nil {
			iotgin.ResFailCode(c, "初始化文件路径失败", 500)
			return nil
		}
		multipartFileName := tempPath + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := iotutil.GetType(multipartFileName)
		if err1 == nil {
			url, err := thirdUpload(config.Global.Oss.UseOss, fileName, multipartFileName)
			if err != nil {
				iotgin.ResFailCode(c, "上传第三方失败", 500)
				return nil
			} else {
				fileResponse := FileResponse{
					Size:     iotutil.GetFileSize(multipartFileName),
					Path:     multipartFileName,
					FullPath: urlPerfix + url,
					Name:     f.Filename,
					Type:     fileType,
				}
				if source == "1" {
					fileResponse.Path = "/static/uploadfile/" + fileName
					fileResponse.FullPath = "/static/uploadfile/" + fileName
				} else {
					fileResponse.Path = fileName
					fileResponse.FullPath = url
				}
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	return multipartFile
}

func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPerfix string) (FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		iotgin.ResFailCode(c, "图片不能为空", 500)
		return FileResponse{}, true
	}
	// 上传文件至指定目录
	guid := uuid.New().String()
	dir, _ := c.GetPostForm("dir")
	fileName := iotutil.ConnectPath(dir, guid+iotutil.GetExt(files.Filename)) // fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(files.Filename))
	//fileName := fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(files.Filename))

	err = iotutil.IsNotExistMkDir(tempPath)
	err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
	if err != nil {
		iotgin.ResFailCode(c, "初始化文件路径失败", 500)
		return FileResponse{}, false
	}
	singleFile := tempPath + fileName
	//if  iotutil.GetFileSize(singleFile) > 10485760 {
	//	iotgin.ResFailCode(c, "上传文件大小超过10m", 500)
	//	return FileResponse{}, true
	//}
	_ = c.SaveUploadedFile(files, singleFile)
	fileKey, _ := iotutil.FileMD5(singleFile)
	//if err != nil {
	//	iotgin.ResFailCode(c, "初始化文件路径失败", 500)
	//	return FileResponse{}, false
	//}
	fileType, _ := iotutil.GetType(singleFile)
	fileResponse = FileResponse{
		Size:     iotutil.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPerfix + singleFile,
		Name:     files.Filename,
		Type:     fileType,
		Key:      fileKey,
	}
	source, _ := c.GetPostForm("source")
	url, err := thirdUpload(config.Global.Oss.UseOss, fileName, singleFile)
	if err != nil {
		iotgin.ResFailCode(c, "上传第三方失败", 500)
		return FileResponse{}, true
	}
	if source == "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	} else {
		fileResponse.Path = fileName
		fileResponse.FullPath = url
	}
	return fileResponse, false
}

func thirdUpload(source string, name string, path string) (string, error) {
	switch source {
	case "ali":
		return ossUpload(name, path)
	case "qiniu":
		return qiniuUpload(name, path)
	case "s3":
		return s3Upload(name, path)
	}
	return "", nil
}

// 阿里云OSS文件上传
func ossUpload(name string, path string) (string, error) {
	cnf := config.Global.Oss.Ali
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,
		AccessKeyID:     cnf.AccessKeyID,
		AccessKeySecret: cnf.AccessKeySecret,
		BucketName:      cnf.BucketName,
	}
	ossType := oss.Setup(file_store.AliYunOSS)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	return url, nil
}

// 亚马逊s3文件上传
func s3Upload(name string, path string) (string, error) {
	cnf := config.Global.Oss.S3
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,
		AccessKeyID:     cnf.AccessKeyID,
		AccessKeySecret: cnf.AccessKeySecret,
		BucketName:      cnf.BucketName,
		Region:          cnf.Region,
	}
	ossType := oss.Setup(file_store.AwsS3Kodo)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	return url, nil
}

// 七牛OSS上传
func qiniuUpload(name string, path string) (string, error) {
	cnf := config.Global.Oss.Qiniu
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,
		AccessKeyID:     cnf.AccessKeyID,
		AccessKeySecret: cnf.AccessKeySecret,
		BucketName:      cnf.BucketName,
	}
	ossType := oss.Setup(file_store.QiNiuKodo)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	return url, nil
}

/*
SaveFileToOSS 保存文件到OSS，用于form表单直接上传文件，参考管理平台代码

	const TestCaseTempPath = "testCase"
	f, err := apis.SaveFileToOSS(c, file, apis.TestCaseTempPath, "xlsx", "xls")
	if err != nil {
		return
	} else {
		req.TplFile = f.FullPath
		req.TplFileName = file.Filename
		req.TplFileSize = file.Size
		break
	}
*/
func SaveFileToOSS(c *gin.Context, f *multipart.FileHeader, savedir string, wantType ...string) (*FileResponse, error) {
	guid := uuid.New().String()
	dir := filepath.Join(tempPath, savedir)
	fileName := fmt.Sprintf("%s/%s", savedir, guid+iotutil.GetExt(f.Filename))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	savefile := filepath.Join(dir, guid+iotutil.GetExt(f.Filename))
	err = c.SaveUploadedFile(f, savefile)
	if err != nil {
		return nil, err
	}
	realtype, err := filetype.MatchFile(savefile)
	realExension := realtype.Extension
	//if realtype.MIME.Subtype != wantType {
	//如果没有设置想要的后缀，则不限制，如果设置则以设备为准  realtype.MIME.Subtype
	if len(wantType) != 0 && !iotutil.ArraysExistsString(wantType, realExension) {
		return nil, errors.New(fmt.Sprintf("文件类型错误,要求%s,实际%s", strings.Join(wantType, "、"), realExension))
	}
	url, err := thirdUpload(config.Global.Oss.UseOss, fileName, savefile)
	if err != nil {
		return nil, err
	}
	return &FileResponse{
		Size:     iotutil.GetFileSize(savefile),
		Path:     fileName,
		FullPath: url,
		Name:     f.Filename,
		Type:     realExension,
	}, nil
}
