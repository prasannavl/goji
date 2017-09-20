package mroute

import (
	"context"

	"github.com/prasannavl/mchain"
	"github.com/prasannavl/mroute/internal"
)

type match struct {
	context.Context
	p Pattern
	h mchain.Handler
}

func (m match) Value(key interface{}) interface{} {
	switch key {
	case internal.Pattern:
		return m.p
	case internal.Handler:
		return m.h
	default:
		return m.Context.Value(key)
	}
}

var _ context.Context = match{}
