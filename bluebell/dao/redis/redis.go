package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	//"google.golang.org/genproto/googleapis/cloud/redis/v1"
)

var client *redis.Client

//连接redis
func Init() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poll_size"),
	})

	_, err = client.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("connect redis err:%v\n", err))
	} else {
		zap.L().Info("zap connect redis successful")
		fmt.Println("connect redis success")

	}

	return nil

}

func Close() {
	_ = client.Close()
}
