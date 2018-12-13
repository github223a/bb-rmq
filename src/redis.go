package src

func RedisInit() {
	// url := fmt.Sprintf("%s:%d", GetConfig().Redis.Host, GetConfig().Redis.Port)
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     url,
	// 	Password: GetConfig().Redis.Password,
	// 	DB:       0,
	// })

	// _, err := client.Ping().Result()
	// core.FailOnError(core, err, "Error on ping redis.", "redist")

	// entities.Redis = entities.NewRedisEntity(client)
	// log.Printf(HEADER_REDIS_MESSAGE+"Redis is starting by url %s", url)
	// Output: PONG <nil>
}
