/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
)


// zonesCmd represents the zones command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var zoneResults ZoneResults
		URL := "https://api.cloudflare.com/client/v4/"
		zonesUrl := URL + "/zones"
		zoneResults.queryAPI(zonesUrl)
		
		zones := zoneResults.Result
		// Paginating results
		numOfPages := zoneResults.ResultInfo.TotalPages
		for i := 2; i <= numOfPages; i++ {
			pageNum := strconv.Itoa(i)
			pagedUrl := zonesUrl + "?page=" + pageNum
			zoneResults.queryAPI(pagedUrl)
			zones = append(zones, zoneResults.Result...)
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		for _, zone := range zones {
			fmt.Fprintln(w, zone.Name + "\t" + zone.ID + "\t" )
		}
		w.Flush()
	},
}

func init() {
	listCmd.AddCommand(zonesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zonesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zonesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Query API
func (z *ZoneResults) queryAPI(apiUrl string) *ZoneResults {
	var c cred
	ApiCred := c.getCred()
	ApiEmail := ApiCred.ApiEmail
	ApiKey := ApiCred.ApiKey
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("X-Auth-Email", ApiEmail)
	req.Header.Add("X-Auth-Key", ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(responseData), &z)

	return z
}