package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case "+":
			n1 := pop(&stack)
			n2 := pop(&stack)
			n3 := n2 + n1
			stack = append(stack, n3)
		case "-":
			n1 := pop(&stack)
			n2 := pop(&stack)
			n3 := n2 - n1
			stack = append(stack, n3)
		case "*":
			n1 := pop(&stack)
			n2 := pop(&stack)
			n3 := n2 * n1
			stack = append(stack, n3)
		case "/":
			n1 := pop(&stack)
			n2 := pop(&stack)
			n3 := n2 / n1
			stack = append(stack, n3)
		default:
			intVar, _ := strconv.Atoi(token)
			stack = append(stack, intVar)
		}
	}

	return stack[0]
}

func pop(s *[]int) int {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return n
}
