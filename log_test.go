// log_test.go
package tinylog

import (
	"bytes"
	"testing"
)

func TestBasicLogging(t *testing.T) {
	var buf bytes.Buffer

	logger := NewLogger(&buf, "TEST")
	logger.Log("Test message")

	logOutput := buf.String()
	expected := "TEST: Test message\n"

	if logOutput != expected {
		t.Errorf("Expected log output to be %q, but got %q", expected, logOutput)
	}
}
