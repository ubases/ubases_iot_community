package iotgin

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

type AppLog struct {
	Id             int64             `json:"id"`
	Account        string            `json:"account"`
	AppKey         string            `json:"appKey"`
	TenantId       string            `json:"tenantId"`
	RegionServerId int64             `json:"regionServerId"`
	LogType        string            `json:"logType"`
	EventName      string            `json:"eventName"`
	Details        map[string]string `json:"details"`
	CreatedAt      time.Time         `json:"createdAt"`
}

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AppLogger(jspub *jetstream.JsPublisherMgr) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		ip := c.ClientIP()
		sys := c.GetHeader("x-sys-info")
		appKey := c.GetHeader("appKey")
		tenantId := c.GetHeader("tenantId")
		regionServerId := c.GetHeader("region")
		account, ok := c.Get("Account")
		if !ok {
			iotlogger.LogHelper.Helper.Errorf("get account: %v", account)
			return
		}
		resp := struct {
			Code int
			Msg  string
		}{}
		if err := json.Unmarshal([]byte(blw.Body.String()), &resp); err != nil {
			iotlogger.LogHelper.Helper.Errorf("json unmarshal response error: %v", err)
			return
		}
		if resp.Msg == "ok" {
			resp.Msg = "success"
		}
		var logType string
		if resp.Code == 0 {
			logType = iotconst.APP_OPERATE_LOG
		} else {
			logType = iotconst.APP_ERROR_LOG
		}
		if event, ok := iotconst.LogEventMap[c.Request.RequestURI]; ok {
			resp.Msg = event
			if err := newAppLog(jspub, account.(string), logType, event, ip, sys, resp.Msg, appKey, tenantId, regionServerId); err != nil {
				iotlogger.LogHelper.Helper.Errorf("new app register log error: %v", err)
				return
			}
		} else {
			// 以下针对uri+请求参数的情况，需要再做一次判断
			for k, v := range iotconst.LogEventMap {
				if strings.Contains(c.Request.RequestURI, k) {
					resp.Msg = v
					if err := newAppLog(jspub, account.(string), logType, v, ip, sys, resp.Msg, appKey, tenantId, regionServerId); err != nil {
						iotlogger.LogHelper.Helper.Errorf("new app register log error: %v", err)
						return
					}
				}
			}
		}
	}
}

func newAppLog(jspub *jetstream.JsPublisherMgr, account, logType, eventName, ip, sys, msg, appKey, tenanntId, regionServerId string) error {
	var regionId int64 = 0
	regionId, _ = iotutil.ToInt64AndErr(regionServerId)
	appLog := AppLog{
		Id:             iotutil.GetNextSeqInt64(),
		Account:        account,
		AppKey:         appKey,
		TenantId:       tenanntId,
		RegionServerId: regionId,
		LogType:        logType,
		EventName:      eventName,
		Details: map[string]string{
			"ip":     ip,
			"system": sys,
			"msg":    msg,
		},
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(appLog)
	if err != nil {
		return err
	}
	pd := &jetstream.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_RECORDS,
		Data:    string(data),
	}
	jspub.PushData(pd)
	iotlogger.LogHelper.Helper.Debugf("subject: %s data: %s", pd.Subject, pd.Data)
	return nil
}
