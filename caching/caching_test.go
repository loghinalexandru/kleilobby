package caching_test

import (
	"testing"
	"time"

	"github.com/loghinalexandru/klei-lobby/caching"
)

func TestCacheWhenKeyIsNotPresent(t *testing.T) {
	t.Parallel()

	var want int

	target := caching.New[int](time.Hour)

	got := target.Get("invalid key")

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	if target.Contains("invalid key") {
		t.Error("key should not be present")
	}
}

func TestCacheWhenKeyIsPresentAndValid(t *testing.T) {
	t.Parallel()

	want := 12

	target := caching.New[int](time.Hour)
	target.Add("key", want)

	got := target.Get("key")

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	if !target.Contains("key") {
		t.Error("missing expected key")
	}
}

func TestCacheWhenKeyIsExpired(t *testing.T) {
	t.Parallel()

	var want int

	target := caching.New[int](time.Nanosecond)
	target.Add("key", 123)

	time.Sleep(time.Nanosecond)

	got := target.Get("key")

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	if target.Contains("key") {
		t.Error("key should not be present")
	}
}

func TestAddWhenKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	want := 123
	target := caching.New[int](time.Hour)

	target.Add("test key", want)
	got := target.Get("test key")

	if got != 123 {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAddWhenKeyExists(t *testing.T) {
	t.Parallel()

	want := 123
	newWant := 24

	target := caching.New[int](time.Hour)

	target.Add("test key", want)
	got := target.Get("test key")

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	target.Add("test key", newWant)
	got = target.Get("test key")

	if got != newWant {
		t.Errorf("want %v, got %v", newWant, got)
	}
}
