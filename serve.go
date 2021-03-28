package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Fatal(http.ListenAndServe(":80", loggerAndHeader(http.FileServer(http.Dir(".")))))
}

func loggerAndHeader(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.RequestURI, ".wasm") {
			w.Header().Add("Content-Type", "application/wasm")
		}
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
