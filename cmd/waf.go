/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

var ZoneId string
var UseRulesetEngine bool

type RulesetRule struct {
		Id          string    `json:"id"`
		Version     string    `json:"version"`
		Action      string    `json:"action"`
		Categories  []string  `json:"categories"`
		Description string    `json:"description"`
		LastUpdated time.Time `json:"last_updated"`
		Ref         string    `json:"ref"`
		Enabled     bool      `json:"enabled"`
}

type Ruleset struct {
	Result struct {
		Id          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Source      string   `json:"source"`
		Kind        string   `json:"kind"`
		Version     string   `json:"version"`
		Rules       []RulesetRule `json:"rules"`
		LastUpdated  time.Time
		Phase        string
	}
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

// wafCmd represents the waf command
var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "List all WAF rules",
	Run: func(cmd *cobra.Command, args []string) {
		var apiUrl string
		// Predefined Cloudflare Managed WAF Rulesets 
		rulesetIds := []string{"efb7b8c949ac4650a09736fc376e9aee", "4814384a9e5d4991b9815dcfc25d2f1f", "c2e184081120413c86c3ab7e14069605"}
		// Predefined Cloudflare Managed Firewall rules packs
		// packageIds := []string{"1e334934fd7ae32ad705667f8c1057aa", "c504870194831cd12c3fc0284f294abb"}
		URL := "https://api.cloudflare.com/client/v4/"
		if UseRulesetEngine {
			var ruleset Ruleset
			var rules []RulesetRule
			for _, rulesetId := range rulesetIds {
				apiUrl = URL + "zones/" + ZoneId + "/rulesets/" + rulesetId
				result := fetchAPI(apiUrl)
				json.Unmarshal([]byte(result), &ruleset)
				rules = append(rules, ruleset.Result.Rules...)
			}
			fmt.Println(rules)
		}else{
			// apiUrl = URL + "zones/" + ZoneId + "/firewall/waf/packages" + PackageId + "/rules"
			// result := fetchAPI(apiUrl)
			// fmt.Println(result)
			fmt.Println("Use old Managed Firewall API to retrieve WAF rules")
		}
	},
}

func init() {
	listCmd.AddCommand(wafCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	wafCmd.PersistentFlags().StringVarP(&ZoneId, "zoneid", "z", "", "Zone ID")
	wafCmd.MarkPersistentFlagRequired("zoneid")
	wafCmd.PersistentFlags().BoolVarP(&UseRulesetEngine, "ruleset", "r", false, "Use Ruleset Engine")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wafCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
