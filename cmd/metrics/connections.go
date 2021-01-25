package metrics

import (
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"sync/atomic"
)

type ConnectionWatcher struct {
	n int64
}

// OnStateChange records open connections in response to connection
// state changes. Set net/http Server.ConnState to this method
// as value.
func (cw *ConnectionWatcher) OnStateChangeForFastHttp(conn net.Conn, state fasthttp.ConnState) {
	switch state {
	case fasthttp.StateNew:
		atomic.AddInt64(&cw.n, 1)
	case fasthttp.StateHijacked, fasthttp.StateClosed:
		atomic.AddInt64(&cw.n, -1)
	}
	ActiveHttpRequests.WithLabelValues().Set(float64(cw.n))
}

func (cw *ConnectionWatcher) OnStateChangeForNetHttp(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		atomic.AddInt64(&cw.n, 1)
	case http.StateHijacked, http.StateClosed:
		atomic.AddInt64(&cw.n, -1)
	}
	ActiveHttpRequests.WithLabelValues().Set(float64(cw.n))
}

// Count returns the number of connections at the time
// the call.
func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))
}
