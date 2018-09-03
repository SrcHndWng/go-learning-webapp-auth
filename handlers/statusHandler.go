package handlers

import "net/http"

// StatusHandler will be invoked when the user calls the /status route.
// It will simply return a string with the message "API is up and running"
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
