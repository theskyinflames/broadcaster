package api

import (
	"fmt"
	"net/http"
)

//go:generate moq -stub -out zmock_basicauth_from_http_test.go -pkg api_test . BasicAuth

// BasicAuth provides basic access authentication based on user/password
type BasicAuth interface {
	Auth(user, pass string) error
}

// BasicAuthFromHTTPRequest provides basic access authentication based on HTTP header
type BasicAuthFromHTTPRequest struct {
	basicAuth BasicAuth
}

// NewBasicAuthFromHTTPRequest is a constructor
func NewBasicAuthFromHTTPRequest(ba BasicAuth) BasicAuthFromHTTPRequest {
	return BasicAuthFromHTTPRequest{basicAuth: ba}
}

// Auth implements Authenticator interface used for the authn middleware
func (b BasicAuthFromHTTPRequest) Auth(r *http.Request) error {
	user, pass, ok := r.BasicAuth()
	if !ok {
		return fmt.Errorf("basic auth header missing")
	}
	return b.basicAuth.Auth(user, pass)
}
