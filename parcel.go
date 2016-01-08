package scoutred

import (
	"time"
)

type Parcel struct {
	//	internal system id
	//	we use an int64 as there are an estimated 3 billion parcels in the US alone
	Id int64 `json:"id"`
	//	sha1 hash of the id used to identify the parcel at the data source
	SourceId string `json:"srcId"`
	//	assessor parcel number
	APN  *string `json:"apn"`
	Data struct {
		Owner struct {
			Name1    *string `json:"name1"`
			Name2    *string `json:"name2"`
			Name3    *string `json:"name3"`
			Address1 *string `json:"address1"`
			Address2 *string `json:"address2"`
			Address3 *string `json:"address3"`
			Address4 *string `json:"address4"`
			Zip      *string `json:"zip"`
		} `json:"owner"`
		Address          Address `json:"address"`
		LegalDescription *string `json:"legalDescription"`
		Subdivision      struct {
			Map  *string `json:"map"`
			Name *string `json:"name"`
		} `json:"subdivision"`
		//	assessor data
		Assessor struct {
			Land         *float64 `json:"land"`
			Improvements *float64 `json:"improvements"`
		} `json:"assessor"`
	} `json:"data"`
	Geom     MultiPolygon `json:"geom"`
	GeomArea float64      `json:"geomArea"`
	Updated  time.Time    `json:"update"`
	Created  time.Time    `json:"created"`
}
