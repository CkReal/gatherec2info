package common

import (
	"fmt"
	"strconv"
)

var regionId string

func GetRegionId() (string, error) {
	var getId string
	var err error

	//Tokyo is Default Region
	result := "ap-northeast-1"

	array := []string{
		"default",
		"US East (N. Virginia)",
		"US West (Oregon)",
		"US West (N. California)",
		"EU (Ireland)",
		"EU (Frankfurt)",
		"Asia Pacific (Singapore)",
		"Asia Pacific (Tokyo)",
		"Asia Pacific (Sydney)",
		"Asia Pacific (Seoul)",
		"Asia Pacific (Mumbai)",
		"South America (São Paulo)",
	}

	for i, r := range array {
		fmt.Printf("[%d] %s\n", i, r)
	}

	fmt.Print("Input Region Id(default:Tokyo)> ")
	fmt.Scanln(&getId)
	fmt.Println("-------------------------------------")

	//No Input
	if getId == "" {
		return result, err
	}

	//Error Handling
	id, err := strconv.Atoi(getId)
	if err != nil {
		return result, fmt.Errorf("[error]Region Id is 0-%d.\n", len(array)-1)
	}

	switch array[id] {
	case "US East (N. Virginia)":
		return "us-east-1", err
	case "US West (Oregon)":
		return "us-west-2", err
	case "US West (N. California)":
		return "us-west-1", err
	case "EU (Ireland)":
		return "eu-west-1", err
	case "EU (Frankfurt)":
		return "eu-central-1", err
	case "Asia Pacific (Singapore)":
		return "ap-southeast-1", err
	case "Asia Pacific (Tokyo)":
		return "ap-northeast-1", err
	case "Asia Pacific (Sydney)":
		return "ap-southeast-2", err
	case "Asia Pacific (Seoul)":
		return "ap-northeast-2", err
	case "Asia Pacific (Mumbai)":
		return "ap-south-1", err
	case "South America (São Paulo)":
		return "sa-east-1", err
	default:
		return result, err
	}
}
