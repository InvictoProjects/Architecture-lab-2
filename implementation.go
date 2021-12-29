package lab2

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type TypeOfNode int

const (
	OPERATOR TypeOfNode = iota
	OPERAND
)

type node struct {
	typeOfNode TypeOfNode
	value      string
	parent     *node
	left       *node
	right      *node
}

// PrefixToPostfix converts an input expression from a prefix notation
// to a postfix entry. First, function validates the input expression.
// If an input string is not a prefix notation then error occurred. In
// general, the function builds a binary tree of the expression, and
// then forms the result.
func PrefixToPostfix(input string) (string, error) {
	isValid := validate(input)
	if isValid != true {
		return "", fmt.Errorf("the input expression contains not allowed symbols")
	}

	root, err := buildExpressionTree(input)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	postfix := getPostfixFromTree(root)
	return postfix, nil
}

func buildExpressionTree(expression string) (*node, error) {
	expression += " "
	var nodeStack []*node
	var prevElem, elem string
	for _, code := range expression {
		char := string(code)
		if char == " " {
			if prevElem == "" {
				if isOperator(elem) {
					node := node{OPERATOR, elem, nil, nil, nil}
					nodeStack = append(nodeStack, &node)
				} else {
					return nil, errors.New("the prefix notation expression was not found")
				}
			} else if isOperator(elem) && isOperator(prevElem) {
				parent := nodeStack[len(nodeStack)-1]
				node := node{OPERATOR, elem, parent, nil, nil}
				parent.left = &node
				nodeStack = append(nodeStack, &node)
			} else if !isOperator(elem) && isOperator(prevElem) {
				parent := nodeStack[len(nodeStack)-1]
				node := node{OPERAND, elem, parent, nil, nil}
				parent.left = &node
				nodeStack = append(nodeStack, &node)
			} else if isOperator(elem) && !isOperator(prevElem) {
				parent := nodeStack[len(nodeStack)-2]
				node := node{OPERATOR, elem, parent, nil, nil}
				parent.right = &node
				nodeStack = append(nodeStack, &node)
			} else {
				parent := nodeStack[len(nodeStack)-2]
				node := node{OPERAND, elem, parent, nil, nil}
				parent.right = &node
				nodeStack = append(nodeStack, &node)
				flag := true
				for flag {
					if len(nodeStack) < 2 {
						flag = false
						continue
					}
					left := nodeStack[len(nodeStack)-2]
					right := nodeStack[len(nodeStack)-1]
					if left.parent == right.parent {
						nodeStack = nodeStack[:len(nodeStack)-2]
					} else {
						flag = false
					}
				}
			}
			prevElem = elem
			elem = ""
			continue
		}
		elem += char
	}
	return nodeStack[0], nil
}

func isOperator(s string) bool {
	flag := false
	switch s {
	case "+", "-", "*", "/", "^":
		flag = true
	}
	return flag
}

func validate(expression string) bool {
	matched, _ := regexp.MatchString(`^[\d\s+\-*/^]+$`, expression)
	return matched
}

func getPostfixFromTree(root *node) string {
	var result strings.Builder
	nodeStack := []*node{root}
	isVisited := make(map[*node]bool)
	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		if (node.left == nil || isVisited[node.left]) &&
			(node.right == nil || isVisited[node.right]) {
			result.WriteString(node.value)
			result.WriteString(" ")
			nodeStack = nodeStack[:len(nodeStack)-1]
			isVisited[node] = true
		} else {
			if !isVisited[node.left] && !isVisited[node.right] {
				if node.value == "+" || node.value == "*" {
					if !isOperator(node.left.value) && isOperator(node.right.value) {
						nodeStack = append(nodeStack, node.right)
						continue
					}
				}
			}
			if isVisited[node.left] {
				nodeStack = append(nodeStack, node.right)
			} else {
				nodeStack = append(nodeStack, node.left)
			}
		}
	}

	return strings.TrimSpace(result.String())
}
