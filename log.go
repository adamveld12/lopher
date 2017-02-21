/*
Package lopher is inspired by Dave Cheney's blog post on logging: https://dave.cheney.net/2015/11/05/lets-talk-about-logging
lopher only supports two logging levels:
1. Things that developers care about when they are developing or debugging software. (DEBUG)
2. Things that users care about when using your software. (INFO)

If you need to log errors, use `Info`.
If you need to log warnings, use `Debug`` or don't log them (most likely no one will care).
If you need log.Fatal, bubble the error up to `main.main` and gracefully exit there; it makes the most sense
and is cleaner than the alternatives.
*/
package lopher

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Logger is a a simple logger interface that only supports two levels: INFO and DEBUG
type Logger interface {
	// Info logs at info level
	Info(...interface{})

	// Infof logs a formatted string at info level
	Infof(string, ...interface{})

	// Debug logs at debug level
	Debug(...interface{})

	// Debugf logs a formatted string at debug level
	Debugf(string, ...interface{})

	// SetFlags sets the log entry prefix flags
	SetFlags(LogFlags)

	// SetOutput sets the output writer for the logger
	SetOutput(io.Writer)

	// SetDebug sets debug mode
	SetDebug(bool)
}

// LogFlags are a set of flags which define the prefix for each log entry
type LogFlags int

// These flags define which text to prefix to each log entry generated by the Logger.
const (
	// Bits or'ed together to control what's printed.
	// There is no control over the order they appear (the order listed
	// here) or the format they present (as described in the comments).
	// The prefix is followed by a colon only when Llongfile or Lshortfile
	// is specified.
	// For example, flags Ldate | Ltime (or LstdFlags) produce,
	//	2009/01/23 01:23:23 message
	// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	LFdate         = 1 << iota                            // the date in the local time zone: 2009/01/23
	LFtime                                                // the time in the local time zone: 01:23:23
	LFmicroseconds                                        // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	LFlongfile                                            // full file name and line number: /a/b/c/d.go:23
	LFshortfile                                           // final file name element and line number: d.go:23. overrides Llongfile
	LFUTC                                                 // if Ldate or Ltime is set, use UTC rather than the local time zone
	LFstdFlags     = LFdate | LFtime | LFUTC | LFlongfile // initial values for the standard logger
	LFNone         = 0
)

// New creates a new Logger
// Loggers are immutable
func New(out io.Writer, debug bool, flags LogFlags) Logger {
	return &Log{sync.Mutex{}, out, debug, flags}
}

// Log is a simple implementation of Logger
type Log struct {
	sync.Mutex
	// Output is the writer to send logs to
	io.Writer
	// Debug when true will enable the DEBUG logging level
	DebugMode bool
	// Flags are a set of flags to use for setting up log time stamps
	Flags LogFlags
}

// SetFlags sets the log flags used by the Logger
func (l *Log) SetFlags(f LogFlags) {
	l.Lock()
	defer l.Unlock()
	l.Flags = f
}

// SetOutput sets the output writer for the Logger
func (l *Log) SetOutput(w io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.Writer = w
}

// SetDebug enables or disables debug printing
func (l *Log) SetDebug(v bool) {
	l.Lock()
	defer l.Unlock()
	l.DebugMode = v
}

// Info logs at info level
func (l *Log) Info(v ...interface{}) {
	l.print("INFO", v...)
}

// Infof logs a formatted string at info level
func (l *Log) Infof(fmtStr string, v ...interface{}) {
	l.print("INFO", fmt.Sprintf(fmtStr, v...))
}

// Debug logs at debug level
func (l *Log) Debug(v ...interface{}) {
	if !l.DebugMode {
		return
	}

	l.print("DEBUG", v...)
}

// Debugf logs a formatted string at debug level
func (l *Log) Debugf(fmtStr string, v ...interface{}) {
	if !l.DebugMode {
		return
	}

	l.print("DEBUG", fmt.Sprintf(fmtStr, v...))
}

func (l *Log) print(level string, v ...interface{}) error {
	l.Lock()
	defer l.Unlock()

	caller := ""
	if l.Flags&(LFshortfile|LFlongfile) != 0 {
		l.Unlock()
		// # of calls back up the call stack
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		l.Lock()

		caller = fmt.Sprintf("%s:%d", file, line)
	}

	header := fmtHeader(l.Flags, caller, time.Now())
	// wipe out all of the new lines for better structured logging
	args := strings.Replace(fmt.Sprint(v...), "\n", " ", -1)
	_, err := fmt.Fprintf(l, "%s[%s] %+v\n", header, level, strings.TrimSpace(args))
	return fmt.Errorf("Could not write to output: %+v", err)
}

func fmtHeader(flags LogFlags, file string, t time.Time) string {
	if flags == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	if flags&LFUTC != 0 {
		t = t.UTC()
	}

	if flags&(LFdate|LFtime|LFmicroseconds) != 0 {
		y, m, d := t.Date()
		buf.WriteString(fmt.Sprintf("%d/%d/%d ", y, m, d))

		if flags&(LFtime) != 0 {
			h, min, s := t.Clock()
			buf.WriteString(fmt.Sprintf("%d:%d:%0d", h, min, s))
			if flags&(LFmicroseconds) != 0 {
				buf.WriteString(fmt.Sprintf(".%d", t.Nanosecond()/1e3))
			}
			buf.WriteString(" ")
		}
	}

	if flags&(LFshortfile|LFlongfile) != 0 {
		if flags&(LFshortfile) != 0 {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}

		buf.WriteString(fmt.Sprintf("%s ", file))
	}

	return buf.String()
}