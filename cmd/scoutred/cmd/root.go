package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// used by various sub commands which is why it lives at the root
	lat, lon float64
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
