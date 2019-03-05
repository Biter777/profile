profile
=======

Simple profiling support package for Go with multi profiling support

[![Build Status](https://travis-ci.org/pkg/profile.svg?branch=master)](https://travis-ci.org/pkg/profile) [![GoDoc](http://godoc.org/github.com/biter777/profile?status.svg)](http://godoc.org/github.com/biter777/profile)


installation
------------

    go get github.com/biter777/profile

usage
-----

Enabling profiling in your application is as simple as one line at the top of your main function. 

By default CPU, Memory, Mutex, Block profiling and Trace out is enabled.

```go
import "github.com/biter777/profile"

func main() {
    defer profile.Start().Stop()
    ...
}
```

options
-------

What to profile is controlled by config value passed to profile.Start. 

By default CPU, Memory, Mutex, Block profiling and Trace out is enabled.

```go
import "github.com/biter777/profile"

func main() {
    // p.Stop() must be called before the program exits to
    // ensure profiling information is written to disk.
    p := profile.Start(profile.CPUProfile, profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
    ...
}
```

Several convenience package level values are provided for cpu, memory, and block (contention) profiling.

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/profile).

contributing
------------

We welcome pull requests, bug fixes and issue reports.

Before proposing a change, please discuss it first by raising an issue.
