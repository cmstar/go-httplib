# httplib - A library for HTTP requests

[![GoDoc](https://pkg.go.dev/badge/github.com/cmstar/go-httplib)](https://pkg.go.dev/github.com/cmstar/go-httplib)
[![Build](https://github.com/cmstar/go-httplib/workflows/Go/badge.svg)](https://github.com/cmstar/go-httplib/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/cmstar/go-httplib/branch/master/graph/badge.svg)](https://codecov.io/gh/cmstar/go-httplib)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![GoVersion](https://img.shields.io/github/go-mod/go-version/cmstar/go-httplib)](https://github.com/cmstar/go-httplib/blob/main/go.mod)

## Features

- Build HTTP request in an easy way.
- The `headers` package provides HTTP header constants.
- Shortcut methods for reading string/binary body directly from an URL.

## Install

```bash
go get -u github.com/cmstar/go-httplib@latest
```

## Quick start

### RequestBuilder

`RequestBuilder` is used to simply build HTTP request.

```go
// Setup RequestBuilder. It supports fluent APIs.
b := httplib.NewBuilder("POST", "http://example.org").
    WithQueries(map[string]any{
        "q2": "v2",
        "q3": 3,
    }).
    WithHeader("x-custom-value", 112233).
    WithForm("f1", "vf1").
    WithForm("f2", "vf2")

// Do request. This will send a request like this:
// ====
// POST http://example.org?q2=v2&q3=3 HTTP/1.1
// Host: example.org
// X-Custom-Value: 112233
// <Some other headers such as User-Agent ...>
//
// f1=vf1&f2=vf2
// ====
resp, err := b.ReadString()

// Upload a file.
file, _ := os.Open("/some/file")
defer file.Close()

b = httplib.NewBuilder("PUT", "http://example.org")
b.SetReaderBody(file)
resp := b.MustDo()
```

### Shortcuts

Send request directly.

```go
// Send a GET request.
stringResponse, err := httplib.Get("http://example.org")

// Send a POST request with custom headers.
customHeaders := map[string]any{
    "Custom-Header":  "v1",
    "Custom-Header2": "v2",
}
binaryBody := []byte("request-body")
binaryResponse, err := httplib.PostBinaryWithHeaders("http://example.org", binaryBody, customHeaders)
```