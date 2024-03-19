package apis

//
//import (
//	"cloud_platform/iot_common/iotgin"
//	"cloud_platform/iot_common/iotoss/file_store"
//	"cloud_platform/iot_common/iotutil"
//	"encoding/base64"
//	"fmt"
//	"io/ioutil"
//	"strings"
//
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//)
//
//type FileResponse struct {
//	Size     int64  `json:"size"`
//	Path     string `json:"path"`
//	FullPath string `json:"full_path"`
//	Name     string `json:"name"`
//	Type     string `json:"type"`
//}
//
//const tempPath = "/usr/local/bat/temp/"
//
//var Filecontroller File
//
//type File struct {
//}
//
//// UploadFile 上传图片
//// @Summary 上传图片
//// @Description 获取JSON
//// @Tags 公共接口
//// @Accept multipart/form-data
//// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
//// @Param file formData file true "file"
//// @Param source formData file true "source （3 七牛）"
//// @Param dir formData file true "dir 目录名"
//// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
//// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
//// @Router /fileStore/uploadOssFile [post]
//// @Security Bearer
//func (e File) UploadFile(c *gin.Context) {
//	tag, _ := c.GetPostForm("type")
//	urlPrefix := fmt.Sprintf("http://%s/", c.Request.Host)
//	var fileResponse FileResponse
//
//	switch tag {
//	case "1": // 单图
//		var done bool
//		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
//		if done {
//			return
//		}
//		iotgin.ResSuccess(c, fileResponse)
//		return
//	case "2": // 多图
//		multipartFile := e.multipleFile(c, urlPrefix)
//		iotgin.ResSuccess(c, multipartFile)
//		return
//	case "3": // base64
//		fileResponse = e.baseImg(c, fileResponse, urlPrefix)
//		iotgin.ResSuccess(c, fileResponse)
//	default:
//		var done bool
//		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
//		if done {
//			return
//		}
//		iotgin.ResSuccess(c, fileResponse)
//		return
//	}
//}
//
//func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
//	files, _ := c.GetPostForm("file")
//	file2list := strings.Split(files, ",")
//	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
//	guid := uuid.New().String()
//
//	dir, _ := c.GetPostForm("dir")
//	fileName := fmt.Sprintf("%s/%s", dir, guid+".jpg")
//
//	err := iotutil.IsNotExistMkDir(tempPath)
//	err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
//	if err != nil {
//		iotgin.ResFailCode(c, "初始化文件路径失败", 500)
//		return FileResponse{}
//	}
//	base64File := tempPath + fileName
//	_ = ioutil.WriteFile(base64File, ddd, 0666)
//	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
//	fileResponse = FileResponse{
//		Size:     iotutil.GetFileSize(base64File),
//		Path:     base64File,
//		FullPath: urlPerfix + base64File,
//		Name:     "",
//		Type:     typeStr,
//	}
//	source, _ := c.GetPostForm("source")
//	url, err := thirdUpload(source, fileName, base64File)
//	if err != nil {
//		iotgin.ResFailCode(c, "上传第三方失败", 500)
//		return fileResponse
//	}
//	if source == "1" {
//		fileResponse.Path = "/static/uploadfile/" + fileName
//		fileResponse.FullPath = "/static/uploadfile/" + fileName
//	} else {
//		fileResponse.Path = fileName
//		fileResponse.FullPath = url
//	}
//	return fileResponse
//}
//
//func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
//	files := c.Request.MultipartForm.File["file"]
//	source, _ := c.GetPostForm("source")
//	dir, _ := c.GetPostForm("dir")
//	var multipartFile []FileResponse
//	for _, f := range files {
//		guid := uuid.New().String()
//		fileName := fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(f.Filename))
//
//		err := iotutil.IsNotExistMkDir(tempPath)
//		err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
//		if err != nil {
//			iotgin.ResFailCode(c, "初始化文件路径失败", 500)
//			return nil
//		}
//		multipartFileName := tempPath + fileName
//		err1 := c.SaveUploadedFile(f, multipartFileName)
//		fileType, _ := iotutil.GetType(multipartFileName)
//		if err1 == nil {
//			url, err := thirdUpload(source, fileName, multipartFileName)
//			if err != nil {
//				iotgin.ResFailCode(c, "上传第三方失败", 500)
//				return nil
//			} else {
//				fileResponse := FileResponse{
//					Size:     iotutil.GetFileSize(multipartFileName),
//					Path:     multipartFileName,
//					FullPath: urlPerfix + url,
//					Name:     f.Filename,
//					Type:     fileType,
//				}
//				if source == "1" {
//					fileResponse.Path = "/static/uploadfile/" + fileName
//					fileResponse.FullPath = "/static/uploadfile/" + fileName
//				} else {
//					fileResponse.Path = fileName
//					fileResponse.FullPath = url
//				}
//				multipartFile = append(multipartFile, fileResponse)
//			}
//		}
//	}
//	return multipartFile
//}
//
//func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPerfix string) (FileResponse, bool) {
//	files, err := c.FormFile("file")
//
//	if err != nil {
//		iotgin.ResFailCode(c, "图片不能为空", 500)
//		return FileResponse{}, true
//	}
//	// 上传文件至指定目录
//	guid := uuid.New().String()
//	dir, _ := c.GetPostForm("dir")
//	fileName := fmt.Sprintf("%s/%s", dir, guid+iotutil.GetExt(files.Filename))
//
//	err = iotutil.IsNotExistMkDir(tempPath)
//	err = iotutil.IsNotExistMkDir(tempPath + "/" + dir)
//	if err != nil {
//		iotgin.ResFailCode(c, "初始化文件路径失败", 500)
//		return FileResponse{}, false
//	}
//	singleFile := tempPath + fileName
//	_ = c.SaveUploadedFile(files, singleFile)
//	fileType, _ := iotutil.GetType(singleFile)
//	fileResponse = FileResponse{
//		Size:     iotutil.GetFileSize(singleFile),
//		Path:     singleFile,
//		FullPath: urlPerfix + singleFile,
//		Name:     files.Filename,
//		Type:     fileType,
//	}
//	source, _ := c.GetPostForm("source")
//	url, err := thirdUpload(source, fileName, singleFile)
//	if err != nil {
//		iotgin.ResFailCode(c, "上传第三方失败", 500)
//		return FileResponse{}, true
//	}
//	if source == "1" {
//		fileResponse.Path = "/static/uploadfile/" + fileName
//		fileResponse.FullPath = "/static/uploadfile/" + fileName
//	} else {
//		fileResponse.Path = fileName
//		fileResponse.FullPath = url
//	}
//	return fileResponse, false
//}
//
//func thirdUpload(source string, name string, path string) (string, error) {
//	switch source {
//	case "2":
//		return ossUpload(name, path)
//	case "3":
//		return qiniuUpload(name, path)
//	}
//	return "", nil
//}
//
//func ossUpload(name string, path string) (string, error) {
//	oss := file_store.ALiYunOSS{}
//	return "", oss.UpLoad(name, path)
//}
//
//func qiniuUpload(name string, path string) (string, error) {
//	oss := file_store.OXS{
//		Endpoint:        "rab4kk5ki.hn-bkt.clouddn.com",
//		AccessKeyID:     "_SRlsiDrTatwIIKLM84nINyCg0T25sA99B8GfTRF",
//		AccessKeySecret: "VnW3PHcMtPMST7XOT2R0yROgsjiQjcULVQ7Az8Co",
//		BucketName:      "aithings",
//	}
//	ossType := oss.Setup(file_store.QiNiuKodo)
//	err := ossType.UpLoad(name, path)
//	if err != nil {
//		return "", err
//	}
//	url := ossType.GetUrl(name)
//	return url, nil //fmt.Sprintf("https://%s/%s", oss.Endpoint, url), nil
//}
//
////注册路由
//func (uc File) RegisterRouter(e *gin.Engine) {
//	webApiPrefix := "/v1/platform/web"
//	r := e.Group(webApiPrefix).Group("/common") //.Use(iotgin.AuthCheck)
//	{
//		r.POST("/fileStore/uploadOssFile", uc.UploadFile)
//	}
//}
