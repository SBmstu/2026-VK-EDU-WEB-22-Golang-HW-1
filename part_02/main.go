package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"calc/calc"
)

func main() {
	var input string

	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
			os.Exit(1)
		}

		input = line
	}

	result, err := calc.Eval(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
