package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	//方式一: 一个API就行(可以使用绝对路径或者相对路径)
	viper.SetConfigFile("./conf/config.yaml") //这个API决定了读取的file文件的路径 与下面的AddConfigPath没有关系

	//方式二: 两个API搭配使用
	// viper.SetConfigName("config") //不需要后缀
	// viper.SetConfigType("yaml") //专门用于从远程获取配置文件信息
	//viper.AddConfigPath("./conf/config.yaml") //.表示当前目录下的配置文件
	err = viper.ReadInConfig() //读取配置文件信息

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v\n", err))
	} else {
		fmt.Println("config file loading successfly")
	}

	//开启监控
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed")
	})

	return
}
