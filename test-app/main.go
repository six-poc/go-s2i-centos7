package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received request from %s for %s", r.RemoteAddr, r.URL)

	h := w.Header()
	h.Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Hello world!\n\n")
	fmt.Fprintf(w, "Go version: %s\n", runtime.Version())
}

func main() {
	flag.Parse()

	log.Info("Starting server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
