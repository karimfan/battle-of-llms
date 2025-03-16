package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// evaluatePrefix evaluates a prefix expression using a stack-based approach.
func evaluatePrefix(expression string) (int, error) {
	// Split input into tokens
	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, errors.New("empty expression")
	}

	// Reverse iterate over tokens and use a stack
	stack := []int{}

	// Process tokens in reverse order
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		// Check if the token is an operator
		if isOperator(token) {
			// Ensure we have at least two operands to apply the operator
			if len(stack) < 2 {
				return 0, errors.New("malformed expression: not enough operands")
			}

			// Pop two operands from the stack
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform the operation
			result, err := applyOperator(token, op1, op2)
			if err != nil {
				return 0, err
			}

			// Push result back onto stack
			stack = append(stack, result)

		} else {
			// Convert token to integer
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, errors.New("invalid token: " + token)
			}

			// Push number onto stack
			stack = append(stack, num)
		}
	}

	// If there's more than one number left in the stack, the expression was malformed
	if len(stack) != 1 {
		return 0, errors.New("malformed expression: too many operands")
	}

	return stack[0], nil
}

// isOperator checks if a token is a valid operator
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// applyOperator performs an arithmetic operation
func applyOperator(op string, lhs, rhs int) (int, error) {
	switch op {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if rhs == 0 {
			return 0, errors.New("division by zero")
		}
		return lhs / rhs, nil
	default:
		return 0, errors.New("unknown operator: " + op)
	}
}

func main() {
	// Test cases
	expressions := []string{
		"+ 3 4",             // 3 + 4 = 7
		"- * 10 2 5",        // (10 * 2) - 5 = 15
		"/ + 6 2 4",         // (6 + 2) / 4 = 2
		"* + 4 5 - 9 3",     // (4 + 5) * (9 - 3) = 9 * 6 = 54
		"- + 7 / 6 3 2",     // (7 + (6 / 3)) - 2 = 7 + 2 - 2 = 7
		"- + 10 5 3",        // (10 + 5) - 3 = 12
		"+ 10 * 2 3",        // 10 + (2 * 3) = 10 + 6 = 16
		"/ 8 0",             // Division by zero error
		"* + 2 3 -",         // Malformed expression
	}

	// Run the prefix calculator for each test case
	for _, expr := range expressions {
		result, err := evaluatePrefix(expr)
		if err != nil {
			fmt.Printf("Expression: %s -> Error: %s\n", expr, err)
		} else {
			fmt.Printf("Expression: %s -> Result: %d\n", expr, result)
		}
	}
}

