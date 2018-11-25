package src

import (
	"./constants"
	"./entities"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func RedisInit() {
	url := fmt.Sprintf("%s:%d", constants.CONFIG.Redis.Host, constants.CONFIG.Redis.Port)
	client := redis.NewClient(&redis.Options {
		Addr: url,
		Password: constants.CONFIG.Redis.Password,
		DB: 0,
	})

	_, err := client.Ping().Result()
	FailOnError(err, "Error on ping redis.", "redist")

	entities.Redis = entities.NewRedisEntity(client)
	log.Printf(constants.HEADER_REDIS_MESSAGE + "Redis is starting by url %s", url)
	// Output: PONG <nil>
}