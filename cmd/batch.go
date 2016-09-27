package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "gatherec2info Batch Mode",
	Long:  `gatherec2info Batch Mode`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(batchCmd)

	//RegionId Setting
	batchCmd.PersistentFlags().StringP("region", "r", "ap-northeast-1", "AWS Region(default: ap-northeast-1)")
	viper.BindPFlag("regionId", batchCmd.PersistentFlags().Lookup("region"))

	//ProfileId Setting
	batchCmd.PersistentFlags().StringP("profile", "p", "", "Profile Id(ex. sample)")
	viper.BindPFlag("profileId", batchCmd.PersistentFlags().Lookup("profile"))

	//OutputFile Setting
	batchCmd.PersistentFlags().StringP("output", "o", "", "OutputFile PATH(default: ${ProfileId}_EC2.tsv)")
	viper.BindPFlag("outputFile", batchCmd.PersistentFlags().Lookup("output"))

}
