/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ronyldo12/secrets/crypt"
	"github.com/ronyldo12/secrets/masterkey"
	"github.com/ronyldo12/secrets/secret_file"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create a session in the current terminal",
	Long:  `Use this command to a session in the current terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		mkey, err := masterkey.Read()
		if err != nil {
			log.Fatal(err)
		}

		validateMasterKeyCanDecrypt(mkey)

		masterkey.Define(mkey)
		fmt.Println("you're logged!")
	},
}

func validateMasterKeyCanDecrypt(mkey string) {
	file_content := secret_file.FileGetContent()
	if file_content == "" {
		return
	}

	decrypted, err := crypt.Decrypt(file_content, masterkey.Prepare(mkey))

	if err != nil {
		log.Fatal(err)
	}

	if json.Valid([]byte(decrypted)) == false {
		fmt.Println("\nInvalid Master Key. Try again!")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
