package client

import (
	"fmt"

	scoutred "github.com/scoutred/scoutred-go"
)

// ZoningByID will fetch a pracel resource by Scoutred ID
func (c Client) ZoningByID(id int64) (*scoutred.Zoning, error) {
	var (
		err  error
		zone *scoutred.Zoning
	)

	//	format our url for this request
	url := fmt.Sprintf("/zoning/%d", id)

	//	make our requeset
	err = scoutred.Call("GET", url, c.Key, nil, &zone)
	if err != nil {
		return nil, err
	}

	return zone, nil
}
