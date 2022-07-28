package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// firewallCmd represents the firewall command
var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "List firewall rules of a given zone",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("firewall called" + ZoneId)
	},
}

func init() {
	listCmd.AddCommand(firewallCmd)
	firewallCmd.PersistentFlags().StringVarP(&ZoneId, "zoneid", "z", "", "Zone ID")
	firewallCmd.MarkPersistentFlagRequired("zoneid")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// firewallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// firewallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
