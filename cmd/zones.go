/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"text/tabwriter"
)

var TableOutput bool

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
		// var zoneResults ZoneResults
		URL := "https://api.cloudflare.com/client/v4/"
		zonesUrl := URL + "/zones"
		result := fetchAPI(zonesUrl)

		// Print output in table format
		if TableOutput {
			zoneResults := ZoneResults{}
			json.Unmarshal([]byte(result), &zoneResults)
			zones := zoneResults.Result
			// Paginating results
			numOfPages := zoneResults.ResultInfo.TotalPages
			for i := 2; i <= numOfPages; i++ {
				pageNum := strconv.Itoa(i)
				pagedUrl := zonesUrl + "?page=" + pageNum
				pagedResult := fetchAPI(pagedUrl)
				json.Unmarshal([]byte(pagedResult), &zoneResults)
				zones = append(zones, zoneResults.Result...)
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			for _, zone := range zones {
				fmt.Fprintln(w, zone.Name + "\t" + zone.ID + "\t" )
			}
			w.Flush()
		}else{
			fmt.Println(result)
		}
	},
}

func init() {
	listCmd.AddCommand(zonesCmd)
	zonesCmd.Flags().BoolVarP(&TableOutput, "table", "t", false, "Set output to table")
}
