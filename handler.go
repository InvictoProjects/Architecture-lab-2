package lab2

import (
	"bytes"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(ch.Reader)
	if err != nil {
		return err
	}

	var input = strings.TrimSuffix(buffer.String(), "\n")

	result, err := PrefixToPostfix(input)
	if err != nil {
		return err
	}

	_, err = io.WriteString(ch.Writer, result)
	if err != nil {
		return err
	}

	return nil
}
