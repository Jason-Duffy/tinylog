package tinylog

import (
	"io"
	"sync"

	"github.com/Jason-Duffy/tinyfmt"
)

// LogLevel represents the severity of the log message.
type LogLevel int

// Define log levels with increasing severity.
const (
	ERROR   LogLevel = 0
	WARNING LogLevel = 1
	INFO    LogLevel = 2
	DEBUG   LogLevel = 3
)

var (
	loggingEnabled = true
	mu             sync.RWMutex
)

// SetLoggingEnabled sets the global logging enabled state.
func SetLoggingEnabled(enabled bool) {
	mu.Lock()
	defer mu.Unlock()
	loggingEnabled = enabled
}

// Logger represents a logger with an output destination.
type Logger struct {
	output      io.Writer
	packageName string
	level       LogLevel
	mu          sync.RWMutex
}

// NewLogger creates a new Logger with the specified output.
func NewLogger(output io.Writer, packageName string, level LogLevel) *Logger {
	return &Logger{
		output:      output,
		packageName: packageName,
		level:       level,
	}
}

// SetLogLevel sets the log level of the logger.
func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Log logs a message to the output.
func (l *Logger) Log(message string) {
	mu.RLock()
	defer mu.RUnlock()
	if !loggingEnabled {
		return
	}
	formattedMessage, _ := tinyfmt.Sprintf("%s > %s\n", l.packageName, message)
	l.output.Write([]byte(formattedMessage))
}

// LogLevel logs a message with a specified level to the output.
func (l *Logger) LogLevel(level LogLevel, message string) {
	mu.RLock()
	defer mu.RUnlock()
	if !loggingEnabled || level > l.level {
		return
	}
	levelStr := []string{"ERROR", "WARNING", "INFO", "DEBUG"}[level]
	formattedMessage, _ := tinyfmt.Sprintf("%s > %s: %s\n", levelStr, l.packageName, message)
	l.output.Write([]byte(formattedMessage))
}
