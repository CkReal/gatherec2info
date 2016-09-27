package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var toolVer = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display This Tools's Version Information",
	Long: `Display This Tools's Version Information
    *.*.*
    | | |__  Patch Version(ex. typo,refactoring)
    | |____  Minor Version(ex. add func)
    |______  Major Version
  `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(toolVer)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
