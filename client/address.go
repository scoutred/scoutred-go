package client

import (
	"net/url"

	scoutred "github.com/scoutred/scoutred-go"
)

// AddressSearch will query Scoutred for addresses or APNs
func (c *Client) AddressSearch(query string) ([]scoutred.Address, error) {
	var addresses []scoutred.Address

	// set the query params so they can be properly encoded
	vals := url.Values{}
	vals.Set("q", query)

	// build our URL and query params
	uri := url.URL{
		Path:     "/search/addresses",
		RawQuery: vals.Encode(),
	}

	// make our requeset
	if err := c.Call("GET", uri.String(), nil, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}
