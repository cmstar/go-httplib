package httplib_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cmstar/go-httplib"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert := assert.New(t)
		s := NewTestServer(http.StatusOK, []byte(""))
		defer s.Close()

		content, err := httplib.Get(s.URL)
		r := s.Request
		assert.NoError(err)
		assert.Equal("/", r.RequestURI)
		assert.Equal("", content)
	})

	t.Run("ok", func(t *testing.T) {
		assert := assert.New(t)
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		path := "/path/?a=1&b=2"
		content, err := httplib.Get(s.URL + path)
		r := s.Request
		assert.NoError(err)
		assert.Equal(path, r.RequestURI)
		assert.Equal(string(DefaultBody), content)
	})

	t.Run("err", func(t *testing.T) {
		assert := assert.New(t)
		s := NewTestServer(http.StatusBadRequest, DefaultBody)
		defer s.Close()

		content, err := httplib.Get(s.URL)
		assert.NotNil(err)
		assert.Contains(err.Error(), "400 Bad Request")
		assert.Equal("", content)
	})
}

func TestGetWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	content, err := httplib.GetWithHeaders(s.URL, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})

	r := s.Request
	assert.NoError(err)
	assert.Contains(r.Header, "X-Header1", "v1")
	assert.Contains(r.Header, "X-Header2", "v2")
	assert.Equal(string(DefaultBody), content)
}

func TestGetBinary(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	path := "/path"
	content, err := httplib.GetBinary(s.URL + path)
	r := s.Request
	assert.NoError(err)
	assert.Equal(path, r.RequestURI)
	assert.Equal(DefaultBody, content)
}

func TestGetBinaryWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	content, err := httplib.GetBinaryWithHeaders(s.URL, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})

	r := s.Request
	assert.NoError(err)
	assert.Contains(r.Header, "X-Header1", "v1")
	assert.Contains(r.Header, "X-Header2", "v2")
	assert.Equal(DefaultBody, content)
}

func TestMustGet(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	_ = httplib.MustGet(s.URL)
	assert.Fail("should panic")
}

func TestMustGetWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		r := s.Request
		assert.Contains(r.Header, "X-Header1", "v1")
		assert.Contains(r.Header, "X-Header2", "v2")

		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	_ = httplib.MustGetWithHeaders(s.URL, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.Fail("should panic")
}

func TestMustGetBinary(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	_ = httplib.MustGetBinary(s.URL)
	assert.Fail("should panic")
}

func TestMustGetBinaryWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		r := s.Request
		assert.Contains(r.Header, "X-Header1", "v1")
		assert.Contains(r.Header, "X-Header2", "v2")

		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	_ = httplib.MustGetBinaryWithHeaders(s.URL, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.Fail("should panic")
}

func TestPost(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert := assert.New(t)
		s := NewTestServer(http.StatusOK, DefaultBody)
		defer s.Close()

		requestBody := "request body"
		path := "/?a=1&a=2"
		content, err := httplib.Post(s.URL+path, requestBody)
		assert.NoError(err)

		r := s.Request
		assert.Equal(path, r.RequestURI)
		assert.NoError(err)
		assert.Equal(requestBody, string(s.Body))
		assert.Equal(string(DefaultBody), content)
	})

	t.Run("err", func(t *testing.T) {
		assert := assert.New(t)
		s := NewTestServer(http.StatusBadRequest, DefaultBody)
		defer s.Close()

		requestBody := "request body"
		content, err := httplib.Post(s.URL, requestBody)
		assert.NotNil(err)
		assert.Contains(err.Error(), "400 Bad Request")
		assert.Equal("", content)
	})
}

func TestPostWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	requestBody := "request body"
	content, err := httplib.PostWithHeaders(s.URL, requestBody, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.NoError(err)

	r := s.Request
	assert.Contains(r.Header, "X-Header1", "v1")
	assert.Contains(r.Header, "X-Header2", "v2")
	assert.Equal(requestBody, string(s.Body))
	assert.Equal(string(DefaultBody), content)
}

func TestPostBinary(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	requestBody := "request body"
	content, err := httplib.PostBinary(s.URL, []byte(requestBody))
	assert.NoError(err)
	assert.Equal(requestBody, string(s.Body))
	assert.Equal(DefaultBody, content)
}

func TestPostBinaryWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusOK, DefaultBody)
	defer s.Close()

	requestBody := "request body"
	content, err := httplib.PostBinaryWithHeaders(s.URL, []byte(requestBody), map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.NoError(err)

	r := s.Request
	assert.Contains(r.Header, "X-Header1", "v1")
	assert.Contains(r.Header, "X-Header2", "v2")
	assert.Equal(requestBody, string(s.Body))
	assert.Equal(DefaultBody, content)
}

func TestMustPost(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	requestBody := "request body"
	_ = httplib.MustPost(s.URL, requestBody)
	assert.Fail("should panic")
}

func TestMustPostWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		r := s.Request
		assert.Contains(r.Header, "X-Header1", "v1")
		assert.Contains(r.Header, "X-Header2", "v2")

		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	requestBody := "request body"
	_ = httplib.MustPostWithHeaders(s.URL, requestBody, map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.Fail("should panic")
}

func TestMustPostBinary(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	requestBody := "request body"
	_ = httplib.MustPostBinary(s.URL, []byte(requestBody))
	assert.Fail("should panic")
}

func TestMustPostBinaryWithHeaders(t *testing.T) {
	assert := assert.New(t)
	s := NewTestServer(http.StatusBadRequest, DefaultBody)
	defer s.Close()

	defer func() {
		r := s.Request
		assert.Contains(r.Header, "X-Header1", "v1")
		assert.Contains(r.Header, "X-Header2", "v2")

		err := recover()
		assert.NotNil(err)
		assert.Contains(fmt.Sprintf("%s", err), "400 Bad Request")
	}()

	requestBody := "request body"
	_ = httplib.MustPostBinaryWithHeaders(s.URL, []byte(requestBody), map[string]any{
		"X-Header1": "v1",
		"X-Header2": "v2",
	})
	assert.Fail("should panic")
}
