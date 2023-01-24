package basicauth

import (
	"errors"
)

// BasicAuth adds basic-auth authentication
type BasicAuth struct {
	credentials Credentials
}

// NewBasicAuth is a constructor
func NewBasicAuth(c Credentials) BasicAuth {
	return BasicAuth{credentials: c}
}

// ErrAuth is an authentication error
var ErrAuth = errors.New("authentication failed")

// Auth makes basic auth authentication from http request
func (ba BasicAuth) Auth(user, pass string) error {
	if user != ba.credentials.User || pass != ba.credentials.Password {
		return ErrAuth
	}
	return nil
}
