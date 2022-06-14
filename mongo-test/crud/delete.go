package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteOne(ctx, query)
	return
}

func deleteMany(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteMany(ctx, query)
	return
}

func DeleteMain() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	query := bson.D{
		{"maths", bson.D{{"$gt", 60}}},
	}

	result, err := deleteOne(client, ctx, "gfg", "marks", query)

	fmt.Println("No.of rows affected by DeleteOne()")
	fmt.Println(result.DeletedCount)

	query = bson.D{
		{"science", bson.D{{"$gt", 0}}},
	}

	result, err = deleteMany(client, ctx, "gfg", "marks", query)

	fmt.Println("No.of rows affected by DeleteMany()")
	fmt.Println(result.DeletedCount)
}
