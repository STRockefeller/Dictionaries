package apicall

import (
	"errors"
	"net/http"
)

type apiCall func() (*http.Response, error)

// This function takes a URL string as input and returns an anonymous function that
// makes an HTTP GET request to that URL when called.
func Get(url string) apiCall {
	return func() (*http.Response, error) {
		return http.Get(url)
	}
}

// This function takes an apiCall function and a parseResponse function as input, and uses the provided
// apiCall to send an HTTP request. It checks for errors, bad status codes, and returns the parsed response
// as type R (a generic type).
func Call[R any](api apiCall, parseResponse func(*http.Response) (R, error)) (R, error) {
	resp, err := api()
	if err != nil {
		return *new(R), err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return *new(R), errors.New("bad status:" + resp.Status)
	}
	return parseResponse(resp)
}
