package leagues

import (
	"encoding/xml"
	"github.com/jmoiron/sqlx"
)

func GetUserTeams(db *sqlx.DB, params map[string]string) ([]byte, error) {
	if params["refresh"] == "true" {
		var result UserTeamApiXml
		userTeams, err := getUserTeam(params["yahoo_access_token"], params["game_id"], params["season"])
		if err != nil {
			return nil, err
		}
		xml.Unmarshal(userTeams, &result)

		if errCreate := saveUserTeams(db, params["yahoo_access_token"], params["user_id"], result); errCreate != nil {
			return nil, errCreate
		}

		leagueInfo, err := TeamSerializer(result)

		return leagueInfo, nil

	} else {
		return genericLeagueErrorSerializer("Does not support database persistence yet.  Set refresh=true in request params")
	}
}

func GetLeagueTeams(db *sqlx.DB, params map[string]string) ([]byte, error) {

	if params["refresh"] == "true" {
		leagueTeams, err := getLeagueTeams(params["yahoo_access_token"], formatLeagueKey(params["game_id"], params["league_key"]))
		if err != nil {
			return nil, err
		}
		return leagueTeams, nil
	}
	//todo: Add handling of persisted LeagueTeams
	return genericLeagueErrorSerializer("Data persistence not supported")
}
