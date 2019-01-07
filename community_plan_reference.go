package scoutred

import "time"

type CommunityPlanReference struct {
	ID              *int32     `json:"id"`
	CommunityPlanID *int32     `json:"-"`
	Description     *string    `json:"description"`
	URL             *string    `json:"url"`
	Created         *time.Time `json:"created"`
	Updated         *time.Time `json:"updated"`
}
