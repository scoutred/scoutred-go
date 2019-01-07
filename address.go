package scoutred

import (
	"time"

	"github.com/scoutred/app/geojson"
)

type Address struct {
	ID           *int64         `json:"id"`
	ParcelID     *int64         `json:"parcelId"`
	Street       Street         `json:"street"`
	Unit         *string        `json:"unit"`
	Postal       *string        `json:"postal"`
	Jurisdiction *string        `json:"jurisdiction"`
	State        *string        `json:"state"`
	Country      *string        `json:"country"`
	Geom         *geojson.Point `json:"geom"`
	Created      *time.Time     `json:"created"`
	Updated      *time.Time     `json:"updated"`
}

type Street struct {
	Number         *int32  `json:"number"`
	NumberFraction *string `json:"numberFraction"`
	PreDirection   *string `json:"preDirection"`
	Name           *string `json:"name"`
	Suffix         *string `json:"suffix"`
	PostDirection  *string `json:"postDirection"`
}
