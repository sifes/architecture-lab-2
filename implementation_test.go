package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    string
	expected string
	hasError bool
	errMsg   string
}

func TestEvaluatePrefix(t *testing.T) {
	tests := []testCase{
		{name: "Addition", input: "+ 3 4", expected: "7", hasError: false, errMsg: ""},
		{name: "Subtraction", input: "- 10 4", expected: "6", hasError: false, errMsg: ""},
		{name: "Multiplication", input: "* 3 4", expected: "12", hasError: false, errMsg: ""},
		{name: "Division", input: "/ 8 2", expected: "4", hasError: false, errMsg: ""},
		{name: "Power", input: "^ 2 3", expected: "8", hasError: false, errMsg: ""},
		
		{name: "Expression 1", input: "* + 6 2 3", expected: "24", hasError: false, errMsg: ""},
		{name: "Expression 2", input: "+ * - 7 3 4 2", expected: "18", hasError: false, errMsg: ""},
		{name: "Expression with power", input: "* + 3 ^ 2 3 4", expected: "44", hasError: false, errMsg: ""},
		{name: "Large expression", input: "* + + 4 / - * 6 2 3 - 2 5 2 2", expected: "6", hasError: false, errMsg: ""},
		
		{name: "Float addition", input: "+ 1.5 2.5", expected: "4", hasError: false, errMsg: ""},
		{name: "Float multiplication", input: "* 2.5 4", expected: "10", hasError: false, errMsg: ""},
		{name: "Float division", input: "/ 5.5 2", expected: "2.75", hasError: false, errMsg: ""},
		
		{name: "Empty input", input: "", expected: "", hasError: true, errMsg: "invalid expression: empty input"},
		{name: "Not enough operands", input: "+ 2", expected: "", hasError: true, errMsg: "invalid expression: not enough operands"},
		{name: "Unsupported symbol", input: "+ 2 a", expected: "", hasError: true, errMsg: "invalid expression: unsupported symbol"},
		{name: "Double operator", input: "-- + 3 2", expected: "", hasError: true, errMsg: "invalid expression: unsupported symbol"},
		{name: "Division by zero", input: "/ 4 0", expected: "", hasError: true, errMsg: "division by zero"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := EvaluatePrefix(tc.input)

			if tc.hasError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func ExampleEvaluatePrefix() {
	res, _ := EvaluatePrefix("+ 2 2")
	fmt.Println(res)
	// Output:
	// 4
}