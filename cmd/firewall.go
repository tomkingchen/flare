package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

type FirewallRule struct {
	Id           string   `json:"id"`
	Paused       bool     `json:"paused"`
	Description  string   `json:"description"`
	Action       string   `json:"action"`
	Filter       struct {
		Id           string   `json:"id"`
		Expression   string   `json:"expression"`
		Paused       bool     `json:"paused"`
		Description  string   `json:"description"`
    } `json:"filter"`
	CreatedOn    time.Time    `json:"created_on"`
	ModifiedOn   time.Time    `json:"modified_on"`
}

type FirewallRulesResult struct {
	Result []FirewallRule
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}

// firewallCmd represents the firewall command
var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "List firewall rules of a given zone",
	Run: func(cmd *cobra.Command, args []string) {
		var firewallRulesResult FirewallRulesResult
		var rules []FirewallRule
		URL := "https://api.cloudflare.com/client/v4/"
		apiUrl := URL + "zones/" + ZoneId + "/firewall/rules"
		result := fetchAPI(apiUrl)
		json.Unmarshal([]byte(result), &firewallRulesResult)
		rules = append(rules, firewallRulesResult.Result...)
		numOfPages := firewallRulesResult.ResultInfo.TotalPages
		for i := 2; i <= numOfPages; i++ {
			pageNum := strconv.Itoa(i)
			pagedUrl := apiUrl + "?page=" + pageNum
			pagedResult := fetchAPI(pagedUrl)
			json.Unmarshal([]byte(pagedResult), &firewallRulesResult)
			rules = append(rules, firewallRulesResult.Result...)
		}
		j, _ := json.Marshal(rules)
		fmt.Println(string(j))
	},
}

func init() {
	listCmd.AddCommand(firewallCmd)
	firewallCmd.PersistentFlags().StringVarP(&ZoneId, "zoneid", "z", "", "Zone ID")
	firewallCmd.MarkPersistentFlagRequired("zoneid")
}
