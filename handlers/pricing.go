package handlers

import (
	"main/views/pricing"
	"net/http"
)

func HandlePricing(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, pricing.Pricing())
}
