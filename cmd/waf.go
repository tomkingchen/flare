/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// wafCmd represents the waf command
var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "List all WAF rules",
	Long: `List all firewall rules, Managed WAF rules together include the new RuleSet Engine rules`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("waf called")
	},
}

func init() {
	listCmd.AddCommand(wafCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wafCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wafCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
