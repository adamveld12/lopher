package lopher

import (
	"bytes"
	"math"
	"os"
	"testing"
	"time"
)

func TestFmtFuncs(t *testing.T) {
	l := New(nil, false, LFNone)
	cases := map[string]struct {
		subject  func(string, ...interface{})
		fmt      string
		input    []interface{}
		expected string
		devMode  bool
	}{
		"Infof":                        {l.Infof, "Hello %s!", []interface{}{"World"}, "[INFO] Hello World!\n", false},
		"Infof trim new line":          {l.Infof, "Hello %s!\n", []interface{}{"World"}, "[INFO] Hello World!\n", false},
		"Debugf with devmode disabled": {l.Debugf, "Hello %s!", []interface{}{"World"}, "", false},
		"Debugf":                       {l.Debugf, "Hello %s!", []interface{}{"World"}, "[DEBUG] Hello World!\n", true},
	}

	for tcn, tc := range cases {
		b := &bytes.Buffer{}
		l.SetDebug(tc.devMode)
		l.SetOutput(b)
		tc.subject(tc.fmt, tc.input...)
		actual := b.String()
		if actual != tc.expected {
			t.Errorf("%s FAILED\n\texpected: \"%s\"\tactual:   \"%s\"", tcn, tc.expected, actual)
		}
	}
}

func TestBaseFuncs(t *testing.T) {
	l := New(nil, false, LFNone)
	cases := map[string]struct {
		subject  func(...interface{})
		input    []interface{}
		expected string
		debug    bool
	}{
		"Info":                    {l.Info, []interface{}{"Hello World!"}, "[INFO] Hello World!\n", false},
		"Info trim new line":      {l.Info, []interface{}{"Hello World!\n"}, "[INFO] Hello World!\n", false},
		"Info trim new lines":     {l.Info, []interface{}{"Hello World!\n\n\n\n"}, "[INFO] Hello World!\n", false},
		"Debug dev mode disabled": {l.Debug, []interface{}{"Hello World!"}, "", false},
		"Debug":                   {l.Debug, []interface{}{"Hello World!"}, "[DEBUG] Hello World!\n", true},
	}

	for tcn, tc := range cases {
		b := &bytes.Buffer{}
		l.SetOutput(b)
		l.SetDebug(tc.debug)
		tc.subject(tc.input...)
		actual := b.String()
		if actual != tc.expected {
			t.Errorf("%s FAILED\n\texpected: \"%s\"\tactual:   \"%s\"", tcn, tc.expected, actual)
		}
	}
}

func Example_package() {
	SetFlags(LFNone)
	started := time.Now()
	Info("App Started.")

	Debug("Taking time measurement...")

	// Setting debug enables the debug level
	SetDebug(true)

	time.Sleep(time.Second)
	s := math.Floor(time.Since(started).Seconds())
	Debugf("App ran for %v seconds.", s)
	Info("App Exiting.")
	// Output:
	// [INFO] App Started.
	// [DEBUG] App ran for 1 seconds.
	// [INFO] App Exiting.

}

func Example() {
	l := New(os.Stdout, false, LFNone)
	started := time.Now()
	l.Info("App Started.")

	l.Debug("Taking time measurement...")

	// Setting debug enables the debug level
	l.SetDebug(true)

	time.Sleep(time.Second)
	s := math.Floor(time.Since(started).Seconds())
	l.Debugf("App ran for %v seconds.", s)
	l.Info("App Exiting.")
	// Output:
	// [INFO] App Started.
	// [DEBUG] App ran for 1 seconds.
	// [INFO] App Exiting.
}
