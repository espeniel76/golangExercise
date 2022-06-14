package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// deleteOne is a user defined function that delete,
// a single document from the collection.
// Returns DeleteResult and an error if any.
func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}

// deleteMany is a user defined function that delete,
// multiple documents from the collection.
// Returns DeleteResult and an error if any.
func deleteMany(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match documents from the collection.
	result, err = collection.DeleteMany(ctx, query)
	return
}

func DeleteMain() {

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	// free resource when main function is returned
	defer close(client, ctx, cancel)

	// This query delete document when the maths
	// field is greater than 60
	query := bson.D{
		{"maths", bson.D{{"$gt", 60}}},
	}

	// Returns result of deletion and error
	result, err := deleteOne(client, ctx, "gfg", "marks", query)

	// print the count of affected documents
	fmt.Println("No.of rows affected by DeleteOne()")
	fmt.Println(result.DeletedCount)

	// This query deletes documents that has
	// science field greater that 0
	query = bson.D{
		{"science", bson.D{{"$gt", 0}}},
	}

	// Returns result of deletion and error
	result, err = deleteMany(client, ctx, "gfg", "marks", query)

	// print the count of affected documents
	fmt.Println("No.of rows affected by DeleteMany()")
	fmt.Println(result.DeletedCount)
}
