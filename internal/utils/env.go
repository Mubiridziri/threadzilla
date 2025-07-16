package utils

import (
	"os"
	"strconv"
)

func GetEnvStr(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		parsedValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}

		return parsedValue
	}

	return defaultValue
}
