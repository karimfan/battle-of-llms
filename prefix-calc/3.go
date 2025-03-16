package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// PrefixCalculator evaluates mathematical expressions in prefix notation.
// It handles the basic operations: +, -, *, / and properly handles multi-digit numbers.
// Example: "+ 2 3" evaluates to 5, "* + 2 3 4" evaluates to 20
type PrefixCalculator struct {
	expression string
	position   int
	tokens     []string
}

// NewPrefixCalculator creates a new calculator with the given expression
func NewPrefixCalculator(expression string) *PrefixCalculator {
	return &PrefixCalculator{
		expression: expression,
		position:   0,
	}
}

// Tokenize breaks the expression into tokens (operators and numbers)
func (pc *PrefixCalculator) Tokenize() error {
	pc.tokens = []string{}
	
	// Trim any leading/trailing whitespace
	expr := strings.TrimSpace(pc.expression)
	if expr == "" {
		return errors.New("empty expression")
	}
	
	i := 0
	for i < len(expr) {
		// Skip whitespace
		if unicode.IsSpace(rune(expr[i])) {
			i++
			continue
		}
		
		// Handle operators
		if isOperator(string(expr[i])) {
			pc.tokens = append(pc.tokens, string(expr[i]))
			i++
			continue
		}
		
		// Handle numbers (including multi-digit)
		if unicode.IsDigit(rune(expr[i])) {
			start := i
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				i++
			}
			pc.tokens = append(pc.tokens, expr[start:i])
			continue
		}
		
		// If we get here, there's an invalid character
		return fmt.Errorf("invalid character at position %d: '%c'", i, expr[i])
	}
	
	return nil
}

// Calculate evaluates the prefix expression and returns the result
func (pc *PrefixCalculator) Calculate() (int, error) {
	if err := pc.Tokenize(); err != nil {
		return 0, err
	}
	
	pc.position = 0
	result, err := pc.evaluateExpression()
	if err != nil {
		return 0, err
	}
	
	// Check if we consumed all tokens
	if pc.position < len(pc.tokens) {
		return 0, errors.New("extra tokens in expression")
	}
	
	return result, nil
}

// evaluateExpression recursively evaluates a prefix expression
func (pc *PrefixCalculator) evaluateExpression() (int, error) {
	if pc.position >= len(pc.tokens) {
		return 0, errors.New("unexpected end of expression")
	}
	
	token := pc.tokens[pc.position]
	pc.position++
	
	// If token is a number, return its value
	if !isOperator(token) {
		return strconv.Atoi(token)
	}
	
	// Evaluate left operand
	left, err := pc.evaluateExpression()
	if err != nil {
		return 0, err
	}
	
	// Evaluate right operand
	right, err := pc.evaluateExpression()
	if err != nil {
		return 0, err
	}
	
	// Perform operation
	switch token {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, errors.New("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", token)
	}
}

// isOperator checks if a token is an operator
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// CalculatePrefix is a convenience function to calculate a prefix expression
func CalculatePrefix(expression string) (int, error) {
	calculator := NewPrefixCalculator(expression)
	return calculator.Calculate()
}

func main() {
	// Test expressions
	expressions := []string{
		"+ 2 3",          // 5
		"* + 2 3 4",      // 20
		"/ + 4 5 3",      // 3
		"- + 8 / 6 3 2",  // 8
		"+ 10 * 2 3",     // 16
		"* + 2 3 - 5 1",  // 20
		"/ 8 + 1 1",      // 4
		"/ 6 0",          // Error: division by zero
		"+ 1",            // Error: unexpected end of expression
		"+ 1 2 3",        // Error: extra tokens in expression
		"& 1 2",          // Error: invalid character
	}
	
	for _, expr := range expressions {
		result, err := CalculatePrefix(expr)
		if err != nil {
			fmt.Printf("Error evaluating '%s': %v\n", expr, err)
		} else {
			fmt.Printf("'%s' = %d\n", expr, result)
		}
	}
}
