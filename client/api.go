package client

import (
	"github.com/scoutred/publisher/sdk/parcel"
)

type API struct {
	// Parcels is the client used to invoke /parcels APIs.
	Parcels *parcel.Client
}

//	set up our various endpoints
func (this *API) Init(key string) {
	this.Parcels = &parcel.Client{Key: key}
}
