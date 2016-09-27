package message

import (
	"fmt"
	"os"
	"strings"
)

func PrintFooter(file string, msg string, batchFlag bool) {
	var flag string

	if batchFlag == true {
		flag = "Y"
	} else {
		fmt.Println(msg)
		fmt.Println("-------------------------------------")
		fmt.Printf("Create [%s] File?(Y/N)> ", file)
		fmt.Scanln(&flag)
	}

	if strings.ToUpper(flag) == "Y" {
		file, err := os.Create(file)

		if err != nil {
			fmt.Println("[error]File Open Error")
			panic(err)
		}

		defer file.Close()
		file.Write(([]byte)(msg))

	} else {
		return
	}
}
