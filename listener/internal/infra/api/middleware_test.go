package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"theskyinflames/core-tech/listener/internal/infra/api"

	"github.com/stretchr/testify/require"
)

func TestAuthnMiddleware(t *testing.T) {
	t.Run(`Given an authenticator that returns an error,
			when it's called,
			then a HTTP code 401 is returned`, func(t *testing.T) {
		var (
			authenticator = &AuthenticatorMock{
				AuthFunc: func(_ *http.Request) error {
					return errors.New("")
				},
			}
			hf = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
			mw = api.AuthnMiddleware(authenticator)
		)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		mw(hf).ServeHTTP(rr, req)

		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run(`Given a middleware authenticator,
			when it's called,
			then the underlying HTTP handler is called`, func(t *testing.T) {
		var (
			authenticator = &AuthenticatorMock{
				AuthFunc: func(_ *http.Request) error {
					return nil
				},
			}
			called bool
			hf     = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
			})
			mw = api.AuthnMiddleware(authenticator)
		)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		mw(hf).ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		require.True(t, called)
	})
}
