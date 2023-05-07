package router

import (
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()

	patternFirst := regexp.MustCompile("/api/test")
	patternSecond := regexp.MustCompile("api/test2")
	handler := func(w http.ResponseWriter, r *http.Request) {}

	router := New(log.Default(),
		WithRoute("GET", patternFirst, handler),
		WithRoute("POST", patternSecond, handler),
	)

	if router == nil || len(router.routes) != 2 {
		t.Fatal("unexpected router returned")
	}

	containsRoute(router.routes, route{method: "GET", pattern: patternFirst}, t)
	containsRoute(router.routes, route{method: "POST", pattern: patternSecond}, t)
}

func TestSetup(t *testing.T) {
	t.Parallel()

	defer func() { _ = recover() }()

	router := New(log.Default())
	mux := http.DefaultServeMux

	router.Setup("/", mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	t.Errorf("router path handler not registered")
}

func TestRouteWithNoRoutes(t *testing.T) {
	t.Parallel()

	router := New(log.Default())
	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002", nil)

	router.handleRequest(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusNotFound {
		t.Error("unexpected status code")
	}
}

func TestRouteWithMatch(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute("GET", regexp.MustCompile("/api/*"), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002/api/test", nil)

	router.handleRequest(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusOK {
		t.Error("unexpected status code")
	}
}

func TestRouteWithNoMatch(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute("GET", regexp.MustCompile("/api/test"), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002", nil)

	router.handleRequest(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusNotFound {
		t.Error("unexpected status code")
	}
}

func TestRouteWithMatchAndContext(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute("GET", regexp.MustCompile("/api/(?P<number>[0-9]+)/*"), spyContextKeyHandler("number", "123", t)),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002/api/123/test", nil)

	router.handleRequest(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusOK {
		t.Error("unexpected status code")
	}
}

func spyContextKeyHandler(key string, want string, t *testing.T) http.HandlerFunc {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if got := r.Context().Value(ContextKey(key)); got != want {
			t.Errorf("context value mismatch (-got +want) %s", cmp.Diff(got, want))
		}
	}
}

func containsRoute(routes map[route]http.HandlerFunc, target route, t *testing.T) {
	t.Helper()

	var contains bool

	for r := range routes {
		if r == target {
			contains = true
		}
	}

	if contains == false {
		t.Errorf("missing expected route: %v %v", target.pattern.String(), target.method)
	}
}
