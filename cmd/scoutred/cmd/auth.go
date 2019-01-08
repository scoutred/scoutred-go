package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/scoutred/scoutred-go/client"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the Scoutred API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		email := cmd.Flag("email").Value.String()
		pw := cmd.Flag("password").Value.String()

		// create a new API client with no API key
		c := client.New("")

		token, err := c.Auth(email, pw)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v", token)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().StringP("email", "e", "", "Email used to register on Scoutred")
	authCmd.MarkFlagRequired("email")
	authCmd.Flags().StringP("password", "p", "", "Password associated with Email on Scoutred")
	authCmd.MarkFlagRequired("password")
}
