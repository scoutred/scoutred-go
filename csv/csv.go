package csv

import (
	"io"

	"github.com/scoutred/scoutred-go/client"
)

type Params struct {
	Header bool
	// Append indicates if the data returned from the API should be appended to the CSV row
	Append bool
	// Reader references where data is coming from.
	Reader io.Reader
	// Writer references where data will be written
	Writer io.Writer
	// Client is a reference to the instantiated Scoutred client
	Client *client.Client
}

type LonLatParams struct {
	// LatIdx is the index of the column in the CSV which represents the latitude of the record
	LatIdx uint64
	// LonIdx is the index of the column in the CSV which represents the longitude of the record
	LonIdx uint64
	// Params are the common params across all CSV jobs
	Params
}
