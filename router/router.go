package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/loghinalexandru/klei-lobby/dst"
	"github.com/loghinalexandru/klei-lobby/dst/model"
)

const (
	allRoute        = "^/api/v1/dst$"
	rowIDRoute      = "^/api/v1/dst/(?P<%v>[a-zA-Z0-9]+)$"
	serverNameRoute = `^/api/v1/dst/(?P<%v>KU_[\-\_\+a-zA-Z0-9]+)/(?P<%v>[a-zA-Z\s0-9]+)$`
)

type router struct {
	log    *log.Logger
	routes map[*regexp.Regexp]http.HandlerFunc
}

func New(logger *log.Logger, dst *dst.Handler) *router {
	routes := make(map[*regexp.Regexp]http.HandlerFunc)

	routes[regexp.MustCompile(allRoute)] = dst.All
	routes[regexp.MustCompile(fmt.Sprintf(rowIDRoute, model.RowID))] = dst.RowID
	routes[regexp.MustCompile(fmt.Sprintf(serverNameRoute, model.HostKU, model.ServerName))] = dst.ServerName

	return &router{
		log:    logger,
		routes: routes,
	}
}

func (r *router) SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", r.route)
}

func (r *router) route(writer http.ResponseWriter, request *http.Request) {
	for route, handler := range r.routes {
		if route.MatchString(request.URL.Path) {
			handler(writer, request.WithContext(r.buildContext(request.URL.Path, route)))
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func (r *router) buildContext(path string, route *regexp.Regexp) context.Context {
	ctx := context.Background()
	groupMatches := route.FindStringSubmatch(path)[1:]
	groupKeys := route.SubexpNames()[1:]

	for i, match := range groupMatches {
		ctx = context.WithValue(ctx, model.PathKey(groupKeys[i]), match)
	}

	return ctx
}
