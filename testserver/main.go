// A minimal plaintext HTTP server for experimenting with Shakti's HTTP parser.
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:              "127.0.0.1:6011",
		ReadHeaderTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header()["Date"] = nil // Keep the response headers minimal.
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Content-Length", "22")
			w.Header().Set("Connection", "close")
			io.WriteString(w, "<h1>Hello, World!</h1>")
		}),
	}

	log.Print("Listening on http://", server.Addr)
	log.Fatal(server.ListenAndServe())
}
