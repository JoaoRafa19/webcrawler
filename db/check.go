package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func VisitedLink(link string) bool {
	client, ctx := getConection()
	defer client.Disconnect(ctx)
	c := client.Database("crawler").Collection("links")
	opts := options.Count().SetLimit(1)
	n, err := c.CountDocuments(
		context.TODO(), 
		bson.D{{Key: "link", Value: link}},
		opts,
	)

	if err != nil {
		panic(err)
	}

	return n > 0

}
