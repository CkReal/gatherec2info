package common

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func GetInfoSlice(msg string, infoSlice []string) []string {
	var infoCSV string
	idSlice := []string{}
	result := []string{}

	for i, j := range infoSlice {
		fmt.Printf("[%d] %s\n", i, j)
	}

	fmt.Print(msg)
	fmt.Scanln(&infoCSV)
	fmt.Println("-------------------------------------")

	reader := csv.NewReader(strings.NewReader(infoCSV))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		idSlice = record
	}

	//Make Result Information
	for _, i := range idSlice {
		id, err := strconv.Atoi(i)

		if err != nil {
			fmt.Printf("[error]Input Information Number\n")
		}

		result = append(result, infoSlice[id])
	}

	return result
}
