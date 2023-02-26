package headers

/* From letter O to T. */

// Origin
//
// Initiates a request for cross-origin resource sharing (asks server for Access-Control-* response fields).
//
// Class: Request field, Standard, Permanent
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

// Permissions-Policy
//
// To allow or disable different features or APIs of the browser.
//
// Class: Response field, Non-standard
//
// Example:
//
//	Permissions-Policy: fullscreen=(), camera=(), microphone=(), geolocation=(), interest-cohort=()
const PermissionsPolicy = "Permissions-Policy"

// Pragma
//
// Implementation-specific fields that may have various effects anywhere along the request-response chain.
//
// Class: Request field, Response field, Standard, Permanent
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

// P3P
//
// This field is supposed to set P3P policy, in the form of P3P:CP="your_compact_policy". However,
// P3P did not take off, most browsers have never fully implemented it, a lot of websites set this field
// with fake policy text, that was enough to fool browsers the existence of P3P policy and grant permissions
// for third party cookies.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	P3P: CP="This is not a P3P policy! See https://en.wikipedia.org/wiki/Special:CentralAutoLogin/P3P for more info."
const P3P = "P3P"

// Preference-Applied
//
// Indicates which Prefer tokens were honored by the server and applied to the processing of the request.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Preference-Applied: return=representation
//
// Standard:
//   - [RFC 7240]
//
// [RFC 7240]: https://datatracker.ietf.org/doc/html/rfc7240
const PreferenceApplied = "Preference-Applied"

// Proxy-Authenticate
//
// Request authentication to access the proxy.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Proxy-Authenticate: Basic
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const ProxyAuthenticate = "Proxy-Authenticate"

// Public-Key-Pins
//
// HTTP Public Key Pinning, announces hash of website's authentic TLS certificate.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Public-Key-Pins: max-age=2592000; pin-sha256="E9CZ9INDbd+2eRQozYqqbQ2yXLVKB9+xcprMF+44U1g=";
//
// Standard:
//   - [RFC 7469]
//
// [RFC 7469]: https://datatracker.ietf.org/doc/html/rfc7469
const PublicKeyPins = "Public-Key-Pins"

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

// Retry-After
//
// If an entity is temporarily unavailable, this instructs the client to try again later.
// Value could be a specified period of time (in seconds) or a HTTP-date.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Example 1: Retry-After: 120
//	Example 2: Retry-After: Fri, 07 Nov 2014 23:59:59 GMT
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const RetryAfter = "Retry-After"

// Refresh
//
// Used in redirection, or when a new resource has been created. This refresh redirects after 5 seconds.
// Header extension introduced by Netscape and supported by most web browsers. Defined by HTML Standard
//
// Class: Response field, Non-standard
//
// Example:
//
//	Refresh: 5; url=http://www.w3.org/pub/WWW/People.html
const Refresh = "Refresh"

// Report-To
//
// Instructs the user agent to store reporting endpoints for an origin.
//
// Class: Response field, Non-standard
//
// Example:
//
//	Report-To: {"group":"csp-endpoint", "max_age":10886400, "endpoints":[{"url":"https-url-of-site-which-collects-reports"}]}
const ReportTo = "Report-To"

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

// Server
//
// A name for the server.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Server: Apache/2.4.1 (Unix)
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Server = "Server"

// Set-Cookie
//
// An HTTP cookie/
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Set-Cookie: UserID=JohnDoe; Max-Age=3600; Version=1
//
// Standard:
//   - [RFC 6265]
//
// [RFC 6265]: https://datatracker.ietf.org/doc/html/rfc6265
const SetCookie = "Set-Cookie"

// Strict-Transport-Security
//
// A HSTS Policy informing the HTTP client how long to cache the HTTPS only policy and whether this applies to subdomains.
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Strict-Transport-Security: max-age=16070400; includeSubDomains
const StrictTransportSecurity = "Strict-Transport-Security"

// Status
//
// CGI header field specifying the status of the HTTP response.
// Normal HTTP responses use a separate "Status-Line" instead, defined by [RFC 9110].
//
// Class: Response field, Non-standard
//
// Example:
//
//	Status: 200 OK
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Status = "Status"

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

// Timing-Allow-Origin
//
// The Timing-Allow-Origin response header specifies origins that are allowed to see values of attributes retrieved
// via features of the Resource Timing API, which would otherwise be reported as zero due to cross-origin restrictions.
//
// Class: Response field, Non-standard
//
// Example:
//
//	Timing-Allow-Origin: *
//	Timing-Allow-Origin: <origin>[, <origin>]*
const TimingAllowOrigin = "Timing-Allow-Origin"

// Trailer
//
// The Trailer general field value indicates that the given set of header fields is present in the trailer of
// a message encoded with chunked transfer coding.
//
// Class: Request field, Response field, Standard, Permanent
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
// Class: Request field, Response field, Standard, Permanent
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

// Tk
//
// Tracking Status header, value suggested to be sent in response to a DNT(do-not-track), possible values:
//   - "!" — under construction
//   - "?" — dynamic
//   - "G" — gateway to multiple parties
//   - "N" — not tracking
//   - "T" — tracking
//   - "C" — tracking with consent
//   - "P" — tracking only if consented
//   - "D" — disregarding DNT
//   - "U" — updated
//
// Class: Response field, Standard, Permanent
//
// Example:
//
//	Tk: ?
const Tk = "Tk"
