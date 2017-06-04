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
	if dest, ok := backend.ActiveBackend.Get(path); ok {
		backend.ActiveBackend.MetricIncrement(path)
		log.WithFields(log.Fields{
			"path":  path,
			"count": backend.ActiveBackend.MetricGet(path),
		}).Debug("GET shortlinks")
		http.Redirect(w, r, dest, http.StatusSeeOther)
	} else {
		http.Error(w, "Unsupported path", http.StatusNotFound)
	}
}
