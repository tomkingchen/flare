/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Id string

// zoneCmd represents the zone command
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Get details of an individual zone",
	Run: func(cmd *cobra.Command, args []string) {
		var zoneSettings ZoneSettings
		URL := "https://api.cloudflare.com/client/v4/"
		zoneSettingsUrl := URL + "/zones/" + Id + "/settings"
		err := getAPI(zoneSettingsUrl, &zoneSettings)
		if err != nil {
			fmt.Printf("Error getting zone settings: %s/n", err.Error())
		}else{
			fmt.Println(zoneSettings.Result)
		}
	},
}

func init() {
	getCmd.AddCommand(zoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	zoneCmd.PersistentFlags().StringVarP(&Id, "id", "i", "", "Zone ID")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
