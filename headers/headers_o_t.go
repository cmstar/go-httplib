package headers

/* From letter O to T. */

// Origin
//
// Initiates a request for cross-origin resource sharing (asks server for Access-Control-* response fields).
//
// Class: Request field, Standard, Permanent: standard
//
// Example:
//
//	Origin: http://www.example-social-network.com
//
// Standard:
//   - [RFC 6454]
//
// [RFC 6454]: https://datatracker.ietf.org/doc/html/rfc6454
const Origin = "Origin"

// Pragma
//
// Implementation-specific fields that may have various effects anywhere along the request-response chain.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Pragma: no-cache
//
// Standard:
//   - [RFC 9111]
//
// [RFC 9111]: https://datatracker.ietf.org/doc/html/rfc9111
const Pragma = "Pragma"

// Prefer
//
// Allows client to request that certain behaviors be employed by a server while processing a request.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Prefer: return=representation
//
// Standard:
//   - [RFC 7240]
//
// [RFC 7240]: https://datatracker.ietf.org/doc/html/rfc7240
const Prefer = "Prefer"

// Proxy-Authorization
//
// Authorization credentials for connecting to a proxy.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Proxy-Authorization: Basic QWxhZGRpbJpVcGVuIHNlc2FtZQ==
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ProxyAuthorization = "Proxy-Authorization"

// Proxy-Connection
//
// Implemented as a misunderstanding of the HTTP specifications. Common because of mistakes in implementations of
// early HTTP versions. Has exactly the same functionality as standard Connection field.
//
// Must not be used with HTTP/2.
//
// Class: Request field, Non-standard
//
// Example:
//
//	Proxy-Connection: keep-alive
const ProxyConnection = "Proxy-Connection"

// Range
//
// Request only part of an entity. Bytes are numbered from 0. See [Byte serving].
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Range: bytes=500-999
//
// Standard:
//   - [RFC 9110]
//
// [Byte serving]: https://en.wikipedia.org/wiki/Byte_serving
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Range = "Range"

// Referer
//
// This is the address of the previous web page from which a link to the currently requested page was followed.
// (The word "referrer" has been misspelled in the RFC as well as in most implementations to the point that it
// has become standard usage and is considered correct terminology)
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Referer: http://en.wikipedia.org/wiki/Main_Page
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Referer = "Referer"

// Save-Data
//
// The Save-Data client hint request header available in Chrome, Opera, and Yandex browsers lets developers
// deliver lighter, faster applications to users who opt-in to data saving mode in their browser.
//
// Class: Request field, Non-standard
//
// Example:
//
//	Save-Data: on
const SaveData = "Save-Data"

// TE
//
// The transfer encodings the user agent is willing to accept: the same values as for the response header field
// Transfer-Encoding can be used, plus the "trailers" value (related to the "chunked" transfer method) to notify
// the server it expects to receive additional fields in the trailer after the last, zero-sized, chunk.
//
// Only trailers is supported in HTTP/2.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	TE: trailers, deflate
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const TE = "TE"

// Trailer
//
// The Trailer general field value indicates that the given set of header fields is present in the trailer of
// a message encoded with chunked transfer coding.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Trailer: Max-Forwards
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Trailer = "Trailer"

// Transfer-Encoding
//
// The form of encoding used to safely transfer the entity to the user. Currently defined methods are:
// chunked, compress, deflate, gzip, identity.
//
// Must not be used with HTTP/2.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Transfer-Encoding: chunked
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const TransferEncoding = "Transfer-Encoding"
