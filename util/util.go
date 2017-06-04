package util

import (
	"net/http"
	"os"
)

/**
 * PathFromRequest extracts the shortened URL path from
 * an HTTP request
 */
func PathFromRequest(r *http.Request) string {
	return r.URL.Path[1:]
}

/**
 * Getenv will get an ENV variable with a default
 */
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
