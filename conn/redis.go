package conn

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func InitRedis() error {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "43.138.135.43:6379",
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	_, err := RedisDB.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败")

	} else {
		fmt.Println("redis连接成功!!!")
	}
	return err
}
