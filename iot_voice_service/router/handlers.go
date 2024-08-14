package routers

import (
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/service"
	"fmt"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-session/session"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	Clients map[string]string
)

func accessTokenExpHandler(w http.ResponseWriter, r *http.Request) (exp time.Duration, err error) {
	return time.Duration(0), nil
}

// AccessTokenExpHandler
func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	iotlogger.LogHelper.Info("userAuthorizeHandler")
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return "", err
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		iotlogger.LogHelper.Info("userAuthorizeHandler ok false ")
		if r.Form == nil {
			r.ParseForm()
		}
		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	iotlogger.LogHelper.Info("userAuthorizeHandler end " + userID)
	return userID, nil
}

func loginHandler(c *gin.Context) {
	iotlogger.LogHelper.Info("loginHandler")
	store, err := session.Start(c.Request.Context(), c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if c.Request.Method == "POST" {
		if c.Request.Form == nil {
			if err := c.Request.ParseForm(); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		result, err := service.NewUserApi().Auth(c)
		if err != nil {
			iotlogger.LogHelper.Error("登录失败: ", err)
			c.JSON(http.StatusOK, ioterrs.ResponseModel{Code: 401, Message: err.Error()})
			return
		}
		iotlogger.LogHelper.Info(fmt.Sprintf("登录成功 %v", result))
		store.Set("LoggedInUserID", result["userid"])
		store.Set("axy_token", result["access_token"])
		store.Save()
		// c.Writer.Header().Set("Location", "/auth")
		// c.Writer.WriteHeader(http.StatusFound)
		c.JSON(http.StatusOK, ioterrs.ResponseModel{Code: 0, Message: "Success"})
		return
	} else {
		v := c.Request.URL.Query()
		for s, x := range v {
			val := ""
			if len(x) > 0 {
				val = x[0]
			}
			store.Set(s, val)
		}
		store.Save()
	}
	//TODO 地址配置
	//outputHTML(c.Writer, c.Request, "/opt/bat/temp/web/login.html")
	outputHTML(c.Writer, c.Request, "./web/login.html")
}

func authHandler(c *gin.Context) {
	iotlogger.LogHelper.Info("authHandler")
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Writer.Header().Set("Location", "/login")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}
	//outputHTML(c.Writer, c.Request, "/opt/bat/temp/web/auth.html")
	outputHTML(c.Writer, c.Request, "./web/auth.html")
}

func authorizeHandler(c *gin.Context) {
	store, err := session.Start(c.Request.Context(), c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	} else {
		//http://127.0.0.1:31003/login?
		//client_id=AlexaTest&
		//response_type=code&
		//state=A2SAAEAEO1Mz-7r4DhB5uLdOZO2lGsB8GRPsYGgD7YP1cTChooe3j9EgvokfDhVydzpMK7FkgadoeMuY1X4iOEiaky3AoMFvMY4D3r4yQZn35wwEsWXKkov47KeGxZhDSsf11XHO3-thJdKTHoeMDOWqJ5ybiJbnVp9wRj2GwKhzCiW720IyCuZv-WIIhj-sK0C5WZEHPR032Muh8Zi--LKGAT6j8sOSW7E04KZsS9OxAkK9WNLGhZSkR3R1OTCYoPcN0LQNauI0fWC3CavkFylSi1r2G8GWcOjjc0vww0dIrLs-zfBToe_LX-oqO3QPz_7TtmQibc_bmHBpqk-KwEo_1zyMTqHgI9bh-BEqdnSCMxwzSeM3bbsVxOj3FBfYC87BMEO-5zkGKDxy2KeI-QjYAM1-I9xQNGeQUOMJBIe2aAUZyaPuNxxG2wGENyCu9f1qgxjg0VDJYtv-8heDVHJnDjSpzWUByHJEvn0ov4FF2Mqeic9TsWWqW9HsbR4e-codlKT-LA0A3hReVYbB67quT-NOjWSMd0egArfZT57yNuZfVPM9GVWhsMi75kzZDbirzRdliPzJMNshKvIwnnsLMBPLqXA1gzVWvi8mwtXBnpwTAGNpsU2GgPQGHUKRhu2_yPm1jp8yfHJ6fZeb3BUsfGuWtZDSrjMulK2CKWbBaYaeodAxMM&
		//scope=profile%3Auser_id&
		//redirect_uri=https%3A%2F%2Fpitangui.amazon.com%2Fapi%2Fskill%2Flink%2FM3U4PMEYLO6DKE
		var (
			client_id, _     = store.Get("client_id")
			response_type, _ = store.Get("response_type")
			state, _         = store.Get("state")
			scope, _         = store.Get("scope")
			redirect_uri, _  = store.Get("redirect_uri")
		)
		form = url.Values{}
		form.Set("client_id", iotutil.ToString(client_id))
		form.Set("response_type", iotutil.ToString(response_type))
		form.Set("state", iotutil.ToString(state))
		form.Set("scope", iotutil.ToString(scope))
		form.Set("redirect_uri", iotutil.ToString(redirect_uri))
	}
	c.Request.Form = form

	store.Delete("ReturnUri")
	store.Save()
	//c.Writer.Header().Set("Location", "/login")
	ginserver.HandleAuthorizeRequest(c)
}

//AccessTokenExpHandler

func handleTokenRequest(c *gin.Context) {
	headers := c.Request.Header
	iotlogger.LogHelper.Info("===> headers：", iotutil.ToString(headers))
	clientId := c.Request.PostForm.Get("client_id")
	for k, v := range c.Request.PostForm {
		iotlogger.LogHelper.Info("===> %s：%s", k, v)
	}
	//TODO 临时兼容代码
	tClientId := c.Request.Form.Get("client_id")
	if clientId == "AlexaTest" || tClientId == "" {
		c.Request.Form = url.Values{}
		clientId := c.Request.PostForm.Get("client_id")
		redirectUri := c.Request.PostForm.Get("redirect_uri")
		grantType := c.Request.PostForm.Get("grant_type")
		code := c.Request.PostForm.Get("code")
		refreshToken := c.Request.PostForm.Get("refresh_token")
		clientSecret := ""
		if v, ok := Clients[clientId]; ok {
			clientSecret = v
		}
		c.Request.Form.Set("client_id", clientId)
		c.Request.Form.Set("redirect_uri", redirectUri)
		c.Request.Form.Set("grant_type", grantType)
		c.Request.Form.Set("client_secret", clientSecret)
		c.Request.Form.Set("code", code) //"NDFHZDGYNMITMWMWYY0ZYTE1LTLHNWETNJZLYZRIZTA5NTBK")
		c.Request.Form.Set("refresh_token", refreshToken)
		iotlogger.LogHelper.Infof("===> form：clientId=%v, redirectUri=%v, grantType=%v, clientSecret=%v, code=%v, refreshToken=%v",
			clientId, redirectUri, grantType, clientSecret, code, refreshToken)
	}
	ginserver.HandleTokenRequest(c)
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

func GetRegionList(c *gin.Context) {
	lang := c.GetHeader("lang")
	if lang == "" {
		lang = "zh"
	}
	ip := c.ClientIP()
	resp, err := service.RegionList(lang, ip)
	if err != nil {
		c.JSON(http.StatusOK, ioterrs.ResponseModel{Code: 401, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ioterrs.ResponseModel{Code: 0, Message: "Success", Data: resp})
}
