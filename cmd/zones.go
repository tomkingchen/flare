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
	Short: "List Cloudflare zones",
	Run: func(cmd *cobra.Command, args []string) {
		var zoneResults ZoneResults
		var zones []Zone
		URL := "https://api.cloudflare.com/client/v4/"
		zonesUrl := URL + "/zones"
		result := fetchAPI(zonesUrl)
		json.Unmarshal([]byte(result), &zoneResults)
		zones = append(zones, zoneResults.Result...)
		// Paginating results
		numOfPages := zoneResults.ResultInfo.TotalPages
		for i := 2; i <= numOfPages; i++ {
			pageNum := strconv.Itoa(i)
			pagedUrl := zonesUrl + "/?page=" + pageNum
			pagedResult := fetchAPI(pagedUrl)
			json.Unmarshal([]byte(pagedResult), &zoneResults)
			zones = append(zones, zoneResults.Result...)
		}
		// Print output in table format
		if TableOutput {
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			for _, zone := range zones {
				fmt.Fprintln(w, zone.Name + "\t" + zone.ID + "\t" )
			}
			w.Flush()
		}else{
			j, _ := json.Marshal(zones)
			fmt.Println(string(j))
		}
	},
}

func init() {
	listCmd.AddCommand(zonesCmd)
	zonesCmd.Flags().BoolVarP(&TableOutput, "table", "t", false, "set output to table format")
}
