package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() (err error) {
	//方式一: 一个API就行(可以使用绝对路径或者相对路径)
	viper.SetConfigFile("./conf/config.yaml") //这个API决定了读取的file文件的路径 与下面的AddConfigPath没有关系

	//方式二: 两个API搭配使用
	// viper.SetConfigName("config") //不需要后缀 是根据文件名找,因此配置文件名不要重名
	//viper.AddConfigPath("./conf/config.yaml") //.表示当前目录下的配置文件
	err = viper.ReadInConfig() //读取配置文件信息

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v\n", err))
	} else {
		fmt.Println("config file loading successfly")
	}

	//将读取配置文件中的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.unmashell failed err:%v\n", err)
	}
	//开启监控
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.unmashell failed err:%v\n", err)
		}
	})

	return
}
