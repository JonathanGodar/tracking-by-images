/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/JonathanGodar/go-web-gin/client/api"
	"github.com/spf13/cobra"
	"github.com/atotto/clipboard"
)

// trackerCmd represents the tracker command
var trackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

var trackerAddCmd = &cobra.Command {
	Use: "add",
	Short: "Cooool dude",
	RunE: func(cmd *cobra.Command, args []string) error {
		me, err := userService.GetMe(cmd.Context(), "")
		if err != nil {
			return err
		}

		resp, err := trackerService.AddTracker(cmd.Context(), api.AddTrackerRequest{
			OwnerID: me.User.ID,
			IsActive: false,
		})

		if err != nil {
			return err
		}

		fullURL := serverURL + resp.Tracker.URL

		fmt.Println("Added tracker. URL:\n", fullURL)

		if err = clipboard.WriteAll(fullURL); err != nil {
			fmt.Println("Could not copy to clipboard, please copy manually")
		} else {
			fmt.Println("The url has been copied to the clipboard")
		}

		fmt.Println("Setup your tracker and then press ENTER to activate it")
		fmt.Scanln()

		if _, err := trackerService.UpdateTracker(cmd.Context(), api.UpdateTrackerRequest {
			ID: resp.Tracker.ID,
			IsActive: true,
		}); err != nil {
			fmt.Println("Could not activate tracker!")
		} else {
			fmt.Println("The tracker is now active!")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(trackerCmd)



	trackerCmd.AddCommand(trackerAddCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
