package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretKey = "your-secret-key"

type loginUser struct {
	name     string
	password string
}

// Password is MD5 hash from 'password'.
var user = loginUser{name: "username", password: "5f4dcc3b5aa765d61d8327deb882cf99"}

// GetTokenHandler authenticate with basic auth, creates token and return.
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Basic Authentication
	userName := basicAuth(r)
	if userName == "" {
		w.WriteHeader(401)
		w.Write([]byte("Authentication failed."))
		return
	}

	// Set algorithm
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Set user, period.
	token.Claims = jwt.MapClaims{
		"user": userName,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// Sigin to token.
	tokenString, err := token.SignedString([]byte(secretKey))
	if err == nil {
		w.Write([]byte(tokenString))
	} else {
		w.Write([]byte("Could not generate token"))
	}
})

func basicAuth(r *http.Request) string {
	userName, password, authOK := r.BasicAuth()
	if authOK == false {
		return ""
	}

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashed := hex.EncodeToString(hasher.Sum(nil))

	if (userName != user.name) || (hashed != user.password) {
		return ""
	}

	return userName
}

func validateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})
	if err != nil {
		log.Printf("raise token error : %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write([]byte("your token is not authorized."))
		return nil, err
	}
	return token, nil
}
