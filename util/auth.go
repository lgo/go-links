package util

import (
	"net/http"
)

/**
 * AuthMethodType to use with the BasicKeyAuth
 * @type {[type]}
 */
type AuthMethodType int

const (
	AuthWithHeader AuthMethodType = iota
	AuthWithQueryParam
)

/**
 * AuthOptions for the authorization
 * AdminKey is the key that is checked against
 * AuthMethod is the method to check for the key
 */
type AuthOptions struct {
	AdminKey   string
	AuthMethod AuthMethodType
}

type handler func(w http.ResponseWriter, r *http.Request)

/**
 * BasicAuth wraps an HTTP handler to add basic authentication
 * Provided with a set of AuthOptions, it will either pass through the handler
 * or provide a 404 error
 */
func BasicAuth(o AuthOptions) func(handler) handler {
	return func(h handler) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			var authParam string
			if o.AuthMethod == AuthWithHeader {
				authParam = r.Header.Get("Golink-Auth")
			} else if o.AuthMethod == AuthWithQueryParam {
				authParam = r.URL.Query().Get("authkey")
			}

			if o.AdminKey == authParam {
				h(w, r)
			} else {
				http.Error(w, "Unsupported path", http.StatusNotFound)
			}
		}
	}
}
