package mongodb

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Read a document in mongodb
func ReadById(client *mongo.Client, collection string, id int) (error, []byte) {
	collectionInstance := client.Database(Database).Collection(collection)

	var result primitive.D
	filter := bson.D{{Key: "id", Value: id}}
	err := collectionInstance.FindOne(context.TODO(), filter).Decode(&result)

	// check for errors in the read
	if err != nil {
		panic(err)
	}

	jsonStr, err := json.Marshal(result.Map())

	if err != nil {
		return err, nil
	}

	return nil, jsonStr
}
