package scoutred

import "time"

type ContractorTrade struct {
	ID      *int32     `json:"id"`
	Name    *string    `json:"name"`
	Updated *time.Time `json:"updated"`
	Created *time.Time `json:"created"`
}
