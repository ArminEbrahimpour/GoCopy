package main

import (
	"fmt"

	"github.com/ArminEbrahimpour/GoCopy/internal/clipboard"
	"github.com/ArminEbrahimpour/GoCopy/internal/stdin"
)

func main() {

	err := stdin.ShowInputData()
	if err != nil {
		fmt.Println(err)
	}
	data, err := stdin.GetInput()

	clipboard.Save(data)

}
