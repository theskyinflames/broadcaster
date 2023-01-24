package api

import "net/http"

//go:generate moq -stub -out zmock_middleware_test.go -pkg api_test . Authenticator

// Authenticator provides authentication
type Authenticator interface {
	Auth(r *http.Request) error
}

// HTTPHandleFuncMiddleware is an http.HandleFunc middleware
type HTTPHandleFuncMiddleware func(http.HandlerFunc) http.HandlerFunc

// AuthnMiddleware is a middleware in charge of checking for the authentication
func AuthnMiddleware(auth Authenticator) HTTPHandleFuncMiddleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := auth.Auth(r); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			hf.ServeHTTP(w, r)
		}
	}
}
