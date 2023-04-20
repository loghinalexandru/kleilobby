package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

type ContextKey string

type routerOpt func(*router)

type router struct {
	log    *log.Logger
	routes map[*regexp.Regexp]http.HandlerFunc
}

func New(logger *log.Logger, opts ...routerOpt) *router {
	r := &router{
		log:    logger,
		routes: make(map[*regexp.Regexp]http.HandlerFunc),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithRoute(route *regexp.Regexp, handler http.HandlerFunc) routerOpt {
	return func(r *router) {
		r.routes[route] = handler
	}
}

func (r *router) Setup(basePath string, mux *http.ServeMux) {
	mux.HandleFunc(basePath, r.route)
}

func (r *router) route(writer http.ResponseWriter, request *http.Request) {
	for route, handler := range r.routes {
		if route.MatchString(request.URL.Path) {
			r.log.Printf("path match: %v", route.String())
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
		ctx = context.WithValue(ctx, ContextKey(groupKeys[i]), match)
	}

	return ctx
}
