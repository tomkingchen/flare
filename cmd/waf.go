/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ZoneId string
var RulesetId string

// wafCmd represents the waf command
var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "List all WAF rules",
	Run: func(cmd *cobra.Command, args []string) {
		var apiUrl string
		URL := "https://api.cloudflare.com/client/v4/"
		if &RulesetId != nil {
			apiUrl = URL + "zones/" + ZoneId + "/rulesets/" + RulesetId
		}else{
			apiUrl = URL + "zones/" + ZoneId + "/"
		}
		result := fetchAPI(apiUrl)
		fmt.Println(result)
	},
}

func init() {
	listCmd.AddCommand(wafCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	wafCmd.PersistentFlags().StringVarP(&ZoneId, "zoneid", "z", "", "Zone ID")
	wafCmd.PersistentFlags().StringVarP(&RulesetId, "rulesetid", "r", "", "Ruleset ID")
	wafCmd.PersistentFlags().StringVarP(&PackageId, "packageid", "p", "", "WAF Package ID")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wafCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
