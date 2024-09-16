package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://theshubhamdev:mongopswd@clovercluster.5tnqf.mongodb.net/"
const dbName = "clovermovie"
const colName = "watchList"

var Collections *mongo.Collection

func init() {

	opts := options.Client().ApplyURI(connectionString)
	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	Collections = client.Database(dbName).Collection(colName)
	// Inserting a sample document to trigger collection creation
	sampleData := bson.D{{Key: "movie", Value: "DDLJ"}, {Key: "watched", Value: false}}

	_, err = Collections.InsertOne(ctx, sampleData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You successfully connected to MongoDB!")
}
