package stdin

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() ([]string, error) {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return data, err
	}
	return data, nil
}

func ShowInputData() error {

	data, err := GetInput()
	if err != nil {
		fmt.Printf("Error from getting input from pipeline: %s", err)
		return err
	}

	for _, val := range data {
		fmt.Println(val)
	}

	return nil
}
