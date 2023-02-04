package drivers

import (
	"context"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBConfig types.IMongoDB

type mongodbProvider struct {
	dsn      string
	database string
	mongodb  *mongo.Client
	DB       *mongo.Database
}

func (m *mongodbProvider) Connect() {
	clientOptions := options.Client()
	clientOptions.ApplyURI(m.dsn)

	var err error
	m.mongodb, err = mongo.NewClient(clientOptions)
	if err != nil {
		lets.LogE("MongoDB: %v", err)
		return
	}

	err = m.mongodb.Connect(context.Background())
	if err != nil {
		lets.LogE("MongoDB: %v", err)
		return
	}

	m.DB = m.mongodb.Database(m.database)
}

// Define MySQL service host and port
func MongoDB() {
	if MongoDBConfig == nil {
		return
	}

	lets.LogI("MongoDB Starting ...")

	mongodb := mongodbProvider{
		dsn:      MongoDBConfig.GetDsn(),
		database: MongoDBConfig.GetDatabase(),
	}
	mongodb.Connect()

	// Inject Gorm into repository
	for _, repository := range MongoDBConfig.GetRepositories() {
		repository.SetDriver(mongodb.DB)
	}
}
