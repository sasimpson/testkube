package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hi world!\n")
	fmt.Fprintf(w, "i am %s\n", hostname)
	fmt.Fprintf(w, "OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	log.Printf("%s: %s, %s", hostname, runtime.GOOS, runtime.GOARCH)
}
