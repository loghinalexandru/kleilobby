package router

import (
	"log"
	"net/http"
	"regexp"

	"github.com/loghinalexandru/klei-lobby/dst"
)

type Router struct {
	log *log.Logger
	dst *dst.Handler
}

func New(logger *log.Logger, dst *dst.Handler) *Router {
	return &Router{
		log: logger,
		dst: dst,
	}
}

func (r *Router) SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", r.switchRouter)
}

// TODO: add tracing & route match logging
func (r *Router) switchRouter(writer http.ResponseWriter, request *http.Request) {
	all := regexp.MustCompile("^/api/v1/dst$")
	rowID := regexp.MustCompile("^/api/v1/dst/([a-zA-Z0-9]+)$")
	serverName := regexp.MustCompile(`^/api/v1/dst/(KU_[\-\_\+a-zA-Z0-9]+)/([a-zA-Z\s0-9]+)$`)

	switch {
	case all.MatchString(request.URL.Path):
		r.dst.All(writer, request)
	case rowID.MatchString(request.URL.Path):
		pathRowID := rowID.FindStringSubmatch(request.URL.Path)[1]
		r.dst.RowID(writer, request, pathRowID)
	case serverName.MatchString(request.URL.Path):
		pathHostKU := serverName.FindStringSubmatch(request.URL.Path)[1]
		pathServerName := serverName.FindStringSubmatch(request.URL.Path)[2]
		r.dst.ServerName(writer, request, pathServerName, pathHostKU)
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}
