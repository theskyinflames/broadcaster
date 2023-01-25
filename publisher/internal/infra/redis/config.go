package redis

import "theskyinflames/core-tech/publisher/internal/helpers"

const defaultDB = 0

// Config is a Redis client config
type Config struct {
	Topic    string
	Addr     string
	Password string
	Db       int
}

// Redis client config variables names
const (
	TopicEnvVar    = "REDIS_TOPIC"
	AddrEnvVar     = "REDIS_ADDR"
	PasswordEnvVar = "REDIS_PASSWORD"
)

// NewConfig is a constructor
func NewConfig() (Config, error) {
	topic, err := helpers.GetEnv(TopicEnvVar)
	if err != nil {
		return Config{}, err
	}
	addr, err := helpers.GetEnv(AddrEnvVar)
	if err != nil {
		return Config{}, err
	}
	password, err := helpers.GetEnv(PasswordEnvVar)
	if err != nil {
		return Config{}, err
	}
	return Config{
		Topic:    topic,
		Addr:     addr,
		Password: password,
		Db:       defaultDB,
	}, nil
}
