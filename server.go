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

	R = gin.New()

	// 注册中间件
	R.Use(gin.Recovery())
	R.Use(AccessMiddleware())
}

func Run() {
	err := R.Run(viper.GetString("server.addr"))
	if err != nil {
		log.Fatal("Error listen")
	}
}
