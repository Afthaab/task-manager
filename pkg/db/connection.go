package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() (*mongo.Database, error) {
	dbPassword := os.Getenv("DB_PASSWORD")

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://mongodb:%s@cluster1.kfqqheq.mongodb.net/?retryWrites=true&w=majority", dbPassword)).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping the MongoDB: %v", err)
	}
	databseConn := client.Database("taskmanager")

	return databseConn, nil

}
