package scoutred

type MultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
	CRS         CRS             `json:"crs"`
}

//	coordinate reference system
type CRS struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"name"`
	} `json:"properties"`
}
