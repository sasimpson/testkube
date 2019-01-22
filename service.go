package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Header().Set("X-Goos", runtime.GOOS)
		w.Header().Set("X-Goarch", runtime.GOARCH)
		w.Header().Set("X-Pod-Name", hostname)
		log.Println(runtime.GOOS, runtime.GOARCH, hostname)
		h.ServeHTTP(w,r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", Middleware(HomeHandler()))
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi world!\n")
	})
}


