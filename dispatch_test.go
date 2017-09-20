package mroute

import (
	"context"
	"net/http"
	"testing"

	"github.com/prasannavl/goerror/httperror"

	"github.com/prasannavl/mchain"
	"github.com/prasannavl/mroute/internal"
)

func TestDispatch(t *testing.T) {
	t.Parallel()

	var d dispatch

	w, r := wr()
	err := d.ServeHTTP(w, r)
	e := err.(httperror.HttpError)
	if e.Code() != 404 {
		t.Errorf("status: expected %d, got %d", 404, w.Code)
	}

	w, r = wr()
	h := mchain.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(123)
		return nil
	})
	ctx := context.WithValue(context.Background(), internal.Handler, h)
	r = r.WithContext(ctx)
	d.ServeHTTP(w, r)
	if w.Code != 123 {
		t.Errorf("status: expected %d, got %d", 123, w.Code)
	}
}
