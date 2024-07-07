// =============================================================================
// Project: tinylog
// File: log.go
// Description: A simple, lightweight logging package for Go, intended for use
// 				on embedded targets.
// Datasheet/Docs:
//
// Author: Jason Duffy
// Created on: 07/07/2024
//
// Copyright: (C) 2024, Jason Duffy
// License: See LICENSE file in the project root for full license information.
// Disclaimer: See DISCLAIMER file in the project root for full disclaimer.
// =============================================================================

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package tinylog

import (
	"io"
	"sync"

	"github.com/Jason-Duffy/tinyfmt"
)

// -------------------------------------------------------------------------- //
//               Public Consts, Structs & Variable Definitions                //
// -------------------------------------------------------------------------- //

// LogLevel represents the severity of the log message.
type LogLevel int

// Logger represents a logger with an output destination.
type Logger struct {
	output      io.Writer
	packageName string
	level       LogLevel
	mu          sync.RWMutex
}

// Define log levels with increasing severity.
const (
	ERROR   LogLevel = 0
	WARNING LogLevel = 1
	INFO    LogLevel = 2
	DEBUG   LogLevel = 3
)

// -------------------------------------------------------------------------- //
//               Private Consts, Structs & Variable Definitions               //
// -------------------------------------------------------------------------- //

var (
	loggingEnabled = true
	mu             sync.RWMutex
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// SetLoggingEnabled sets the global logging enabled state.
func SetLoggingEnabled(enabled bool) {
	mu.Lock()
	defer mu.Unlock()
	loggingEnabled = enabled
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
