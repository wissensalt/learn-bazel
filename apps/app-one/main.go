package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from app-one"))
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return nil
	}

	return r
}

func main() {
	fmt.Println("Starting App One...")
	NewRouter()
}
