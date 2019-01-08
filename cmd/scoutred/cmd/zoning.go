package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

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
		c := client.New(cmd.Flag("key").Value.String())

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
	zoningCmd.Flags().StringP("key", "k", "", "your Scoutred API key (required)")
	zoningCmd.MarkFlagRequired("key")
}
