package handlers

import "net/http"

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
