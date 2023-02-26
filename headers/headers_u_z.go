package headers

/* From letter U to Z. */

// User-Agent
//
// The user agent string of the user agent.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const UserAgent = "User-Agent"

// Upgrade
//
// Ask the server to upgrade to another protocol.
//
// Must not be used in HTTP/2.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Upgrade: h2c, HTTPS/1.3, IRC/6.9, RTA/x11, websocket
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Upgrade = "Upgrade"

// Upgrade-Insecure-Requests
//
// Tells a server which (presumably in the middle of a HTTP -> HTTPS migration) hosts mixed content that
// the client would prefer redirection to HTTPS and can handle Content-Security-Policy: upgrade-insecure-requests
// Must not be used with HTTP/2
//
// Class: Request field, Non-standard
//
// Example:
//
//	Upgrade-Insecure-Requests: 1
const UpgradeInsecureRequests = "Upgrade-Insecure-Requests"

// Via
//
// Informs the server of proxies through which the request was sent.
//
// Class: Request field, Standard, Permanent
//
// Example:
//
//	Via: 1.0 fred, 1.1 example.com (Apache/1.1)
//
// Standard:
//   - [RFC 9110]
//
// [RFC 9110]: https://datatracker.ietf.org/doc/html/rfc9110
const Via = "Via"

// Warning
//
// A general warning about possible problems with the entity body.
//
// Class: Request field, Standard, Obsolete[21]
//
// Example:
//
//	Warning: 199 Miscellaneous warning
//
// Standard:
//   - [RFC 7234]
//   - [RFC 9111]
//
// [RFC 7234]: https://datatracker.ietf.org/doc/html/rfc7234
// [RFC 9111]: https://datatracker.ietf.org/doc/html/rfc9111
const Warning = "Warning"

// X-Requested-With
//
// Mainly used to identify Ajax requests (most JavaScript frameworks send this field with value of XMLHttpRequest);
// also identifies Android apps using WebView
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Requested-With: XMLHttpRequest
const XRequestedWith = "X-Requested-With"

// X-Forwarded-For
//
// A de facto standard for identifying the originating IP address of a client connecting to a web server through
// an HTTP proxy or load balancer. Superseded by Forwarded header.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Forwarded-For: client1, proxy1, proxy2
//	X-Forwarded-For: 129.78.138.66, 129.78.64.103
const XForwardedFor = "X-Forwarded-For"

// X-Forwarded-Host
//
// A de facto standard for identifying the original host requested by the client in the Host HTTP request header,
// since the host name and/or port of the reverse proxy (load balancer) may differ from the origin server handling the request.
// Superseded by Forwarded header.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Forwarded-Host: en.wikipedia.org:8080
//	X-Forwarded-Host: en.wikipedia.org
const XForwardedHost = "X-Forwarded-Host"

// X-Forwarded-Proto
//
// A de facto standard for identifying the originating protocol of an HTTP request, since a reverse proxy (or a load balancer)
// may communicate with a web server using HTTP even if the request to the reverse proxy is HTTPS. An alternative form of
// the header (X-ProxyUser-Ip) is used by Google clients talking to Google servers. Superseded by Forwarded header.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Forwarded-Proto: https
const XForwardedProto = "X-Forwarded-Proto"

// X-Http-Method-Override
//
// Requests a web application to override the method specified in the request (typically POST) with the method given in
// the header field (typically PUT or DELETE). This can be used when a user agent or firewall prevents PUT or DELETE methods
// from being sent directly (note that this is either a bug in the software component, which ought to be fixed, or an intentional
// configuration, in which case bypassing it may be the wrong thing to do).
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-HTTP-Method-Override: DELETE
const XHttpMethodOverride = "X-Http-Method-Override"

// X-ATT-DeviceId
//
// Allows easier parsing of the MakeModel/Firmware that is usually found in the User-Agent String of AT&T Devices
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Att-DeviceId: GT-P7320/P7320XX5PG
const XAttDeviceId = "X-ATT-DeviceId"

// X-Wap-Profile
//
// Links to an XML file on the Internet with a full description and details about the device currently connecting.
// In the example to the right is an XML file for an AT&T Samsung Galaxy S2.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Wap-Profile: http://wap.samsungmobile.com/uaprof/SGH-I777.xml
const XWapProfile = "X-Wap-Profile"

// X-UIDH
//
// Server-side deep packet insertion of a unique ID identifying customers of Verizon Wireless; also known as "perma-cookie" or "supercookie"
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-UIDH: ...
const XUIDH = "X-UIDH"

// X-Csrf-Token
//
// Used to prevent cross-site request forgery. Alternative header names are: X-CSRFToken and X-XSRF-TOKEN
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Csrf-Token: i8XNjC4b8KVok4uw5RftR38Wgp2BFwQl
const XCsrfToken = "X-Csrf-Token"

// X-Request-ID
//
// Correlates HTTP requests between a client and server.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Request-ID: f058ebd6-02f7-4d3f-942e-904344e8cde5
const XRequestID = "X-Request-ID"

// X-Correlation-ID
//
// Correlates HTTP requests between a client and server.
//
// Class: Request field, Non-standard
//
// Example:
//
//	X-Correlation-ID: f058ebd6-02f7-4d3f-942e-904344e8cde5
const XCorrelationID = "X-Correlation-ID"
