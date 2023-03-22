package main

import (
	"log"
	"net/http"
	"os"

	"github.com/loghinalexandru/klei-lobby/dst"
	"github.com/loghinalexandru/klei-lobby/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	dstHandler := dst.NewHandler(logger)
	dstHandler.SetupRoutes(mux)

	logger.Println("Server starting...")
	server.New(mux).ListenAndServe()
}
