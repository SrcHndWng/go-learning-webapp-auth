package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
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

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// Our API is going to consist of three routes
	// /status - which we will call to make sure that our API is up and running
	// /products - which will retrieve a list of products that the user can leave feedback on
	// /products/{slug}/feedback - which will capture user feedback on products
	r.Handle("/status", appHandlers.StatusHandler).Methods("GET")
	r.Handle("/products", appHandlers.ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", appHandlers.AddFeedbackHandler).Methods("POST")
	// GetToken
	r.Handle("/get-token", appHandlers.GetTokenHandler).Methods("GET")

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
