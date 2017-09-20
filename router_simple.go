// +build mroute_router_simple

package mroute

import (
	"net/http"

	"github.com/prasannavl/mchain"
)

/*
This is the simplest correct router implementation for mroute.
*/

type router []route

type route struct {
	Pattern
	mchain.Handler
}

func (rt *router) add(p Pattern, h mchain.Handler) {
	*rt = append(*rt, route{p, h})
}

func (rt *router) route(r *http.Request) *http.Request {
	for _, route := range *rt {
		if r2 := route.Match(r); r2 != nil {
			return r2.WithContext(&match{
				Context: r2.Context(),
				p:       route.Pattern,
				h:       route.Handler,
			})
		}
	}
	return r.WithContext(&match{Context: r.Context()})
}
