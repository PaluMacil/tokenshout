package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	// set port from default or command flag
	var port string
	flag.StringVar(&port, "p", "12345", "port to serve tokenshout")
	flag.Parse()
	// create multiplexer from standard library and handle endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>PONG!!!</h1>"))
	})
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
		// time from when the connection is accepted to when the request body is fully read
		ReadTimeout: 4 * time.Second,
		// time from the end of the request header read to the end of the response write
		WriteTimeout: 6 * time.Second,
	}
	// serve
	log.Println("Now serving tokenshout on port", port)
	log.Fatal(srv.ListenAndServe())
}
