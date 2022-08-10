package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Database = "splitsplat"

// Create a document in mongodb
func Create(client *mongo.Client, collection string, data interface{}) error {
	bsonData, err := bson.Marshal(data)

	if err != nil {
		return err
	}

	firstCollection := client.Database(Database).Collection(collection)

	// insert the bson object using InsertOne()
	result, err := firstCollection.InsertOne(context.TODO(), bsonData)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}

	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
	return nil
}
