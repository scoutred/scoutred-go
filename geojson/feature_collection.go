//	package provides go structs for GeoJSON data
package geojson

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string       `json:"type"`
	Geometry   MultiPolygon `json:"geometry"`
	Properties struct {
		ID  int64   `json:"id"`
		APN *string `json:"apn"`
	} `json:"properties"`
}
