package endpoint

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/xLegoz/go-links/backend"
	"io/ioutil"
	"net/http"
)

type goLinkPostRequestStruct struct {
	URL string `json:"url"`
}

type goLinkPostResponseStruct struct {
	Success bool `json:"success"`
}

type goLinkGetStruct struct {
	URL     string `json:"url"`
	Metrics uint   `json:"metrics"`
}

func writeJSON(w http.ResponseWriter, body interface{}) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

/**
 * AdminApiLinksHandler
 */
func AdminApiLinksHandler(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)["link"]
	switch r.Method {
	case "GET":
		url, err := backend.ActiveBackend.Get(link)
		if err != nil {
			http.Error(w, "Short link not found.", http.StatusNotFound)
			return
		}
		writeJSON(w, goLinkGetStruct{
			URL:     url,
			Metrics: backend.ActiveBackend.MetricGet(link),
		})
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var t goLinkPostRequestStruct
		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.WithFields(log.Fields{
			"shortlink": link,
			"url":       t.URL,
		}).Info("Set shortlink")
		err = backend.ActiveBackend.Store(link, t.URL)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, goLinkPostResponseStruct{
			Success: true,
		})
	default:
		// Give an error message.
	}
}
