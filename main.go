package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var links = make(map[string]string)
var metrics = make(map[string]uint)

var linkLock sync.RWMutex
var metrixLock sync.Mutex

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func metricCount(url string) {
	metrics[url] += 1
	// record with GAnalytics if possible, or will the target website record?
}

func GetGolink(r *http.Request) (string, bool) {
	if dest, ok := links[r.URL.Path[1:]]; ok {
		return dest, true
	} else {
		return "", false
	}
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type GoLinkStruct struct {
	URL string `json:"url"`
}

func AdminApiLinksHandler(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)["link"]
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET there, I love %s!", link)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var t GoLinkStruct
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		log.WithFields(log.Fields{
			"shortlink": link,
			"url":       t.URL,
		}).Info("Set shortlink")
		links[link] = t.URL
	default:
		// Give an error message.
	}
}

func GolinkHandler(w http.ResponseWriter, r *http.Request) {
	if dest, ok := GetGolink(r); ok {
		// metricCount(newUrl)
		http.Redirect(w, r, dest, http.StatusSeeOther)
	} else {
		http.Error(w, "Unsupported path", http.StatusNotFound)
	}
}

func main() {
	port, err := strconv.Atoi(getenv("PORT", "8080"))
	if err != nil {
		log.WithFields(log.Fields{
			"port": getenv("PORT", "8080"),
		}).Fatal("Invalid PORT environment variable: expected integer.")
		os.Exit(1)
	}

	env := getenv("GOLINK_ENV", "dev")

	log.WithFields(log.Fields{
		"port": port,
		"env":  env,
	}).Info("Starting go-link HTTP server")

	var addr string

	if env == "dev" {
		addr = fmt.Sprintf("127.0.0.1:%d", port)
	} else if env == "prod" {
		addr = fmt.Sprintf("0.0.0.0:%d", port)
	}

	router := mux.NewRouter()
	router.HandleFunc("/admin/dashboard", AdminDashboardHandler).Methods("GET")
	router.HandleFunc("/admin/api/links/{link}", AdminApiLinksHandler).Methods("GET", "POST")
	router.NotFoundHandler = http.HandlerFunc(GolinkHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
