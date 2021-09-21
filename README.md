# go-metrics

[![Go Reference](https://pkg.go.dev/badge/github.com/ronelliott/go-metrics.svg)](https://pkg.go.dev/github.com/ronelliott/go-metrics) [![Go Report Card](https://goreportcard.com/badge/github.com/ronelliott/go-metrics)](https://goreportcard.com/report/github.com/ronelliott/go-metrics)

go-metrics is a [Go](https://golang.org/) package that provides ...

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/ronelliott/go-metrics
```

### Usage

Import the package into your project.

```go
import "github.com/ronelliott/go-metrics"
```

Construct a new Metrics instance which can then be used ...

```go
m := metrics.New(time.Minute)
m.Update("bucket", 100, time.Now())
```

See Documentation below for more detailed information.


## Documentation

**NOTICE**: This library and the Discord API are unfinished.
Because of that there may be major changes to library in the future.

The code is documented and is the only documentation available. Both GoDoc and
GoWalker (below) present that information in a nice format.

- [![Go Reference](https://pkg.go.dev/badge/github.com/ronelliott/go-metrics.svg)](https://pkg.go.dev/github.com/ronelliott/go-metrics)
- [![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/ronelliott/go-metrics)
