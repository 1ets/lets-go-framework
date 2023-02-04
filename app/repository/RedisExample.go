package repository

import (
	"context"
	"encoding/json"
	"lets-go-framework/app/models"
	"lets-go-framework/lets"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define repository.
var RedisExample = &redisExample{}

// Repository user.
type redisExample struct {
	db *redis.Client
}

// Implement types.IMySQLRepository.
// Mandatory.
func (tbl *redisExample) SetDriver(db *redis.Client) {
	tbl.db = db
}

// Save user data.
func (tbl *redisExample) SaveUser(user *models.User) {
	ttl := time.Duration(3000) * time.Second
	tx := tbl.db.Set(context.Background(), "user", user, ttl)
	if err := tx.Err(); err != nil {
		lets.LogE("unable to SET data. error: %v", err)
		return
	}
	lets.LogI("set operation success")
}

// Get user data.
func (tbl *redisExample) GetUser() (user models.User) {
	tx := tbl.db.Get(context.Background(), "user")
	if err := tx.Err(); err != nil {
		lets.LogE("unable to GET data. error: %v", err)
		return
	}

	result, err := tx.Result()
	if err != nil {
		lets.LogE("unable to GET data. error: %v", result)
		return
	}

	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		lets.LogE("Bind: %s", err.Error())
	}

	return
}
