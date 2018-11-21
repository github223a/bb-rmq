package entities

import (
	"github.com/go-redis/redis"
)

var Redis *RedisStruct

type RedisStruct struct {
	IsAlive bool
	Client *redis.Client
}

func NewRedisEntity(client *redis.Client) *RedisStruct {
	return &RedisStruct {
		IsAlive: true,
		Client: client,
	}
}