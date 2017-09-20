package mroute

import (
	"net/http"

	"github.com/prasannavl/goerror/httperror"

	"github.com/prasannavl/mchain"
	"github.com/prasannavl/mroute/internal"
)

type dispatch struct{}

func (d dispatch) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	h := ctx.Value(internal.Handler)
	if h == nil {
		return httperror.New(http.StatusNotFound, "route not found", false)
	}
	return h.(mchain.Handler).ServeHTTP(w, r)
}
