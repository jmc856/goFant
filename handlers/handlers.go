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
	"gofant/api"
	"github.com/jmoiron/sqlx"
	"gofant/authorization"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main")
}

func handleYahooLogin(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return yahoo.HandleYahooAuth(env, params)
}

func handleYahooCallback(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return yahoo.AuthorizeRequest(env, r, params)
}

func handleLogin(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return users.LoginUsername(env.DB, params)
}

func createUser(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return users.CreateUser(env, params)
}

func updateUser(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return users.UpdateUser(env, params)
}

func getUserTeams(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return leagues.GetUserTeams(env, params)
}

func getLeagueTeams(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return leagues.GetLeagueTeams(env, params)
}

func getRoster(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetRoster(env, params)
}

func getPositionTypes(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetPositionTypes(env, params)
}

func getStatCategories(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return rosters.GetStatCategories(env, params)
}

func createTransaction(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return transactions.CreateTransaction(env, params)
}

func getTransaction(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return transactions.GetTransaction(params)
}

func listTransaction(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return transactions.ListTransactions(env.DB, params)
}

func acceptTransaction(env *api.Env, w http.ResponseWriter, r *http.Request, params map[string]string) ([]byte, error) {
	return transactions.AcceptTransaction(env, params)
}

func Handlers(db *sqlx.DB) *mux.Router {

	// Initialise our app-wide environment with the services/info we need.
	env := &api.Env{
		DB: db,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleMain)
	r.HandleFunc("/get-token", authorization.GetTokenHandler).Methods("GET")
	r.Handle("/login", api.Handler{ env, api.Validator{api.ValidateLogin}, handleLogin})
	r.Handle("/YahooLogin",  api.Handler{env, api.Validator{yahoo.ValidateYahooLogin}, handleYahooLogin})
	r.Handle("/YahooCallback", api.Handler{env, api.Validator{yahoo.ValidateYahooCallback}, handleYahooCallback})

	r.Handle("/users/create", api.Handler{ env, api.Validator{users.ValidateCreateUser}, createUser})
	r.Handle("/users/update", api.Handler{env, api.Validator{users.ValidateUpdateUser}, updateUser})
	r.Handle("/users/getTeams", api.Handler{ env, api.Validator{api.ValidateUserTeams}, getUserTeams})

	r.Handle("/leagues/getLeagueTeams", api.Handler{ env, api.Validator{api.ValidateLeagueTeams}, getLeagueTeams})

	r.Handle("/rosters/getRoster", api.Handler{ env, api.Validator{api.ValidateGetRoster}, getRoster})
	r.Handle("/rosters/getpositiontypes", api.Handler{ env, api.Validator{rosters.ValidateGetPositionTypes}, getPositionTypes})
	r.Handle("/rosters/getstatcategories", api.Handler{ env, api.Validator{rosters.ValidateGetStatCategories}, getStatCategories})

	r.Handle("/transactions/create", api.Handler{ env, api.Validator{transactions.ValidateCreateTransaction}, createTransaction})
	r.Handle("/transactions/get", api.Handler{ env, api.Validator{transactions.ValidateGetTransaction}, getTransaction})
	r.Handle("/transactions/list", api.Handler{env, api.Validator{transactions.ValidateListTransaction}, listTransaction})
	r.Handle("/transactions/accept", api.Handler{env, api.Validator{transactions.ValidateAcceptTransaction}, acceptTransaction})

	return r
}
