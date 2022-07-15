/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/JonathanGodar/go-web-gin/client/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth [email] [password]",
	Short: "Authenticates with a backend",
	Long: `Takes an email and password which are then used to get an access token from the specified server.

	The access token will be saved into the configuration file
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		password := args[1]
		client := api.New(serverURL + "oto/")

		userService := api.NewUserService(client)

		resp, err := userService.GetAccessToken(context.Background(),
			api.GetAccessTokenRequest{
				Email:    email,
				Password: password,
			})

		if err != nil {
			return err
		}

		viper.Set("accessToken", resp.Token)
		if err := viper.WriteConfig(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Args = cobra.ExactValidArgs(2)

	// cobra.NoArgs(authCmd, authCmd.)

	// authCmd.Args = cobra.NoArgs()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
