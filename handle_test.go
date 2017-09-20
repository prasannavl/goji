package mroute

import (
	"net/http"
	"testing"

	"github.com/prasannavl/mchain"
)

func TestHandle(t *testing.T) {
	t.Parallel()

	m := NewMux()
	called := false
	fn := func(w http.ResponseWriter, r *http.Request) error {
		called = true
		return nil
	}
	m.Handle(boolPattern(true), mchain.HandlerFunc(fn))

	w, r := wr()
	m.ServeHTTP(w, r)
	if !called {
		t.Error("expected handler to be called")
	}
}

func TestHandleFunc(t *testing.T) {
	t.Parallel()

	m := NewMux()
	called := false
	fn := func(w http.ResponseWriter, r *http.Request) error {
		called = true
		return nil
	}
	m.HandleFunc(boolPattern(true), fn)

	w, r := wr()
	m.ServeHTTP(w, r)
	if !called {
		t.Error("expected handler to be called")
	}
}
