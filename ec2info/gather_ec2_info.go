package ec2info

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GatherEC2Info(resp *ec2.DescribeInstancesOutput, gatherInfoSlice []string) (string, error) {
	var result string
	var err error
	var allFlag bool
	var msg string

	//################################
	//Make Result Header Message
	//################################
	for _, getInfo := range gatherInfoSlice {
		if getInfo == "ALL" {
			allFlag = true
		}
	}
	//All Info
	if allFlag == true {
		for i, ec2Info := range ec2InfoSlice {
			if i == 0 {
				//All is skip
				continue
			} else if i == 1 {
				msg = ec2Info
			} else {
				msg = msg + "\t" + ec2Info
			}
		}
		//Other Info
	} else {
		for i, ec2Info := range gatherInfoSlice {
			if i == 0 {
				msg = ec2Info
			} else {
				msg = msg + "\t" + ec2Info
			}
		}
	}

	result = msg + "\n"

	//################################
	//Make Result Body Message
	//################################
	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			ec2Map := make(map[string]string)
			//var msg string

			//################################
			//Tag Info Input
			//################################
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					ec2Map["Tags.Name"] = *t.Value
				}
			}

			//################################
			//SecurityGroups Info Input
			//################################
			for num, sg := range i.SecurityGroups {
				if num == 0 {
					ec2Map["SecurityGroups"] = *sg.GroupName
				} else {
					ec2Map["SecurityGroups"] = ec2Map["SecurityGroups"] + "," + *sg.GroupName
				}
			}

			//################################
			//Other Info Input
			//################################
			ec2Map["InstanceId"] = *i.InstanceId

			if i.PublicIpAddress != nil {
				ec2Map["PublicIpAddress"] = *i.PublicIpAddress
				ec2Map["PublicDnsName"] = *i.PublicDnsName
			} else {
				ec2Map["PublicIpAddress"] = "-"
				ec2Map["PublicDnsName"] = "-"
			}

			if i.PrivateIpAddress != nil {
				ec2Map["PrivateIpAddress"] = *i.PrivateIpAddress
				ec2Map["PrivateDnsName"] = *i.PrivateDnsName
			} else {
				ec2Map["PrivateIpAddress"] = "-"
				ec2Map["PrivateDnsName"] = "-"
			}

			if i.VpcId != nil {
				ec2Map["VpcId"] = *i.VpcId
				ec2Map["SubnetId"] = *i.SubnetId
			} else {
				ec2Map["VpcId"] = "-"
				ec2Map["SubnetId"] = "-"
			}

			ec2Map["ImageId"] = *i.ImageId
			ec2Map["KeyName"] = *i.KeyName
			ec2Map["InstanceType"] = *i.InstanceType
			ec2Map["AvailabilityZone"] = *i.Placement.AvailabilityZone
			ec2Map["State"] = *i.State.Name
			ec2Map["RootDeviceType"] = *i.RootDeviceType
			ec2Map["VirtualizationType"] = *i.VirtualizationType
			ec2Map["LaunchTime"] = i.LaunchTime.String()

			for i, getInfo := range gatherInfoSlice {
				if getInfo == "ALL" {
					//ALL Info Case
					for j, ec2Info := range ec2InfoSlice {
						if j == 0 {
							//All is skip
							continue
						} else if j == 1 {
							msg = ec2Map[ec2Info]
						} else {
							msg = msg + "\t" + ec2Map[ec2Info]
						}
					}
					break
				} else {
					if i == 0 {
						msg = ec2Map[getInfo]
					} else {
						msg = msg + "\t" + ec2Map[getInfo]
					}
				}
			}
			result = result + msg + "\n"
		}
	}
	return result, err
}
