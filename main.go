package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
)

var staticDir = os.Getenv("TESTKUBE_STATICDIR")

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Header().Set("X-Goos", runtime.GOOS)
		w.Header().Set("X-Goarch", runtime.GOARCH)
		w.Header().Set("X-Pod-Name", hostname)
		log.Println(runtime.GOOS, runtime.GOARCH, hostname)
		h.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", middleware(homeHandler()))
	r.Handle("/static", middleware(http.StripPrefix(strings.TrimRight("/static/", "/"), http.FileServer(http.Dir(staticDir)))))
	http.ListenAndServe(":80", r)
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi world!\n")
	})
}
