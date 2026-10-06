package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"calc/calc"
)

const (
	codeSuccess = 0
	codeFailure = 1
)


func main() {
	os.Exit(process())
}

func process() int {
	var input string

	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			return codeFailure
		}

		input = string(data)
	}

	result, err := calc.Eval(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return codeFailure
	}

	fmt.Println(result)

	return codeSuccess
}
