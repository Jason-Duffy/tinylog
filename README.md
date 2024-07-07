# TinyLog

TinyLog is a lightweight logging library for Go, designed to be efficient and easy to use. It leverages the TinyFmt library for formatting log messages.

## Goals

- **Lightweight**: Minimal overhead and dependencies.
- **Easy to use**: Simple API for common logging tasks.
- **Efficient**: Optimized for performance and low resource usage.

## Requirements

- Go 1.20 or later.
- TinyFmt library (`github.com/Jason-Duffy/tinyfmt`).

## Installation

To install TinyLog, you can use `go get`:

```sh
go get github.com/Jason-Duffy/tinylog
```

## Usage

Here's a quick example of how to use TinyLog in your project:

```go
package main

import (
    "github.com/Jason-Duffy/tinylog"
    "os"
)

func main() {
    // Create a new logger instance
    logger := tinylog.NewLogger(os.Stdout)

    // Log a message
    err := logger.Log("Hello, %s! The answer is %d.", "world", 42)
    if err != nil {
        println("Error logging message:", err.Error())
    }
}
```

## API

### Logger

#### NewLogger

```go
func NewLogger(output *os.File) *Logger
```

Creates a new `Logger` instance that writes to the specified output file.

#### Log

```go
func (l *Logger) Log(format string, args ...interface{}) error
```

Logs a formatted message using the specified format and arguments. The format specifiers follow the conventions used in TinyFmt's `Sprintf`.

## Examples

### Basic Logging

```go
package main

import (
    "github.com/Jason-Duffy/tinylog"
    "os"
)

func main() {
    // Create a logger that writes to stdout
    logger := tinylog.NewLogger(os.Stdout)

    // Log messages
    logger.Log("This is a simple log message.")
    logger.Log("Hello, %s!", "world")
    logger.Log("Value: %d", 42)
}
```

### Logging to a File

```go
package main

import (
    "github.com/Jason-Duffy/tinylog"
    "os"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        println("Error creating log file:", err.Error())
        return
    }
    defer file.Close()

    // Create a logger that writes to a file
    logger := tinylog.NewLogger(file)
    logger.Log("This log message goes to a file.")
}
```

### Logging to Both Stdout and a File

```go
package main

import (
    "github.com/Jason-Duffy/tinylog"
    "os"
    "io"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        println("Error creating log file:", err.Error())
        return
    }
    defer file.Close()

    // Create a multi-writer that writes to both stdout and a file
    multiWriter := io.MultiWriter(os.Stdout, file)

    // Create a logger that writes to both stdout and the file
    logger := tinylog.NewLogger(multiWriter)
    logger.Log("This log message goes to both stdout and the file.")
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [TinyFmt](https://github.com/Jason-Duffy/tinyfmt) for the formatting library used in TinyLog.
