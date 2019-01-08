package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scoutred",
	Short: "Scoutred",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
func init() {
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "your Scoutred API key (required)")
	rootCmd.MarkFlagRequired("key")
}

func setAPIKey() {
	if key == "" {
		log.Fatal("key required. include the --key flag (i.e. scoutred --key [your API key] [command]")
	}

	c = client.New(key)
}
*/
