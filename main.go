package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/loghinalexandru/kleilobby/dst"
	"github.com/loghinalexandru/kleilobby/router"
	"github.com/loghinalexandru/kleilobby/server"
)

const (
	allRoute        = "^/api/v1/dst$"
	rowIDRoute      = "^/api/v1/dst/(?P<%v>[a-zA-Z0-9]+)$"
	serverNameRoute = `^/api/v1/dst/(?P<%v>KU_[\-\_\+a-zA-Z0-9]+)/(?P<%v>[a-zA-Z\s0-9]+)$`
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	dstHandler := dst.NewHandler(logger)

	router := router.New(logger,
		router.WithRoute("GET", regexp.MustCompile(allRoute), dstHandler.All),
		router.WithRoute("GET", regexp.MustCompile(fmt.Sprintf(rowIDRoute, dst.RowID)), dstHandler.RowID),
		router.WithRoute("GET", regexp.MustCompile(fmt.Sprintf(serverNameRoute, dst.HostKU, dst.ServerName)), dstHandler.ServerName),
		router.WithRoute("HEAD", regexp.MustCompile(fmt.Sprintf(serverNameRoute, dst.HostKU, dst.ServerName)), dstHandler.Exists))

	router.Setup("/", mux)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger.Println("Server starting...")
	server.New(mux).ListenAndServe()
}
