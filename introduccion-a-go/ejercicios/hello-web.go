package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle(`/saluda/`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola %s en FLISoL", r.RequestURI[8:])
	}))

	http.Handle(`/despide/`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola %s en FLISoL", r.RequestURI[8:])
	}))

	http.ListenAndServe(`:8090`, nil)
}
