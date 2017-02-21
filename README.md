# Lopher
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/adamveld12/lopher)

A dead simple log library, inspired by [Dave Cheney's blog post on logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

Only two log levels exist: Info and Debug. Info is always enabled, Debug is for development time.

Included are the same flags that are available with the standard library log package, with an API that
should feel familiar to you if you've used the standard library `log` package at all.

## How to use 

```go
package main

import (
    "time"
    "os"
    "gopkg.in/adamveld12/lopher.v1"
)

func main(){
    // make a new instance:
    l := lopher.New(os.Stdout, false, lopher.LFstdFlags)

    started := time.Now()

    // mm/dd/yyyy hh:mm:ss /file/path/main.go [INFO] App started.\n
    l.Info("App started.")

    // Debug entries don't print unless you set DebugMode to true
    l.DebugMode(true)

    l.Debug("App is doing stuff")
    l.Debugf("App ran for %+v", time.Since(started))

    // Also included are package level functions with the same API as an instance
    lopher.Info("App Exiting")
}
```

## Credits

Dave Cheney for being dope

## License

MIT