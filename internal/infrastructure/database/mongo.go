package database

import (
	"github.com/arfanxn/welding/internal/infrastructure/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// NewMongoClientFromConfig creates a new MongoDB client instance using the provided configuration.
// It initializes the client with the MongoURI from the config and establishes a connection.
func NewMongoClientFromConfig(cfg *config.Config) (*mongo.Client, error) {
	mongoURI := cfg.MongoURI

	// Set client options using the URI from config
	clientOpts := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewMongoDatabaseFromClientAndConfig returns a *mongo.Database for the configured DB name.
func NewMongoDatabaseFromClientAndConfig(client *mongo.Client, cfg *config.Config) (*mongo.Database, error) {
	mongoDatabase := cfg.MongoDB
	return client.Database(mongoDatabase), nil
}
