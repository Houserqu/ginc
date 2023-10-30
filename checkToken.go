package ginc

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := viper.GetString("api_token")

		if token == c.Query("token") || token == c.GetHeader("token") {
			c.Next()
			return
		}

		c.AbortWithStatus(403)
	}
}
