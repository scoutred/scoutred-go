package geojson

import "encoding/json"

type Point struct {
	CRS         CRS       `json:"crs"`
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

//	NewPoint instantiates a new 4326 point with the provided lat lon
func NewPoint(lat, lon float64) *Point {
	return &Point{
		CRS:         *NewCRS("EPSG:4326"),
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}
}

func (p Point) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(b)
}
