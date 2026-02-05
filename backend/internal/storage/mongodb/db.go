package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New(dsn string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Failed to connect mongodb")
		return nil, err
	}
	fmt.Println("Connected to mongodb")

	return client, nil
}
