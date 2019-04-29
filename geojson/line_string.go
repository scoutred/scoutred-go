package geojson

import "encoding/json"

type LineString struct {
	CRS         CRS         `json:"crs"`
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

//	NewLineString instantiates a new 4326 point with the provided lat lon
func NewLineString(coords [][]float64) *LineString {
	return &LineString{
		CRS:         *NewCRS("EPSG:4326"),
		Type:        "LineString",
		Coordinates: coords,
	}
}

func (p LineString) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(b)
}
