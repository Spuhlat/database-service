package utils

import (
	"net/http"
	"net/url"
)

func flattenParamMap(originalMap url.Values) map[string][]string {
	resultMap := make(map[string][]string)

	for key, val := range originalMap {
		if len(val) > 0 {
			resultMap[key] = []string{val[0]}
		} else {
			resultMap[key] = make([]string, 1)
		}
	}
	return resultMap
}

func GetRequestParams(r *http.Request) map[string][]string {
	return flattenParamMap(r.URL.Query())
}
