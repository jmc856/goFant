package handlers

import (
	"fmt"
	"net/http"
	"gofant/yahoo"
	"gofant/users"
	"gofant/transactions"
	"gofant/leagues"
	"gofant/rosters"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gofant/authorization"
	"gofant/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
)

func handleMain(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Main")
}

func handleYahooLogin(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return yahoo.HandleYahooAuth(env, params)
}

func handleYahooCallback(env *api.Env, _ http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return yahoo.AuthorizeRequest(env, r, params)
}

func handleLogin(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return users.Login(env.DB, params)
}

func getUser(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return users.GetUser(env.DB, params)
}

func createUser(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return users.CreateUser(env.DB, params)
}

func updateUser(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return users.UpdateUser(env.DB, params)
}

func listUserTeams(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return leagues.GetUserTeams(env.DB, params)
}

func listLeaguesTeams(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return leagues.GetLeagueTeams(env.DB, params)
}

func getRoster(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetRoster(env.DB, params)
}

func getPositionTypes(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetPositionTypes(env.DB, params)
}

func getStatCategories(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetStatCategories(env.DB, params)
}

func createTransaction(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return transactions.CreateTransaction(env.DB, params)
}

func getTransaction(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return transactions.GetTransaction(env.DB, params)
}

func listTransaction(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return transactions.ListTransactions(env.DB, params)
}

func editTransaction(env *api.Env, _ http.ResponseWriter, _ *http.Request, params map[string]string) ([]byte, error) {
	return transactions.EditTransaction(env.DB, params)
}

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func Handlers(db *sqlx.DB) *mux.Router {

	// Initialise our app-wide environment with the services/info we need.
	env := &api.Env{
		DB: db,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleMain)
	r.HandleFunc("/get-token", authorization.GetTokenHandler).Methods("GET")

	r.Handle("/login", api.Handler{ env, api.Validator{api.ValidateLogin}, handleLogin}).Methods("POST")
	r.Handle("/YahooLogin",  api.Handler{env, api.Validator{yahoo.ValidateYahooLogin}, handleYahooLogin}).Methods("POST")
	r.Handle("/YahooCallback", api.Handler{env, api.Validator{yahoo.ValidateYahooCallback}, handleYahooCallback}).Methods("POST")

	r.Handle("/users/{userId:[0-9]+}", jwtMiddleware.Handler(api.Handler{ env, api.Validator{api.ValidateGetReq}, getUser})).Methods("GET")
	r.Handle("/users", jwtMiddleware.Handler(api.Handler{ env, api.Validator{users.ValidateCreateUser}, createUser})).Methods("POST")
	r.Handle("/users/{userId:[0-9]+}", jwtMiddleware.Handler(api.Handler{env, api.Validator{users.ValidateUpdateUser}, updateUser})).Methods("PUT")
	r.Handle("/users/teams", jwtMiddleware.Handler(api.Handler{ env, api.Validator{api.ValidateUserTeams}, listUserTeams})).Methods("GET")

	r.Handle("/leagues/teams", jwtMiddleware.Handler(api.Handler{ env, api.Validator{api.ValidateLeagueTeams}, listLeaguesTeams})).Methods("GET")

	r.Handle("/rosters", jwtMiddleware.Handler(api.Handler{ env, api.Validator{api.ValidateGetRoster}, getRoster})).Methods("GET")
	r.Handle("/rosters/positions", jwtMiddleware.Handler(api.Handler{ env, api.Validator{rosters.ValidateRosterPositions}, getPositionTypes})).Methods("GET")
	r.Handle("/rosters/stats/categories", jwtMiddleware.Handler(api.Handler{ env, api.Validator{rosters.ValidateStatCategories}, getStatCategories})).Methods("GET")

	r.Handle("/transactions", jwtMiddleware.Handler(api.Handler{ env, api.Validator{transactions.ValidateCreateTransaction}, createTransaction})).Methods("POST")
	r.Handle("/transactions/{txId:[0-9]+}",jwtMiddleware.Handler( api.Handler{ env, api.Validator{api.ValidateGetReq}, getTransaction})).Methods("GET")
	r.Handle("/transactions", jwtMiddleware.Handler(api.Handler{env, api.Validator{transactions.ValidateListTransaction}, listTransaction})).Methods("GET")
	r.Handle("/transactions/{txId:[0-9]+}", jwtMiddleware.Handler(api.Handler{env, api.Validator{transactions.ValidateEditTransaction}, editTransaction})).Methods("PUT")

	return r
}
