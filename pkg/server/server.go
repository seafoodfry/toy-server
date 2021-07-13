package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// ServerAddress is the address our server will listen on.
	ServerAddress = "0.0.0.0:8080"
	// ServerShutdownTimeout is how long we are willing to wait for a graceful
	// termination.
	ServerShutdownTimeout = 15 * time.Second
)

// Init initializes our HTTP server.
func Init() (*http.Server, <-chan struct{}) {
	mux := http.NewServeMux()
	//mux.HandleFunc("/api/echo", echoHandler)
	mux.Handle("/metrics", promhttp.Handler())

	// We want explicit timeuts to prevent connections from sticking around
	// indefinetely. See
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         ServerAddress,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// We'll create a channel so we can communicate with the rest of the app when
	// a termination signal has been received.
	stop := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			syscall.SIGINT,  // Ctrl+C.
			syscall.SIGTERM, // Termination Request.
			syscall.SIGSEGV, // FullDerp.
		)
		sig := <-c
		log.Warningf("Signal (%v) detected, shutting down", sig)

		// We'll try and shut down the server and will let the process run only
		// until the 'ServerShutdownTimeout'.
		ctx, cancel := context.WithTimeout(context.Background(), ServerShutdownTimeout)
		defer cancel()

		// Since we are shutting down we will disable HTTP keep-alives.
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v", err)
		}
		close(stop)
	}()

	return server, stop
}
