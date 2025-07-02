package clipboard

import (
	"fmt"

	"golang.design/x/clipboard"
)

func Save(data []string) {

	err := clipboard.Init()
	if err != nil {
		fmt.Printf("error in initialization of clipboard: %s", err)
		return
	}

	for _, val := range data {
		clipboard.Write(clipboard.FmtText, []byte(val))
	}

}

func Paste() {

	err := clipboard.Init()
	if err != nil {
		fmt.Printf("error in initialization of clipboard: %s", err)
		return
	}

	clipboard.Read(clipboard.FmtText)
}
