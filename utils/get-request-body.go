package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func GetRequestBody[T any](body io.ReadCloser) T {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}

	var message T
	json.Unmarshal([]byte(bodyBytes), &message)

	return message
}
