package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml") //专门用于从远程获取配置文件信息
	//viper.AddConfigPath(".")    //.表示当前目录下的配置文件
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
