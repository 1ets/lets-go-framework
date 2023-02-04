package types

import (
	"github.com/go-redis/redis/v8"
)

type IRedisRepository interface {
	SetDriver(*redis.Client)
}
