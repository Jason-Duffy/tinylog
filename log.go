package tinylog

import (
	"os"

	"github.com/Jason-Duffy/tinyfmt"
)

// Logger is a simple logger struct
type Logger struct {
	output *os.File
}

// NewLogger creates a new Logger instance
func NewLogger(output *os.File) *Logger {
	return &Logger{output: output}
}

// Log logs a formatted message
func (l *Logger) Log(format string, args ...interface{}) error {
	return tinyfmt.PrintToIO(l.output, format, args...)
}
