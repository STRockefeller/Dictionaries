package utils

import (
	"encoding/json"
	"net/http"
)

// ParseResponseBody is a function that takes in an HTTP response and returns a decoded result of type R and an error (if any).
// The generic type R enables parsing response bodies into different types by the caller.
func ParseResponseBody[R any](resp *http.Response) (R, error) {
	var result R
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
