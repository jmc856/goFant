package leagues

import (
	"fmt"
	"errors"
	"strings"
	"gofant/api"
	"encoding/xml"
	"github.com/jmoiron/sqlx"
)

var AllGameCodes = []string{"nfl", "mlb", "nhl", "nba"}

func getUserTeam(accessToken string, gameId string, season string) ([]byte, error) {
	var format = "xml"

	if !_checkElement(gameId, AllGameCodes) {
		return genericLeagueErrorSerializer("Season not supported")
	}

	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/users;use_login=1/games;season='%s';game_keys=%s/teams?format=%s", season, gameId, format)

	contents, err := api.DoRequest(string(apiUrl), accessToken)

	if err != nil {
		errors.New("I dno")
	}

	return contents, nil
}

func saveUserTeams(db *sqlx.DB, accessToken, userId string, userTeams UserTeamApiXml) error {
	var teamStruct Team
	var league League
	teams := userTeams.Users[0].Games[0].Teams

	for _, team := range teams {
		gameId, leagueKey, teamId := parseTeamKey(team.TeamKey)

		// Check that league exists with teams leagueKey
		rows, _ := db.Queryx(GET_LEAGUES, formatLeagueKey(gameId, leagueKey))

		if rows.Next() == false {
			leagueStruct, err := getOrCreateLeague(db, formatLeagueKey(gameId, leagueKey), accessToken)
			if err != nil {
				return err
			}
			if dbErr := db.QueryRowx(CREATE_OR_UPDATE_TEAM, team.Name, leagueStruct.ID, teamId, team.TeamKey, gameId).StructScan(&teamStruct); dbErr != nil {
				return api.ApiError{
					Status: "2001",
					Message: "CREATE TEAM ERROR",
				}
			}

		} else {
			if dbErr := rows.StructScan(&league); dbErr != nil {
				return dbErr
			}
			if err := db.QueryRowx(CREATE_OR_UPDATE_TEAM, team.Name, league.ID, teamId, team.TeamKey, gameId).StructScan(&teamStruct); err != nil {
				return api.ApiError{
					Status: "2001",
					Message: "CREATE TEAM ERROR",
				}
			}
		}
		// Create user_team join table row
		if _, err := db.Queryx(CREATE_OR_UPDATE_USER_TEAM, userId, teamStruct.ID); err != nil {
			return api.ApiError{
				Status: "2001",
				Message: "CREATE TEAM ERROR",
			}
		}
	}
	return nil
}

func getLeagueTeams(accessToken string, leagueKey string) ([]byte, error) {
	var Result LeagueTeamsApiXml

	var dataFormat = "xml"
	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/league/%s/teams?format=%s", leagueKey, dataFormat)

	contents, err := api.DoRequest(string(apiUrl), accessToken)

	if err != nil {
		return nil, err
	}

	err2 := xml.Unmarshal(contents, &Result)

	if err2 != nil {
		return genericLeagueErrorSerializer("Failed while unmarshalling into LeagueTeamsApi")
	}
	return LeagueTeamsSerializerXml(Result)
}

func getOrCreateLeague(db *sqlx.DB, leagueKey, accessToken string) (League, error) {
	var leagueResult LeagueApiXml

	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/league/%s", leagueKey)
	contents, err := api.DoRequest(string(apiUrl), accessToken)
	if err != nil {
		return League{}, err
	}

	err2 := xml.Unmarshal(contents, &leagueResult)
	if err2 != nil {
		return League{}, nil
	}

	league := leagueResult.League
	var leagueStruct League
	err3 := db.QueryRowx(CREATE_OR_UPDATE_LEAGUE, league.Name, league.LeagueKey, league.CurrentWeek, league.DraftStatus, league.LeagueType,
				 league.NumTeams, league.Season, league.GameCode).StructScan(&leagueStruct)
	fmt.Println(err3)
	return leagueStruct, nil
}

func _checkElement(s string, arr []string) bool {
	for _, b := range arr {
		if strings.ToLower(s) == b {
			return true
		}
	}
	return false
}

// Takes Yahoo league
func parseTeamKey(team_key string) (string, string, string) {
	s := strings.Split(team_key, ".")
	league, _, leagueId, _, teamKey := s[0], s[1], s[2], s[3], s[4]
	return league, leagueId, teamKey

}

func formatLeagueKey(gameId string, leagueKey string) string {
	return fmt.Sprintf("%s.l.%s", gameId, leagueKey)
}
