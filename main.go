package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/loghinalexandru/kleilobby/dst"
	"github.com/loghinalexandru/kleilobby/server"
)

func main() {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))

	mux := http.NewServeMux()
	dstHandler := dst.NewHandler(logger)

	mux.HandleFunc("GET /api/v1/dst", dstHandler.All)
	mux.HandleFunc("GET /api/v1/dst/{rowID}", dstHandler.RowID)
	mux.HandleFunc("GET /api/v1/dst/{hostKU}/{serverName}", dstHandler.ServerName)
	mux.HandleFunc("HEAD /api/v1/dst/{hostKU}/{serverName}", dstHandler.ServerName)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger.Info("Server starting...")
	_ = server.New(mux).ListenAndServe()
}
