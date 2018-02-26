package authorization

import (
	"net/http"
)


var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	jwt := CreateJWT()
	w.Write([]byte(jwt))
})
