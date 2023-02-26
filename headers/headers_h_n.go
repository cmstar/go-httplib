package headers

/* From letter H to N. */

// Host
//
// The domain name of the server (for virtual hosting), and the TCP port number on which the server is listening.
// The port number may be omitted if the port is the standard port for the service requested.
//
// Mandatory since HTTP/1.1. If the request is generated directly in HTTP/2, it should not be used.
//
// Class: Request field, Standard
//
// Example:
//
//	Host: en.wikipedia.org:8080
//	Host: en.wikipedia.org
const Host = "Host"

// HTTP2-Settings
//
// A request that upgrades from HTTP/1.1 to HTTP/2 MUST include exactly one HTTP2-Setting header field.
// The HTTP2-Settings header field is a connection-specific header field that includes parameters that
// govern the HTTP/2 connection, provided in anticipation of the server accepting the request to upgrade.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	HTTP2-Settings: token64
const HTTP2Settings = "HTTP2-Settings"

// If-Match
//
// Only perform the action if the client supplied entity matches the same entity on the server.
// This is mainly for methods like PUT to only update a resource if it has not been modified since the user last updated it.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	If-Match: "737060cd8c284d8af7ad3082f209582d"
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const IfMatch = "If-Match"

// If-Modified-Since
//
// Allows a 304 Not Modified to be returned if content is unchanged.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	If-Modified-Since: Sat, 29 Oct 1994 19:43:31 GMT
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const IfModifiedSince = "If-Modified-Since"

// If-None-Match
//
// Allows a 304 Not Modified to be returned if content is unchanged, see [HTTP ETag].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	If-None-Match: "737060cd8c284d8af7ad3082f209582d"
//
// Standard:
//   - [RFC 9110]
//
// [HTTP ETag]: https://en.wikipedia.org/wiki/HTTP_ETag
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const IfNoneMatch = "If-None-Match"

// If-Range
//
// If the entity is unchanged, send me the part(s) that I am missing; otherwise, send me the entire new entity.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	If-Range: "737060cd8c284d8af7ad3082f209582d"
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const IfRange = "If-Range"

// If-Unmodified-Since
//
// Only send the response if the entity has not been modified since a specific time.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	If-Unmodified-Since: Sat, 29 Oct 1994 19:43:31 GMT
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const IfUnmodifiedSince = "If-Unmodified-Since"

// Max-Forwards
//
// Limit the number of times the message can be forwarded through proxies or gateways.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Max-Forwards: 10
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const MaxForwards = "Max-Forwards"
