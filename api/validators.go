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
	type RosterRequest struct {
		Username 		string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 		string			`validate:"min=8"`
		AccessToken 	string			`validate:"nonzero"`
		Game_Id 		string			`validate:"nonzero"`
		League_Key      string			`validate:"nonzero"`
		Team_Id      	string			`validate:"nonzero"`
		Refresh     	string			`validate:"regexp=^true$"`
	}
	var request RosterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(request); errs != nil {
		return nil, errs
	}
	return map[string]string{
		"Username":    		request.Username,
		"Password":    		request.Password,
		"access_token": 	request.AccessToken,
		"game_id":      	request.Game_Id,
		"league_key":      	request.League_Key,
		"team_id":        	request.Team_Id,
		"Refresh":     		request.Refresh,
	}, nil
}

func ValidateUserTeams(r *http.Request) (map[string]string, error) {
	type UserTeamsRequest struct {
		Username 		string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 		string			`validate:"min=8"`
		AccessToken 	string			`validate:"nonzero"`
		Game_Id 		string			`validate:"nonzero"`
		Season 			string			`validate:"nonzero"`
		UserId          string          `validate:"nonzero" json:"user_id"`
		Refresh     	string			`validate:"regexp=^true$"`
	}
	var request UserTeamsRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(request); errs != nil {
		return nil, errs
	}
	return map[string]string{
		"Username":    		request.Username,
		"Password":    		request.Password,
		"access_token":		request.AccessToken,
		"game_id":      	request.Game_Id,
		"Season":      		request.Season,
		"user_id":			request.UserId,
		"Refresh":     		request.Refresh,
	}, nil
}

func ValidateLeagueTeams(r *http.Request) (map[string]string, error) {
	type UserTeamsRequest struct {
		Username 	string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 	string			`validate:"min=8"`
		AccessToken string			`validate:"nonzero"`
		Game_Id 	string			`validate:"nonzero"`
		League_Key  string			`validate:"nonzero"`
		Refresh     string			`validate:"regexp=^true$"`
	}
	var request UserTeamsRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(request); errs != nil {
		return nil, errs
	}
	return map[string]string{
		"Username":    		request.Username,
		"Password":    		request.Password,
		"access_token": 	request.AccessToken,
		"game_id":      	request.Game_Id,
		"league_key":   	request.League_Key,
		"Refresh":     		request.Refresh,
	}, nil
}
