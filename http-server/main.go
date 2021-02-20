package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/test200", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	m.HandleFunc("/test500", func(w http.ResponseWriter, r *http.Request) {
		// simulate 500 eror code
		w.WriteHeader(500)
	})
	const addr = ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: m,
	}
	log.Infof("listening on '%s'", addr)
	log.Fatal(server.ListenAndServe())
}
