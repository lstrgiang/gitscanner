package redis

import (
	redis "github.com/go-redis/redis/v7"
	"github.com/hibiken/asynq"
)

type Redis struct {
	client *redis.Client
}

func (r Redis) GetClient() *redis.Client {
	return r.client
}

func CreateRedisClient(location string, useCluster bool) (*Redis, error) {
	options, err := redis.ParseURL(location)
	if err != nil {
		return nil, err
	}

	options.Username = ""
	client := redis.NewClient(options)
	return &Redis{client: client}, nil

}
func NewAsyncRedisClient(location string) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr: location,
	}
}

func NewAsynqClient(redisLocation string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisLocation, // Redis server address
	})
}
