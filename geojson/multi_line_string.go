package geojson

import "encoding/json"

type MultiLineString struct {
	CRS         CRS           `json:"crs"`
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

//	NewLineString instantiates a new 4326 point with the provided lat lon
func NewMultiLineString(coords [][][]float64) *MultiLineString {
	return &MultiLineString{
		CRS:         *NewCRS("EPSG:4326"),
		Type:        "MultiLineString",
		Coordinates: coords,
	}
}

func (p MultiLineString) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(b)
}
