package api_test

import (
	"net/http"
	"testing"

	"theskyinflames/core-tech/listener/internal/infra/api"
	"theskyinflames/core-tech/listener/internal/infra/basicauth"

	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	t.Run(`Given a basic authentication from HTTP header,
		when the basic auth header is not passed,
		then an error is returned`, func(t *testing.T) {
		url := "https://localhost:80"
		method := "GET"

		req, _ := http.NewRequest(method, url, nil)

		auth := api.NewBasicAuthFromHTTPRequest(nil)
		require.Error(t, auth.Auth(req))
	})

	t.Run(`Given a basic authentication from HTTP header with a basic authenticator that returns an error,
		when it's called,
		then an error is returned`, func(t *testing.T) {
		url := "https://localhost:80"
		method := "GET"

		req, _ := http.NewRequest(method, url, nil)
		req.Header.Add("Authorization", "Basic dXNlcjpwd2Q=")

		auth := api.NewBasicAuthFromHTTPRequest(&BasicAuthMock{
			AuthFunc: func(_, _ string) error {
				return basicauth.ErrAuth
			},
		})
		require.ErrorIs(t, auth.Auth(req), basicauth.ErrAuth)
	})

	t.Run(`Given a basic authentication from HTTP header,
		when it's called,
		then no error is returned`, func(t *testing.T) {
		url := "https://localhost:80"
		method := "GET"

		req, _ := http.NewRequest(method, url, nil)
		req.Header.Add("Authorization", "Basic dXNlcjpwd2Q=")

		auth := api.NewBasicAuthFromHTTPRequest(&BasicAuthMock{
			AuthFunc: func(_, _ string) error {
				return nil
			},
		})
		require.NoError(t, auth.Auth(req))
	})
}
