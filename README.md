# Go Scoutred
The official Go client for the Scoutred API. This client can be used as a Go package or via the command line.

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
	apiKey := "8da87f..." // your API key

	// create a new Scoutred API client
	c := client.New(apiKey)

	srp, err := c.ParcelByID(22219)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", srp)	
}
```
