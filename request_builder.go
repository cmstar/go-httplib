package httplib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestBuilder is used to simply build HTTP request.
type RequestBuilder struct {
	Method  string
	baseUrl string

	query  url.Values
	header http.Header

	// Can be string/[]byte/io.Reader/url.Values .
	body any
}

// NewBuilder creates a new instance of RequestBuilder.
func NewBuilder(method string, baseUrl string) *RequestBuilder {
	return &RequestBuilder{
		Method:  method,
		baseUrl: baseUrl,

		query:  make(url.Values),
		header: make(http.Header),
	}
}

// URL returns the whole URL, includes all added query strings.
func (x *RequestBuilder) URL() string {
	uri := x.baseUrl
	queryString := x.query.Encode()

	if queryString != "" {
		hasQuestionMark := false
		for i := 0; i < len(uri); i++ {
			if uri[i] == '?' {
				hasQuestionMark = true
				break
			}
		}

		if hasQuestionMark {
			uri += "&"
		} else {
			uri += "?"
		}

		uri += queryString
	}

	return uri
}

// WithQuery appends a query string to the URL.
//
// If the given value is not a string, it will be converted to a string.
func (x *RequestBuilder) WithQuery(name string, value any) *RequestBuilder {
	s := x.toString(value)
	x.query.Add(name, s)
	return x
}

// WithQuery appends a group of query strings to the URL.
//
// If a value in the map is not a string, it will be converted to a string.
func (x *RequestBuilder) WithQueries(values map[string]any) *RequestBuilder {
	if values == nil {
		return x
	}

	for k, v := range values {
		s := x.toString(v)
		x.query.Add(k, s)
	}
	return x
}

// WithForm appends a parameter to the form,
// and sets the header Content-Type to 'application/x-www-form-urlencoded'.
//
// If the given value is not a string, it will be converted to a string.
//
// If the current body set is not a form, it will be replaced.
func (x *RequestBuilder) WithForm(name string, value any) *RequestBuilder {
	q := x.ensureForm()
	s := x.toString(value)
	q.Add(name, s)
	return x
}

// WithForms appends a group of parameters to the form,
// and sets the header Content-Type to 'application/x-www-form-urlencoded'.
//
// If a value in the map is not a string, it will be converted to a string.
//
// If the current body set is not a form, it will be replaced.
func (x *RequestBuilder) WithForms(values map[string]any) *RequestBuilder {
	q := x.ensureForm()

	if values == nil {
		return x
	}

	for k, v := range values {
		s := x.toString(v)
		q.Add(k, s)
	}
	return x
}

// WithHeader appends a HTTP header. The name will be converted to the Header-Naming-Style.
//
// If the given value is not a string, it will be converted to a string.
func (x *RequestBuilder) WithHeader(name string, value any) *RequestBuilder {
	s := x.toString(value)
	x.header.Add(name, s)
	return x
}

// WithHeader appends a group of HTTP headers. The name will be converted to the Header-Naming-Style.
//
// If a value in the map is not a string, it will be converted to a string.
func (x *RequestBuilder) WithHeaders(values map[string]any) *RequestBuilder {
	if values == nil {
		return x
	}

	for k, v := range values {
		s := x.toString(v)
		x.header.Add(k, s)
	}
	return x
}

// SetStringBody set a string as the request body. If another body was set, it will be replaced.
func (x *RequestBuilder) SetStringBody(body string) *RequestBuilder {
	x.body = body
	return x
}

// SetBinaryBody set a slice of bytes as the request body. If another body was set, it will be replaced.
func (x *RequestBuilder) SetBinaryBody(body []byte) *RequestBuilder {
	x.body = body
	return x
}

// SetBinaryBody set a io.Reader as the request body. If another body was set, it will be replaced.
//
// The reader will be closed if it implements io.ReadCloser.
func (x *RequestBuilder) SetReaderBody(reader io.Reader) *RequestBuilder {
	x.body = reader
	return x
}

// Build returns a instance of http.Request, which includes the options set in other methods.
//
// If the body is string/[]byte or one of one of strings.Reader/bytes.Buffer/bytes.Reader,
// Build() can be called multiple times.
//
// If the body is io.Read and is not one of strings.Reader/bytes.Buffer/bytes.Reader, once the request
// is performed, the reader reached the end and is not reusable, you can call SetReaderBody() again to
// setup a new body.
func (x *RequestBuilder) Build() (*http.Request, error) {
	uri := x.URL()
	body := x.buildBody()
	request, err := http.NewRequest(x.Method, uri, body)
	if err != nil {
		return nil, err
	}

	request.Header = x.header
	return request, nil
}

// Do executes the HTTP request.
func (x *RequestBuilder) Do() (*http.Response, error) {
	request, err := x.Build()
	if err != nil {
		return nil, err
	}

	response, err := new(http.Client).Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReadBinary executes the HTTP request, if the status code of the response is 200 OK,
// it returns the whole response body as a slice of byte; otherwise returns an error.
//
// If you need to get the body when the status code is not 200 OK, call Do().
func (x *RequestBuilder) ReadBinary() ([]byte, error) {
	res, err := x.Do()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	return io.ReadAll(res.Body)
}

// ReadString executes the HTTP request, if the status code of the response is 200 OK,
// it returns the whole response body as a string; otherwise returns an error.
//
// If you need to get the body when the status code is not 200 OK, call Do().
func (x *RequestBuilder) ReadString() (string, error) {
	res, err := x.ReadBinary()
	return string(res), err
}

// MustDo is the panic version of Do().
func (x *RequestBuilder) MustDo() *http.Response {
	res, err := x.Do()
	if err != nil {
		panic(err)
	}
	return res
}

// MustReadBinary is the panic version of ReadBinary().
func (x *RequestBuilder) MustReadBinary() []byte {
	res, err := x.ReadBinary()
	if err != nil {
		panic(err)
	}
	return res
}

// MustReadString is the panic version of ReadString().
func (x *RequestBuilder) MustReadString() string {
	res, err := x.ReadString()
	if err != nil {
		panic(err)
	}
	return res
}

func (x *RequestBuilder) buildBody() io.Reader {
	if x.body != nil {
		switch v := x.body.(type) {
		case string:
			return strings.NewReader(v)

		case []byte:
			return bytes.NewBuffer(v)

		case io.Reader:
			return v

		case url.Values:
			q := v.Encode()
			return strings.NewReader(q)
		}
	}

	return nil
}

func (x *RequestBuilder) toString(v any) string {
	if v == nil {
		return ""
	}

	switch t := v.(type) {
	case string:
		return t
	case fmt.Stringer:
		return t.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (x *RequestBuilder) ensureForm() url.Values {
	x.header.Set("Content-Type", "application/x-www-form-urlencoded")

	if values, ok := x.body.(url.Values); ok {
		return values
	}

	values := make(url.Values)
	x.body = values
	return values
}
