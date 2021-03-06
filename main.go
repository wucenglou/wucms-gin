package main

import (
	"wucms-gin/core"
	"wucms-gin/global"

	"github.com/gin-gonic/gin"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8010") // listen and serve on 0.0.0.0:8080
}
