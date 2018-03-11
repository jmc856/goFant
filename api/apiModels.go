package api

import (
	"net/http"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
	"strings"
	"github.com/gorilla/mux"
)

type YahooApiError struct {
	Description      string `xml:"description"`
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

type ApiError struct {
	error
	Status string          `json:"status"`
	ErrorString string	   `json:"error"`
	Message string		   `json:"message"`
	Result interface{}     `json:"result"`
}
// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

func (ae ApiError) Convert() []byte {
	js, _ := json.Marshal(ae)
	return js
}
// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

func (ae ApiError) Error() string {
	return ae.error.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

type Validator struct {
	Validate func(r *http.Request) (map[string]string, error)
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	Validator
	H func(e *Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error)
}

// A (simple) example of our application-wide configuration.
type Env struct {
	DB   *sqlx.DB
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println(fmt.Sprintf("Call to %s from %s - %d bytes", r.URL, r.Host, r.ContentLength))
	params, validationErr := h.Validator.Validate(r)
	if validationErr != nil {
		jsonErr, _ := genericValidationError(validationErr.Error())
		returnJson(w, jsonErr)
		return
	}

	// Gets path variables from request and adds to params
	vars := mux.Vars(r)
	for k, v := range vars {
		params[k] = v
	}

	accessToken := r.Header.Get("Authorization")
	if accessToken != "" {
		splitToken := strings.Split(accessToken, "Bearer")
		accessToken = splitToken[1][1:]
		username, _ := getUserFromAccessToken(h.Env.DB, accessToken)
		params["request_username"] = username
	}

	jsonResult, err := h.H(h.Env, w, r, params)
	switch r.Method {
		case "POST":
			w.WriteHeader(http.StatusCreated)
		case "DELETE":
			w.WriteHeader(http.StatusNoContent)
	}
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())

		case ApiError:
			returnJson(w, e.Convert())
			return
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			log.Println(e)
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	} else {
		returnJson(w, jsonResult)
		return
	}
}

func returnJson(w http.ResponseWriter, json []byte) int {
	result, err := w.Write(json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return result
}

func getUserFromAccessToken(db *sqlx.DB, accessToken string) (string, error) {

	var username string
	if dbErr := db.QueryRowx(selectUserFromAccessToken, accessToken).Scan(&username); dbErr != nil {
		fmt.Println(dbErr)
		return "", dbErr
	}
	return username, nil
}

const selectUserFromAccessToken = `SELECT u.username FROM users u JOIN user_credentials uc ON u.id=uc.user_id WHERE uc.access_token=$1`