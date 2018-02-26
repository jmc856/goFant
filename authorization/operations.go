package authorization

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	)

func CreateJWT() string {
	/* Set up a global string for our secret */
	var mySigningKey = []byte("secret")

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "gofant.client"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}
