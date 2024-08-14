package alexa

import (
	"bytes"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/entitys"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	AlexaDiscoverDevicesResponse = "Alexa.Discovery"
	AlexaDeleteReport            = "DeleteReport"
	AlexaAddOrUpdateReport       = "AddOrUpdateReport"
	AcceptGrant                  = "Alexa.Authorization"
)

func RunController(res *entitys.DirectiveRequet, data []byte, userId, token string) error {
	cachedAlexaTokenInfo(userId, token, getCorrleationToken(res), *res)
	switch res.Directive.Header.Namespace {
	case AlexaDiscoverDevicesResponse:
		DiscoverDevices(res, userId, token)
	case AcceptGrant:
		GetAlexaToken(userId, res)
	case AlexaDeleteReport:
	case AlexaAddOrUpdateReport:
	default:
		return SetProperties(res, data, userId, token)
	}
	return nil
}

func AddRefreshTokenUserId(userId string) {
	if RefreshTokenUserIds == nil {
		RefreshTokenUserIds = &sync.Map{}
	}
	RefreshTokenUserIds.Store(userId, 1)
}

func GetAlexaToken(userId string, res *entitys.DirectiveRequet) (string, error) {
	//res.Directive.Payload = entitys.AlexaControlPropertiesContext{}
	res.Directive.Header.Namespace = "Alexa.Authorization"
	res.Directive.Header.Name = "AcceptGrant.Response"
	var apiUrl = res.Directive.Header.AlexaAuthTokenUrl
	var payload entitys.AlexaVoiceGrantPayload
	err := iotutil.StructToStructErr(res.Directive.Payload, &payload)
	if err != nil {
		return "", err
	}
	requestData := map[string]interface{}{
		"grant_type":    "authorization_code", //payload.Grant.Type,
		"code":          payload.Grant.Code,
		"client_id":     res.Directive.Header.AlexaClientId,
		"client_secret": res.Directive.Header.AlexaClientSecret,
	}
	jsonStr, _ := json.Marshal(requestData)
	iotlogger.LogHelper.Helper.Infof("alexa get token paramaters: url: %s, res:%s ", apiUrl, string(jsonStr))
	// 创建 HTTP POST 请求
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	// 读取响应体并处理
	defer resp.Body.Close()

	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		iotlogger.LogHelper.Helper.Debug("向Alexa获取Token报错: err： ", err.Error())
		return "", err
	}
	iotlogger.LogHelper.Helper.Debug("向Alexa获取Token响应: ", string(resbody))
	resq, err := iotutil.JsonToMapErr(string(resbody))
	if err != nil {
		return "", err
	}
	//logger.info(f"access_token: {lwa_tokens['access_token']}")
	//logger.info(f"refresh_token: {lwa_tokens['refresh_token']}")
	//logger.info(f"token_type: {lwa_tokens['token_type']}")
	//logger.info(f"expires_in: {lwa_tokens['expires_in']}")
	accessToken := resq["access_token"]
	if accessToken == "" {
		return "", errors.New(string(resbody))
	}
	var alexaToken map[string]interface{} = resq
	alexaToken["updateTokenTime"] = time.Now().Unix()
	cmd := iotredis.GetClient().HMSet(context.Background(), fmt.Sprintf(iotconst.AlexaVoiceUserTokenKey, userId), alexaToken)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return "", nil
}

func GetAlexaRefreshToken(userId string) (string, error) {
	tokenInfo, err := getCachedAlexaTokenInfo(userId)
	if err != nil {
		return "", err
	}
	refreshToken := tokenInfo["refresh_token"]
	alexaClientId := tokenInfo["alexaClientId"]
	alexaClientSecret := tokenInfo["alexaClientSecret"]
	apiUrl := tokenInfo["alexaAuthTokenUrl"]
	//是否需要判断上次token刷新的时间
	//updateTokenTime := tokenInfo["updateTokenTime"]
	//if updateTokenTime != "" {
	//	tokenTime, _ := iotutil.ToInt64AndErr(updateTokenTime)
	//	if tokenTime > time.Now().Add(-30*time.Minute).Unix() {
	//		return "", nil
	//	}
	//}

	if apiUrl == "" || alexaClientSecret == "" || alexaClientId == "" || refreshToken == "" {
		return "", errors.New("参数异常")
	}
	requestData := map[string]interface{}{
		"grant_type":    "refresh_token", //payload.Grant.Type,
		"refresh_token": refreshToken,
		"client_id":     alexaClientId,
		"client_secret": alexaClientSecret,
	}
	jsonStr, _ := json.Marshal(requestData)
	iotlogger.LogHelper.Helper.Infof("alexa get refresh token paramaters: url: %s, res:%s ", apiUrl, string(jsonStr))
	// 创建 HTTP POST 请求
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	// 读取响应体并处理
	defer resp.Body.Close()

	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		iotlogger.LogHelper.Helper.Debug("向Alexa获取refresh_token报错: err： ", err.Error())
		return "", err
	}
	iotlogger.LogHelper.Helper.Debug("向Alexa获取refresh_token响应: ", string(resbody))
	resq, err := iotutil.JsonToMapErr(string(resbody))
	if err != nil {
		return "", err
	}
	accessToken := resq["access_token"]
	if accessToken == "" {
		return "", errors.New(string(resbody))
	}
	var alexaToken map[string]interface{} = resq
	alexaToken["updateTokenTime"] = time.Now().Unix()
	cmd := iotredis.GetClient().HMSet(context.Background(), fmt.Sprintf(iotconst.AlexaVoiceUserTokenKey, userId), alexaToken)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return "", nil
}

func getCorrleationToken(res *entitys.DirectiveRequet) string {
	correlToken := ""
	if res != nil {
		if res.Directive.Header.CorrelationToken != "" {
			correlToken = res.Directive.Header.CorrelationToken
		}
	}
	return correlToken
}

func cachedAlexaTokenInfo(userId string, token string, aToken string, res entitys.DirectiveRequet) error {
	//缓存设备信息
	var alexaToken map[string]interface{} = map[string]interface{}{}
	if token != "" {
		alexaToken["token"] = token
	}
	if aToken != "" {
		alexaToken["correlationToken"] = aToken
	}
	if res.Directive.Header.AlexaClientId != "" {
		alexaToken["alexaClientId"] = res.Directive.Header.AlexaClientId
		alexaToken["alexaClientSecret"] = res.Directive.Header.AlexaClientSecret
		alexaToken["alexaAuthTokenUrl"] = res.Directive.Header.AlexaAuthTokenUrl
		alexaToken["alexaEventUrl"] = res.Directive.Header.AlexaEventUrl
	}
	if len(alexaToken) == 0 {
		return nil
	}
	iotredis.GetClient().HMSet(context.Background(), fmt.Sprintf(iotconst.AlexaVoiceUserTokenKey, userId), alexaToken)
	//增加用户缓存逻辑
	iotredis.GetClient().HMSet(context.Background(), iotconst.AlexaVoiceAllUserIdKey, map[string]interface{}{
		userId: 1,
	})
	RefreshTokenUserIds.Store(userId, 1)
	return nil
}

func getCachedAlexaTokenInfo(userId string) (map[string]string, error) {
	cmd := iotredis.GetClient().HGetAll(context.Background(), fmt.Sprintf(iotconst.AlexaVoiceUserTokenKey, userId))
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd.Val(), nil
}
