package apis

import (
	"bytes"
	"cloud_platform/iot_cloud_api_service/controls/config/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
	"os"
	"path"
)

type CommonController struct {
}

var Commoncontroller CommonController

var imageSavePath = "D://upload/upload/temp/"

//文件上传
//oss文件上传

// @Summary 本地文件下载
// @Description 本地文件下载
// @Tags 文件
// @Accept application/json
// @Param filename path string true "文件名称"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /common/local/getfile [post]
func (CommonController) GetLocalFile(c *gin.Context) {
	fileName := c.Param("fileName")
	bucket := c.Param("bucket")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	var path = fmt.Sprintf("%s%s/%s", imageSavePath, bucket, fileName)
	c.File(path)
}

// @Summary 本地文件上传
// @Description 本地文件上传
// @Tags 文件
// @Accept application/json
// @Param filename path string true "文件名称"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /common/local/uploadfile [post]
func (CommonController) UploadLocalFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		iotgin.ResFailCode(c, "upload image failed ", -1)
		return
	}
	//network_guide
	bucket := c.PostForm("bucket")

	//获取文件后缀
	headerFileName := header.Filename
	ext := path.Ext(headerFileName)

	var buffer bytes.Buffer
	buffer.WriteString(iotutil.Uuid())
	buffer.WriteString(ext)
	filename := buffer.String()

	//创建目录
	var path = fmt.Sprintf("%s%s/", imageSavePath, bucket)
	iotutil.CheckAndCreateFolder(path)

	var filepath string = fmt.Sprintf("%s%s", path, filename)
	out, err := os.Create(filepath)
	if err != nil {
		iotgin.ResFailCode(c, "upload images failed，"+err.Error(), -1)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		iotgin.ResFailCode(c, "copy failed copy  file failed", -1)
		return
	}
	iotgin.ResSuccess(c, filename)
}

func (CommonController) RegionList(c *gin.Context) {
	s := services.SysAreaService{}
	list, err := s.QueryRegionList()
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, list)
}
