package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ArminEbrahimpour/GoCopy/internal/clipboard"
	"github.com/ArminEbrahimpour/GoCopy/internal/stdin"
)

func main() {

	var copy bool = false
	var paste bool = false

	flag.BoolVar(&copy, "c", false, "copy to the clipboard")
	flag.BoolVar(&paste, "p", false, "paste to the clipboard")

	flag.Parse()

	/*
		err := stdin.ShowInputData()
		if err != nil {
			fmt.Println(err)
		}
	*/
	data, err := stdin.GetInput()

	if err != nil {
		fmt.Println(err)
	}
	clipboard.Start()

	if copy {
		clipboard.Save(data)
	}

	if paste {
		clipboard.Paste()
	}

	// making the process proceed
	time.Sleep(2 * time.Minute)
}
