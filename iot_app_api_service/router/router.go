package routers

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_common/iotlogger"
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iottrace"

	commonRouters "cloud_platform/iot_app_api_service/controls/common/router"
	communityRouters "cloud_platform/iot_app_api_service/controls/community/router"
	devRouters "cloud_platform/iot_app_api_service/controls/dev/router"
	documentRouters "cloud_platform/iot_app_api_service/controls/document/router"
	intelligenceRouters "cloud_platform/iot_app_api_service/controls/intelligence/router"
	messageRouters "cloud_platform/iot_app_api_service/controls/message/router"
	productRouters "cloud_platform/iot_app_api_service/controls/product/router"
	appUpgradeRouters "cloud_platform/iot_app_api_service/controls/upgrade/router"
	userRouters "cloud_platform/iot_app_api_service/controls/user/router"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //生产模式启动

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		url := c.Request.URL.Path
		host := c.Request.Host

		c.JSON(404, gin.H{"code": 2, "msg": "未找到请求路由的处理函数" + url + " ," + host})
		c.Abort()
	})

	router.Use(iotgin.GinLogger())

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//链路追踪
	router.Use(iottrace.TracerWrapper)

	//服务限流，限流参数maxBurstSize
	if config.Global.Service.Httpqps < 10000 || config.Global.Service.Httpqps > 100000 {
		config.Global.Service.Httpqps = 10000
	}
	router.Use(iotgin.Limiter(config.Global.Service.Httpqps))

	//客户端IP限流，注意配置QPS参数
	if config.Global.Service.IPLimitRequest > 10000 {
		config.Global.Service.IPLimitRequest = 10000
	}
	if config.Global.Service.IPLimitRequest < 5 {
		config.Global.Service.IPLimitRequest = 5
	}
	err := iotgin.SetupIPRateLimiter(config.Global.Service.IPLimitRequest)
	if err != nil {
		fmt.Println(err)
	}
	router.Use(iotgin.LimitMiddleware())

	//fixme Prometheus，subsystem和自定义参数根据实际需要定义
	var p = iotgin.NewPrometheus("api")
	p.Use(router)

	// 跨域
	//router.Use(func(ctx *gin.Context) {
	//	ctx.Header("Access-Control-Allow-Origin", "*")                                                        //跨域
	//	ctx.Header("Access-Control-Allow-Headers", "token,Content-Type,appKey,tenantId,region,tz,x-sys-info") //必须的请求头
	//	ctx.Header("Access-Control-Allow-Methods", "OPTIONS,POST,GET")                                        //接收的请求方法
	//})
	router.Use(CORS())

	router.Use(gin.CustomRecovery(ControlRecovery))

	//业务相关的路由配置
	commonRouters.RegisterRouter(router)
	userRouters.RegisterRouter(router)
	productRouters.RegisterRouter(router)
	documentRouters.RegisterRouter(router)
	messageRouters.RegisterRouter(router)
	appUpgradeRouters.RegisterRouter(router)
	intelligenceRouters.RegisterRouter(router)
	devRouters.RegisterRouter(router)
	communityRouters.RegisterRouter(router)

	controls.RefreshProductCache()
	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, User-Agent, Referer, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, lang, appKey, tenantId, region, tz, x-sys-info, appPushId")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
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
