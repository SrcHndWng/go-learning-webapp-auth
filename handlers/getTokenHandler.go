package handlers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretKey = "your-secret-key"

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// アルゴリズムの指定
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// ユーザ、有効期限を設定
	token.Claims = jwt.MapClaims{
		"user": "Guest",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// トークンに対して署名の付与
	tokenString, err := token.SignedString([]byte(secretKey))
	if err == nil {
		w.Write([]byte(tokenString))
	} else {
		w.Write([]byte("Could not generate token"))
	}
})

func validateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write([]byte("your token is not authorized."))
		return nil, err
	}
	return token, nil
}
