package client

import (
	"fmt"

	scoutred "github.com/scoutred/scoutred-go"
)

// ParcelByID will fetch a pracel resource by Scoutred ID
func (c Client) ParcelByID(id int64) (*scoutred.Parcel, error) {
	var (
		err    error
		parcel *scoutred.Parcel
	)

	//	format our url for this request
	url := fmt.Sprintf("/parcels/%d", id)

	//	make our requeset
	err = scoutred.Call("GET", url, c.Key, nil, &parcel)
	if err != nil {
		return nil, err
	}

	return parcel, nil
}
