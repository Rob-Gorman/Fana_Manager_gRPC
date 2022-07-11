package data

// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/mongo#Database.CreateCollection

import (
	"context"
	"sovereign/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() (results []bson.M) {
	uri := utils.DBURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	utils.HandleErr(err, "Error initializing Mongo client.")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	utils.HandleErr(err, "Error connecting to Mongo client after initialization")
	defer client.Disconnect(ctx)

	db := client.Database(utils.DBName())
	flags := db.Collection("flags")
	query, err := flags.Find(ctx, bson.D{})
	utils.HandleErr(err, "Cannot query the collection")

	query.All(ctx, &results)
	return results
}
