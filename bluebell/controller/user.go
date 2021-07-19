package controller

import (
	//"bluebell/logic"
	"bluebell/logic"
	"bluebell/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignupHandler(c *gin.Context) {
	//1. 获取参数和参数校验  绑定参数
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//绑定参数有误
		zap.L().Error("Sign up with fause valid", zap.Error(err))

		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	zap.L().Info("Sign up success")

	//2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("Sign up failed, err:%v\n", zap.Error(err))
		return
	}

	return

	//3.返回响应
	//ResponseSuccess(c, nil)
}

//用户登录逻辑
func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验  new关键字对结构体实例化,得到的是结构体的地址
	p := new(model.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login Failed to bind JSON ", zap.Error(err))
	}
	//登录有误
	if err := logic.Login(p); err != nil {
		zap.L().Error("用户登录有误", zap.Error(err))
		//返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})

		return
	}

	//登录成功 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
