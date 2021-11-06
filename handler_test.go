package lab2

import (
	. "gopkg.in/check.v1"
	"strings"
)

type WriterStub struct {
	OutString string
}

func (stub *WriterStub) Write(p []byte) (int, error) {
	stub.OutString = string(p)
	return 0, nil
}

func (s *MySuite) TestComputeHandler(c *C) {
	reader := strings.NewReader("+ 5 * - 4 2 3")
	writer := &WriterStub{}

	handler := ComputeHandler{Reader: reader, Writer: writer}
	err := handler.Compute()

	c.Assert(err, IsNil)
	c.Assert(writer.OutString, Equals, "4 2 - 3 * 5 +")
}

func (s *MySuite) TestComputeHandler_ReturnError(c *C) {
	reader := strings.NewReader("not valid expression")
	writer := &WriterStub{}

	handler := ComputeHandler{Reader: reader, Writer: writer}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
