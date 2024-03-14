package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(user, pass string) *mongo.Client {
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// conn := fmt.Sprintf("mongodb+srv://%s:%s@bingo.upejgm6.mongodb.net/?retryWrites=true&w=majority&appName=Bingo", user, pass)
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	return client
}
