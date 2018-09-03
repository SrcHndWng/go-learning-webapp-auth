package handlers

import (
	"net/http"

	"github.com/SrcHndWng/go-learning-webapp-auth/auth"
)

// GetTokenHandler authenticate with basic auth, creates token and return.
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, isLogin := auth.Login(r)
	if !isLogin {
		w.WriteHeader(401)
		w.Write([]byte("Authentication failed."))
		return
	}

	tokenString, err := auth.CreateToken(user)
	if err == nil {
		w.Write([]byte(tokenString))
	} else {
		w.Write([]byte("Could not generate token"))
	}
})
