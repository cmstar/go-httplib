// Package headers package provides constants of HTTP headers.
//
// The list of HTTP headers is taken from https://en.wikipedia.org/wiki/List_of_HTTP_header_fields .
package headers

/* From letter A to G. */

// Accept-CH
//
// Requests HTTP Client Hints.
//
// Class: Response field, Standard, Experimental
//
// Example:
//
//	Accept-CH: UA, Platform
//
// Standard:
//   - [RFC 8942]
//
// [RFC 8942]: https://datatracker.ietf.org/doc/html/rfc8942
const AcceptCH = "Accept-CH"

// Access-Control-Allow-Origin
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Access-Control-Allow-Origin: *
//
// Standard:
//   - [RFC 7480]
//
// [cross-origin resource sharing]: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlAllowOrigin = "Access-Control-Allow-Origin"

// Access-Control-Allow-Credentials
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard
//
// Standard:
//   - [RFC 7480]
//
// [cross-origin resource sharing]: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlAllowCredentials = "Access-Control-Allow-Credentials"

// Access-Control-Expose-Headers
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard
//
// Standard:
//   - [RFC 7480]
//
// [cross-origin resource sharing]: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlExposeHeaders = "Access-Control-Expose-Headers"

// Access-Control-Max-Age
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard
//
// Standard:
//   - [RFC 7480]
//
// [cross-origin resource sharing]: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlMaxAge = "Access-Control-Max-Age"

// Access-Control-Allow-Methods
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard
//
// Standard:
//   - [RFC 7480]
//
// [cross-origin resource sharing]: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlAllowMethods = "Access-Control-Allow-Methods"

// Access-Control-Allow-Headers
//
// Specifying which web sites can participate in [cross-origin resource sharing].
//
// Class: Response field, Standard
//
// Standard:
//   - [RFC 7480]
//
// [RFC 7480]: https://datatracker.ietf.org/doc/html/rfc7480
const AccessControlAllowHeaders = "Access-Control-Allow-Headers"

// Accept-Patch
//
// Specifies which patch document formats this server supports.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Accept-Patch: text/example;charset=utf-8
//
// Standard:
//   - [RFC 5789]
//
// [RFC 5789]: https://datatracker.ietf.org/doc/html/rfc5789
const AcceptPatch = "Accept-Patch"

// Accept-Ranges
//
// What partial content range types this server supports via byte serving.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Accept-Ranges: bytes
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const AcceptRanges = "Accept-Ranges"

// Age
//
// The age the object has been in a proxy cache in seconds.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Age: 12
//
// Standard:
//   - [RFC 9111]
//
// [RFC 9111]: https://datatracker.ietf.org/doc/html/rfc9111
const Age = "Age"

// Allow
//
// Valid methods for a specified resource. To be used for a 405 Method not allowed
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Allow: GET, HEAD
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Allow = "Allow"

// Alt-Svc
//
// A server uses "Alt-Svc" header (meaning Alternative Services) to indicate that its resources can also
// be accessed at a different network location (host or port) or using a different protocol
//
// When using HTTP/2, servers should instead send an ALTSVC frame.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Alt-Svc: http/1.1="http2.example.com:8001"; ma=7200
const AltSvc = "Alt-Svc"

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
// Request: Used to specify directives that must be obeyed by all caching mechanisms along the request-response chain.
//
// Response: Tells all caching mechanisms from server to client whether they may cache this object. Measured in seconds.
//
// Class: Request field, Response field, Standard, Permanent
//
// Example:
//
//	Cache-Control: no-cache
//	Cache-Control: max-age=3600
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
// Class: Request field, Response field, Standard, Permanent
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
// Class: Request field, Response field, Standard, Permanent
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

// Content-MD5
//
// A Base64-encoded binary MD5 sum of the content of the body.
//
// Class: Request field, Response field, Standard, Obsolete
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
// Request: The Media type (MIME) of the body of the request (used with POST and PUT requests).
//
// Response: The Media type of the body.
//
// Class: Request field, Response field, Standard, Permanent
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
// Class: Request field, Standard, Permanent
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

// Content-Disposition
//
// An opportunity to raise a "File Download" dialogue box for a known MIME type with binary format or suggest
// a filename for dynamic content. Quotes are necessary with special characters.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Content-Disposition: attachment; filename="name.ext"
//
// Standard:
//   - [RFC 2616]
//   - [RFC 4021]
//   - [RFC 6266]
//
// [RFC 2616]: https://datatracker.ietf.org/doc/html/rfc2616
// [RFC 4021]: https://datatracker.ietf.org/doc/html/rfc4021
// [RFC 6266]: https://datatracker.ietf.org/doc/html/rfc6266
const ContentDisposition = "Content-Disposition"

// Content-Language
//
// The natural language or languages of the intended audience for the enclosed content.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Content-Language: da
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentLanguage = "Content-Language"

// Content-Location
//
// An alternate location for the returned data.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Content-Location: /index.htm
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentLocation = "Content-Location"

// Content-Range
//
// Where in a full body message this partial message belongs.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Content-Range: bytes 21010-47021/47022
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ContentRange = "Content-Range"

// Content-Security-Policy
//
// Content Security Policy definition.
//
// Class: Response field, Non-standard
//
// Example:
//
//	X-WebKit-CSP: default-src 'self'
const ContentSecurityPolicy = "Content-Security-Policy"

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

// Delta-Base
//
// Specifies the delta-encoding entity tag of the response.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Delta-Base: "abc"
//
// Standard:
//   - [RFC 3229]
//
// [RFC 3229]: https://datatracker.ietf.org/doc/html/rfc3229
const DeltaBase = "Delta-Base"

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

// Expect-CT
//
// Notify to prefer to enforce Certificate Transparency.
//
// Class: Response field, Non-standard
//
// Example:
//
//	Expect-CT: max-age=604800, enforce, report-uri="https://example.example/report"
const ExpectCT = "Expect-CT"

// ETag
//
// An identifier for a specific version of a resource, often a message digest.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	ETag: "737060cd8c284d8af7ad3082f209582d"
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ETag = "ETag"

// Expires
//
// Gives the date/time after which the response is considered stale (in "HTTP-date" format as defined by [RFC 9110])
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Expires: Thu, 01 Dec 1994 16:00:00 GMT
//
// Standard:
//   - [RFC 9111]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
// [RFC 9111]: https://datatracker.ietf.org/doc/html/rfc9111
const Expires = "Expires"

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
// Non-standard header field used by Microsoft applications and load-balancers.
//
// Class: Request field, Non-standard
//
// Example:
//
//	Front-End-Https: on
const FrontEndHttps = "Front-End-Https"
