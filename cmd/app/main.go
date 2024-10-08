package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	router.Get("/foo", handleFoo)
	fmt.Println("hello world!")

	listenAddr := ":3000"
	slog.Info("HTTPS Server Started", listenAddr)
	http.ListenAndServe(listenAddr, router)
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
