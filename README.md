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


## Command Line Interface (CLI)
This repository also has a CLI which can be downloaded from the release page. Here are some examples of usage: 

### Decorate a CSV list with zoning designations and output to a new file

```sh
./scoutred zoning csv --key [your API key] -i [incoming CSV file] --lon [CSV column to map to lat] --lat [CSV column to map to lat] --header --append -o [file to output to]
```


### Parcels

```sh
$ ./scoutred parcels -key [your API key] --lon -117.235974 --lat 32.952645
```