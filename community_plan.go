package scoutred

import (
	"time"

	"github.com/scoutred/scoutred-go/geojson"
)

type CommunityPlan struct {
	ID          *int32                   `json:"id"`
	Name        *string                  `json:"name"`
	Description *string                  `json:"description"`
	References  []CommunityPlanReference `json:"references"`
	Bounds      geojson.Polygon          `json:"bounds"`
	Geohash     *string                  `json:"geohash"`
	Created     *time.Time               `json:"created"`
	Updated     *time.Time               `json:"updated"`
}
