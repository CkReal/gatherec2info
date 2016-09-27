package ec2info

import (
	"fmt"
	"github.com/CkReal/gatherec2info/common"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ExecEC2(sess *session.Session, batchFlag bool) (string, error) {
	var result string
	var err error
	var inputSlice []string

	ec2DescribeInstances, err := Initialize(sess)

	if err != nil {
		return result, fmt.Errorf("[error]Failed Gather EC2 Information.")
	}

	//####################################
	//Input Information
	//####################################
	if batchFlag == true {
		//batch mode -> All Information Output
		inputSlice = append(inputSlice, "ALL")
	} else {
		inputSlice = common.GetInfoSlice(
			"Input Request Information,comma delimited(ex. 1,2) > ",
			ec2InfoSlice,
		)
	}

	//Gathering EC2 Information
	result, err = GatherEC2Info(ec2DescribeInstances, inputSlice)

	if err != nil {
		return result, err
	}

	return result, err
}
