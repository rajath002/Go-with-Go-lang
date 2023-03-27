package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/rajath002/gomongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:Welcome@demo@user@cluster-demo-project.hahillw.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"

const collectionName = "watchlist"

// !MOST IMPORTANT
var collection *mongo.Collection

// Connect with MongoDB
func init() {
	// Client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance is ready

	fmt.Println("Collection instance is ready")
}

// MONGODB  helpers - file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in the db with id: ", inserted.InsertedID)
}
