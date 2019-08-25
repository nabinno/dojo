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

// ValidBracesStack ________________________________________
type ValidBracesStack struct {
	top  *validBracesElement
	size int
}

type validBracesElement struct {
	value interface{}
	next  *validBracesElement
}

// Len ....
func (stack *ValidBracesStack) Len() int {
	return stack.size
}

// Push ....
func (stack *ValidBracesStack) Push(value interface{}) {
	stack.top = &validBracesElement{value, stack.top}
	stack.size++
}

// Pop ....
func (stack *ValidBracesStack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

// ValidBraces ________________________________________
func ValidBraces(paren string) (rc bool) {
	stack := new(ValidBracesStack)
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
