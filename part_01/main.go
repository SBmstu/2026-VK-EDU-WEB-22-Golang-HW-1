package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"uniq/types"
	"uniq/uniq"
	"uniq/utils"
)

var options types.Options

func init() {
	initFlags()
}

func initFlags() {
	flag.BoolVar(&options.Count, types.FlagCount, false, types.DescCount)
	flag.BoolVar(&options.Repeated, types.FlagRepeated, false, types.DescRepeated)
	flag.BoolVar(&options.Unique, types.FlagUnique, false, types.DescUnique)
	flag.BoolVar(&options.IgnoreCase, types.FlagIgnoreCase, false, types.DescIgnoreCase)
	flag.IntVar(&options.SkipFields, types.FlagSkipFields, 0, types.DescSkipFields)
	flag.IntVar(&options.SkipChars, types.FlagSkipChars, 0, types.DescSkipChars)
	flag.Parse()
}
 
func main() {
	var options types.Options

	if err := utils.ValidateOptions(options); err != nil {
		fmt.Fprintf(os.Stderr, "Error! Validate options: %v\n", err)

		flag.Usage()

		os.Exit(1)
	}

	args := flag.Args()
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error! Opening input file: %v\n", err)
			os.Exit(1)
		}

		defer f.Close()

		input = f
	}

	if len(args) > 1 {
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error! Creating output file: %v\n", err)
			os.Exit(1)
		}

		defer f.Close()

		output = f
	}

	lines, err := readLines(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading: %v\n", err)
		os.Exit(1)
	}

	result, err := uniq.Uniq(lines, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error! Uniq(): %v\n", err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(output)

	defer writer.Flush()

	for _, line := range result {
		fmt.Fprintln(writer, line)
	}
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
