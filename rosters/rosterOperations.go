package rosters

import (
	"fmt"
	"encoding/xml"
	"gofant/api"
	"github.com/jmoiron/sqlx"
)

func getTeamRoster(accessToken string, leagueTeamKey string) ([]byte, error) {
	var Result RosterApiXml

	var dataFormat = "xml"
	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/team/%s/roster/players?format=%s", leagueTeamKey, dataFormat)
	contents, err := api.DoRequest(string(apiUrl), accessToken)

	if err != nil {
		return nil, err
	}

	err2 := xml.Unmarshal(contents, &Result)

	if err2 != nil {
		return genericRosterErrorSerializer("Failed while unmarshalling into RosterApiXml")
	}
	roster, err3 := TeamRosterSerializer(Result)
	if err3 != nil {
		return genericRosterErrorSerializer("Failed while serializing roster api")
	}

	return roster, nil
}

// Result example: 371.l.102542.t.7
func makeLeagueTeamKey(league string, game_id string, team_id string) string {
	return fmt.Sprintf("%s.l.%s.t.%s", game_id, league, team_id)
}


func getPlayer(accessToken string, player_id string) ([]byte, error) {
	//.p.{player_id}
	// Example: pnfl.p.5479 or 223.p.5479
	// https://fantasysports.yahooapis.com/fantasy/v2/player/371.p.<id>
	return []byte("Not supported"), nil
}

func getAndSavePositionTypes(db *sqlx.DB, accessToken, gameKey string) error {
	var dataFormat = "xml"
	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/game/%s/position_types?format=%s", gameKey, dataFormat)
	contents, err := api.DoRequest(string(apiUrl), accessToken)
	if err != nil {
		return api.ApiError{
			Status: "1234",
			Message: fmt.Sprintf("Failed on api call to %s", apiUrl),
		}
	}
	var Result PositionTypeApiXml
	if err := xml.Unmarshal(contents, &Result); err != nil {
		return err
	}
	game := Result.Game
	gameKey2 := game.GameKey
	gameId := game.GameId
	season := game.Season
	for _, pos := range game.PositionTypes {
		if _, err := db.Exec(INSERT_POSITION_TYPE, gameKey2, gameId, season, pos.Type, pos.DisplayName); err != nil {
			//TODO Handle error
		}
	}
	return nil
}

func getAndSaveStatCategories(db *sqlx.DB, accessToken, gameKey string) error {
	var dataFormat = "xml"
	var apiUrl = fmt.Sprintf(api.BaseYahooApi+ "/game/%s/stat_categories?format=%s", gameKey, dataFormat)
	contents, err := api.DoRequest(string(apiUrl), accessToken)
	if err != nil {
		return api.ApiError{
			Status: "1234",
			Message: fmt.Sprintf("Failed on api call to %s", apiUrl),
		}
	}
	var Result StatCategoryApiXml
	if err := xml.Unmarshal(contents, &Result); err != nil {
		return err
	}
	game := Result.Game
	gameKey2 := game.GameKey
	gameId := game.GameId
	season := game.Season
	for _, stat := range game.StatCategories.Stats {
		for _, position := range stat.PositionTypes {

			var positionTypeId int
			if err := db.QueryRowx(GET_POSITION_TYPE_ID, string(season), string(gameKey2), string(position)).Scan(&positionTypeId); err != nil {
				return api.ApiError{
					Status: "2001",
					Message: "",
				}
			}

			if _, err := db.Exec(INSERT_STAT_CATEGORY, gameKey2, gameId, season, stat.StatId, stat.Name,
			stat.DisplayName, stat.SortOrder, positionTypeId); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}
