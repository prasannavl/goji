package mroute

import (
	"net/http"
	"testing"

	"github.com/prasannavl/mchain"
)

func expectSequence(t *testing.T, ch chan string, seq ...string) {
	for i, str := range seq {
		if msg := <-ch; msg != str {
			t.Errorf("[%d] expected %s, got %s", i, str, msg)
		}
	}
}

func TestMiddleware(t *testing.T) {
	t.Parallel()

	m := NewMux()
	ch := make(chan string, 10)
	m.Use(func(h mchain.Handler) mchain.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) error {
			ch <- "before one"
			err := h.ServeHTTP(w, r)
			ch <- "after one"
			return err
		}
		return mchain.HandlerFunc(fn)
	})
	m.Use(func(h mchain.Handler) mchain.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) error {
			ch <- "before two"
			err := h.ServeHTTP(w, r)
			ch <- "after two"
			return err
		}
		return mchain.HandlerFunc(fn)
	})
	m.Handle(boolPattern(true), mchain.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ch <- "handler"
		return nil
	}))

	m.ServeHTTP(wr())

	expectSequence(t, ch, "before one", "before two", "handler", "after two", "after one")
}

func makeMiddleware(ch chan string, name string) mchain.Middleware {
	return func(h mchain.Handler) mchain.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) error {
			ch <- "before " + name
			err := h.ServeHTTP(w, r)
			ch <- "after " + name
			return err
		}
		return mchain.HandlerFunc(fn)
	}
}

func TestMiddlewareReconfigure(t *testing.T) {
	t.Parallel()

	m := NewMux()
	ch := make(chan string, 10)
	m.Use(makeMiddleware(ch, "one"))
	m.Use(makeMiddleware(ch, "two"))
	m.Handle(boolPattern(true), mchain.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ch <- "handler"
		return nil
	}))

	w, r := wr()
	m.ServeHTTP(w, r)

	expectSequence(t, ch, "before one", "before two", "handler", "after two", "after one")

	m.Use(makeMiddleware(ch, "three"))

	w, r = wr()
	m.ServeHTTP(w, r)

	expectSequence(t, ch, "before one", "before two", "before three",
		"handler", "after three", "after two", "after one")
}
