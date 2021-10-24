package lab2

import "fmt"


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

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	input += " "
	var nodeStack []*node
	var prevElem string
	elem := ""
	for _, char := range input {
		if string(char) == " " {
			if prevElem == "" {
				if isOperator(elem) {
					node := node{OPERATOR, elem, nil, nil, nil}
					nodeStack = append(nodeStack, &node)
				} else {
					return "", fmt.Errorf("The prefix notation expression was not found")
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
		elem += string(char)
	}

	var result string
	isVisited := make(map[*node]bool)
	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		if (node.left == nil || isVisited[node.left]) &&
			(node.right == nil || isVisited[node.right]) {
			result += node.value
			result += " "
			nodeStack = nodeStack[:len(nodeStack)-1]
			isVisited[node] = true
		} else {
			if isVisited[node.left] {
				nodeStack = append(nodeStack, node.right)
			} else {
				nodeStack = append(nodeStack, node.left)
			}
		}
	}

	return result, nil
}

func isOperator(s string) bool {
	flag := false
	switch s {
	case "+":
		flag = true
	case "-":
		flag = true
	case "*":
		flag = true
	case "/":
		flag = true
	default:
		flag = false
	}
	return flag
}
