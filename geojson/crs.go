package geojson

//	coordinate reference system
type CRS struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"name"`
	} `json:"properties"`
}

func NewCRS(proj string) *CRS {
	crs := CRS{
		Type: "Name",
	}

	crs.Properties.Name = proj

	return &crs
}
