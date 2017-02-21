package lopher

import (
	"io"
	"os"
)

var defaultLogger = New(os.Stdout, false, LFstdFlags)

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(fmt string, v ...interface{}) {
	defaultLogger.Infof(fmt, v...)
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugf(fmt string, v ...interface{}) {
	defaultLogger.Debugf(fmt, v...)
}

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func SetFlags(f LogFlags) {
	defaultLogger.SetFlags(f)
}

func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}
