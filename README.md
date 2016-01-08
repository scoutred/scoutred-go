# Go ScoutRED
The official Go client for the ScoutRED API

## Installation

```
go get github.com/scoutred/scoutred-go
```

## Usage

### Parcels

```go
import (
	"log"

	"github.com/scoutred/scoutred-go"
	"github.com/scoutred/scoutred-go/client"
)

//	create a new ScoutRED API client
scoutred := client.API{}
scoutred.Init(apiKey)

srp, err := scoutred.Parcels.GetBySrcId(*sgp.Apn)
if err != nil {
	log.Fatal(err)
}

log.Printf("%+v", srp)
```
