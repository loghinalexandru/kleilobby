package router

import (
	"log"
	"net/http"
	"regexp"

	"github.com/loghinalexandru/klei-lobby/handlers"
)

type Router struct {
	log *log.Logger
	dst *handlers.DontStarveTogether
}

func New(logger *log.Logger, dst *handlers.DontStarveTogether) *Router {
	return &Router{
		log: logger,
		dst: dst,
	}
}

func (r *Router) SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", r.switchRouter)
}

// TODO: parse needed values from URI
func (r *Router) switchRouter(writer http.ResponseWriter, request *http.Request) {
	all := regexp.MustCompile("^/all$")
	rowId := regexp.MustCompile("^/[a-zA-Z0-9]+$")

	switch {
	case all.MatchString(request.URL.Path):
		r.dst.All(writer, request)
	case rowId.MatchString(request.URL.Path):
		r.dst.RowId(writer, request)
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}
