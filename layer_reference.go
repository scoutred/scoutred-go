package scoutred

import "time"

type LayerReference struct {
	ID          *int32     `json:"id"`
	SourceID    *int32     `json:"sourceId"`
	LayerID     *int32     `json:"layerId"`
	Description *string    `json:"description"`
	URL         *string    `json:"url"`
	Created     *time.Time `json:"created"`
	Updated     *time.Time `json:"updated"`
}
