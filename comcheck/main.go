package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	arr := []string{"01", "galaxy"}

	for _, v := range args {
		if v == arr[0] || v == arr[1] || v == arr[0]+" "+arr[1] || v == arr[1]+" "+arr[0] {
			fmt.Println("Alert!!!")
			break
		}
	}
}
