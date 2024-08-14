package apis

import (
	"bytes"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

var OemAppIntroducecontroller OemAppIntroduceController

var serviceIntroduce apiservice.OemAppIntroduceService

type OemAppIntroduceController struct {
}

// 用户协议,隐私政策,关于我们,新增
func (OemAppIntroduceController) OemAppIntroduceCheckVersion(c *gin.Context) {
	var req entitys.OemAppIntroduceVersionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceCheckVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 用户协议,隐私政策,关于我们,新增
func (OemAppIntroduceController) OemAppIntroduceAdd(c *gin.Context) {
	var req entitys.OemAppIntroduceSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceAdd(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 用户协议,隐私政策,关于我们,修改
func (OemAppIntroduceController) OemAppIntroduceUpdate(c *gin.Context) {
	var req entitys.OemAppIntroduceSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceUpdate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// OemAppIntroduceCopy 用户协议,隐私政策,关于我们,复制
func (OemAppIntroduceController) OemAppIntroduceCopy(c *gin.Context) {
	var req entitys.OemAppIntroduceCopyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.AppId == 0 {
		iotgin.ResBadRequest(c, "appId")
		return
	}
	if req.NewVersion == "" {
		iotgin.ResBadRequest(c, "newVersion")
		return
	}
	if req.OldVersion == "" {
		iotgin.ResBadRequest(c, "oldVersion")
		return
	}
	err = serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceCopy(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 用户协议,隐私政策,关于我们,修改
func (OemAppIntroduceController) OemAppIntroduceStatusEnable(c *gin.Context) {
	var req entitys.OemAppIntroduceStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceStatusEnable(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 用户协议,隐私政策,关于我们,获取详细
func (OemAppIntroduceController) OemAppIntroduceDetail(c *gin.Context) {
	var req entitys.OemAppIntroduceDetailReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceDetail(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// Voice doc
func (OemAppIntroduceController) OemAppVoiceIntroduceDetail(c *gin.Context) {
	var req entitys.OemAppIntroduceDetailReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.Version = ""
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceDetail(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 用户协议,隐私政策,关于我们,获取详细
func (OemAppIntroduceController) OemAppIntroduceList(c *gin.Context) {
	var req entitys.OemAppIntroduceListReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, _, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 用户协议,隐私政策,关于我们,获取详细 url地址
func (OemAppIntroduceController) OemAppIntroduceUrlList(c *gin.Context) {
	code := c.Query("contentType")
	if code == "" {
		iotgin.ResBadRequest(c, "contentType")
		return
	}
	appId := c.Query("appId")
	if appId == "" {
		iotgin.ResBadRequest(c, "appId")
		return
	}
	//appKeyStr := strings.Split(appKey, ".")[0]
	req := entitys.OemAppIntroduceListReq{
		AppId:       appId,
		ContentType: iotutil.ToInt32(code),
	}

	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceLinkList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppIntroduceController) OemAppIntroduceTemplateUrlList(c *gin.Context) {
	id := c.Query("contentType")
	if id == "" {
		iotgin.ResBadRequest(c, "contentType")
		return
	}
	tid := iotutil.ToInt(id)
	res, err := serviceIntroduce.SetContext(controls.WithOpenUserContext(c)).OemAppIntroduceTemplateLink(tid)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 协议链接html
func (OemAppIntroduceController) GetIntroduceByAppHtml(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	ret, err := serviceIntroduce.SetContext(controls.WithUserContext(c)).OemAppIntroduceDetailById(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	title := ""
	if ret.ContentType == 1 {
		title = "用户协议"
	} else if ret.ContentType == 2 {
		title = "隐私政策"
	} else if ret.ContentType == 3 {
		title = "关于我们"
	}

	indextmpl := strings.Join([]string{iotconst.GetTemplatesDir(), "index.tmpl"}, string(filepath.Separator))
	tmp, err := template.ParseFiles(indextmpl)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, gin.H{
		"title":   title,
		"content": ret.Content,
	}); err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//c.HTML(http.StatusOK, "", buf.String())

	c.Writer.WriteString(buf.String())
}

// 协议模板链接打开
func (OemAppIntroduceController) GetIntroduceTtemplateByAppHtml(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	ret, err := serviceIntroduce.SetContext(controls.WithUserContext(c)).OemAppIntroduceTemplateDetailById(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	title := ""
	if ret.ContentType == 1 {
		title = "用户协议"
	} else if ret.ContentType == 2 {
		title = "隐私政策"
	} else if ret.ContentType == 3 {
		title = "关于我们"
	}
	indextmpl := strings.Join([]string{iotconst.GetTemplatesDir(), "index.tmpl"}, string(filepath.Separator))
	tmp, err := template.ParseFiles(indextmpl)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, gin.H{
		"title":   title,
		"content": ret.Content,
	}); err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//c.HTML(http.StatusOK, "", buf.String())

	c.Writer.WriteString(buf.String())
}
