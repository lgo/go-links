package endpoint

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/xLegoz/go-links/backend"
	"io/ioutil"
	"net/http"
)

type goLinkPostStruct struct {
	URL string `json:"url"`
}

type goLinkGetStruct struct {
	URL     string `json:"url"`
	Metrics uint   `json:"metrics"`
}

/**
 * AdminApiLinksHandler
 */
func AdminApiLinksHandler(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)["link"]
	switch r.Method {
	case "GET":
		url, ok := backend.ActiveBackend.Get(link)
		if !ok {
			http.Error(w, "Short link not found.", http.StatusNotFound)
			return
		}

		body, err := json.Marshal(goLinkGetStruct{
			URL:     url,
			Metrics: backend.ActiveBackend.MetricGet(link),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var t goLinkPostStruct
		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.WithFields(log.Fields{
			"shortlink": link,
			"url":       t.URL,
		}).Info("Set shortlink")
		backend.ActiveBackend.Store(link, t.URL)
	default:
		// Give an error message.
	}
}
