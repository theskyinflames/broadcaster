package basicauth

import "theskyinflames/core-tech/listener/internal/helpers"

// Credentials are api credentials
type Credentials struct {
	User     string
	Password string
}

// Config is api config
type Config struct {
	Credentials Credentials
}

// Auth environment variables' names
const (
	AuthUserEnvVar     = "AUTH_USER"
	AuthPasswordEnvVar = "AUTH_PASSWORD"
)

// NewConfig is a constructor
func NewConfig() (Config, error) {
	user, err := helpers.GetEnv(AuthUserEnvVar)
	if err != nil {
		return Config{}, err
	}
	pass, err := helpers.GetEnv(AuthPasswordEnvVar)
	if err != nil {
		return Config{}, err
	}
	return Config{
		Credentials: Credentials{
			User:     user,
			Password: pass,
		},
	}, nil
}
