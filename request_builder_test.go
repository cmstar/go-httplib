package httplib_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/cmstar/go-httplib"
	"github.com/stretchr/testify/assert"
)

var DefaultBody = []byte("default body")

type TestServer struct {
	*httptest.Server

	Request *http.Request // Stores the last request.
	Body    []byte        // Stores the body of the  last request.
}

func NewTestServer(status int, body []byte) *TestServer {
	ts := &TestServer{}
	ts.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts.Request = r
		ts.Body, _ = io.ReadAll(r.Body)

		w.WriteHeader(status)
		w.Write(body)
	}))
	return ts
}

func TestRequestBuilder_URL(t *testing.T) {
	t.Run("url-without-query", func(t *testing.T) {
		urlBase := "http://temp.org"
		b := httplib.NewBuilder("POST", urlBase)
		assert.Equal(t, urlBase, b.URL())

		b.WithQuery("b", 22)
		assert.Equal(t, urlBase+"?b=22", b.URL())

		b.WithQueries(nil)
		assert.Equal(t, urlBase+"?b=22", b.URL())

		b.WithQueries(map[string]any{
			"c": 33,
			"d": "&", // Should be escaped.
		})
		b.WithQuery("e", nil)
		assert.Equal(t, urlBase+"?b=22&c=33&d=%26&e=", b.URL())

		b.WithQuery("f", reflect.Struct) // Stringer.
		assert.Equal(t, urlBase+"?b=22&c=33&d=%26&e=&f=struct", b.URL())
	})

	t.Run("url-with-query", func(t *testing.T) {
		urlBase := "http://temp.org/path?a=1"
		b := httplib.NewBuilder("POST", urlBase)

		assert.Equal(t, urlBase, b.URL())
		b.WithQuery("b", 22)
		assert.Equal(t, urlBase+"&b=22", b.URL())

		b.WithQueries(map[string]any{
			"a": 3, // Merged into a.
		})
		b.WithQuery("a", 4)
		assert.Equal(t, urlBase+"&a=3&a=4&b=22", b.URL())
	})
}

func TestRequestBuilder_WithHeader(t *testing.T) {
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	b := httplib.NewBuilder("GET", s.URL)
	b.WithHeader("h1-header", "v1")
	b.WithHeader("h2-header", "v2")
	b.WithHeaders(map[string]any{
		"h3-header": 3,
		"h4-header": 4,
	})
	b.WithHeaders(nil)

	content, err := b.ReadString()
	r := s.Request
	assert := assert.New(t)
	assert.NoError(err)

	// All headers are converted to Header-Naming-Style .
	assert.Contains(r.Header, "H1-Header", "v1")
	assert.Contains(r.Header, "H2-Header", "v2")
	assert.Contains(r.Header, "H3-Header", "3")
	assert.Contains(r.Header, "H4-Header", "4")
	assert.Equal(string(DefaultBody), content)
}

func TestRequestBuilder_WithForm(t *testing.T) {
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	b := httplib.NewBuilder("POST", s.URL).
		SetStringBody("should be overwrite by WithForm").
		WithQuery("a", 1).
		WithQuery("b", 2).
		WithHeader("h", "v").
		WithForm("a", 3).
		WithForm("b", "?"). // Need escape.
		WithForms(nil).
		WithForms(map[string]any{
			"c": "cc",
			"d": "中文",
		})

	content, err := b.ReadBinary()
	r := s.Request
	assert := assert.New(t)
	assert.NoError(err)
	assert.Contains(r.Header, "H", "v")
	assert.Contains(r.Header, "H", "v")
	assert.Equal("/?a=1&b=2", r.RequestURI)
	assert.Equal("a=3&b=%3F&c=cc&d=%E4%B8%AD%E6%96%87", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))
}

func TestRequestBuilder_SetStringBody(t *testing.T) {
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	b := httplib.NewBuilder("POST", s.URL).
		WithForm("a", 1).
		SetBinaryBody([]byte("override")).
		SetStringBody("override").
		SetStringBody("body") // Override others.

	content, err := b.ReadBinary()
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal("body", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))

	// The body is   reusable.
	content, err = b.ReadBinary()
	assert.NoError(err)
	assert.Equal("body", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))
}

func TestRequestBuilder_SetBinaryBody(t *testing.T) {
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	b := httplib.NewBuilder("POST", s.URL).
		WithForm("a", 1).
		SetStringBody("override").
		SetBinaryBody([]byte("override")).
		SetBinaryBody([]byte("body")) // Override others.

	content, err := b.ReadBinary()
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal("body", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))

	// The body is   reusable.
	content, err = b.ReadBinary()
	assert.NoError(err)
	assert.Equal("body", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))
}

func TestRequestBuilder_SetReaderBody(t *testing.T) {
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	b := httplib.NewBuilder("POST", s.URL).
		WithForm("a", 1).
		SetStringBody("override").
		SetBinaryBody([]byte("override")).
		SetReaderBody(strings.NewReader("body")) // Override others.

	content, err := b.ReadBinary()
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal("body", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))

	// The body is not reusable.
	content, err = b.ReadBinary()
	assert.NoError(err)
	assert.Equal("", string(s.Body))
	assert.Equal(string(DefaultBody), string(content))
}

func TestRequestBuilder_Do(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		b := httplib.NewBuilder("POST", s.URL).SetStringBody("body")

		resp, err := b.Do()
		assert := assert.New(t)
		assert.NoError(err)
		assert.Equal("body", string(s.Body))

		assert.NotNil(resp)
		assert.Equal(http.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(string(DefaultBody), string(respBody))
	})

	t.Run("err-bad-method", func(t *testing.T) {
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		b := httplib.NewBuilder("\t", s.URL)
		resp, err := b.Do()
		assert := assert.New(t)
		assert.Error(err)
		assert.Contains(err.Error(), "invalid method")
		assert.Nil(resp)
	})

	t.Run("err-bad-server", func(t *testing.T) {
		b := httplib.NewBuilder("GET", "wrong url")

		resp, err := b.Do()
		assert := assert.New(t)
		assert.Error(err)
		assert.Nil(resp)
	})
}

func TestRequestBuilder_MustDo(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		b := httplib.NewBuilder("POST", s.URL)

		resp := b.MustDo()
		assert := assert.New(t)
		assert.Equal("", string(s.Body))

		assert.NotNil(resp)
		assert.Equal(http.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(string(DefaultBody), string(respBody))
	})

	t.Run("panic", func(t *testing.T) {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
		}()

		b := httplib.NewBuilder("GET", "wrong url")
		b.MustDo()
	})
}

func TestRequestBuilder_MustReadBinary(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		b := httplib.NewBuilder("POST", s.URL)

		resp := b.MustReadBinary()
		assert := assert.New(t)
		assert.Equal("", string(s.Body))
		assert.Equal(string(DefaultBody), string(resp))
	})

	t.Run("panic", func(t *testing.T) {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
		}()

		b := httplib.NewBuilder("GET", "wrong url")
		b.MustReadBinary()
	})
}

func TestRequestBuilder_MustReadString(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		b := httplib.NewBuilder("POST", s.URL)

		resp := b.MustReadString()
		assert := assert.New(t)
		assert.Equal("", string(s.Body))
		assert.Equal(string(DefaultBody), resp)
	})

	t.Run("panic", func(t *testing.T) {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
		}()

		b := httplib.NewBuilder("GET", "wrong url")
		b.MustReadString()
	})
}
