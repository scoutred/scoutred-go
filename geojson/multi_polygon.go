package geojson

type MultiPolygon struct {
	Type        string          `json:"type"`
	BoundingBox []float64       `json:"bbox,omitempty"`
	Coordinates [][][][]float64 `json:"coordinates"`
	CRS         CRS             `json:"crs,omitempty"`
}
