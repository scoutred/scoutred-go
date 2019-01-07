package scoutred

import (
	"time"

	"github.com/scoutred/app/geojson"
)

type Zoning struct {
	ID           *int32             `json:"id"`
	Jurisdiction *Jurisdiction      `json:"jurisdiction"`
	Designation  *string            `json:"designation"`
	Description  *string            `json:"description"`
	Bounds       geojson.Polygon    `json:"bounds"`
	Geohash      *string            `json:"geohash"`
	Regulations  []ZoningRegulation `json:"regulations"`
	References   []ZoningReference  `json:"references"`
	Created      *time.Time         `json:"created"`
	Updated      *time.Time         `json:"updated"`
}
