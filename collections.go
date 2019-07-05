package scoutred

import "time"

type Collection struct {
	// ID is the unique ID of the collection
	ID *int32 `json:"id"`
	// OwnerUserID is the ID of the collection owner
	OwnerUserID *int32 `json:"ownerUserId"`
	// Count indicates the number of parcels in a collection
	Count int32 `json:"count"`
	// Name is the display name of the collection
	Name *string `json:"name"`
	// Created is the timestamp of when the collection was created
	Created *time.Time `json:"created"`
	// Updated is the timestamp of when the collection was last updated
	Updated *time.Time `json:"updated"`
}
