package leagues

import (
	"gofant/users"
	"encoding/xml"
	"github.com/jmoiron/sqlx"
)

func GetUserTeams(db *sqlx.DB, params map[string]string) ([]byte, error) {
	if params["Refresh"] == "true" {
		var result UserTeamApiXml
		userTeams, err := getUserTeam(params["access_token"], params["game_id"], params["Season"])
		if err != nil {
			return nil, err
		}
		xml.Unmarshal(userTeams, &result)

		if errCreate := saveUserTeams(db, params["access_token"], params["user_id"], result); errCreate != nil {
			return nil, errCreate
		}

		leagueInfo, err := TeamSerializer(result)

		return leagueInfo, nil

	} else {
		return genericLeagueErrorSerializer("Does not support database peristence yet.  Set refresh=true in request params")
	}
}

func GetLeagueTeams(db *sqlx.DB, params map[string]string) ([]byte, error) {

	_, userErr := users.GetUserFromPassword(db, params["username"], params["password"])
	if userErr != nil {
		return nil, userErr
	}

	if params["Refresh"] == "true" {
		leagueTeams, err := getLeagueTeams(params["access_token"], formatLeagueKey(params["game_id"], params["league_key"]))
		if err != nil {
			return nil, err
		}

		return leagueTeams, nil
	}
	//todo: Add handling of persisted LeagueTeams

	return genericLeagueErrorSerializer("Data persistence not supported")
}
