package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gatherec2info",
	Short: "This tool is gathering EC2 Information.",
	Long:  "This tool is gathering EC2 Information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()
}
