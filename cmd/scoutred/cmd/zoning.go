package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
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
		z, err := c.ZoningByID(zoningID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(os.Stdout).Encode(z)
	},
}

func init() {
	rootCmd.AddCommand(zoningCmd)

	zoningCmd.Flags().Int64Var(&zoningID, "id", 0, "The ID of the zoning record to fetch")
}
