package routers

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_demo_api_service/config"
	systemrouters "cloud_platform/iot_demo_api_service/controls/system/router"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iottrace"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //生产模式启动

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 2, "msg": "未找到请求路由的处理函数"})
		c.Abort()
	})

	router.Use(iotgin.GinLogger())

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

	var p = iotgin.NewPrometheus("api")
	p.Use(router)

	// 跨域
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")                            //跨域
		ctx.Header("Access-Control-Allow-Headers", "token,Content-Type")          //必须的请求头
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS,POST,GET,PUT,DELETE") //接收的请求方法
	})

	router.Use(gin.CustomRecovery(ControlRecovery))

	//业务相关的路由配置
	systemrouters.RegisterRouter(router)
	return router
}

func LoadHtmlRouter(router *gin.Engine) {
	router.LoadHTMLGlob(strings.Join([]string{iotconst.GetTemplatesDir(), "*"}, string(filepath.Separator)))
}

func ControlRecovery(c *gin.Context, err interface{}) {
	if err != nil {
		iotlogger.LogHelper.Errorf("%s,crash:%v", c.Request.RequestURI, err)
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
