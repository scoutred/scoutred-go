package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
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
		p, err := c.ParcelByID(parcelID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(os.Stdout).Encode(p)
	},
}

func init() {
	rootCmd.AddCommand(parcelsCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	parcelsCmd.Flags().Int64Var(&parcelID, "id", 0, "The ID of the parcel record to fetch")
}
