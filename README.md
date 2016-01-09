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

func main(){
	apiKey := "8da87f..." //	your API key

	//	create a new ScoutRED API client
	api := client.API{}
	api.Init(apiKey)

	srp, err := api.Parcels.GetBySrcId(*sgp.Apn)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", srp)	
}
```
