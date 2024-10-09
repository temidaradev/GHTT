package main

import (
	"log"
	"log/slog"
	"main/handlers"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Get("/foo", handlers.Make(handlers.HandleFoo))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTPS Server Started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
