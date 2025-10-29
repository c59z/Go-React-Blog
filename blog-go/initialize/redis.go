package initialize

import (
	"blog-go/global"
	"os"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func ConnectRedis() redis.Client {
	redisConfig := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		global.Log.Error("Failed to connect to Redis:", zap.Error(err))
		os.Exit(-1)
	}

	return *client
}
