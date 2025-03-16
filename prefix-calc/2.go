package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// prefixCalculate evaluates a prefix expression and returns the result.
func prefixCalculate(expression string) (int, error) {
	tokens := strings.Fields(expression)
	stack := make([]int, 0)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isNumber(token) {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			}
		} else {
			return 0, errors.New("invalid token")
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

// isNumber checks if a token is a number.
func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

// isOperator checks if a token is an operator.
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func main() {
	expression := "/ + 4 5 9"
	result, err := prefixCalculate(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
