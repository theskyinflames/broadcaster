package helpers

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", fmt.Errorf("env variable %s not found", key)
	}
	return value, nil
}
