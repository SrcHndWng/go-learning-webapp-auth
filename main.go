package main

import (
	"net/http"
	"os"

	appHandlers "github.com/SrcHndWng/go-learning-webapp-auth/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.Handle("/status", appHandlers.StatusHandler).Methods("GET")
	r.Handle("/products", appHandlers.ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", appHandlers.AddFeedbackHandler).Methods("POST")
	r.Handle("/get-token", appHandlers.GetTokenHandler).Methods("GET")

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
