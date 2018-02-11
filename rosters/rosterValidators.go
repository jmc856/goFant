package rosters

import (
	"net/http"
	"encoding/json"
	"gopkg.in/validator.v2"
)

func ValidateGetPositionTypes(r *http.Request) (map[string]string, error) {
	type Req struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"   json:"password"`
		AccessToken			string			`validate:"nonzero" json:"access_token"`
		GameKey				string			`validate:"nonzero" json:"game_key"`
		Refresh		  		string			`json:"refresh"`
	}

	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(req); errs != nil {
		return nil, errs
	}
	params := map[string]string{
		"access_token": req.AccessToken,
		"username":     req.Username,
		"password":     req.Password,
		"refresh": 		req.Refresh,
		"game_key":     req.GameKey,
	}
	return params, nil
}

func ValidateGetStatCategories(r *http.Request) (map[string]string, error) {
	type Req struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"   json:"password"`
		AccessToken			string			`validate:"nonzero" json:"access_token"`
		GameKey				string			`validate:"nonzero" json:"game_key"`
		Refresh		  		string			`json:"refresh"`
	}

	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(req); errs != nil {
		return nil, errs
	}
	params := map[string]string{
		"access_token": req.AccessToken,
		"username":     req.Username,
		"password":     req.Password,
		"refresh": 		req.Refresh,
		"game_key":     req.GameKey,
	}
	return params, nil
}