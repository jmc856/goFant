package rosters

import (
	"net/http"
	"gofant/api"
	"gopkg.in/validator.v2"
)

func ValidateRosterPositions(r *http.Request) (map[string]string, error) {
	type Filter struct {
		GameKey		[]string `validate:"nonzero"`
		Refresh     []string
	}

	gameKey, _ := r.URL.Query()["game_key"]
	refresh, _ := r.URL.Query()["refresh"]

	filter := Filter{GameKey: gameKey, Refresh: refresh}
	if errs := validator.Validate(filter); errs != nil {
		return nil, errs
	}
	return api.RequestToMap(r), nil
}

func ValidateStatCategories(r *http.Request) (map[string]string, error) {
	type Filter struct {
		GameKey		[]string `validate:"nonzero"`
		Refresh     []string
	}

	gameKey, _ := r.URL.Query()["game_key"]
	refresh, _ := r.URL.Query()["refresh"]

	filter := Filter{GameKey: gameKey, Refresh: refresh}
	if errs := validator.Validate(filter); errs != nil {
		return nil, errs
	}
	return api.RequestToMap(r), nil
}