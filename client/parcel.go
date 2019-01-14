package client

import (
	"fmt"
	"net/url"

	scoutred "github.com/scoutred/scoutred-go"
)

// ParcelByID will fetch a pracel resource by Scoutred ID
func (c *Client) ParcelByID(id int64) (*scoutred.Parcel, error) {
	var (
		parcel *scoutred.Parcel
	)

	//	format our url for this request
	uri := fmt.Sprintf("/parcels/%d", id)

	//	make our requeset
	if err := c.Call("GET", uri, nil, &parcel); err != nil {
		return nil, err
	}

	return parcel, nil
}

// ParcelByLonLat will fetch pracel resources by lon / lat value.
// The result set can include more than a single parcel in certain
// situations (i.e. condo towers, duplexes, etc. )
func (c *Client) ParcelsByLonLat(lon, lat float64) ([]scoutred.Parcel, error) {
	var (
		parcels []scoutred.Parcel
	)

	uri := url.URL{
		Path: "/parcels",
	}

	// build query params
	q := uri.Query()
	q.Set("lon", fmt.Sprintf("%v", lon))
	q.Set("lat", fmt.Sprintf("%v", lat))

	// set query params on the uri
	uri.RawQuery = q.Encode()

	//	make our requeset
	if err := c.Call("GET", uri.String(), nil, &parcels); err != nil {
		return nil, err
	}

	return parcels, nil
}
