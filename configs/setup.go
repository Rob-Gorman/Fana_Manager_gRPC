package configs

import (
	"fmt"
	"sovereign/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectClient() *mongo.Client {
	uri := DBURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	utils.HandleErr(err, "Error initializing Mongo client.")

	ctx, _ := utils.StandardContext()
	// defer cancel()

	err = client.Connect(ctx)
	utils.HandleErr(err, "Error connecting to Mongo client after initialization")

	err = client.Ping(ctx, nil)
	utils.HandleErr(err, "Error pinging Mongo Client")

	fmt.Println("Connected to MongoDB")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := DBName()
	db := client.Database(dbName)
	return db.Collection(collectionName)
}
