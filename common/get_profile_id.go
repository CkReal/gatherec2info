package common

import (
	"fmt"
)

func GetProfileId() (string, error) {
	var getId string
	var result string
	var err error

	fmt.Print("Input Profile Id(ex. sample)> ")
	fmt.Scanln(&getId)
	fmt.Println("-------------------------------------")

	//Error Handling
	if getId == "" {
		return result, fmt.Errorf("[error]Input Profile Id.")
	}

	result = getId
	return result, err
}
