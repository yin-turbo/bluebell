package main

import (
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/router"
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed , err: %v\n", err)
		return
	}
	//加载日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	zap.L().Debug("log init success")

	//初始化数据库
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	//注册路由
	router.Setup()

}
