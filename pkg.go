package lopher

import (
	"io"
	"os"
)

var dLogger = New(os.Stderr, false, "", LFstdFlags)

// Info logs at info level
func Info(v ...interface{}) {
	dLogger.Info(v...)
}

// Infof logs a formatted string at info level
func Infof(fmt string, v ...interface{}) {
	dLogger.Infof(fmt, v...)
}

// Debug logs at debug level
func Debug(v ...interface{}) {
	dLogger.Debug(v...)
}

// Debugf logs a formatted string at debug level
func Debugf(fmt string, v ...interface{}) {
	dLogger.Debugf(fmt, v...)
}

// SetOutput sets the output writer for the logger
func SetOutput(w io.Writer) {
	dLogger.SetOutput(w)
}

// SetFlags sets the log entry prefix flags
func SetFlags(f LogFlags) {
	dLogger.SetFlags(f)
}

// SetDebug sets debug mode
func SetDebug(debug bool) {
	dLogger.SetDebug(debug)
}
