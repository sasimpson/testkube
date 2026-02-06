package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var staticDir = os.Getenv("TESTKUBE_STATICDIR")
var sslDir = os.Getenv("TESTKUBE_SSLDIR")
var sslEnabled = os.Getenv("TESTKUBE_SSLENABLED")

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Header().Set("X-Goos", runtime.GOOS)
		w.Header().Set("X-Goarch", runtime.GOARCH)
		w.Header().Set("X-Pod-Name", hostname)
		h.ServeHTTP(w, r)
	})
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	r := http.NewServeMux()
	r.Handle("/", middleware(homeHandler()))

	if staticDir != "" {
		r.Handle("/static", middleware(http.StripPrefix(strings.TrimRight("/static/", "/"), http.FileServer(http.Dir(staticDir)))))
	}

	if sslEnabled == "true" {
		sslServer := &http.Server{Addr: ":8443", Handler: r}
		tlsCert := fmt.Sprintf("%s/tls.crt", sslDir)
		tlsKey := fmt.Sprintf("%s/tls.key", sslDir)

		slog.Info("Starting HTTPS server", "port", 8443, "tls_cert", tlsCert, "tls_key", tlsKey)
		sslServer.ListenAndServeTLS(tlsCert, tlsKey)
	}

	server := &http.Server{
		Addr: ":8080", Handler: r,
	}
	slog.Info("Starting HTTP server", "port", 8080)
	server.ListenAndServe()

}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hi world! I'm %s!\n", hostname)
	})
}
