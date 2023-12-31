package db


import (
	"context"
	
)

func Insert(collection string, data interface {}) error {
	client, ctx := getConection()
	defer client.Disconnect(ctx)

	c := client.Database("crawler").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)

	return err

}