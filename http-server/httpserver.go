package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// twoHundred
func twoHundred(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "HTTP Endpoint OK!"}`))
}

// fiveHundred
func fiveHundred(w http.ResponseWriter, r *http.Request) {
	// simulate 500 eror code
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "HTTP Endpoint Internal Error!"}`))
}
func router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/test200", twoHundred)
	m.HandleFunc("/test500", fiveHundred)
	return m
}
func main() {
	const addr = ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: router(),
	}
	log.Infof("listening on '%s'", addr)
	log.Fatal(server.ListenAndServe())
}
