package main

import (
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/invictoProjects/Architecture-lab-2"
	"io"
	"os"
	"strings"
	lab2 "github.com/invictoProjects/Architecture-lab-2"
)

var inputExpression string
var inputFile string
var outputFile string

func main() {
	flag.StringVar(&inputExpression, "e", "", "Expression to compute")
	flag.StringVar(&inputFile, "f", "", "File to read expression from")
	flag.StringVar(&outputFile, "o", "", "File to write output to")
	flag.Parse()

	reader, writer, err := getReaderAndWriter()
	if err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, "error occurred: ", err)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	handler := &lab2.ComputeHandler{Reader: reader, Writer: writer}
	err = handler.Compute()
	if err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, "error occurred: ", err)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}

func getReaderAndWriter() (io.Reader, io.Writer, error) {
	if inputExpression != "" && inputFile != "" {
		return nil, nil, errors.New("both expression and file to read given")
	} else if inputExpression == "" && inputFile == "" {
		return nil, nil, errors.New("no expression is given")
	}

	var reader io.Reader
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, nil, errors.New("error reading file")
		}
		if err != nil {
			return nil, nil, err
		}
		reader = file
	} else {
		reader = strings.NewReader(inputExpression)
	}

	var writer io.Writer
	if outputFile != "" {
		file, err := os.OpenFile(outputFile, os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, nil, errors.New("error opening file to write output to")
		}
		if err != nil {
			return nil, nil, err
		}
		writer = file
	} else {
		writer = os.Stdout
	}
	return reader, writer, nil
}
