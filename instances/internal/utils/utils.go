package utils

import (
	"encoding/json"
	"net/http"
)

func ParseResponse[R any](resp *http.Response) (R, error) {
	var result R
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
