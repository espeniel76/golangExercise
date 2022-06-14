package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func MainConnect() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017?maxPoolSize=20")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	ping(client, ctx)
}
