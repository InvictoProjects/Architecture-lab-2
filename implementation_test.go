package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestBob(c *C) {
	c.Assert("bob", Equals, "bob")
}

func (s *MySuite) TestPrefixToPostfixAddition(c *C) {
	res, err := PrefixToPostfix("+ 5 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 3 +")
}

func (s *MySuite) TestPrefixToPostfixSubtraction(c *C) {
	res, err := PrefixToPostfix("- 5 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 3 -")
}

func (s *MySuite) TestPrefixToPostfixMultiplication(c *C) {
	res, err := PrefixToPostfix("* 5 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 3 *")
}

func (s *MySuite) TestPrefixToPostfixDivision(c *C) {
	res, err := PrefixToPostfix("/ 5 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 3 /")
}

func (s *MySuite) TestPrefixToPostfixExponent(c *C) {
	res, err := PrefixToPostfix("^ 5 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 3 ^")
}

func (s *MySuite) TestPrefixToPostfixComplexExpressions(c *C) {
	res1, err1 := PrefixToPostfix("/ + + 2 - / 4 * 5 - 5 * 7 10 5 / 6 1 7")
	c.Assert(err1, IsNil)
	c.Assert(res1, Equals, "4 5 7 10 * - 5 * / 5 - 2 + 6 1 / + 7 /")

	res2, err2 := PrefixToPostfix("- * + 3 - 5 7 - 4 + 2 / 5 4 / + 1 * 3 7 2")
	c.Assert(err2, IsNil)
	c.Assert(res2, Equals, "5 7 - 3 + 4 5 4 / 2 + - * 3 7 * 1 + 2 / -")

	res3, err3 := PrefixToPostfix("+ 5 * - 4 2 3")
	c.Assert(err3, IsNil)
	c.Assert(res3, Equals, "4 2 - 3 * 5 +")
}

func (s *MySuite) TestPrefixToPostfixEmptyString(c *C) {
	res, err := PrefixToPostfix("")
	c.Assert(err, NotNil)
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestPrefixToPostfixInvalidCharacters(c *C) {
	res1, err1 := PrefixToPostfix("# 5 5")
	c.Assert(err1, NotNil)
	c.Assert(res1, Equals, "")

	res2, err2 := PrefixToPostfix("+ 5 * @ 4 2 3")
	c.Assert(err2, NotNil)
	c.Assert(res2, Equals, "")

	res3, err3 := PrefixToPostfix("# 1 % a @ .")
	c.Assert(err3, NotNil)
	c.Assert(res3, Equals, "")
}

func (s *MySuite) TestPrefixToPostfixWithoutOperators(c *C) {
	res1, err1 := PrefixToPostfix("1 2")
	c.Assert(err1, NotNil)
	c.Assert(res1, Equals, "")

	res2, err2 := PrefixToPostfix(" 1 2")
	c.Assert(err2, NotNil)
	c.Assert(res2, Equals, "")
}

// ExamplePrefixToPostfix is an example of using the PrefixToPostfix function
func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 5 * - 4 2 1")
	fmt.Println(res)
}
