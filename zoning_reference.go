package scoutred

import "time"

type ZoningReference struct {
	ID          *int32     `json:"id"`
	ZoningID    *int32     `json:"zoningId"`
	Description *string    `json:"description"`
	URL         *string    `json:"url"`
	Created     *time.Time `json:"created"`
	Updated     *time.Time `json:"updated"`
}
