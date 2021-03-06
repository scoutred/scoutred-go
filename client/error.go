package client

import (
	"encoding/json"
	"errors"
)

var ErrNotFound = errors.New("resource not found")

type Error struct {
	Msg string `json:"error"`
	// HTTP Status Code associated with the error
	StatusCode int `json:"statusCode"`
}

//	serialize the error for printing as a string
func (e Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
