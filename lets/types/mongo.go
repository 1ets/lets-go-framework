package types

import (
	"fmt"
)

const (
	MONGODB_DSN      = "mongodb://localhost:27017"
	MONGODB_DATABASE = "default"
)

type IMongoDB interface {
	GetDsn() string
	GetDatabase() string
	GetRepositories() []IMongoDBRepository
}

type MongoDB struct {
	Dsn          string
	Database     string
	Repositories []IMongoDBRepository
}

func (r *MongoDB) GetDsn() string {
	if r.Dsn == "0" {
		fmt.Println("Configs MongoDB: MONGODB_DSN is not set in .env file, using default configuration.")
		return MONGODB_DATABASE
	}
	return r.Dsn
}

func (r *MongoDB) GetDatabase() string {
	if r.Database == "0" {
		fmt.Println("Configs MongoDB: MONGODB_DATABASE is not set in .env file, using default configuration.")
		return MONGODB_DATABASE
	}
	return r.Database
}

func (r *MongoDB) GetRepositories() []IMongoDBRepository {
	return r.Repositories
}
