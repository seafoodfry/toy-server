/*
  Here we implement some middleware for obtaining info about requests.
  However, if we implemen the app with a distributed tracing library
  (i.e., Honeycomb) we get this for free and we get a platform for studying,
  and making sense of our traffic.

  This is just for show purposes.

  If you are interested in seeing how this can be used in real lie checkout:
  https://github.com/honeycombio/beeline-go/blob/main/wrappers/hnynethttp/nethttp.go
*/
package server

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	log "github.com/sirupsen/logrus"
)

// tracerMiddleware keeps track of useful info for all the incoming requests to
// the HTTP handlers it wraps.
func tracerMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// If we call the handler this way we can obtain some basic metrics that we
		// can log.
		m := httpsnoop.CaptureMetrics(h, w, r)

		// And keep track of the request in the logs.
		log.WithFields(log.Fields{
			"method":    r.Method,
			"url":       r.URL,
			"code":      m.Code,
			"duration":  m.Duration,
			"bytes":     m.Written,
			"referer":   r.Header.Get("Referer"),
			"userAgent": r.Header.Get("User-Agent"),
		}).Debug("")
	}

	return http.HandlerFunc(fn)
}
