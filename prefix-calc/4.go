package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// isOperator checks if a token is a valid operator.
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// isNumber checks if a token is a valid number.
func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

// evaluate evaluates a prefix expression.
func evaluate(tokens []string) (int, error) {
	if len(tokens) == 0 {
		return 0, errors.New("empty expression")
	}

	// Pop the first token from the list.
	token := tokens[0]
	tokens = tokens[1:]

	if isNumber(token) {
		// If the token is a number, return it as an integer.
		return strconv.Atoi(token)
	} else if isOperator(token) {
		// If the token is an operator, recursively evaluate the next two operands.
		lhs, err := evaluate(tokens)
		if err != nil {
			return 0, err
		}

		rhs, err := evaluate(tokens)
		if err != nil {
			return 0, err
		}

		// Perform the operation.
		switch token {
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
			return 0, fmt.Errorf("invalid operator: %s", token)
		}
	} else {
		// If the token is neither a number nor an operator, return an error.
		return 0, fmt.Errorf("invalid token: %s", token)
	}
}

// prefixCalculate parses and evaluates a prefix expression.
func prefixCalculate(input string) (int, error) {
	// Split the input string into tokens.
	tokens := strings.Fields(input)

	// Evaluate the prefix expression.
	return evaluate(tokens)
}

func main() {
	// Test cases.
	expressions := []string{
		"+ 1 2",               // 1 + 2 = 3
		"- 10 5",              // 10 - 5 = 5
		"* 3 4",               // 3 * 4 = 12
		"/ 20 4",              // 20 / 4 = 5
		"+ * 2 3 4",           // (2 * 3) + 4 = 10
		"/ + 10 2 3",          // (10 + 2) / 3 = 4
		"+ 100 200",           // 100 + 200 = 300
		"/ 10 0",              // Division by zero
		"* abc 10",            // Invalid token
		"+ 1 2 3",             // Extra tokens
		"",                    // Empty input
	}

	for _, expr := range expressions {
		result, err := prefixCalculate(expr)
		if err != nil {
			fmt.Printf("Error evaluating '%s': %v\n", expr, err)
		} else {
			fmt.Printf("Result of '%s' = %d\n", expr, result)
		}
	}
}
