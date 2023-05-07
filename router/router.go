package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type ContextKey string

type routerOpt func(*router)

type route struct {
	pattern *regexp.Regexp
	method  string
}

type router struct {
	log    *log.Logger
	routes map[route]http.HandlerFunc
}

func New(logger *log.Logger, opts ...routerOpt) *router {
	r := &router{
		log:    logger,
		routes: make(map[route]http.HandlerFunc),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithRoute(method string, regex *regexp.Regexp, handler http.HandlerFunc) routerOpt {
	return func(r *router) {
		r.routes[route{pattern: regex, method: method}] = handler
	}
}

func (r *router) Setup(basePath string, mux *http.ServeMux) {
	mux.HandleFunc(basePath, r.handleRequest)
}

func (r *router) handleRequest(writer http.ResponseWriter, request *http.Request) {
	for route, handler := range r.routes {
		if route.pattern.MatchString(request.URL.Path) && strings.EqualFold(request.Method, route.method) {
			r.log.Printf("path match: %v %v", strings.ToUpper(route.method), route.pattern.String())
			handler(writer, request.WithContext(r.buildContext(request.URL.Path, route.pattern)))
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
		ctx = context.WithValue(ctx, ContextKey(groupKeys[i]), match)
	}

	return ctx
}
