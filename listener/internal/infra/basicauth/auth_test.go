package basicauth_test

import (
	"testing"

	"theskyinflames/core-tech/listener/internal/infra/basicauth"

	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	var (
		user        = "user"
		pwd         = "pwd"
		credentials = basicauth.Credentials{
			User:     user,
			Password: pwd,
		}
	)

	t.Run(`Given a wrong pair of user/password,
		when it's called,
		then an error is returned`, func(t *testing.T) {
		auth := basicauth.NewBasicAuth(credentials)
		require.ErrorIs(t, auth.Auth("", ""), basicauth.ErrAuth)
	})

	t.Run(`Given the right pair of user/password,
		when it's called,
		then no error is returned`, func(t *testing.T) {
		auth := basicauth.NewBasicAuth(credentials)
		require.NoError(t, auth.Auth(user, pwd))
	})
}
