package auth

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

// LoginUser describes User.
type LoginUser struct {
	Name     string
	Password string
}

// Password is MD5 hash from 'password'.
var onlyUser = LoginUser{Name: "username", Password: "5f4dcc3b5aa765d61d8327deb882cf99"}

// Login executes authentication.
func Login(r *http.Request) (LoginUser, bool) {
	userName, password, authOK := r.BasicAuth()
	if authOK == false {
		return LoginUser{}, false
	}

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashed := hex.EncodeToString(hasher.Sum(nil))

	if (userName != onlyUser.Name) || (hashed != onlyUser.Password) {
		return LoginUser{}, false
	}

	return onlyUser, true
}
