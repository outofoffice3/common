package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

type Logger interface {
	SetLogLevel(level LogLevel)
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// LogLevel represents different log levels.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

// ConsoleLogger is a simple logger implementation that logs to the console.
type ConsoleLogger struct {
	mu       sync.Mutex
	logLevel LogLevel
	writer   io.Writer
}

// NewConsoleLogger creates a new ConsoleLogger instance with the specified log level.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger {
	return &ConsoleLogger{
		logLevel: logLevel,
		writer:   os.Stdout,
	}
}

// SetLogLevel sets the log level for the logger.
func (l *ConsoleLogger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

// log logs a message at the specified log level.
func (l *ConsoleLogger) log(level LogLevel, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.logLevel {
		return
	}

	pc, _, _, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc)
	if function != nil {
		file, _ := function.FileLine(pc)
		funcName := path.Base(file)
		method := path.Base(function.Name())

		message := fmt.Sprintf("[%s](%s)[%s] : %s", LogLevelToString(level), funcName, method, fmt.Sprintf(format, args...))
		log.Println(message)
	} else {
		message := fmt.Sprintf("[%s]: %s", LogLevelToString(level), fmt.Sprintf(format, args...))
		log.Println(message)
	}
}

// LogLevelToString converts the LogLevel to a string.
func LogLevelToString(level LogLevel) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Debugf logs a debug message.
func (l *ConsoleLogger) Debugf(format string, args ...interface{}) {
	l.log(LogLevelDebug, format, args...)
}

// Infof logs an info message.
func (l *ConsoleLogger) Infof(format string, args ...interface{}) {
	l.log(LogLevelInfo, format, args...)
}

// Warnf logs a warning message.
func (l *ConsoleLogger) Warnf(format string, args ...interface{}) {
	l.log(LogLevelWarning, format, args...)
}

// Errorf logs an error message.
func (l *ConsoleLogger) Errorf(format string, args ...interface{}) {
	l.log(LogLevelError, format, args...)
}
