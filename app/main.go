package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("App running on port 5000 sfsadf")

	r := router()

	http.ListenAndServe(":5000", r)
}

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
