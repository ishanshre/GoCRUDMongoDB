package connections

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Defining a interface and its method signatures
type DbInterface interface {
	GetProductCollection() *mongo.Collection
}

// defining a mongo db struct
type DB struct {
	Client *mongo.Client
}

// return product collections from the mongo db
func (db *DB) GetProductCollection() *mongo.Collection {
	return db.Client.Database("myDB").Collection("products")
}

// establishes a connection to mongodb
func ConnectToNoSql(dsn string) (DbInterface, error) {
	client, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client: client,
	}, nil
}

// return a mongo db client
func NewDatabase(dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, err
}
