package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"suzaku/pkg/common/config"
)

var RedisClient *redisClient

type redisClient struct {
	client *redis.Client
	Prefix        string
}

func init() {
	var (
		client *redis.Client
		err    error
	)
	// 创建连接池
	client = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address[0],
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Db,
	})
	// 判断是否能够链接到数据库
	_, err = client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	RedisClient = &redisClient{client,config.Config.Redis.Prefix}
}
