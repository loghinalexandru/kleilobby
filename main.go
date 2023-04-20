package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/loghinalexandru/klei-lobby/dst"
	"github.com/loghinalexandru/klei-lobby/router"
	"github.com/loghinalexandru/klei-lobby/server"
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
		router.WithRoute(regexp.MustCompile(allRoute), dstHandler.All),
		router.WithRoute(regexp.MustCompile(fmt.Sprintf(rowIDRoute, dst.RowID)), dstHandler.RowID),
		router.WithRoute(regexp.MustCompile(fmt.Sprintf(serverNameRoute, dst.HostKU, dst.ServerName)), dstHandler.ServerName))

	router.SetupRouter("/", mux)

	logger.Println("Server starting...")
	server.New(mux).ListenAndServe()
}
