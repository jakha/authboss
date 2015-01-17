package authboss

import "net/http"

const (
	// SessionKey is the primarily used key by authboss.
	SessionKey = "uid"
	// HalfAuthKey is used for sessions that have been authenticated by
	// the remember module. This serves as a way to force full authentication
	// by denying half-authed users acccess to sensitive areas.
	HalfAuthKey = "halfauth"
)

// ClientStorer should be able to store values on the clients machine. Cookie and
// Session storers are built with this interface.
type ClientStorer interface {
	Put(key, value string)
	Get(key string) (string, bool)
	Del(key string)
}

// CookieStoreMaker is used to create a cookie storer from an http request. Keep in mind
// security considerations for your implementation, Secure, HTTP-Only, etc flags.
type CookieStoreMaker func(http.ResponseWriter, *http.Request) ClientStorer

// SessionStoreMaker is used to create a session storer from an http request.
// It must be implemented to satisfy certain modules (auth, remember primarily).
// It should be a secure storage of the session. This means if it represents a cookie-based session
// storage these cookies should be signed in order to prevent tampering, or they should be encrypted.
type SessionStoreMaker func(http.ResponseWriter, *http.Request) ClientStorer