package envutil

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// GetString returns the value of the environment variable or the default value.
// It trims leading and trailing whitespace from the environment value.
func GetString(key, defaultValue string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return defaultValue
}

// GetInt returns the value of the environment variable as an int or the default value.
// It trims leading and trailing whitespace from the environment value.
func GetInt(key string, defaultValue int) int {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err == nil {
		return i
	}
	log.Printf("Warning: envutil.GetInt invalid value for %s: %q (using default %d)", key, val, defaultValue)
	return defaultValue
}

// GetBool returns the value of the environment variable as a bool or the default value
func GetBool(key string, defaultValue bool) bool {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultValue
	}
	v := strings.ToLower(val)
	switch v {
	case "true", "1":
		return true
	case "false", "0":
		return false
	}
	log.Printf("Warning: envutil.GetBool invalid value for %s: %q (using default %v)", key, val, defaultValue)
	return defaultValue
}
