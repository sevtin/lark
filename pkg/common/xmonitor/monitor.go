package xmonitor

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

/*
http://127.0.0.1:7302/debug/pprof/
http://127.0.0.1:7302/metrics
*/
func RunMonitor(port int) {
	go func(p int) {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":"+strconv.Itoa(p), nil)
	}(port)
}
