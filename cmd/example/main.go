package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	//lab2 "github.com/invictoProjects/architecture-lab-2"
	"io"
	"os"
)

var inputExpression string
var inputFile string
var outputFile string

func main() {
	flag.StringVar(&inputExpression, "e", "", "Expression to compute")
	flag.StringVar(&inputFile, "f", "", "File to read expression from")
	flag.StringVar(&outputFile, "o", "", "File to write output to")
	flag.Parse()

	reader, writer, error := getReaderAndWriter()
	if error != nil {
		fmt.Fprintln(os.Stderr, "error occured: ", error)
		return
	}

	b := make([]byte, 8)
	reader.Read(b)
	writer.Write(b)
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//handler := &lab2.ComputeHandler{
	//	reader: readFrom,
	//	writer:  writeTo,
	//	}
	//	err := handler.Compute()
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
		err = file.Close()
		if err != nil {
			return nil, nil, err
		}
		reader = file
	} else {
		reader = strings.NewReader(inputExpression)
	}

	var writer io.Writer
	if outputFile != "" {
		file, err := os.Open(outputFile)
		if err != nil {
			return nil, nil, errors.New("error opening file to write output to")
		}
		err = file.Close()
		if err != nil {
			return nil, nil, err
		}
		writer = file
	} else {
		writer = os.Stdout
	}
	return reader, writer, nil
}
