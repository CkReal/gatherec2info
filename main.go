package main

import (
	"fmt"
	"os"

	"github.com/CkReal/gatherec2info/cmd"
	"github.com/CkReal/gatherec2info/common"
	"github.com/CkReal/gatherec2info/ec2info"
	"github.com/CkReal/gatherec2info/message"

	"github.com/spf13/viper"
)

func main() {
	var batchFlag bool
	var result string

	//####################################
	//Check Args
	//####################################
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if len(os.Args) != 1 {
		if os.Args[1] == "help" || os.Args[1] == "--help" {
			return
		} else if os.Args[1] == "version" {
			return
		} else if os.Args[1] == "batch" {
			if len(os.Args) == 2 {
				err :=
					"Input Profile Id and Output File PATH.\n" +
						">BatchMode Help\n" +
						"$ gatherec2info batch --help"
				fmt.Println(err)
				return
			} else {
				if os.Args[2] == "help" || os.Args[2] == "--help" {
					return
				} else {
					batchFlag = true
				}
			}
		}
	}

	//####################################
	//Check Args for Batch Mode
	//####################################
	regionId := viper.GetString("regionId")
	profileId := viper.GetString("profileId")
	outputFile := viper.GetString("outputFile")

	if batchFlag == true {
		if profileId == "" || outputFile == "" {
			err :=
				"Input Profile Id and Output File PATH.\n" +
					">BatchMode Help\n" +
					"$ gatherec2info batch --help"
			fmt.Println(err)
			return
		}
	}

	//####################################
	//Display Start Message
	//####################################
	if batchFlag == false {
		message.PrintHeader()
	}

	//####################################
	//Input Profile Id
	//####################################
	if batchFlag == false {
		var err error
		profileId, err = common.GetProfileId()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//####################################
	//Input Region Id
	//####################################
	if batchFlag == false {
		var err error
		regionId, err = common.GetRegionId()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//####################################
	//Get Session
	//####################################
	sess, err := common.GetSession(profileId, regionId)

	if err != nil {
		fmt.Println(err)
		return
	}

	//####################################
	//Gather EC2 Information
	//####################################
	result, err = ec2info.ExecEC2(sess, batchFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	//####################################
	//Output Result File
	//####################################
	if batchFlag == false {
		outputFile = profileId + "_EC2.tsv"
	}
	message.PrintFooter(outputFile, result, batchFlag)
}
