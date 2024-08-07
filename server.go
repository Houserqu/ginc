package ginc

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var R *gin.Engine

func InitServer() {
	InitConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	R = gin.New()

	if viper.GetString("ip_header") != "" {
		R.TrustedPlatform = viper.GetString("ip_header")
	}

	// 注册中间件
	R.Use(gin.Recovery())
	R.Use(AccessMiddleware())
	if viper.GetBool("cors.enable") {
		R.Use(CorsMiddleware())
	}
	if viper.GetString("api_token") != "" {
		R.Use(CheckToken())
	}

	// 监控
	if viper.GetBool("metrics") {
		Metrics.Use(R)
	}
}

func Run() {
	err := R.Run(viper.GetString("server.addr"))
	if err != nil {
		log.Fatal("Error listen")
	}
}
