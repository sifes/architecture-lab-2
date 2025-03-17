package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sifes/architecture-lab-2"
)

func main() {
	expressionFlag := flag.String("e", "", "Expression to evaluate")
	inputFileFlag := flag.String("f", "", "File containing the expression to evaluate")
	outputFileFlag := flag.String("o", "", "File to write the result to (optional)")

	flag.Parse()

	if (*expressionFlag != "" && *inputFileFlag != "") || (*expressionFlag == "" && *inputFileFlag == "") {
		fmt.Fprintln(os.Stderr, "Error: You must specify either -e or -f, but not both")
		flag.Usage()
		os.Exit(1)
	}

	var reader io.Reader
	if *expressionFlag != "" {
		reader = strings.NewReader(*expressionFlag)
	} else {
		file, err := os.Open(*inputFileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	var writer io.Writer = os.Stdout
	if *outputFileFlag != "" {
		file, err := os.Create(*outputFileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	}

	handler := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
