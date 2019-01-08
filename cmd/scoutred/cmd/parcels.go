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
	Short: "Fetch parcel records",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(cmd.Flag("key").Value.String())

		p, err := c.ParcelByID(parcelID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(os.Stdout).Encode(p)
	},
}

func init() {
	parcelsCmd.Flags().Int64Var(&parcelID, "id", 0, "The ID of the parcel record to fetch")
	parcelsCmd.Flags().StringP("key", "k", "", "your Scoutred API key (required)")
	parcelsCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(parcelsCmd)
}
