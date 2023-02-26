// Package headers package provides constants of HTTP headers.
//
// The list of HTTP headers is taken from https://en.wikipedia.org/wiki/List_of_HTTP_header_fields .
package headers

/* From letter A to G. */

// A-IM
//
// Acceptable instance-manipulations for the request.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	A-IM: feed
//
// Standard:
//   - [RFC 3229]
//
// [RFC 3229]: https://datatracker.ietf.org/doc/html/rfc3229
const AIM = "A-IM"

// Accept
//
// Media type(s) that is/are acceptable for the response. See [Content negotiation].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Accept: text/html
//
// Standard:
//   - [RFC 9110]
//
// [Content negotiation]: https://en.wikipedia.org/wiki/Content_negotiation
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Accept = "Accept"

// Accept-Charset
//
// Character sets that are acceptable.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Accept-Charset: utf-8
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const AcceptCharset = "Accept-Charset"

// Accept-Datetime
//
// Acceptable version in time.
//
// Class: Request field, Standard, Provisional
//
// Example:
//
//	Accept-Datetime: Thu, 31 May 2007 20:35:00 GMT
//
// Standard:
//   - [RFC 7089]
//
// [RFC 7089]: https://datatracker.ietf.org/doc/html/rfc7089
const AcceptDatetime = "Accept-Datetime"

// Accept-Encoding
//
// List of acceptable encodings. See [HTTP compression].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Accept-Encoding: gzip, deflate
//
// Standard:
//   - [RFC 9110]
//
// [HTTP compression]: https://en.wikipedia.org/wiki/HTTP_compression
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const AcceptEncoding = "Accept-Encoding"

// Accept-Language
//
// List of acceptable human languages for response. See [Content negotiation].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Accept-Language: en-US
//
// Standard:
//   - [RFC 9110]
//
// [Content negotiation]: https://en.wikipedia.org/wiki/Content_negotiation
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const AcceptLanguage = "Accept-Language"

// Access-Control-Request-Method
//
// Initiates a request for cross-origin resource sharing with Origin.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Access-Control-Request-Method: GET
const AccessControlRequestMethod = "Access-Control-Request-Method"

// Access-Control-Request-Headers
//
// Initiates a request for cross-origin resource sharing with Origin.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Access-Control-Request-Method: GET
const AccessControlRequestHeaders = "Access-Control-Request-Headers"

// Authorization
//
// Authentication credentials for HTTP authentication.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Authorization: Basic QWxhZGRpbApVcGVuIHNlc2FtZQ==
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Authorization = "Authorization"

// Cache-Control
//
// Used to specify directives that must be obeyed by all caching mechanisms along the request-response chain.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Cache-Control: no-cache
//
// Standard:
//   - [RFC 9111]
//
// [RFC 9111]: https://datatracker.ietf.org/doc/html/rfc9111
const CacheControl = "Cache-Control"

// Connection
//
// Control options for the current connection and list of hop-by-hop request fields.
//
// Must not be used with HTTP/2.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Connection: keep-alive
//	Connection: Upgrade
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Connection = "Connection"

// Content-Encoding
//
// The type of encoding used on the data. See [HTTP compression].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Content-Encoding: gzip
//
// Standard:
//   - [RFC 9110]
//
// [HTTP compression]: https://en.wikipedia.org/wiki/HTTP_compression
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentEncoding = "Content-Encoding"

// Content-Length
//
// The length of the request body in octets (8-bit bytes).
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Content-Length: 348
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentLength = "Content-Length"

// Content-MD5
//
// A Base64-encoded binary MD5 sum of the content of the request body.
//
// Class: Request field, Standard, Obsolete
//
// Example:
//
//	Content-MD5: Q2hlY2sgSW50ZWdYaXR5IQ==
//
// Standard:
//   - [RFC 1544]
//   - [RFC 1864]
//   - [RFC 4021]
//
// [RFC 1544]: https://datatracker.ietf.org/doc/html/rfc1544
// [RFC 1864]: https://datatracker.ietf.org/doc/html/rfc1864
// [RFC 4021]: https://datatracker.ietf.org/doc/html/rfc4021
const ContentMD5 = "Content-MD5"

// Content-Type
//
// The Media type of the body of the request (used with POST and PUT requests).
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Content-Type: application/x-www-form-urlencoded
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentType = "Content-Type"

// Cookie
//
// An HTTP cookie previously sent by the server with Set-Cookie.
//
// Class: Request field, Standard, Permanent: standard
//
// Example:
//
//	Cookie: $Version=1; Skin=new;
//
// Standard:
//   - [RFC 2965]
//   - [RFC 6265]
//
// [RFC 2965]: https://datatracker.ietf.org/doc/html/rfc2965
// [RFC 6265]: https://datatracker.ietf.org/doc/html/rfc6265
const Cookie = "Cookie"

// Correlation-ID
//
// Correlates HTTP requests between a client and server.
//
// Class: Request field, Non-standard
//
// Example:
//
//	Correlation-ID: f058ebd6-02f7-4d3f-942e-904344e8cde5
const CorrelationID = "Correlation-ID"

// Date
//
// The date and time at which the message was originated (in "HTTP-date" format as defined by
// [RFC 9110]: HTTP Semantics, section 5.6.7 "Date/Time Formats").
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Date: Tue, 15 Nov 1994 08:12:31 GMT
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Date = "Date"

// DNT
//
// Requests a web application to disable their tracking of a user. This is Mozilla's version of the
// X-Do-Not-Track header field (since Firefox 4.0 Beta 11). Safari and IE9 also have support for this field.
// On March 7, 2011, a draft proposal was submitted to IETF.
// The W3C Tracking Protection Working Group is producing a specification.
//
// Class: Request field, Non-standard
//
// Example:
//
//	DNT: 1 (Do Not Track Enabled)
//	DNT: 0 (Do Not Track Disabled)
const DNT = "DNT"

// Expect
//
// Indicates that particular server behaviors are required by the client.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Expect: 100-continue
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Expect = "Expect"

// Forwarded
//
// Disclose original information of a client connecting to a web server through an HTTP proxy.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Forwarded: for=192.0.2.60;proto=http;by=203.0.113.43 Forwarded: for=192.0.2.43, for=198.51.100.17
//
// Standard:
//   - [RFC 7239]
//
// [RFC 7239]: https://datatracker.ietf.org/doc/html/rfc7239
const Forwarded = "Forwarded"

// From
//
// The email address of the user making the request.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	From: user@example.com
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const From = "From"

// Front-End-Https
//
// # Non-standard header field used by Microsoft applications and load-balancers
//
// Class: Request field, Non-standard
//
// Example:
//
//	Front-End-Https: on
const FrontEndHttps = "Front-End-Https"
