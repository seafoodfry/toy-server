package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// requestsTotal keeps track of all of the requests made to our "echo"
	// endpoint.
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "echo",
			Name:      "requests_total",
			Help:      "Total http requests to our echo JSON api",
		},
		[]string{
			// HTTP status code.
			"code",
		},
	)
)

// RegisterMetrics registers all metrics for the server.
func RegisterMetrics() {
	prometheus.MustRegister(requestsTotal)
}
