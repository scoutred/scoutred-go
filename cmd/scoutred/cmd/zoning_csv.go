package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/scoutred/scoutred-go/client"
	"github.com/scoutred/scoutred-go/csv"
)

var (
	lonIdx    uint64
	latIdx    uint64
	csvHeader bool
	csvAppend bool
)

// csvCmd represents the csv command
var zoningCsvCmd = &cobra.Command{
	Use:   "csv",
	Short: "fetching zoning data in CSV format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		client := client.New(cmd.Flag("key").Value.String())

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

		switch {
		case lonIdx != 0 && latIdx != 0:
			// setup our params
			csvParams := csv.LonLatParams{
				LonIdx: lonIdx,
				LatIdx: latIdx,
				Params: csv.Params{
					Append: csvAppend,
					Header: csvHeader,
					Reader: in,
					Writer: out,
					Client: client,
				},
			}

			if err := csv.ZoningByLonLat(csvParams); err != nil {
				log.Fatal(err)
			}

		default:
			log.Fatal("missing CSV indexes to work with")
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

	zoningCsvCmd.Flags().Uint64Var(&lonIdx, "lon", 0, "the column index of the CSV for the longitude value")
	zoningCsvCmd.Flags().Uint64Var(&latIdx, "lat", 0, "the column index of the CSV for the latitude value")

	zoningCsvCmd.Flags().BoolVar(&csvHeader, "header", false, "if the CSV has a header set this flag to true")
	zoningCsvCmd.Flags().BoolVar(&csvAppend, "append", false, "append Scoutred data to the incoming CSV file")

}
