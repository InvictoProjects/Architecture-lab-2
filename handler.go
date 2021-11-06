package lab2

import (
	"bytes"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	reader io.Reader
	writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	return nil
}
