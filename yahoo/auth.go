package yahoo

import (
	"net/http"
	"fmt"
	"golang.org/x/oauth2"
	"os"
	"gofant/api"
	"gofant/users"
	"io/ioutil"
	"encoding/json"
	"gofant/models"
)

var YahooEndpoints = oauth2.Endpoint{
	AuthURL:  "https://api.login.yahoo.com/oauth2/request_auth",
	TokenURL: "https://api.login.yahoo.com/oauth2/get_token",
}

var (
		yahooOauthConfig = &oauth2.Config{
		RedirectURL:    "http://127.0.0.1/YahooCallback",
		ClientID:       os.Getenv("client_id_gofant"),
		ClientSecret:   os.Getenv("client_secret_gofant"),
		Endpoint:       YahooEndpoints,
	}
)

// Some random string, random for each request
func getOauthStateFromUser(env *api.Env, username string) string {
	var state string
	env.DB.QueryRowx("SELECT state from users WHERE username=$1", username).Scan(&state)
	return state
}

func getUserFromAuthState(env *api.Env, state string) (users.User, error) {
	var user users.User
	err := env.DB.QueryRowx("SELECT * from users WHERE state=$1", state).StructScan(&user)
	return user, err
}

func AuthorizeRequest(env *api.Env, r *http.Request, _ map[string]string) ([]byte, error) {
	state := r.FormValue("state")
	user, userErr := getUserFromAuthState(env, state)
	if userErr != nil {
		return nil, userErr
	}
	code := r.FormValue("code")
	token, err := yahooOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	guid := token.Extra("xoauth_yahoo_guid")

	response, err := users.GetUserProfileFromYahoo(fmt.Sprintf("%s", guid), token.AccessToken)
	defer response.Body.Close()
	userContents, err := ioutil.ReadAll(response.Body)

	return users.CreateUserProfileAndCredentials(env.DB, user, userContents, token)
}

func HandleYahooAuth(env *api.Env, params map[string]string) ([]byte, error) {
	url := yahooOauthConfig.AuthCodeURL(getOauthStateFromUser(env, params["username"]))
	result := map[string]interface{}{"auth_url": url}
	return json.Marshal(models.ApiResult{Status: "0", Result: result})
}
