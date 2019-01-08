package client

import (
	"fmt"

	scoutred "github.com/scoutred/scoutred-go"
)

// ZoningByID will fetch a pracel resource by Scoutred ID

func (c *Client) ZoningByID(id int64) (*scoutred.Zoning, error) {
	var (
		zone *scoutred.Zoning
	)

	//	format our url for this request
	url := fmt.Sprintf("/zoning/%d", id)

	//	make our requeset
	if err := c.Call("GET", url, nil, &zone); err != nil {
		return nil, err
	}

	return zone, nil
}
