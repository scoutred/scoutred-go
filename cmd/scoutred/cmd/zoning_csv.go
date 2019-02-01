package cmd

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/scoutred/scoutred-go/client"
)

var (
	lonIndex, latIndex uint64
	csvHeader          bool
	csvAppend          bool
)

// zoningHeader are the additional zoning headers that will be added
var zoningHeader = []string{
	"designation",
}

// csvCmd represents the csv command
var zoningCsvCmd = &cobra.Command{
	Use:   "csv",
	Short: "fetching zoning data in CSV format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		c := client.New(cmd.Flag("key").Value.String())

		// file to read from
		in, err := os.Open(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal(err)
		}

		var out io.Writer
		if cmd.Flag("output").Value.String() != "" {
			// file to write to
			out, err = os.Create(cmd.Flag("output").Value.String())
			if err != nil {
				log.Fatal(err)
			}
		} else {
			out = os.Stdout
		}

		// CSV reader
		r := csv.NewReader(in)

		// CSV writer
		w := csv.NewWriter(out)

		// iterate the records
		first := true
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			// skip processing the header
			if csvHeader && first {
				first = false

				// decorate the first record with additional headers if we're in append mode
				if csvAppend {
					record = append(record, zoningHeader...)
				} else {
					record = zoningHeader
				}

				if err := w.Write(record); err != nil {
					log.Fatalln("error writing record to csv:", err)
				}

				// Write any buffered data to the underlying writer
				w.Flush()
				if err := w.Error(); err != nil {
					log.Fatal(err)
				}
				continue
			}

			switch {
			case lonIndex != 0 && latIndex != 0:
				// error checking
				if len(record) < int(lonIndex) {
					log.Fatal("invlaid index provided for --lat")
				}
				if len(record) < int(latIndex) {
					log.Fatal("invlaid index provided for --lat")
				}

				lon, err := strconv.ParseFloat(record[lonIndex], 64)
				if err != nil {
					log.Fatalf("invalid float value for --lon: %v", err)
				}

				lat, err := strconv.ParseFloat(record[latIndex], 64)
				if err != nil {
					log.Fatalf("invalid float value for --lat: %v", err)
				}

				zoning, err := c.ZoningByLonLat(lon, lat)
				if err != nil {
					// check if err is a client error
					apiErr, ok := err.(client.Error)
					if !ok {
						log.Fatal(err)
					}

					switch apiErr.StatusCode {
					case http.StatusNotFound:
						// log not found error?
						continue
					default:
						log.Fatal(err)
					}
				}

				if zoning != nil && zoning.Designation != nil {
					// this should map to zoningHeader
					zoningVals := []string{*zoning.Designation}

					if csvAppend {
						record = append(record, zoningVals...)
					} else {
						record = []string{*zoning.Designation}
					}
				}

				if err := w.Write(record); err != nil {
					log.Fatalln("error writing record to csv:", err)
				}

				// write any buffered data to the underlying writer
				w.Flush()
				if err := w.Error(); err != nil {
					log.Fatal(err)
				}
			default:
				log.Fatal("missing CSV indexes to work with")
			}
		}
	},
}

func init() {
	zoningCmd.AddCommand(zoningCsvCmd)

	// flags
	zoningCsvCmd.Flags().StringP("input", "i", "", "CSV file to use as input")
	// TODO(arolek): remove this requirement and support streaming from stdin
	zoningCsvCmd.MarkFlagRequired("input")
	zoningCsvCmd.Flags().StringP("output", "o", "", "CSV file name to write to")

	zoningCsvCmd.Flags().Uint64Var(&lonIndex, "lon", 0, "the column index of the CSV for the longitude value")
	zoningCsvCmd.Flags().Uint64Var(&latIndex, "lat", 0, "the column index of the CSV for the latitude value")

	zoningCsvCmd.Flags().BoolVar(&csvHeader, "header", false, "if the CSV has a header set this flag to true")
	zoningCsvCmd.Flags().BoolVar(&csvAppend, "append", false, "append Scoutred data to the incoming CSV file")

}
