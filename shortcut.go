package httplib

/*
Shortcut to send HTTP requests.
*/

/* ==== GET ==== */

// Get sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Returns an error if the status code is not 200 OK.
func Get(uri string) (string, error) {
	req := NewBuilder("GET", uri)
	return req.ReadString()
}

// GetWithHeaders sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Returns an error if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func GetWithHeaders(uri string, headers map[string]any) (string, error) {
	req := NewBuilder("GET", uri).WithHeaders(headers)
	return req.ReadString()
}

// GetBinary sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Returns an error if the status code is not 200 OK.
func GetBinary(uri string) ([]byte, error) {
	req := NewBuilder("GET", uri)
	return req.ReadBinary()
}

// GetBinaryWithHeaders sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Returns an error if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func GetBinaryWithHeaders(uri string, headers map[string]any) ([]byte, error) {
	req := NewBuilder("GET", uri).WithHeaders(headers)
	return req.ReadBinary()
}

// MustGet sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Panics if the status code is not 200 OK.
func MustGet(uri string) string {
	req := NewBuilder("GET", uri)
	return req.MustReadString()
}

// MustGetWithHeaders sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Panics if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func MustGetWithHeaders(uri string, headers map[string]any) string {
	req := NewBuilder("GET", uri).WithHeaders(headers)
	return req.MustReadString()
}

// MustGetBinary sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Panics if the status code is not 200 OK.
func MustGetBinary(uri string) []byte {
	req := NewBuilder("GET", uri)
	return req.MustReadBinary()
}

// MustGetBinaryWithHeaders sends a GET request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Panics if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func MustGetBinaryWithHeaders(uri string, headers map[string]any) []byte {
	req := NewBuilder("GET", uri).WithHeaders(headers)
	return req.MustReadBinary()
}

/* ==== POST ==== */

// Post sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Returns an error if the status code is not 200 OK.
func Post(uri string, body string) (string, error) {
	req := NewBuilder("POST", uri).SetStringBody(body)
	return req.ReadString()
}

// PostWithHeaders sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Returns an error if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func PostWithHeaders(uri string, body string, headers map[string]any) (string, error) {
	req := NewBuilder("POST", uri).WithHeaders(headers).SetStringBody(body)
	return req.ReadString()
}

// PostBinary sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Returns an error if the status code is not 200 OK.
func PostBinary(uri string, body []byte) ([]byte, error) {
	req := NewBuilder("POST", uri).SetBinaryBody(body)
	return req.ReadBinary()
}

// PostBinaryWithHeaders sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Returns an error if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func PostBinaryWithHeaders(uri string, body []byte, headers map[string]any) ([]byte, error) {
	req := NewBuilder("POST", uri).WithHeaders(headers).SetBinaryBody(body)
	return req.ReadBinary()
}

// MustPost sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Panics if the status code is not 200 OK.
func MustPost(uri string, body string) string {
	req := NewBuilder("POST", uri).SetStringBody(body)
	return req.MustReadString()
}

// MustPostWithHeaders sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a string.
// Panics if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func MustPostWithHeaders(uri string, body string, headers map[string]any) string {
	req := NewBuilder("POST", uri).WithHeaders(headers).SetStringBody(body)
	return req.MustReadString()
}

// MustPostBinary sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Panics if the status code is not 200 OK.
func MustPostBinary(uri string, body []byte) []byte {
	req := NewBuilder("POST", uri).SetBinaryBody(body)
	return req.MustReadBinary()
}

// MustPostBinaryWithHeaders sends a POST request.
// If the status code of the response is 200 OK, returns the whole response body as a slice of byte.
// Panics if the status code is not 200 OK.
//
// A map gives out the HTTP headers. If the map is nil, it is ignored.
// All header names will be converted to the Header-Naming-Style.
//
func MustPostBinaryWithHeaders(uri string, body []byte, headers map[string]any) []byte {
	req := NewBuilder("POST", uri).WithHeaders(headers).SetBinaryBody(body)
	return req.MustReadBinary()
}
