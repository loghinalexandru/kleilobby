package dst

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/loghinalexandru/klei-lobby/caching"
	"github.com/loghinalexandru/klei-lobby/dst/models"
)

func TestGetAllWithNoError(t *testing.T) {
	t.Parallel()

	want := []models.ViewModel{
		{
			HostKU:     "test_host",
			ServerName: "test_server",
		},
	}

	target := &service{
		logger: log.Default(),
		client: newTestClient(func(req *http.Request) *http.Response {

			if strings.Contains(req.URL.Path, "eu-central-1") {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(getMockData(t))),
					Header:     make(http.Header),
				}
			}

			return nil
		}),
	}

	got, err := target.GetAll("eu-central-1")

	if err != nil {
		t.Fatalf("unexpected error encountered: %v", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetAllWithHttpError(t *testing.T) {
	t.Parallel()

	target := &service{
		logger: log.Default(),
		client: newTestClient(func(req *http.Request) *http.Response {
			return nil
		}),
	}

	_, err := target.GetAll("bogus region")

	if err == nil {
		t.Error("expected error missing")
	}
}

func TestGetByRowIDWithNoError(t *testing.T) {
	t.Parallel()

	want := models.ViewModel{
		HostKU:     "test_host",
		ServerName: "test_server",
	}

	target := &service{
		logger: log.Default(),
		client: newTestClient(func(req *http.Request) *http.Response {
			payload, _ := io.ReadAll(req.Body)

			if strings.Contains(req.URL.Host, "eu-central-1") &&
				strings.Contains(string(payload), "custom-row-id") &&
				strings.Contains(string(payload), "custom-token") {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(getMockData(t))),
					Header:     make(http.Header),
				}
			}

			return nil
		}),
	}

	got, err := target.GetByRowID("custom-token", "eu-central-1", "custom-row-id")

	if err != nil {
		t.Fatalf("unexpected error encountered: %v", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetByRowIDWithHttpError(t *testing.T) {
	t.Parallel()

	target := &service{
		logger: log.Default(),
		client: newTestClient(func(req *http.Request) *http.Response {
			return nil
		}),
	}

	_, err := target.GetByRowID("bogus token", "bogus region", "bogus id")

	if err == nil {
		t.Error("expected error missing")
	}
}

func TestGetByServerNameAndHostWithNoError(t *testing.T) {
	t.Parallel()

	want := models.ViewModel{
		HostKU:     "test_host",
		ServerName: "test_server",
	}

	target := &service{
		logger: log.Default(),
		cache:  caching.New[models.ViewModel](time.Hour),
		client: newTestClient(func(req *http.Request) *http.Response {

			if strings.Contains(req.URL.Path, "eu-central-1") || strings.Contains(req.URL.Host, "eu-central-1") {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(getMockData(t))),
					Header:     make(http.Header),
				}
			}

			return nil
		}),
	}

	got, err := target.GetByServerNameAndHost("custom-token", "eu-central-1", "test_server", "test_host")

	if err != nil {
		t.Fatalf("unexpected error encountered: %v", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetByServerNameAndHostWithHttpError(t *testing.T) {
	t.Parallel()

	target := &service{
		logger: log.Default(),
		cache:  caching.New[models.ViewModel](time.Hour),
		client: newTestClient(func(req *http.Request) *http.Response {
			return nil
		}),
	}

	_, err := target.GetByServerNameAndHost("custom-token", "eu-central-1", "test_server", "test_host")

	if err == nil {
		t.Error("expected error missing")
	}
}

func getMockData(t *testing.T) string {
	t.Helper()

	return `{"GET" : [
		{
			"host" : "test_host",
			"name" : "test_server"
		}]
	}`
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
