package scoutred

import (
	"time"

	"github.com/scoutred/app/geojson"
)

type Overlay struct {
	ID           *int32           `json:"id"`
	Name         *string          `json:"name"`
	SourceID     *int32           `json:"sourceId"`
	LayerID      *int32           `json:"layerId"`
	Jurisdiction *Jurisdiction    `json:"jurisdiction"`
	Description  *string          `json:"description"`
	Details      []OverlayDetail  `json:"details"`
	Bounds       geojson.Polygon  `json:"bounds"`
	Geohash      *string          `json:"geohash"`
	References   []LayerReference `json:"references"`
	Created      *time.Time       `json:"created"`
	Updated      *time.Time       `json:"updated"`
}

type OverlayDetail struct {
	Name  *string     `json:"name"`
	Value interface{} `json:"value"`
}
