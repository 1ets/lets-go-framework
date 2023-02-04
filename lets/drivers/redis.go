package drivers

import (
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"

	"github.com/go-redis/redis/v8"
)

var RedisConfig types.IRedis

type redisProvider struct {
	dsn      string
	password string
	database int
	redis    *redis.Client
}

func (m *redisProvider) Connect() {
	m.redis = redis.NewClient(&redis.Options{
		Addr:     m.dsn,
		Password: m.password,
		DB:       m.database,
	})
}

// Define MySQL service host and port
func Redis() {
	if RedisConfig == nil {
		return
	}

	lets.LogI("MySQL Starting ...")

	redis := redisProvider{
		dsn:      RedisConfig.GetDsn(),
		password: RedisConfig.GetPassword(),
		database: RedisConfig.GetDatabase(),
	}
	redis.Connect()

	// Inject Gorm into repository
	for _, repository := range RedisConfig.GetRepositories() {
		repository.SetDriver(redis.redis)
	}
}
