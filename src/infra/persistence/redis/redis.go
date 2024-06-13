package redis

import (
	"context"
	"wallet/src/infra/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewRedisClient(conf config.RedisConf, logger *logrus.Logger) (*redis.Client, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(ctx).Result()

	if err != nil {
		logger.Printf("cant connect Redis :  %s", err)
	}

	return client, nil
}
