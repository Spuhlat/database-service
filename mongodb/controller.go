package mongodb

import (
	"context"
	"main/utils"
	"net/http"
)

type CreateBody struct {
	Database   string
	Collection string
	Data       interface{}
}

func Controller(req *http.Request) error {
	client := Connect()

	if req.Method == "POST" {
		reqBody := utils.GetRequestBody[CreateBody](req.Body)
		collection := reqBody.Collection
		data := reqBody.Data

		if err := Create(client, collection, data); err != nil {
			return err
		}
	}

	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
