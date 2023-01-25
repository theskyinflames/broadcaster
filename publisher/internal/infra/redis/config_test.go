package redis_test

import (
	"os"
	"testing"

	"theskyinflames/core-tech/publisher/internal/infra/redis"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Run(`Given an environment with all variables defined,
			when it's called,
			then a new configuration is returned`, func(t *testing.T) {
		os.Setenv(redis.TopicEnvVar, "topic")
		os.Setenv(redis.AddrEnvVar, "addr")
		os.Setenv(redis.PasswordEnvVar, "pwd")

		c, err := redis.NewConfig()
		require.NoError(t, err)
		require.Equal(t, "topic", c.Topic)
		require.Equal(t, "addr", c.Addr)
		require.Equal(t, "pwd", c.Password)
	})
}
