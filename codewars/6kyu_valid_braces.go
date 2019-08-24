package main

import (
	"strings"
)

// PARENMAP ....
var PARENMAP = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

const (
	// OPENPAREN ....
	OPENPAREN = "({["
	// CLOSEPAREN ....
	CLOSEPAREN = ")}]"
)

// Stack ________________________________________
type Stack struct {
	top  *element
	size int
}

type element struct {
	value interface{}
	next  *element
}

// Len ....
func (stack *Stack) Len() int {
	return stack.size
}

// Push ....
func (stack *Stack) Push(value interface{}) {
	stack.top = &element{value, stack.top}
	stack.size++
}

// Pop ....
func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

// ValidBraces ________________________________________
func ValidBraces(paren string) (rc bool) {
	stack := new(Stack)
	if len(paren) > 0 {
		rc = true
		for _, p := range strings.Split(paren, "") {
			if strings.Index(OPENPAREN, p) != -1 {
				stack.Push(p)
				continue
			}

			if strings.Index(CLOSEPAREN, p) != -1 {
				if stack.Len() > 0 {
					lastParen := stack.Pop().(string)
					if PARENMAP[lastParen] == p {
						continue
					}
				}
				rc = false
				break
			}
		}
	}

	if rc && stack.Len() == 0 {
		rc = true
	} else {
		rc = false
	}
	return
}
