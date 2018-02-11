package leagues

import (
	"encoding/json"
	"gofant/models"
)

func genericLeagueErrorSerializer(error string) ([]byte, error) {
	league := map[string]interface{}{
	}

	result := models.GenericApiError{
		Status: "1",
		Error: error,
		Result: league,
	}

	return json.Marshal(result)
}

func TeamSerializer(UserTeams UserTeamApiXml) ([]byte, error) {
	userTeamList := make([]map[string]interface{}, 0)
	game := UserTeams.Users[0].Games[0]
	teams := UserTeams.Users[0].Games[0].Teams

	for _, team := range teams {
		gameId, leagueKey, teamId := parseTeamKey(team.TeamKey)

		TeamResult := map[string]interface{} {
			"name":     team.Name,
			"team_key": team.TeamKey,
			"league_key": leagueKey,
			"team_id": teamId,
			"game_id": gameId,
			"season": game.Season,
			"current_week": team,
			"faab_balance": team.FaabBalance,
			"num_moves": team.NumOfMoves,
			"num_trades": team.NumOfTrades,
			"display_position": team.WaiverPriority,
			"managers": team.Managers,
			"image_url": team.TeamLogos,
		}
		userTeamList = append(userTeamList, TeamResult)
	}

	Result := map[string]interface{}{
		"teams": userTeamList,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}

func LeagueTeamsSerializerXml(LeagueTeamXml LeagueTeamsApiXml) ([]byte, error) {
	teamList := make([]map[string]interface{}, 0)
	teams := LeagueTeamXml.League.Teams
	for _, team := range teams {
		TeamResult := map[string]interface{} {
			"name":     team.Name,
			"team_key": team.TeamKey,
			"faab_balance": team.FaabBalance,
			"num_moves": team.NumOfMoves,
			"num_trades": team.NumOfTrades,
			"display_position": team.WaiverPriority,
			"managers": team.Managers,
			"image_url": team.TeamLogos,
		}
		teamList = append(teamList, TeamResult)
	}

	Result := map[string]interface{}{
		"name":             LeagueTeamXml.League.Name,
		"league_key":       LeagueTeamXml.League.LeagueKey,
		"current_week":     LeagueTeamXml.League.CurrentWeek,
		"season": 			LeagueTeamXml.League.Season,
		"teams": 			teamList,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}