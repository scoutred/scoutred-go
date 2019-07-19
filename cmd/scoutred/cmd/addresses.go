package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/scoutred/scoutred-go/client"
)

// addressesCmd represents the addresses command
var addressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "search addresses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		query := cmd.Flag("query").Value.String()

		c := client.New(cmd.Flag("key").Value.String())

		addrs, err := c.AddressSearch(query)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(os.Stdout).Encode(addrs)
	},
}

func init() {
	addressesCmd.Flags().StringP("query", "q", "", "the address query to use for searching")
	addressesCmd.MarkFlagRequired("query")
	addressesCmd.PersistentFlags().StringP("key", "k", "", "your Scoutred API key (required)")
	addressesCmd.MarkPersistentFlagRequired("key")

	rootCmd.AddCommand(addressesCmd)
}
