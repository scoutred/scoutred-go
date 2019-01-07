package scoutred

import "time"

type Jurisdiction struct {
	ID      *int32     `json:"id"`
	Name    *string    `json:"name"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}
