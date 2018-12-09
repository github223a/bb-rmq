package src

import (
	"fmt"
	"log"

	"./entities"
	"github.com/go-redis/redis"
)

func RedisInit() {
	url := fmt.Sprintf("%s:%d", CONFIG.Redis.Host, CONFIG.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: CONFIG.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	FailOnError(err, "Error on ping redis.", "redist")

	entities.Redis = entities.NewRedisEntity(client)
	log.Printf(HEADER_REDIS_MESSAGE+"Redis is starting by url %s", url)
	// Output: PONG <nil>
}
