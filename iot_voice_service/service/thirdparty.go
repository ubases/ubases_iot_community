package service

import (
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ThirdParty struct {
}

func NewThirdPartyApi() *ThirdParty {
	s := &ThirdParty{}
	return s
}

func (s *ThirdParty) GetList() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

func (s *ThirdParty) GetThirdLoginJson(c *gin.Context) {
	appKey := s.GetAppKeyByHost(c)
	if appKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	iotgin.ResJSON(c, http.StatusOK, s.GetDeveloperApp(appKey))
}

func (s *ThirdParty) GetAppKeyByHost(c *gin.Context) string {
	host := c.Request.Host
	arr := strings.Split(host, ".")
	//三级域名
	appKey := ""
	if len(arr) > 3 {
		appKey = arr[0]
	}
	iotlogger.LogHelper.Info("host: ", host, ", appKey:", appKey)
	return appKey
}

// GetDeveloperApp 获取公版的APP第三方登录信息apple-app-site-association
func (s *ThirdParty) GetDeveloperApp(appKey string) map[string]interface{} {
	//TODO　获取APP信息，TeamId和Ios包名
	appInfo, err := rpcclient.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{AppKey: appKey})
	if err != nil {
		return nil
	}
	if appInfo.Code != 200 {
		return nil
	}
	appData := appInfo.Data[0]
	//查询开发者APP信息
	return map[string]interface{}{
		"applinks": map[string]interface{}{
			"apps": []map[string]interface{}{},
			"details": []map[string]interface{}{
				{
					"appID": fmt.Sprintf("%s.%s", appData.IosTeamId, appData.IosPkgName),
					"paths": []string{
						fmt.Sprintf("/ai%s/wechat/*", appKey),
						fmt.Sprintf("/ai%s/qq/*", appKey),
						fmt.Sprintf("/ai%s/alipay/*", appKey),
						fmt.Sprintf("/ai%s/weibo/*", appKey),
						fmt.Sprintf("/ai%s/facebook/*", appKey),
						fmt.Sprintf("/ai%s/linkedin/*", appKey),
						fmt.Sprintf("/ai%s/twitter/*", appKey),
						fmt.Sprintf("/ai%s/line/*", appKey),
						fmt.Sprintf("/ai%s/google/*", appKey),
						fmt.Sprintf("/ai%s/douyin/*", appKey),
					},
				},
			},
		},
	}
}
