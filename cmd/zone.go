package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Zone Id
var Id string

// zoneCmd represents the zone command
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Get details of an individual zone",
	Run: func(cmd *cobra.Command, args []string) {
		URL := "https://api.cloudflare.com/client/v4/"
		zoneSettingsUrl := URL + "zones/" + Id + "/settings"
		result := fetchAPI(zoneSettingsUrl)
		fmt.Println(result)
	},
}

func init() {
	getCmd.AddCommand(zoneCmd)
	zoneCmd.PersistentFlags().StringVarP(&Id, "id", "i", "", "Zone ID")
}
