package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		filename := os.Args[1]
		file, error := os.ReadFile(filename)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Print(string(file))
	}
}
