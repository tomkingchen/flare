package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// certverificationCmd represents the certverification command
var certverificationCmd = &cobra.Command{
	Use:   "certverification",
	Short: "Get certificate verification details",
	Run: func(cmd *cobra.Command, args []string) {
		URL := "https://api.cloudflare.com/client/v4/"
		certVeriUrl := URL + "zones/" + ZoneId + "/ssl/verification"
		result := fetchAPI(certVeriUrl)
		fmt.Println(result)
	},
}

func init() {
	getCmd.AddCommand(certverificationCmd)
	certverificationCmd.PersistentFlags().StringVarP(&ZoneId, "zoneid", "z", "", "Zone ID")
	certverificationCmd.MarkPersistentFlagRequired("zoneid")
}
