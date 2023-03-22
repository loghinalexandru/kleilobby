package server

import (
	"net/http"
	"time"
)

func New(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         ":3002",
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
		Handler:      mux,
	}
}
