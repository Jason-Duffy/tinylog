// log.go
package tinylog

import (
	"io"

	"github.com/Jason-Duffy/tinyfmt"
)

// Logger represents a logger with an output destination.
type Logger struct {
	output io.Writer
	name   string
}

// NewLogger creates a new Logger with the specified output.
func NewLogger(output io.Writer, name string) *Logger {
	return &Logger{
		output: output,
		name:   name,
	}
}

// Log logs a message to the output.
func (l *Logger) Log(message string) {
	var prefix = l.name
	formattedMessage, _ := tinyfmt.Sprintf("%s: %s\n", prefix, message)
	l.output.Write([]byte(formattedMessage))
	print(formattedMessage) // Debug print
}
