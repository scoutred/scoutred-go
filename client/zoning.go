package client

import (
	"fmt"
	"net/url"

	scoutred "github.com/scoutred/scoutred-go"
)

// ZoningByID will fetch a pracel resource by Scoutred ID
func (c *Client) ZoningByID(id int64) (*scoutred.Zoning, error) {
	var zone *scoutred.Zoning

	//	format our url for this request
	url := fmt.Sprintf("/zoning/%d", id)

	//	make our requeset
	if err := c.Call("GET", url, nil, &zone); err != nil {
		return nil, err
	}

	return zone, nil
}

// ZoningByLonLat will fetch a zoning resources by lon / lat value.
func (c *Client) ZoningByLonLat(lon, lat float64) (*scoutred.Zoning, error) {
	var zoning scoutred.Zoning

	uri := url.URL{
		Path: "/zoning",
	}

	// build query params
	q := uri.Query()
	q.Set("lon", fmt.Sprintf("%v", lon))
	q.Set("lat", fmt.Sprintf("%v", lat))

	// set query params on the uri
	uri.RawQuery = q.Encode()

	//	make our requeset
	if err := c.Call("GET", uri.String(), nil, &zoning); err != nil {
		return nil, err
	}

	return &zoning, nil
}
