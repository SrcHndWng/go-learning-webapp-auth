package handlers

import (
	"net/http"
)

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: add token validate.
	// status 401 unauthorized.
	w.WriteHeader(401)
	w.Write([]byte("your token is not authorized."))

	// TODO: redo comment out.
	// Here we are converting the slice of products to JSON
	// payload, _ := json.Marshal(products)
	// w.Write([]byte(payload))
})
