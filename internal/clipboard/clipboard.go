package clipboard

import (
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

func Start() {

	err := clipboard.Init()

	if err != nil {
		fmt.Println("error in initialization of clipboard: %s", err)
		return
	}
}

func Save(data []string) {
	/*
		for _, val := range data {
			what := clipboard.Write(clipboard.FmtText, []byte(val))
			fmt.Println(what)
		}
	*/
	joined := strings.Join(data, "\n")
	clipboard.Write(clipboard.FmtText, []byte(joined))

}

func Paste() {
	data := clipboard.Read(clipboard.FmtText)
	fmt.Println(string(data))
}
