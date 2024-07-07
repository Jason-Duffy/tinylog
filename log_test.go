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

func TestLogger_GlobalToggle(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", DEBUG)

	tests := []struct {
		enabled  bool
		logLevel LogLevel
		message  string
		expected string
	}{
		{false, DEBUG, "This log should not appear", ""},
		{true, DEBUG, "This log should appear", "DEBUG > testModule: This log should appear\n"},
	}

	for _, tt := range tests {
		buf.Reset()
		SetLoggingEnabled(tt.enabled)
		logger.LogLevel(tt.logLevel, tt.message)
		if buf.String() != tt.expected {
			t.Errorf("Expected %q but got %q", tt.expected, buf.String())
		}
	}
}

func TestLogger_SetLogLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", ERROR)

	tests := []struct {
		initialLevel LogLevel
		newLevel     LogLevel
		logLevel     LogLevel
		message      string
		expected     string
	}{
		{ERROR, ERROR, ERROR, "Error message", "ERROR > testModule: Error message\n"},
		{ERROR, WARNING, WARNING, "Warning message", "WARNING > testModule: Warning message\n"},
		{WARNING, INFO, INFO, "Info message", "INFO > testModule: Info message\n"},
		{INFO, DEBUG, DEBUG, "Debug message", "DEBUG > testModule: Debug message\n"},
	}

	for _, tt := range tests {
		logger.SetLogLevel(tt.initialLevel)
		buf.Reset()
		logger.SetLogLevel(tt.newLevel)
		logger.LogLevel(tt.logLevel, tt.message)
		if buf.String() != tt.expected {
			t.Errorf("Expected %q but got %q", tt.expected, buf.String())
		}
	}
}

func TestLogger_LogLevel_ChangeLevelMultipleTimes(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", INFO)

	tests := []struct {
		newLevel LogLevel
		logLevel LogLevel
		message  string
		expected string
	}{
		{ERROR, INFO, "Info message", ""},
		{ERROR, ERROR, "Error message", "ERROR > testModule: Error message\n"},
		{DEBUG, DEBUG, "Debug message", "DEBUG > testModule: Debug message\n"},
	}

	for _, tt := range tests {
		logger.SetLogLevel(tt.newLevel)
		buf.Reset()
		logger.LogLevel(tt.logLevel, tt.message)
		if buf.String() != tt.expected {
			t.Errorf("Expected %q but got %q", tt.expected, buf.String())
		}
	}
}

func TestLogger_GlobalToggle_AfterLogLevelChange(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "testModule", DEBUG)

	tests := []struct {
		enabled  bool
		logLevel LogLevel
		message  string
		expected string
	}{
		{false, DEBUG, "This log should not appear", ""},
		{true, ERROR, "Error message", "ERROR > testModule: Error message\n"},
		{false, ERROR, "This log should not appear", ""},
	}

	for _, tt := range tests {
		buf.Reset()
		SetLoggingEnabled(tt.enabled)
		logger.LogLevel(tt.logLevel, tt.message)
		if buf.String() != tt.expected {
			t.Errorf("Expected %q but got %q", tt.expected, buf.String())
		}
	}
}
