package iotgin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotlogger"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		//begin for debug
		data := make([]byte, 0)
		reqMethod := c.Request.Method
		if reqMethod == "POST" || reqMethod == "PUT" {
			_, err := c.MultipartForm()
			//如果是文件上传，则不能将body进行重新赋值
			if err == nil {
				fmt.Println("文件上传")
			} else {
				//为了打印日志，将数据存储到body中
				data, _ = c.GetRawData()
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			}
		} else {
			data, _ = c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		//end for debug
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		//path := c.Request.URL.Path
		path := c.Request.RequestURI

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		iotlogger.LogHelper.Infof("| %3d | %13v | %15s | %s  %s |", statusCode, latency, clientIP, method, path)

		//begin for debug
		iotlogger.LogHelper.Debug("Authorization=", c.Request.Header.Get("Authorization"))
		if len(data) > 0 {
			iotlogger.LogHelper.Debug(string(data))
		}
		//end for debug
	}
}

// 完整输出输入和输出参数
func GinLoggerAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		//begin for debug
		data := make([]byte, 0)
		reqMethod := c.Request.Method
		if reqMethod == "POST" || reqMethod == "PUT" {
			_, err := c.MultipartForm()
			//如果是文件上传，则不能将body进行重新赋值
			if err == nil {
				fmt.Println("文件上传")
			} else {
				//为了打印日志，将数据存储到body中
				data, _ = c.GetRawData()
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			}
		} else {
			data, _ = c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}

		//打印Write数据
		strBody := ""
		var blw bodyLogWriter
		//if we need to log res body
		blw = bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		//end for debug
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		//path := c.Request.URL.Path
		path := c.Request.RequestURI

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		iotlogger.LogHelper.Infof("| %3d | %13v | %15s | %s  %s |", statusCode, latency, clientIP, method, path)

		//begin for debug
		if len(data) > 0 {
			iotlogger.LogHelper.Debug(string(data))
		}

		if strings.Index(path, "/download") == -1 {
			strBody = strings.Trim(blw.bodyBuf.String(), "\n")
			//if len(strBody) > MAX_PRINT_BODY_LEN {
			//	strBody = strBody[:(MAX_PRINT_BODY_LEN - 1)]
			//}
			iotlogger.LogHelper.Info("response=", strBody)
		}
		//end for debug
	}
}

const MAX_PRINT_BODY_LEN = 3000

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}
