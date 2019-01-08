package client

import (
	"encoding/json"
)

type Error struct {
	Msg        string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

//	serialize the error for printing as a string
func (this *Error) Error() string {
	ret, _ := json.Marshal(this)
	return string(ret)
}
