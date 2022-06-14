package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

func UpdateMany(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateMany(ctx, filter, update)
	return
}

func MainUpdate() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	filter := bson.D{{"maths", bson.D{{"$lt", 100}}}}
	update := bson.D{
		{"$set", bson.D{
			{"maths", 100},
		}},
	}
	result, err := UpdateOne(client, ctx, "gfg",
		"marks", filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("update single document")
	fmt.Println(result.ModifiedCount)

	filter = bson.D{
		{"computer", bson.D{{"$lt", 100}}},
	}
	update = bson.D{
		{"$set", bson.D{
			{"computer", 100},
		}},
	}

	result, err = UpdateMany(client, ctx, "gfg", "marks", filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("update multiple document")
	fmt.Println(result.ModifiedCount)
}
