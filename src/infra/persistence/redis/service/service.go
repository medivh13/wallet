package article

import (
	"time"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type ServRedis struct {
	Rdb *redis.Client
}

type ServRedisInt interface {
	SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	GetData(ctx context.Context, key string) (string, error)
}

func NewServRedis(rdb *redis.Client) *ServRedis {
	return &ServRedis{
		Rdb: rdb,
	}
}

func (p *ServRedis) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	rds := p.Rdb.Set(context.Background(), key, value, ttl)

	return rds.Err()
}

func (p *ServRedis) GetData(ctx context.Context, key string) (string, error) {
	dataRedis, err := p.Rdb.Get(context.Background(), key).Result()

	return dataRedis, err
}
