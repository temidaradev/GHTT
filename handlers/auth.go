package handlers

import (
	"main/views/auth"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Signup())
}
