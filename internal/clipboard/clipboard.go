package clipboard

import (
	"fmt"
)

type clipboard interface{}

func NewClipboard() *clipboard {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("error in initialization of clipboard")
		return nil
	}

	return &clipboard
}

func Save(data []string) {

	for _, val := range data {
		c.Write(c.FmtText, []byte(val))
	}

}

func (c *clipboard) Paste() {

	c.Read(c.FmtText)
}
