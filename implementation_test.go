package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePrefix(t *testing.T) {
	result, err := EvaluatePrefix("+ 3 4")
	assert.NoError(t, err)
	assert.Equal(t, "7", result)
}

func TestEvaluatePrefixPower(t *testing.T) {
	result, err := EvaluatePrefix("^ 2 4")
	assert.NoError(t, err)
	assert.Equal(t, "16", result)
}

func TestEvaluatePrefixMedium(t *testing.T) {
	result, err := EvaluatePrefix("* + 6 2 3")
	assert.NoError(t, err)
	assert.Equal(t, "24", result)
}

func TestEvaluatePrefixComplex(t *testing.T) {
	result, err := EvaluatePrefix("+ * - 7 3 4 2")
	assert.NoError(t, err)
	assert.Equal(t, "18", result)
}

func TestEvaluatePrefixComplexPower(t *testing.T) {
	result, err := EvaluatePrefix("* + 3 ^ 2 3 4")
	assert.NoError(t, err)
	assert.Equal(t, "44", result)
}

func TestEvaluatePrefixBig(t *testing.T) {
	result, err := EvaluatePrefix("* + + 4 / - * 6 2 3 - 2 5 2 2")
	assert.NoError(t, err)
	assert.Equal(t, "6", result)
}

func TestEvaluatePrefixEmptyError(t *testing.T) {
	_, err := EvaluatePrefix("")
	assert.Error(t, err)
	assert.Equal(t, "invalid expression: empty input", err.Error())
}

func TestEvaluatePrefixOperandsError(t *testing.T) {
	_, err := EvaluatePrefix("+ 2")
	assert.Error(t, err)
	assert.Equal(t, "invalid expression: not enough operands", err.Error())
}

func TestEvaluatePrefixSymbolError1(t *testing.T) {
	_, err := EvaluatePrefix("-- + 3 2")
	assert.Error(t, err)
	assert.Equal(t, "invalid expression: unsupported symbol", err.Error())
}

func TestEvaluatePrefixSymbolError2(t *testing.T) {
	_, err := EvaluatePrefix("+ 2 a")
	assert.Error(t, err)
	assert.Equal(t, "invalid expression: unsupported symbol", err.Error())
}

func ExampleEvaluatePrefix() {
	res, _ := EvaluatePrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 4
}