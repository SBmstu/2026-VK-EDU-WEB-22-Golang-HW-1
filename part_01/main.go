package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"uniq/uniq"
)

func main() {
	var options uniq.Options
	flag.BoolVar(&options.Count, "c", false, "count repeating lines")
	flag.BoolVar(&options.Repeated, "d", false, "show only repeated lines")
	flag.BoolVar(&options.Unique, "u", false, "show only unique lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.IntVar(&options.SkipFields, "f", 0, "skip first N fields")
	flag.IntVar(&options.SkipChars, "s", 0, "skip first N chars")
	flag.Parse()

	if (options.Count && options.Repeated) || (options.Count && options.Unique) || (options.Repeated && options.Unique) {
		fmt.Fprintln(os.Stderr, "Error! Options -c, -d, -u are mutually exclusive")

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
		fmt.Fprintf(os.Stderr, "Error! %v\n", err)
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
