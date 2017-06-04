package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/xLegoz/go-links/backend"
	"github.com/xLegoz/go-links/endpoint"
	"github.com/xLegoz/go-links/util"
	"net/http"
	"os"
	"strconv"
	"time"
)

var AdminKey string
var Environment string
var Addr string
var LogLevel string

func init() {
	// get environment
	Environment = util.Getenv("GOLINK_ENV", "dev")

	// set logger level
	LogLevel = util.Getenv("GOLINK_LOGLEVEL", "info")
	if LogLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else if LogLevel == "warning" {
		log.SetLevel(log.WarnLevel)
	} else if LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	}

	// get port
	port, err := strconv.Atoi(util.Getenv("PORT", "8080"))
	if err != nil {
		log.WithFields(log.Fields{
			"port": util.Getenv("PORT", "8080"),
		}).Fatal("Invalid PORT environment variable: expected integer.")
		os.Exit(1)
	}

	AdminKey = util.Getenv("ADMIN_KEY", "")
	if len(AdminKey) < 16 {
		log.WithFields(log.Fields{
			"AdminKey": AdminKey,
		}).Fatal("Invalid ADMIN_KEY environment variable: needs to be at least 16 characters.")
		os.Exit(1)
	}

	if Environment == "dev" {
		Addr = fmt.Sprintf("127.0.0.1:%d", port)
	} else if Environment == "prod" {
		Addr = fmt.Sprintf("0.0.0.0:%d", port)
	}
}

func main() {
	log.WithFields(log.Fields{
		"Addr":        Addr,
		"Environment": Environment,
		"LogLevel":    LogLevel,
	}).Info("Starting go-link HTTP server")

	backend.ActiveBackend = &backend.Dict{}
	backend.ActiveBackend.Start()

	router := mux.NewRouter()

	// Authorized via. middleware
	var headerAuth = util.AuthOptions{AdminKey: AdminKey, AuthMethod: util.AuthWithHeader}
	var queryAuth = util.AuthOptions{AdminKey: AdminKey, AuthMethod: util.AuthWithQueryParam}
	router.HandleFunc("/admin/dashboard", util.BasicAuth(queryAuth)(endpoint.AdminDashboardHandler)).Methods("GET")
	router.HandleFunc("/admin/api/links/{link}", util.BasicAuth(headerAuth)(endpoint.AdminApiLinksHandler)).Methods("GET", "POST")

	router.NotFoundHandler = http.HandlerFunc(endpoint.GolinkHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    Addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
