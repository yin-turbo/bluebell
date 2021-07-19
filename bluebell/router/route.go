package router

import (
	"bluebell/controller"
	"bluebell/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//用户注册
	r.POST("/signup", controller.SignupHandler)

	//用户登录
	r.POST("/login",controller.LoginHandler)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	return r
}
