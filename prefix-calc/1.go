package main

import (
        "fmt"
        "strconv"
        "strings"
)

func prefixCalculator(expression string) (int, error) {
        tokens := strings.Fields(expression)
        stack := make([]int, 0)

        for i := len(tokens) - 1; i >= 0; i-- {
                token := tokens[i]

                if num, err := strconv.Atoi(token); err == nil {
                        stack = append(stack, num)
                } else {
                        if len(stack) < 2 {
                                return 0, fmt.Errorf("invalid expression: insufficient operands for operator %s", token)
                        }
                        operand1 := stack[len(stack)-1]
                        stack = stack[:len(stack)-1]
                        operand2 := stack[len(stack)-1]
                        stack = stack[:len(stack)-1]

                        switch token {
                        case "+":
                                stack = append(stack, operand1+operand2)
                        case "-":
                                stack = append(stack, operand1-operand2)
                        case "*":
                                stack = append(stack, operand1*operand2)
                        case "/":
                                if operand2 == 0 {
                                        return 0, fmt.Errorf("division by zero")
                                }
                                stack = append(stack, operand1/operand2)
                        default:
                                return 0, fmt.Errorf("invalid operator: %s", token)
                        }
                }
        }

        if len(stack) != 1 {
                return 0, fmt.Errorf("invalid expression: too many operands or too few operators")
        }

        return stack[0], nil
}

func main() {
        expressions := []string{
                "+ 5 3",
                "- 10 4",
                "* 2 6",
                "/ 8 2",
                "+ * 2 3 4",
                "- / 10 2 3",
                "+ - 10 5 * 2 3",
                "/ 5 0",
                "+ 1",
                "+ 1 2 3",
                "a 1 2",
                "+ 1 b",
        }

        for _, expr := range expressions {
                result, err := prefixCalculator(expr)
                if err != nil {
                        fmt.Printf("Expression: %s, Error: %s\n", expr, err)
                } else {
                        fmt.Printf("Expression: %s, Result: %d\n", expr, result)
                }
        }

}
