package mongodb

import (
	"context"
	"main/utils"
	"net/http"
	"strconv"
)

type CreateBody struct {
	Collection string
	Data       interface{}
}

type ReadBody struct {
	Collection string
}

func Controller(req *http.Request) (error, []byte) {
	client := Connect()

	if req.Method == "POST" {
		reqBody := utils.GetRequestBody[CreateBody](req.Body)

		collection := reqBody.Collection
		data := reqBody.Data

		if err := Create(client, collection, data); err != nil {
			return err, nil
		}

		return nil, nil
	}

	if req.Method == "GET" {
		values := utils.GetRequestParams(req)
		id, err := strconv.Atoi(values["id"][0])

		if err != nil {
			return err, nil
		}

		collection := values["collection"][0]

		err, data := ReadById(client, collection, id)
		if err != nil {
			return err, nil
		}

		return nil, data

	}

	if err := client.Disconnect(context.TODO()); err != nil {
		return err, nil
	}
	return nil, nil
}
