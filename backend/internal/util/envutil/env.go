package envutil

import (
	"os"
	"strconv"
	"strings"
)

// GetString returns the value of the environment variable or the default value
func GetString(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

// GetInt returns the value of the environment variable as an int or the default value
func GetInt(key string, defaultValue int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultValue
}

// GetBool returns the value of the environment variable as a bool or the default value
func GetBool(key string, defaultValue bool) bool {
	v := strings.ToLower(os.Getenv(key))
	switch v {
	case "1", "t", "true", "on", "y", "yes":
		return true
	case "0", "f", "false", "off", "n", "no":
		return false
	}
	return defaultValue
}
