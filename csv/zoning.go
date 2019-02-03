package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/scoutred/scoutred-go/client"
	"github.com/scoutred/scoutred-go/log"
)

// zoningHeader are the additional zoning headers that will be added
var zoningHeader = []string{
	"designation",
}

func ZoningByLonLat(params LonLatParams) error {
	// CSV reader
	r := csv.NewReader(params.Reader)

	// CSV writer
	w := csv.NewWriter(params.Writer)

	// iterate the records
	first := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// skip processing the header
		if params.Header && first {
			first = false

			// decorate the first record with additional headers if we're in append mode
			if params.Append {
				record = append(record, zoningHeader...)
			} else {
				record = zoningHeader
			}

			if err := w.Write(record); err != nil {
				return fmt.Errorf("error writing record to csv: %v", err)
			}

			// Write any buffered data to the underlying writer
			w.Flush()
			if err := w.Error(); err != nil {
				return err
			}
			continue
		}

		// error checking
		if len(record) < int(params.LonIdx) {
			return fmt.Errorf("invlaid index provided for longitude")
		}
		if len(record) < int(params.LatIdx) {
			return fmt.Errorf("invlaid index provided for latitude")
		}

		// value parsing
		lon, err := strconv.ParseFloat(record[params.LonIdx], 64)
		if err != nil {
			return fmt.Errorf("invalid float value for longitude: %v", err)
		}

		lat, err := strconv.ParseFloat(record[params.LatIdx], 64)
		if err != nil {
			return fmt.Errorf("invalid float value for latitude: %v", err)
		}

		// make the API request
		zoning, err := params.Client.ZoningByLonLat(lon, lat)
		if err != nil {
			// check if err is a client error
			apiErr, ok := err.(client.Error)
			if !ok {
				return err
			}

			switch apiErr.StatusCode {
			case http.StatusNotFound:
				// log the warning but don't stop processing
				log.Warn.Printf("record not found for LonLat value (%v, %v)", lon, lat)
			default:
				return err
			}
		}

		if zoning != nil && zoning.Designation != nil {
			// this should map to zoningHeader
			zoningVals := []string{*zoning.Designation}

			if params.Append {
				record = append(record, zoningVals...)
			} else {
				record = []string{*zoning.Designation}
			}
		}

		if err := w.Write(record); err != nil {
			return fmt.Errorf("error writing record to csv: %v", err)
		}

		// write any buffered data to the underlying writer
		w.Flush()

		// check for error
		if err := w.Error(); err != nil {
			return err
		}
	}

	return nil
}
