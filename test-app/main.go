package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

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
