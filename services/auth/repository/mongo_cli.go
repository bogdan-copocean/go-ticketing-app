package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongo() AuthRepository {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// client, connErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://auth-mongo-srv:27017"))
	client, connErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:30090"))
	if connErr != nil {
		fmt.Printf("[connErr]: %v\n", connErr.Error())
		panic(connErr)
	}

	if pingErr := client.Ping(ctx, readpref.Primary()); pingErr != nil {
		fmt.Printf("[pingErr]: %v\n", pingErr)
		panic(pingErr)
	}

	collection := client.Database("auth").Collection("auth")

	return NewAuthRepository(client, collection)

}
