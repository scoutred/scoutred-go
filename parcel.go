package scoutred

import (
	"time"

	"github.com/scoutred/app/geojson"
)

type Parcel struct {
	ID             *int64  `json:"id"`
	StateFIPSCode  *string `json:"stateFIPS"`
	CountyFIPSCode *string `json:"countyFIPS"`
	//	assessor parcel number
	APN     *string  `json:"apn"`
	Address *Address `json:"address"`
	Owner   struct {
		Name1    *string `json:"name1"`
		Name2    *string `json:"name2"`
		Name3    *string `json:"name3"`
		Address1 *string `json:"address1"`
		Address2 *string `json:"address2"`
		Address3 *string `json:"address3"`
		Address4 *string `json:"address4"`
		Postal   *string `json:"postal"`
	} `json:"owner"`
	LegalDescription *string `json:"legalDescription"`
	Subdivision      struct {
		Map  *string `json:"name"`
		Name *string `json:"map"`
	} `json:"subdivision"`
	Assessor struct {
		Land         *int32 `json:"land"`
		Improvements *int32 `json:"improvements"`
	} `json:"assessor"`
	Structure struct {
		ConstructionYear *int32   `json:"effectiveYearBuilt"`
		SFLiving         *int32   `json:"livingSF"`
		SFUsable         *int32   `json:"usableSF"`
		Units            *int32   `json:"units"`
		Bedrooms         *int32   `json:"bedrooms"`
		Bathrooms        *float64 `json:"bathrooms"`
	} `json:"structure"`
	//	the area of the geometry representation of the parcel
	//	this can be used for lot size, but is not the recorded lot size
	//	but rather the area of the geometry in the projection it was created in
	GeomArea *float64 `json:"geomArea"`
	//	State         State         `json:"state"`
	//	County        County        `json:"county"`
	CommunityPlan *CommunityPlan  `json:"communityPlan"`
	Zoning        []Zoning        `json:"zoning"`
	Overlays      []Overlay       `json:"overlays"`
	Permits       []Permit        `json:"permits"`
	Bounds        geojson.Polygon `json:"bounds"`
	Geohash       *string         `json:"geohash"`
	Created       time.Time       `json:"created"`
	Updated       time.Time       `json:"updated"`
	Unlocked      bool            `json:"unlocked"`
}
