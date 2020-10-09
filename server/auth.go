package server

import (
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "supersecret"

func JWTChecker(next http.Handler) http.Handler {
	checker := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		},
		SigningMethod:       jwt.SigningMethodHS256,
		EnableAuthOnOptions: true,
		Extractor: jwtmiddleware.FromFirst(
			jwtmiddleware.FromAuthHeader,
			jwtmiddleware.FromParameter("access_token"), // TODO: unsafe but needed for browser websocket auth
		),
	})
	return checker.Handler(next)

}

func CreateToken(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return signer.SignedString([]byte(JWT_SECRET))
}
