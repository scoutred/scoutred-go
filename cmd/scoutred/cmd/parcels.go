package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/scoutred/scoutred-go/client"
)

var (
	parcelID int64
)

// parcelsCmd represents the parcels command
var parcelsCmd = &cobra.Command{
	Use:   "parcels",
	Short: "fetch parcel records using various query params",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(cmd.Flag("key").Value.String())

		// switch on command param combos
		switch {
		case parcelID != 0:
			p, err := c.ParcelByID(parcelID)
			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(os.Stdout).Encode(p)
		case lon != 0.0 && lat != 0.0:
			ps, err := c.ParcelsByLonLat(lon, lat)
			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(os.Stdout).Encode(ps)
		default:
			log.Fatal("must provide an id or a valid lon and lat combo")
		}
	},
}

func init() {
	parcelsCmd.Flags().Int64Var(&parcelID, "id", 0, "the ID of the parcel record to fetch")
	parcelsCmd.Flags().Float64Var(&lon, "lon", 0.0, "the longitude to query")
	parcelsCmd.Flags().Float64Var(&lat, "lat", 0.0, "the latitude to query")
	parcelsCmd.PersistentFlags().StringP("key", "k", "", "your Scoutred API key (required)")
	parcelsCmd.MarkPersistentFlagRequired("key")

	rootCmd.AddCommand(parcelsCmd)
}
