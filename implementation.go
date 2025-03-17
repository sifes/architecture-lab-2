package lab2

import (
  "fmt"
  "math"
  "strconv"
  "strings"
)

type operationFunc func(a, b float64) float64

func EvaluatePrefix(expression string) (string, error) {
  tokens := strings.Fields(expression)
  
  var stack []float64
  
  operations := map[string]operationFunc{
    "+": func(a, b float64) float64 { return a + b },
    "-": func(a, b float64) float64 { return a - b },
    "*": func(a, b float64) float64 { return a * b },
    "/": func(a, b float64) float64 { return a / b },
    "^": func(a, b float64) float64 { return math.Pow(a, b) },
  }

  for _, token := range tokens {
    _, err := strconv.ParseFloat(token, 64)
    if err == nil {
      continue
    }
    
    if _, exists := operations[token]; !exists {
      return "", fmt.Errorf("invalid expression: unsupported symbol")
    }
  }
  
  for i := len(tokens) - 1; i >= 0; i-- {
    token := tokens[i]
    
    if num, err := strconv.ParseFloat(token, 64); err == nil {
      stack = append(stack, num)
      continue
    }
    
    if len(stack) < 2 {
      return "", fmt.Errorf("invalid expression: not enough operands")
    }
    
    a := stack[len(stack)-1]
    b := stack[len(stack)-2]
    stack = stack[:len(stack)-2]

		operation, exists := operations[token]
    if !exists {
      return "", fmt.Errorf("unknown operator: %s", token)
    }
    
    if token == "/" && b == 0 {
      return "", fmt.Errorf("division by zero")
    }
    
    result := operation(a, b)
    stack = append(stack, result)
  }
  
  if len(stack) == 0 {
    return "", fmt.Errorf("invalid expression: empty input")
  }
  
  result := stack[0]
  if math.Floor(result) == result {
    return fmt.Sprintf("%.0f", result), nil
  }
  
  return fmt.Sprintf("%.2f", result), nil
}
