package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
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

	//初始化redis
	if err := redis.Init(); err != nil {
		zap.L().Error("zap init redis failed, err:%v\n", zap.Error(err))
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	//注册路由
	r := router.Setup(settings.Conf.Mode)
	err := r.Run()
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
