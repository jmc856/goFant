package rosters

import (
	"errors"
	"strings"
	"fmt"
	"gofant/api"
)

func GetRoster(_ *api.Env, params map[string]string) ([]byte, error) {

	leagueTeamKey := makeLeagueTeamKey(params["league_key"], params["game_id"], params["team_id"])

	if strings.ToLower(params["Refresh"]) == "true" {
		//var result RosterApi
		userRoster, err := getTeamRoster(params["access_token"], leagueTeamKey)
		if err != nil {
			return nil, err
		}

		return userRoster, nil

	} else {
		userRoster, err := genericRosterErrorSerializer("Database persistence not supported yet. Must use refresh=true in request")
		if err != nil {
			return nil, errors.New(fmt.Sprint("Error in genericRosterErrorSerializer: %s", err.Error()))
		}
		return userRoster, nil
	}
}


func GetPositionTypes(env *api.Env, params map[string]string) ([]byte, error) {

	// Refresh database
	if params["refresh"] == "true" {
		if err := getAndSavePositionTypes(env.DB, params["access_token"], params["game_key"]); err != nil {
			return nil, api.ApiError{}
		}
	}
	// Return current database
	var positionTypes []PositionType
	if dbErr := env.DB.Select(&positionTypes, GET_POSITION_TYPES); dbErr != nil {
		return nil, dbErr
	}
	return RosterPositionSerializer(positionTypes)

}

func GetStatCategories(env *api.Env, params map[string]string) ([]byte, error) {
	// Refresh database
	if params["refresh"] == "true" {
		if err := getAndSaveStatCategories(env.DB, params["access_token"], params["game_key"]); err != nil {
			return nil, api.ApiError{}
		}
	}
	// Return current database
	var statCategories []StatCategory
	if dbErr := env.DB.Select(&statCategories, GET_STAT_CATEGORIES); dbErr != nil {
		return nil, dbErr
	}
	return StatCategorySerializer(statCategories)}