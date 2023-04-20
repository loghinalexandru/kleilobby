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

	keyFirst := regexp.MustCompile("firstTestKey")
	keySecond := regexp.MustCompile("secondTestKey")
	handler := func(w http.ResponseWriter, r *http.Request) {}

	router := New(log.Default(),
		WithRoute(keyFirst, handler),
		WithRoute(keySecond, handler),
	)

	if router == nil || len(router.routes) != 2 {
		t.Fatal("unexpected router returned")
	}

	if router.routes[keyFirst] == nil {
		t.Errorf("missing first route entry")
	}

	if router.routes[keySecond] == nil {
		t.Errorf("missing second route entry")
	}
}

func TestSetup(t *testing.T) {
	t.Parallel()

	defer func() { _ = recover() }()

	router := New(log.Default(), nil)
	mux := http.DefaultServeMux

	router.Setup("/", mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	t.Errorf("router path handler not registered")
}

func TestRouteWithNoRoutes(t *testing.T) {
	t.Parallel()

	router := New(log.Default(), nil)
	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002", nil)

	router.route(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusNotFound {
		t.Error("unexpected status code")
	}
}

func TestRouteWithMatch(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute(regexp.MustCompile("/api/*"), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002/api/test", nil)

	router.route(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusOK {
		t.Error("unexpected status code")
	}
}

func TestRouteWithNoMatch(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute(regexp.MustCompile("/api/test"), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002", nil)

	router.route(spyResponse, spyRequest)

	if spyResponse.Result().StatusCode != http.StatusNotFound {
		t.Error("unexpected status code")
	}
}

func TestRouteWithMatchAndContext(t *testing.T) {
	t.Parallel()

	router := New(log.Default(),
		WithRoute(regexp.MustCompile("/api/(?P<number>[0-9]+)/*"), spyContextKeyHandler("number", "123", t)),
	)

	spyResponse := httptest.NewRecorder()
	spyRequest := httptest.NewRequest("GET", "http://localhost:3002/api/123/test", nil)

	router.route(spyResponse, spyRequest)

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
