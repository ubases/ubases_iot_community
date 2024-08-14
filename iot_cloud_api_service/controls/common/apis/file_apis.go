package apis

import (
	"archive/zip"
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotoss/file_store"
	"cloud_platform/iot_common/iotstruct"
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

var tempPath = iotconst.GetWorkTempDir() + string(filepath.Separator)

const ControlPanel = "controlPanel"
const TestCaseTempPath = "testCase"

var Filecontroller File

type File struct {
}

// UploadFile 上传文件
// @Summary 上传文件
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
	var fileResponse iotstruct.FileResponse

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
		fileResponse = e.base64Img(c, fileResponse, urlPrefix)
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

// UploadImage 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图片，2：多图片, 3：base64图片)
// @Param file formData file true "file"
// @Param dir formData file true "dir 目录名"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadImages [post]
// @Security Bearer
func (e File) UploadImage(c *gin.Context) {
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse iotstruct.FileResponse

	switch tag {
	case "1": // 单个文件
		files, err := c.FormFile("file")
		if err != nil {
			iotgin.ResFailCode(c, "图片不能为空", 500)
			return
		}
		dir, _ := c.GetPostForm("dir")
		dirPath := fmt.Sprintf("%s%s/", tempPath, dir)
		fileResponse, err := SaveFileToOSS(c, files, dirPath, "png", "jpg", "jpeg")
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		iotgin.ResSuccess(c, fileResponse)
		return
	case "2": // 多个文件
		multipartFile := e.multipleFile(c, urlPrefix)
		iotgin.ResSuccess(c, multipartFile)
		return
	case "3": // base64图片文件
		fileResponse = e.base64Img(c, fileResponse, urlPrefix)
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

// base64 image图片上传
func (e File) base64Img(c *gin.Context, fileResponse iotstruct.FileResponse, urlPerfix string) iotstruct.FileResponse {
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
		return iotstruct.FileResponse{}
	}
	base64File := tempPath + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = iotstruct.FileResponse{
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

func (e File) multipleFile(c *gin.Context, urlPerfix string) []iotstruct.FileResponse {
	files := c.Request.MultipartForm.File["file"]
	//source, _ := c.GetPostForm("source")
	dir, _ := c.GetPostForm("dir")
	dirPath := fmt.Sprintf("%s%s/", tempPath, dir)
	var multipartFile []iotstruct.FileResponse
	for _, f := range files {
		fileResponse, err := SaveFileToOSS(c, f, dirPath, "png", "jpg", "jpeg")
		if err != nil {
			iotgin.ResErrCli(c, err)
			break
		}
		multipartFile = append(multipartFile, *fileResponse)
	}
	return commonGlobal.SaveAttachmentRecord(dir, multipartFile...)
}

func SaveFileToOSS(c *gin.Context, f *multipart.FileHeader, savedir string, wantType ...string) (*iotstruct.FileResponse, error) {
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

	//检查面板是否包含index.html
	if savedir == ControlPanel && len(wantType) == 1 && wantType[0] == "zip" {
		err = VerifyPanel(savefile)
		if err != nil {
			return nil, err
		}
	}
	//读取md5
	md5, _ := iotutil.FileMD5(savefile)
	realtype, err := filetype.MatchFile(savefile)
	realExension := realtype.Extension
	//if realtype.MIME.Subtype != wantType {
	//如果没有设置想要的后缀，则不限制，如果设置则以设备为准  realtype.MIME.Subtype
	if len(wantType) != 0 && !iotutil.ArraysExistsString(wantType, realExension) {
		return nil, errors.New(fmt.Sprintf("文件类型错误,要求%s,实际%s", strings.Join(wantType, "、"), realExension))
	}
	url, err := Upload(config.Global.Oss.UseOss, fileName, savefile)
	if err != nil {
		return nil, err
	}
	fileRes := commonGlobal.SaveAttachmentRecord(dir, iotstruct.FileResponse{
		Size:     iotutil.GetFileSize(savefile),
		Path:     fileName,
		FullPath: url,
		Key:      md5,
		Name:     f.Filename,
		Type:     realExension,
	})
	return &fileRes[0], nil
}

// 上传文件至静态的OSS(公共的桶)
func SaveFileToQiniuStaticOSS(c *gin.Context, f *multipart.FileHeader, savedir string, wantType ...string) (*iotstruct.FileResponse, error) {
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
	//读取md5
	md5, _ := iotutil.FileMD5(savefile)
	realtype, err := filetype.MatchFile(savefile)
	realExension := realtype.Extension
	//if realtype.MIME.Subtype != wantType {
	//如果没有设置想要的后缀，则不限制，如果设置则以设备为准  realtype.MIME.Subtype
	if len(wantType) != 0 && !iotutil.ArraysExistsString(wantType, realExension) {
		return nil, errors.New(fmt.Sprintf("文件类型错误,要求%s,实际%s", strings.Join(wantType, "、"), realExension))
	}
	url, err := UploadStatic(config.Global.Oss.UseOss, fileName, savefile)
	if err != nil {
		return nil, err
	}

	fileRes := commonGlobal.SaveAttachmentRecord(dir, iotstruct.FileResponse{
		Size:     iotutil.GetFileSize(savefile),
		Path:     fileName,
		FullPath: url,
		Key:      md5,
		Name:     f.Filename,
		Type:     realExension,
	})
	return &fileRes[0], nil
}

func (e File) singleFile(c *gin.Context, fileResponse iotstruct.FileResponse, urlPerfix string) (iotstruct.FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		iotgin.ResFailCode(c, "图片不能为空", 500)
		return iotstruct.FileResponse{}, true
	}
	// 上传文件至指定目录
	guid := uuid.New().String()
	dir, _ := c.GetPostForm("dir")
	fileName := iotutil.ConnectPath(dir, guid+iotutil.GetExt(files.Filename)) // fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(files.Filename))

	//err = iotutil.IsNotExistMkDir(tempPath)
	err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
	if err != nil {
		iotgin.ResFailCode(c, "初始化文件路径失败", 500)
		return iotstruct.FileResponse{}, false
	}
	singleFile := tempPath + fileName
	err = c.SaveUploadedFile(files, singleFile)
	if err != nil {
		iotgin.ResFailCode(c, "文件存储出错", 500)
		return iotstruct.FileResponse{}, false
	}
	fileKey, err := iotutil.FileMD5(singleFile)
	if err != nil {
		iotgin.ResFailCode(c, "获取MD5失败", 500)
		return iotstruct.FileResponse{}, false
	}
	fileType, _ := iotutil.GetType(singleFile)
	fileResponse = iotstruct.FileResponse{
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
		return iotstruct.FileResponse{}, true
	}
	if source == "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	} else {
		fileResponse.Path = fileName
		fileResponse.FullPath = url
	}
	return commonGlobal.SaveAttachmentRecord(dir, fileResponse)[0], false
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

func Upload(source string, name string, path string) (string, error) {
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

func UploadStatic(source string, name string, path string) (string, error) {
	switch source {
	case "ali":
		return ossUpload(name, path)
	case "qiniu":
		return QiniuStaticUpload(name, path)
	case "s3":
		return s3Upload(name, path)
	}
	return "", nil
}

func ossUpload(name string, path string) (string, error) {
	cnf := config.Global.Oss.Ali
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,        //"https://iot-aithings-public.s3.amazonaws.com",
		AccessKeyID:     cnf.AccessKeyID,     //"AKIA2K2HXFHOE6NNLAWC",
		AccessKeySecret: cnf.AccessKeySecret, //"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",
		BucketName:      cnf.BucketName,      //"iot-aithings-public",
	}
	ossType := oss.Setup(file_store.AliYunOSS)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	return url, nil
}

func s3Upload(name string, path string) (string, error) {
	cnf := config.Global.Oss.S3
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,        //"https://iot-aithings-public.s3.amazonaws.com",
		AccessKeyID:     cnf.AccessKeyID,     //"AKIA2K2HXFHOE6NNLAWC",
		AccessKeySecret: cnf.AccessKeySecret, //"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",
		BucketName:      cnf.BucketName,      //"iot-aithings-public",
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

// 该为公共资源
func qiniuUpload(name string, path string) (string, error) {
	cnf := config.Global.Oss.Qiniu
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,        // "rab4kk5ki.hn-bkt.clouddn.com",
		AccessKeyID:     cnf.AccessKeyID,     //"_SRlsiDrTatwIIKLM84nINyCg0T25sA99B8GfTRF",
		AccessKeySecret: cnf.AccessKeySecret, //"VnW3PHcMtPMST7XOT2R0yROgsjiQjcULVQ7Az8Co",
		BucketName:      cnf.BucketName,      //   "aithings",
	}
	ossType := oss.Setup(file_store.QiNiuKodo)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	return url, nil
}

func QiniuStaticUpload(name string, path string) (string, error) {
	cnf := config.Global.Oss.Qiniu
	oss := file_store.OXS{
		Endpoint:        cnf.Endpoint,        // "rab4kk5ki.hn-bkt.clouddn.com",
		AccessKeyID:     cnf.AccessKeyID,     //"_SRlsiDrTatwIIKLM84nINyCg0T25sA99B8GfTRF",
		AccessKeySecret: cnf.AccessKeySecret, //"VnW3PHcMtPMST7XOT2R0yROgsjiQjcULVQ7Az8Co",
		BucketName:      cnf.BucketName,      //   "aithings",
	}
	ossType := oss.Setup(file_store.QiNiuKodo)
	err := ossType.UpLoad(name, path)
	if err != nil {
		return "", err
	}
	url := ossType.GetPublicUrl(name)
	//url = strings.Replace(url, "http://", "https://", 1)
	return url, nil //fmt.Sprintf("https://%s/%s", oss.Endpoint, url), nil
}

func VerifyPanel(file string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == "index.html" {
			return nil
		}
	}
	return errors.New("控制面板缺少index.html")
}
