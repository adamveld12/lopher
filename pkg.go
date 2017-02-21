package lopher

import (
	"io"
	"os"
)

var defaultLogger = New(os.Stdout, false, LFstdFlags)

// Info logs at info level
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof logs a formatted string at info level
func Infof(fmt string, v ...interface{}) {
	defaultLogger.Infof(fmt, v...)
}

// Debug logs at debug level
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Debugf logs a formatted string at debug level
func Debugf(fmt string, v ...interface{}) {
	defaultLogger.Debugf(fmt, v...)
}

// SetOutput sets the output writer for the logger
func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

// SetFlags sets the log entry prefix flags
func SetFlags(f LogFlags) {
	defaultLogger.SetFlags(f)
}

// SetDebug sets debug mode
func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}
