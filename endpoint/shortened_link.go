package endpoint

import (
	log "github.com/sirupsen/logrus"
	"github.com/xLegoz/go-links/backend"
	"github.com/xLegoz/go-links/util"
	"net/http"
)

/**
 * GolinkHandler
 */
func GolinkHandler(w http.ResponseWriter, r *http.Request) {
	path := util.PathFromRequest(r)
	if destinationURL, err := backend.ActiveBackend.Get(path); err == nil {
		backend.ActiveBackend.MetricIncrement(path)
		log.WithFields(log.Fields{
			"path":  path,
			"count": backend.ActiveBackend.MetricGet(path),
		}).Debug("GET shortlinks")
		http.Redirect(w, r, destinationURL, http.StatusSeeOther)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
