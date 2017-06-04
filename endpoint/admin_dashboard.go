package endpoint

import (
	"fmt"
	"net/http"
)

/**
 * AdminDashboardHandler
 * @type {String}
 */
func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
