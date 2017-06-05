package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/xLegoz/go-links/backend"
	"github.com/xLegoz/go-links/endpoint"
	"github.com/xLegoz/go-links/util"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var AdminKey string
var Environment string
var Addr string
var LogLevel string
var Store string

func init() {
	// set logger level
	switch LogLevel = util.Getenv("GOLINK_LOGLEVEL", "info"); LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	default:
		log.WithFields(log.Fields{
			"GOLINKS_STORE": Store,
		}).Fatal("Invalid GOLINKS_LOGLEVEL environment variable: expected one of 'debug', 'info', 'warning'")
		os.Exit(1)
	}

	// get port
	port, err := strconv.Atoi(util.Getenv("PORT", "8080"))
	if err != nil {
		log.WithFields(log.Fields{
			"port": util.Getenv("PORT", "8080"),
		}).Fatal("Invalid PORT environment variable: expected integer.")
		os.Exit(1)
	}

	// get secret admin key
	AdminKey = util.Getenv("ADMIN_KEY", "")
	if len(AdminKey) < 16 {
		log.WithFields(log.Fields{
			"AdminKey": AdminKey,
		}).Fatal("Invalid ADMIN_KEY environment variable: needs to be at least 16 characters.")
		os.Exit(1)
	}

	// get the backend store settings
	switch Store = util.Getenv("GOLINKS_STORE", "dict"); Store {
	case "dict":
		backend.ActiveBackend = &backend.Dict{}
	case "redis":
		backend.ActiveBackend = setupRedis()
	default:
		log.WithFields(log.Fields{
			"GOLINKS_STORE": Store,
		}).Fatal("Invalid GOLINKS_STORE environment variable: either 'dict' or 'redis'")
		os.Exit(1)
	}

	// get environment
	switch Environment = util.Getenv("GOLINKS_ENV", "dev"); Environment {
	case "dev":
		Addr = fmt.Sprintf("127.0.0.1:%d", port)
	case "prod":
		Addr = fmt.Sprintf("0.0.0.0:%d", port)
	default:
		log.WithFields(log.Fields{
			"GOLINKS_ENV": Environment,
		}).Fatal("Invalid GOLINKS_ENV environment variable: either 'dev' or 'prod'")
		os.Exit(1)
	}
}

func setupRedis() *backend.Redis {
	// default to standard redis port on localhost
	redisURL := util.Getenv("REDIS_URL", "redis://h:@localhost:6379")
	parts := strings.Split(redisURL, "@")
	redisAddress := parts[1]
	parts = strings.Split(parts[0], ":")
	redisPassword := parts[2]

	return &backend.Redis{
		Options: &redis.Options{
			Addr:     redisAddress,
			Password: redisPassword,
			DB:       0, // use default DB
		},
	}
}

func main() {
	log.WithFields(log.Fields{
		"Addr":        Addr,
		"Environment": Environment,
		"LogLevel":    LogLevel,
		"Store":       Store,
	}).Info("Starting go-link HTTP server")

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
