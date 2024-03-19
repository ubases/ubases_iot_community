package routers

import (
	"cloud_platform/iot_smart_speaker_service/config"
	"cloud_platform/iot_smart_speaker_service/service"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"fmt"
	"net/http"
	"time"

	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"

	oredis "cloud_platform/iot_smart_speaker_service/redis"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iottrace"
)

var (
	manager     *manage.Manager
	clientStore *store.ClientStore
	expire      = time.Hour * 24 * 7
)

func InitOAuth2() error {
	manager = manage.NewDefaultManager()

	defaultAuthorizeCodeTokenCfg := &manage.Config{AccessTokenExp: expire, RefreshTokenExp: expire + time.Hour, IsGenerateRefresh: true}
	manager.SetAuthorizeCodeTokenCfg(defaultAuthorizeCodeTokenCfg)
	//manager.SetAuthorizeCodeExp(time.Hour * 24 * 30)
	// token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     config.Global.Redis.Addrs[0],
		DB:       15,
		Password: config.Global.Redis.Password,
	}))

	manager.SetValidateURIHandler(func(baseURI, redirectURI string) error {
		return nil
	})

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(config.Global.Jwt.SigningKey), jwt.SigningMethodHS256))

	clientStore = store.NewClientStore()

	list, err := common.NewVoiceApi().GetList()
	if err != nil {
		iotlogger.LogHelper.Error("oauth server init read voice error: ", err)
		return err
	}
	iotlogger.LogHelper.Helper.Debug("voice list: ", list)
	Clients = make(map[string]string)
	for _, v := range list {
		client_id := v["clientId"].(string)
		client_secret := v["clientScrect"].(string)
		clientStore.Set(client_id, &models.Client{
			ID:     client_id,
			Secret: client_secret,
			Domain: "https://127.0.0.1", //TODO 配置
		})
		Clients[client_id] = client_secret
	}
	//固定测试数据
	var otherClients = map[string]string{
		"XiaoaiTest":  "XiaoaiTest123",
		"TianmaoTest": "TianmaoTest123",
		"XiaomiTest":  "XiaomiTest123",
		"AlexaTest":   "AlexaTest123",
		"GoogleTest":  "GoogleTest123",
	}
	for k, v := range otherClients {
		clientStore.Set(k, &models.Client{
			ID:     k,
			Secret: v,
			Domain: "https://127.0.0.1",
		})
		Clients[k] = v
	}
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
	ginserver.SetUserAuthorizationHandler(userAuthorizeHandler)
	ginserver.SetAccessTokenExpHandler(accessTokenExpHandler)
	//ginserver.SetTokenType("bearer")

	ginserver.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		iotlogger.LogHelper.Info("SetInternalErrorHandler Internal Error:", err.Error())
		return
	})

	ginserver.SetResponseErrorHandler(func(re *errors.Response) {
		iotlogger.LogHelper.Info("SetResponseErrorHandler Response Error:", re.Error.Error())
	})

	return nil
}

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //生产模式启动

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 2, "msg": "未找到请求路由的处理函数"})
		c.Abort()
	})
	//router.Use(gin.CustomRecovery(ControlRecovery))
	router.Use(iotgin.GinLoggerAll())

	//链路追踪
	router.Use(iottrace.TracerWrapper)

	//fixme 服务限流，限流参数maxBurstSize，需要放在配置文件中，根据实际情况进行配置
	router.Use(iotgin.Limiter(100000))

	//fixme 客户端IP限流，注意配置QPS参数
	err := iotgin.SetupIPRateLimiter(100000)
	if err != nil {
		fmt.Println(err)
	}
	router.Use(iotgin.LimitMiddleware())

	//fixme Prometheus，subsystem和自定义参数根据实际需要定义
	var p = iotgin.NewPrometheus("api")
	p.Use(router)

	router.Use(CORS())
	//router.Use(iotgin.GinLoggerAll())

	//业务相关的路由配置
	tpSvc := service.ThirdParty{}
	router.GET("/.well-known/apple-app-site-association", tpSvc.GetThirdLoginJson)
	router.Any("/login", loginHandler)
	router.GET("/regionList", GetRegionList)
	router.Any("/auth", authHandler)

	auth := router.Group("/oauth")
	{
		auth.Any("/authorize", authorizeHandler)
		auth.Any("/token", handleTokenRequest)
	}

	api := router.Group("/api")
	{
		// api.Use(ginserver.HandleTokenVerify())
		// api.GET("/validate", func(c *gin.Context) {
		// 	ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
		// 	if exists {
		// 		c.JSON(http.StatusOK, ti)
		// 		return
		// 	}
		// 	c.String(http.StatusOK, "not found")
		// })
		api.Any("/googleFulfillment", GoogleFulfillment)
		api.Any("/TMiotGateWay", TMiotGateWay) //预留，之后删除
		api.Any("/tianmaoIotGateWay", TMiotGateWay)
		api.Any("/xiaomiIoTGateWay", xiaoaiRouter.GateWay) //xiaomiRouter.GateWay
		//api.Any("/xmIoTGateWay", xiaoaiRouter.GateWayTest) //预留，之后删除
		api.Any("/xiaoaiIoTGateWay", xiaoaiRouter.GateWay)
		api.Any("/xiaoaiIoTGateWayV2", xiaoaiRouter.GateWayV2)
		api.Any("/xiaoaiIoTGateWayTest", xiaoaiRouter.GateWayTestAccount) //测试账号的方式，voice_test_user字典中配置xiaoai
		api.Any("/xiaoaiIoTGateWayTestV2", xiaoaiRouter.GateWayTestAccountV2)
		api.Any("/xiaoaiIoTGateWayDemo", xiaoaiRouter.GateWayTest) //Demo测试，可以使用它进行上线；
		api.Any("/alexaIoTGateWay", alexaRouter.GateWay)
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, User-Agent, Referer, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, lang, appKey, tenantId, x-sys-info")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}

func ControlRecovery(c *gin.Context, err interface{}) {
	if err != nil {
		iotlogger.LogHelper.Errorf("%s,crash:%v", c.Request.RequestURI, err)
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func SetClientInfo(clientId, clientSecret, domain string) error {
	return clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: domain,
	})
}
