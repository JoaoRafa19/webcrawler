package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConection() (client *mongo.Client, ctx context.Context) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	return
}
