/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JonathanGodar/go-web-gin/client/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var serverURL string

// This code is reeeeeallly bad but its just for fun so EH
var userService *api.UserService 
var trackerService *api.TrackerService 

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// PostRunE: func(cmd *cobra.Command, args []string) error {
	// 	println("Writing config")
	// 	return viper.WriteConfig()
	// },
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")
	rootCmd.PersistentFlags().StringVar(&serverURL, "serverUrl", "http://localhost:8080", "Which server url should be used")
	viper.BindPFlag("serverUrl", rootCmd.Flags().Lookup("serverUrl"))

	initConfig()

	println("token:", viper.GetString("accessToken"))

	client := api.New(serverURL + "/oto/")
	if token := viper.GetString("accessToken"); token != ""{
		client.BeforeRequest = func(req *http.Request) error {
			log.Println("Adding header")
			req.Header.Add("Authorization", "Bearer " + token)
			return nil
		}
	}


	trackerService = api.NewTrackerService(client)
	userService = api.NewUserService(client)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".simple-tracker-cli")
		viper.SetConfigType("json")
	}

	// viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SafeWriteConfig()
}
