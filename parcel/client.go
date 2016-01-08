package parcel

import (
	"bytes"
	"encoding/json"

	"github.com/scoutred/publisher/sdk"
)

// Client is used to invoke /parcels APIs.
type Client struct {
	Key string
}

func (this Client) GetBySrcId(srcId string) (parcel *sdk.Parcel, err error) {
	//	make our requeset
	err = sdk.Call("GET", "/parcels/src-id/"+srcId, this.Key, nil, &parcel)

	return
}

//	can be used to create and update a parcel
func (this Client) UpdateBySrcId(srcId string, p *sdk.Parcel) (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	//	make our requeset
	err = sdk.Call("PUT", "/parcels/src-id/"+srcId, this.Key, bytes.NewBuffer(data), nil)

	return
}

func (this Client) DeleteBySrcId(srcId string) (err error) {
	//	make our requeset
	err = sdk.Call("DELETE", "/parcels/src-id/"+srcId, this.Key, nil, nil)

	return
}
