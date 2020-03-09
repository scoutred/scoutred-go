package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	scoutred "github.com/scoutred/scoutred-go"
)

// Auth allows the user to login with their credentials
// and have the system return back an API token
func (c *Client) Auth(email, password string) (string, error) {
	path := c.ApiUrl + "/login"

	// fill in the struct details
	auth := scoutred.Auth{
		Email:    email,
		Password: password,
	}

	// setup buffer to store encoded JSON
	var buf bytes.Buffer

	// encode JSON
	if err := json.NewEncoder(&buf).Encode(auth); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", path, &buf)
	if err != nil {
		return "", err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")

	//	make our request to the API
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var token scoutred.Token

	// TODO (arolek): expand out response code handling
	switch res.StatusCode {
	case http.StatusOK:
		//	parse our response JSON into the provided struct
		if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
			return "", err
		}

	default: // currently only error handling
		var apiError Error
		if err = json.NewDecoder(res.Body).Decode(&apiError); err != nil {
			//	TODO (arolek): this needs to be of type scoutred.Error
			return "", err
		}

		apiError.StatusCode = res.StatusCode

		return "", &apiError
	}

	return token.Token, nil
}
