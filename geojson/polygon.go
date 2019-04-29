package geojson

/*
{
	"type": "Polygon",
	"coordinates": [
		[
			[-117.180183060805, 32.6963180459813],
			[-117.180183060805, 32.7305263087481],
			[-117.147086641189, 32.7305263087481],
			[-117.147086641189, 32.6963180459813],
			[-117.180183060805, 32.6963180459813]
		]
	]
}
*/

type Polygon struct {
	CRS         CRS           `json:"crs"`
	BoundingBox []float64     `json:"bbox,omitempty"`
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
