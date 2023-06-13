package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/ronyldo12/secrets/masterkey"
	"github.com/ronyldo12/secrets/secret"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add SECRET_NAME",
	Short: "Add a new secret",
	Long: `You can add a new secrets
	examples:
	secrets add API_AUTH_CLIENT
	secrets add API_AUTH_SECRET
    `,
	Run: func(cmd *cobra.Command, args []string) {
		masterkey.Check()

		if len(args) == 0 {
			log.Fatal("you should inform the secret name: secrets add SECRET_NAME")
		}
		fmt.Print("type the value: ")
		byteValue, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return
		}
		value := string(byteValue)
		secretsData := secret.GetSecretsData()
		secretsData.Secrets[args[0]] = value
		secretsData.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//addCmd.PersistentFlags().String("SECRET_NAME", "", "The secret name, i.e: API_AUTH_CLIENT")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
