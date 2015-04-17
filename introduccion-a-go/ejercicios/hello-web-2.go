package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

func main() {
	t, _ := template.ParseFiles(`flisol.tmpl`)

	http.Handle(`/saluda/`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args := map[string]interface{}{
			`nombre`: r.RequestURI[8:],
			`accion`: `Hola`,
		}

		t.Execute(w, args)
	}))

	http.Handle(`/despide/`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args := map[string]interface{}{
			`nombre`: r.RequestURI[9:],
			`accion`: `Adiós`,
		}

		t.Execute(w, args)
	}))

	http.ListenAndServe(`:8090`, nil)
}
