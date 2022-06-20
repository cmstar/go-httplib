# httplib - A library for HTTP requests

[![GoDoc](https://pkg.go.dev/badge/github.com/cmstar/go-httplib)](https://pkg.go.dev/github.com/cmstar/go-httplib)
[![Go](https://github.com/cmstar/go-httplib/workflows/Go/badge.svg)](https://github.com/cmstar/go-httplib/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/cmstar/go-httplib/branch/master/graph/badge.svg)](https://codecov.io/gh/cmstar/go-httplib)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)

## RequestBuilder

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

// Do request.
resp, err := b.ReadString()

// Upload a file.
file, _ := os.Open("/some/file")
b = httplib.NewBuilder("PUT", "http://example.org")
b.SetReaderBody(file)
resp := b.MustDo()
```

## Shortcuts

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