package scoutred

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiUrl = "http://192.168.99.100:8080/api"
	//	apiUrl = "https://dev.scoutred.com/api"
)

// Backend is an interface for making calls against a ScoutRED service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, body io.Reader, v interface{}) error
}

var httpClient = &http.Client{}

// Call is the Backend.Call implementation for invoking ScoutRED APIs.
func Call(method, path, key string, body io.Reader, v interface{}) (err error) {

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
	req.Header.Add("Authorization", "BEARER "+key)

	//	check for request payload
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	//	make our request to the API
	res, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	//	read our response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	//	TODO: expand out response code handling
	if res.StatusCode != 200 {
		var apiError Error
		if err = json.Unmarshal(resBody, &apiError); err != nil {
			return err
		}

		apiError.StatusCode = res.StatusCode

		return &apiError
	}

	//	parse our response JSON into the provided struct
	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return
}
