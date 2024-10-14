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

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/signup", handlers.Make(handlers.HandleSignup))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTPS Server Started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
