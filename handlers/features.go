package handlers

import (
	"main/views/features"
	"net/http"
)

func HandleFeatures(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, features.Features())
}
