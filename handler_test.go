package lab2

import (
	"bytes"
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestComputeHandlerValidInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"+ 2 3", "5\n"},
		{"* + 2 3 5", "25\n"},
		{"/ 10 2", "5\n"},
		{"+ + 4 / - * 6 2 3 2 5", "13.50\n"},
	}
	
	for _, tc := range testCases {
		var outputBuffer bytes.Buffer
		
		handler := &ComputeHandler{
			Reader: strings.NewReader(tc.input),
			Writer: &outputBuffer,
		}
		
		err := handler.Compute()
		
		assert.NoError(t, err)
		
		assert.Equal(t, tc.expected, outputBuffer.String())
	}
}

func TestComputeHandlerInvalidInput(t *testing.T) {
	testCases := []struct {
		input       string
		errorSubstr string
	}{
		{"", "empty input"},
		{"+ 2", "not enough operands"},
		{"$ 2 3", "unsupported symbol"},
		{"+ 2 invalid", "unsupported symbol"},
		{"/ 5 0", "division by zero"},
	}
	
	for _, tc := range testCases {
		var outputBuffer bytes.Buffer
		
		handler := &ComputeHandler{
			Reader: strings.NewReader(tc.input),
			Writer: &outputBuffer,
		}
		
		err := handler.Compute()
		
		assert.Error(t, err)
		
		assert.Contains(t, err.Error(), tc.errorSubstr)
		
		assert.Empty(t, outputBuffer.String())
	}
}

type ErrorReader struct{}

func (r *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, assert.AnError
}

func TestComputeHandlerReaderError(t *testing.T) {
	var outputBuffer bytes.Buffer
	
	handler := &ComputeHandler{
		Reader: &ErrorReader{},
		Writer: &outputBuffer,
	}
	
	err := handler.Compute()
	
	assert.Error(t, err)
	
	assert.Contains(t, err.Error(), "error reading input")
	
	assert.Empty(t, outputBuffer.String())
}