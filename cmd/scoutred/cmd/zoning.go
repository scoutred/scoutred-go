package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	scoutred "github.com/scoutred/scoutred-go"
	"github.com/scoutred/scoutred-go/client"
)

var (
	zoningID int64
)

// zoningCmd represents the zoning command
var zoningCmd = &cobra.Command{
	Use:   "zoning",
	Short: "Fetch zoning records",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
			z   *scoutred.Zoning
		)

		c := client.New(cmd.Flag("key").Value.String())

		switch {
		case zoningID != 0:
			z, err = c.ZoningByID(zoningID)
			if err != nil {
				log.Fatal(err)
			}
		case lon != 0.0 && lat != 0.0:
			z, err = c.ZoningByLonLat(lon, lat)
			if err != nil {
				log.Fatal(err)
			}

		default:
			log.Fatal("must provide an id or a valid lon and lat combo")
		}

		// encode and write
		json.NewEncoder(os.Stdout).Encode(z)
	},
}

func init() {
	zoningCmd.Flags().Int64Var(&zoningID, "id", 0, "The ID of the zoning record to fetch")
	zoningCmd.Flags().Float64Var(&lon, "lon", 0.0, "the longitude to query")
	zoningCmd.Flags().Float64Var(&lat, "lat", 0.0, "the latitude to query")
	zoningCmd.PersistentFlags().StringP("key", "k", "", "your Scoutred API key (required)")
	zoningCmd.MarkPersistentFlagRequired("key")

	rootCmd.AddCommand(zoningCmd)
}
