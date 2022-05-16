package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Db *mongo.Database
var EventsCollection *mongo.Collection
var UsersCollection *mongo.Collection

func Init() func() {
	// Establish mongodb connection
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		panic(err)
	}
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// Define mongodb database + collections
	Db = Client.Database("schej-it")
	EventsCollection = Db.Collection("events")
	UsersCollection = Db.Collection("users")

	// Return a function to close the connection
	return func() {
		Client.Disconnect(ctx)
	}
}
