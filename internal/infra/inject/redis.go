package inject

import (
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/redis"
)

const (
	keyRedisClient     = "inject_redis_client"
	keyRedisAsynClient = "inject_redis_asyn_client"
)

func SetRedisClient(e echo.Context, client *redis.Redis) {
	e.Set(keyRedisClient, client)
}

func RedisClient(e echo.Context) *redis.Redis {
	val := e.Get(keyRedisClient)
	if val == nil {
		e.Logger().Panic("nil db")
	}

	return val.(*redis.Redis)
}

func SetAsynqClient(e echo.Context, client *asynq.Client) {
	e.Set(keyRedisAsynClient, client)
}

func AsynqClient(e echo.Context) *asynq.Client {
	val := e.Get(keyRedisAsynClient)
	if val == nil {
		e.Logger().Panic("nil db")
	}

	return val.(*asynq.Client)
}
