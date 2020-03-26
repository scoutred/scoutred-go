package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func New(key string) *Client {
	return &Client{
		key,
		DefaultApiUrl,
		&http.Client{},
	}
}

// Client has various methods for fetching resources from Scoutred's API
type Client struct {
	// Key is the API key for the instance of the client
	Key string
	// ApiUrl is the url of the api server make requests to. This is
	// useful for testing against local or dev servers
	ApiUrl string
	// embed an http.Client so we have access to all it's methods
	*http.Client
}

const (
	DefaultApiUrl = "https://scoutred.com/api"
)

// Caller is an interface for making calls against a Scoutred service.
// This interface exists to enable mocking during testing if needed.
type Caller interface {
	Call(method, path string, body io.Reader, v interface{}) error
}

// Call is the implementation for invoking Scoutred APIs. On success
// the response is marshalled into the the provided interface{}
func (c *Client) Call(method, path string, body io.Reader, v interface{}) error {
	// check to make sure our API endpoint starts with "/"
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// build the full endpoint
	path = c.ApiUrl + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	// auth header required for request
	req.Header.Add("Authorization", fmt.Sprintf("BEARER %v", c.Key))

	// check for request payload
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// make our request to the API
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// TODO (arolek): expand out response code handling
	switch res.StatusCode {
	case http.StatusOK:
		// parse our response JSON into the provided struct
		if v != nil {
			return json.NewDecoder(res.Body).Decode(v)
		}

	case http.StatusNotFound:
		return ErrNotFound

	default:
		var apiError Error

		byt, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(byt, &apiError); err != nil {
			apiError.Msg = string(byt)
		}

		apiError.StatusCode = res.StatusCode

		return apiError
	}

	return nil
}
