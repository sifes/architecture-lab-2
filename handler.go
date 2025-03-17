package lab2

import (
  "fmt"
  "io"
  "io/ioutil"
  "strings"
)

type ComputeHandler struct {
  Reader io.Reader
  Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
  data, err := ioutil.ReadAll(ch.Reader)
  if err != nil {
    return fmt.Errorf("error reading input: %w", err)
  }
  expression := strings.TrimSpace(string(data))
  result, err := EvaluatePrefix(expression)
  if err != nil {
    return err
  }
  _, err = fmt.Fprintln(ch.Writer, result)
  return err
}