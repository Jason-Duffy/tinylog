package tinylog

import (
	"bytes"
	"testing"
)

func TestLogger_Log(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", DEBUG)

	// Test Log function
	logger.Log("This is a test log message")

	expected := "testModule > This is a test log message\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}

func TestLogger_LogLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", DEBUG)

	tests := []struct {
		level    LogLevel
		message  string
		expected string
	}{
		{DEBUG, "Debug message", "DEBUG > testModule: Debug message\n"},
		{INFO, "Info message", "INFO > testModule: Info message\n"},
		{WARNING, "Warning message", "WARNING > testModule: Warning message\n"},
		{ERROR, "Error message", "ERROR > testModule: Error message\n"},
	}

	for _, tt := range tests {
		buf.Reset()
		logger.LogLevel(tt.level, tt.message)
		if buf.String() != tt.expected {
			t.Errorf("Expected %q but got %q", tt.expected, buf.String())
		}
	}

	// Test log levels filtering
	buf.Reset()
	logger = NewLogger(&buf, "testModule", WARNING)
	logger.LogLevel(DEBUG, "This should not be logged")
	logger.LogLevel(INFO, "This should not be logged either")
	if buf.String() != "" {
		t.Errorf("Expected no output but got %q", buf.String())
	}

	buf.Reset()
	logger.LogLevel(WARNING, "This should be logged")
	expected := "WARNING > testModule: This should be logged\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}

// TestLogger_GlobalToggle tests the global toggle for enabling or disabling logging.
func TestLogger_GlobalToggle(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", DEBUG)

	// Disable logging globally
	SetLoggingEnabled(false)
	logger.Log("This log should not appear")
	if buf.String() != "" {
		t.Errorf("Expected no output but got %q", buf.String())
	}

	// Enable logging globally
	SetLoggingEnabled(true)
	logger.Log("This log should appear")
	expected := "testModule > This log should appear\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}

	// Test LogLevel function with global toggle
	buf.Reset()
	SetLoggingEnabled(false)
	logger.LogLevel(INFO, "This log should not appear")
	if buf.String() != "" {
		t.Errorf("Expected no output but got %q", buf.String())
	}

	SetLoggingEnabled(true)
	logger.LogLevel(INFO, "This log should appear")
	expected = "INFO > testModule: This log should appear\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}
