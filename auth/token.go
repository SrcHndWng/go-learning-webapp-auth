package auth

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretKey = "your-secret-key"

// CreateToken creates signed token.
func CreateToken(user LoginUser) (tokenString string, err error) {
	// Set algorithm
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Set user, period.
	token.Claims = jwt.MapClaims{
		"user": user.Name,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// Sigin to token.
	tokenString, err = token.SignedString([]byte(secretKey))

	return
}

// ValidateToken valid token in request.
func ValidateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
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
