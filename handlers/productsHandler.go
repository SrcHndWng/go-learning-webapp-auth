package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SrcHndWng/go-learning-webapp-auth/auth"
)

// ProductsHandler will be called when the user makes a GET request to the /products endpoint.
// This handler will return a list of products available for users to review.
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if _, err := auth.ValidateToken(w, r); err != nil {
		return
	}

	// Here we are converting the slice of products to JSON
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(products)
	w.Write([]byte(payload))
})
