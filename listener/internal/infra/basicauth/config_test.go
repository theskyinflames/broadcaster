package basicauth_test

import (
	"os"
	"testing"

	"theskyinflames/core-tech/listener/internal/infra/basicauth"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Run(`Given an environment with all variables defined,
			when it's called,
			then a new configuration is returned`, func(t *testing.T) {
		os.Setenv(basicauth.AuthUserEnvVar, "user")
		os.Setenv(basicauth.AuthPasswordEnvVar, "pwd")

		c, err := basicauth.NewConfig()
		require.NoError(t, err)
		require.Equal(t, "user", c.Credentials.User)
		require.Equal(t, "pwd", c.Credentials.Password)
	})
}
