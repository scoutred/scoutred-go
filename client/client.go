package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func New(key string) *Client {
	return &Client{
		key,
		&http.Client{},
	}
}

// Client has various methods for fetching resources from Scoutred's API
type Client struct {
	// Key is the API key for the instance of the client
	Key string

	// embed an http.Client so we have access to all it's methods
	*http.Client
}

const (
	apiUrl = "https://scoutred.com/api"
)

// Caller is an interface for making calls against a Scoutred service.
// This interface exists to enable mocking during testing if needed.
type Caller interface {
	Call(method, path string, body io.Reader, v interface{}) error
}

// Call is the implementation for invoking Scoutred APIs.
func (c *Client) Call(method, path string, body io.Reader, v interface{}) (err error) {
	//	check to make sure our API endpoint starts with "/"
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	//	build the full endpoint
	path = apiUrl + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return
	}

	//	auth header required for request
	req.Header.Add("Authorization", fmt.Sprintf("BEARER %v", c.Key))

	//	check for request payload
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	//	make our request to the API
	res, err := c.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// TODO (arolek): expand out response code handling
	switch res.StatusCode {
	case http.StatusOK:
		//	parse our response JSON into the provided struct
		if v != nil {
			return json.NewDecoder(res.Body).Decode(v)
		}
	default: // currently only error handling
		var apiError Error
		if err = json.NewDecoder(res.Body).Decode(&apiError); err != nil {
			//	TODO: this needs to be of type scoutred.Error
			return err
		}

		apiError.StatusCode = res.StatusCode

		return &apiError
	}

	return
}
