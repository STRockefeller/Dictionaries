package testutils

import (
	"io/ioutil"
	"net/http"
	"strings"

	"bou.ke/monkey"
)

func PatchHttpRequest(url string, mockStatusCode int, mockResponseBody string) func() func() {
	return func() func() {
		monkey.Patch(http.Get, func(urlCheck string) (*http.Response, error) {
			// Check if the URL matches the desired URL
			if urlCheck == url {
				// Create a new response with the desired data body and status code
				return &http.Response{
					StatusCode: mockStatusCode,
					Body:       ioutil.NopCloser(strings.NewReader(mockResponseBody)),
				}, nil
			} else {
				// Call the real http.Get function for all other URLs
				return http.DefaultClient.Get(url)
			}
		})
		return func() {
			monkey.Unpatch(http.Get)
		}
	}
}
