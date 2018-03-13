package api

import (
	"gopkg.in/validator.v2"
	"net/http"
	"encoding/json"
)


func ValidateGetReq(r *http.Request) (map[string]string, error) {
	return map[string]string{}, nil
}

func genericValidationError(error string) ([]byte, error) {
	valid := map[string]interface{}{}

	result := ApiError{
		Status: "1",
		Message: error,
		Result: valid,
	}
	return json.Marshal(result)
}


// Validates and parses request parameters of /login
// Returns map of request payload or error in bytes
func ValidateLogin(r *http.Request) (map[string]string, error) {
	type UserRequest struct {
		Username		string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password		string			`validate:"min=8"`
	}
	var request UserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(request); errs != nil {
		return nil, errs
	}
	return map[string]string{"Username": request.Username, "Password": request.Password}, nil
}

// Validates /roster/getRoster
// Returns map of request payload or error in bytes
func ValidateGetRoster(r *http.Request) (map[string]string, error) {
	type Filter struct {
		GameId    []string `validate:"nonzero"`
		LeagueKey []string `validate:"nonzero"`
		TeamId    []string `validate:"nonzero"`
		Season    []string `validate:"nonzero"`
	}

	teamId, _ := r.URL.Query()["team_id"]
	gameId, _ := r.URL.Query()["game_id"]
	leagueKey, _ := r.URL.Query()["league_key"]
	season, _ := r.URL.Query()["season"]

	filter := Filter{GameId: gameId, Season: season, TeamId: teamId, LeagueKey: leagueKey}
	if errs := validator.Validate(filter); errs != nil {
		return nil, errs
	}
	return RequestToMap(r), nil
}

func RequestToMap(r *http.Request) map[string]string {
	m := map[string]string{}
	for k, v := range r.Form {
		m[k] = v[0]
	}
	return m
}

func ValidateUserTeams(r *http.Request) (map[string]string, error) {
	type Filter struct {
		GameId  []string `validate:"nonzero"`
		Season  []string `validate:"nonzero"`
	}

	gameId, _ := r.URL.Query()["game_id"]
	season, _ := r.URL.Query()["season"]

	filter := Filter{GameId: gameId, Season: season}
	if errs := validator.Validate(filter); errs != nil {
		return nil, errs
	}
	return RequestToMap(r), nil
}

func ValidateLeagueTeams(r *http.Request) (map[string]string, error) {
	type Filter struct {
		GameId    []string `validate:"nonzero"`
		LeagueKey []string `validate:"nonzero"`
		Season    []string `validate:"nonzero"`
	}

	gameId, _ := r.URL.Query()["game_id"]
	leagueKey, _ := r.URL.Query()["league_key"]
	season, _ := r.URL.Query()["season"]

	filter := Filter{GameId: gameId, Season: season, LeagueKey: leagueKey}
	if errs := validator.Validate(filter); errs != nil {
		return nil, errs
	}
	return RequestToMap(r), nil
}
