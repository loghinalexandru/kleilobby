package main

import (
	"log"
	"net/http"
	"os"

	"github.com/loghinalexandru/kleilobby/dst"
	"github.com/loghinalexandru/kleilobby/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	dstHandler := dst.NewHandler(logger)

	mux.HandleFunc("GET /api/v1/dst", dstHandler.All)
	mux.HandleFunc("GET /api/v1/dst/{rowID}", dstHandler.RowID)
	mux.HandleFunc("GET /api/v1/dst/{hostKU}/{serverName}", dstHandler.ServerName)
	mux.HandleFunc("HEAD /api/v1/dst/{hostKU}/{serverName}", dstHandler.ServerName)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger.Println("Server starting...")
	server.New(mux).ListenAndServe()
}
