# Logger 

`logger` is a simple and flexible logging package for Go. It provides a customizable logger interface and a default ConsoleLogger implementation for logging messages to the console.

## Features

- Multiple log levels: Debug, Info, Warning, Error.
- Control over the log level to filter messages.
- Log message details, including function names and source file information.
- Flexible and easy-to-use API.
## Installation

You can install the logger package using the go get command:

``` go 
go get github.com/outofoffice3/common/logger
```
## Usage

Import the `logger` package into your Go code:
``` go
import "github.com/outofoffice3/common/logger"
```

Here's an example of how to create a logger, set the log level, and log messages:

``` go
package main

import (
	"github.com/outofoffice3/common/logger"
)

func main() {
	// Create a new logger with the desired log level (e.g., LogLevelDebug).
	log := logger.NewConsoleLogger(logger.LogLevelDebug)

	// Log messages at various levels.
	log.Debugf("This is a debug message.")
	log.Infof("This is an info message.")
	log.Warnf("This is a warning message.")
	log.Errorf("This is an error message.")
}
```

## API

### - NewConsoleLogger(logLevel LogLevel) *ConsoleLogger
- Creates a new ConsoleLogger instance with the specified log level.

### - SetLogLevel(level LogLevel)
- Sets the log level for the logger.

### Logging Methods
The logger provides methods for logging messages at different levels:

- Debugf(format string, args ...interface{}): Log a debug message.
- Infof(format string, args ...interface{}): Log an info message.
- Warnf(format string, args ...interface{}): Log a warning message.
- Errorf(format string, args ...interface{}): Log an error message.
