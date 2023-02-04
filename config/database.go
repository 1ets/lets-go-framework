package config

import (
	"lets-go-framework/app/repository"
	"lets-go-framework/lets/drivers"
	"lets-go-framework/lets/types"
	"os"
)

func Database() {
	drivers.MySQLConfig = &types.MySQL{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		Charset:   "utf8mb4",
		ParseTime: "True",
		Loc:       "Local",
		Debug:     true,
		Repositories: []types.IMySQLRepository{
			repository.User,
		},
		EnableMigration: true,
	}

	drivers.RedisConfig = &types.Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
		Debug:    false,
		Repositories: []types.IRedisRepository{
			repository.RedisExample,
		},
	}
}
