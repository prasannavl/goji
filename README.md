mroute
====

[![GoDoc](https://godoc.org/github.com/prasannavl/mroute?status.svg)](https://godoc.org/github.com/prasannavl/mroute) [![Build Status](https://travis-ci.org/prasannavl/mroute.svg?branch=master)](https://travis-ci.org/prasannavl/mroute)

mroute is a HTTP request multiplexer, similar to [`net/http.ServeMux`][servemux].
It compares incoming requests to a list of registered [Patterns][pattern], and
dispatches to the [mchain.Handler][handler] that corresponds to the first matching
Pattern. mroute also supports [Middleware][middleware] (composable shared
functionality applied to every request) and uses the standard
[`context`][context] package to store request-scoped values.

**This is a fork of [goji](https://goji.io) adapted to the mchain philosophy of returning errors.**

[servemux]: https://golang.org/pkg/net/http/#ServeMux
[pattern]: https://godoc.org/github.com/prasannavl/mroute#Pattern
[handler]: https://godoc.org/github.com/prasannavl/mchain
[middleware]: https://godoc.org/github.com/prasannavl/mroute#Mux.Use
[context]: https://golang.org/pkg/context


## Quick Start


```go
package main

import (
        "fmt"
        "net/http"

        "github.com/prasannavl/mroute"
        "github.com/prasannavl/mroute/pat"
)

func hello(w http.ResponseWriter, r *http.Request) error {
        name := pat.Param(r, "name")
        fmt.Fprintf(w, "Hello, %s!", name)
        return nil
}

func main() {
        mux := mroute.NewMux()
        mux.HandleFunc(pat.Get("/hello/:name"), hello)

        http.ListenAndServe("localhost:8000", mux)
}
```

Please refer to [mroute's GoDoc Documentation][godoc] for a full API reference.

[godoc]: https://godoc.org/github.com/prasannavl/mroute


## Related links

`mchain`: https://github.com/prasannavl/mchain  
`goji/goji`: https://goji.io 

## Credits

Thanks to `goji` router - mroute is an adaptation of this fantastic work.