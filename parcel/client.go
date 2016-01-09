package parcel

import (
	"bytes"
	"encoding/json"

	"github.com/scoutred/scoutred-go"
)

// Client is used to invoke /parcels APIs.
type Client struct {
	Key string
}

//	fetch parcel by ScoutRED id
func (this Client) GetById(id int64) (parcel *scoutred.Parcel, err error) {
	//	make our requeset
	err = scoutred.Call("GET", "/parcels/src-id/"+id, this.Key, nil, &parcel)

	return
}

//	fetch parcel by id generated from source data import
func (this Client) GetBySrcId(srcId string) (parcel *scoutred.Parcel, err error) {
	//	make our requeset
	err = scoutred.Call("GET", "/parcels/src-id/"+srcId, this.Key, nil, &parcel)

	return
}

//	update parcel by id generated from source data import
func (this Client) UpdateBySrcId(srcId string, p *scoutred.Parcel) (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	//	make our requeset
	err = scoutred.Call("PUT", "/parcels/src-id/"+srcId, this.Key, bytes.NewBuffer(data), nil)

	return
}

//	delete parcel by id generated from source data import
func (this Client) DeleteBySrcId(srcId string) (err error) {
	//	make our requeset
	err = scoutred.Call("DELETE", "/parcels/src-id/"+srcId, this.Key, nil, nil)

	return
}
