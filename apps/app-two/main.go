package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from app-two"))
	})
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		return nil
	}

	return r
}

func main() {
	fmt.Println("Starting App Two...")
	NewRouter()
}
