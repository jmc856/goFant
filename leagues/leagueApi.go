package leagues

import (
	"gofant/models"
	"gofant/users"
	"encoding/xml"
	"gofant/api"
)

func GetUserTeams(env *api.Env, params map[string]string) ([]byte, error) {
	if params["Refresh"] == "true" {
		var result UserTeamApiXml
		userTeams, err := getUserTeam(params["access_token"], params["game_id"], params["Season"])
		if err != nil {
			return nil, err
		}
		xml.Unmarshal(userTeams, &result)

		if errCreate := saveUserTeams(env.DB, params["access_token"], params["user_id"], result); errCreate != nil {
			return nil, errCreate
		}

		leagueInfo, err := TeamSerializer(result)

		return leagueInfo, nil

	} else {
		return genericLeagueErrorSerializer("Does not support database peristence yet.  Set refresh=true in request params")
	}
}

func GetLeagueTeams(req map[string]string) ([]byte, error) {
	var user users.User

	db := models.OpenDataBase()
	defer db.Close()

	// Retrieve user from database
	db.Where(&users.User{Username: req["Username"],
		Password: req["Password"],
	}).First(&user)

	// Retrieve team from database
	var team Team
	db.Model(&user).Related(&team)

	if req["Refresh"] == "true" {
		leagueTeams, err := getLeagueTeams(req["access_token"], formatLeagueKey(req["game_id"], req["league_key"]))
		if err != nil {
			return nil, err
		}

		return leagueTeams, nil
	}
	//todo: Add handling of persisted LeagueTeams

	return genericLeagueErrorSerializer("Data persistence not supported")
}
