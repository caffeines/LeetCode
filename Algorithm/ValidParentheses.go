package main

import (
	"fmt"
)

var stack []string

func push(elem string) {
	stack = append(stack, elem)
}

func pop() string {
	var top string
	if len(stack) > 0 {
		n := len(stack) - 1
		top = stack[n] // Top element
		stack[n] = ""
		stack = stack[:n] // Pop
	}
	return top
}

func top() string {
	var top string
	if len(stack) > 0 {
		n := len(stack) - 1
		top = stack[n] // Top element
	}
	return top
}

func isValid(s string) bool {
	for _, c := range s {
		_s := string(c)
		matched := false
		if _s == ")" && top() == "(" {
			pop()
			matched = true
		}
		if _s == "}" && top() == "{" {
			pop()
			matched = true
		}
		if _s == "]" && top() == "[" {
			pop()
			matched = true
		}
		if !matched {
			push(_s)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{[]}"))
}
