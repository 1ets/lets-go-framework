package types

import "go.mongodb.org/mongo-driver/mongo"

type IMongoDBRepository interface {
	SetDriver(*mongo.Database)
}
