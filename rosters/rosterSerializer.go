package rosters

import (
	"gofant/models"
	"encoding/json"
	"gofant/api"
)


func genericRosterErrorSerializer(error string) ([]byte, error) {
	roster := map[string]interface{}{}

	result := models.GenericApiError{
		Status: api.UnknownError,
		Error: error,
		Result: roster,
	}

	return json.Marshal(result)
}


func TeamRosterSerializer(RosterXml RosterApiXml) ([]byte, error) {
	playerList := make([]map[string]interface{}, 0)
	players := RosterXml.Team.Roster.Players
	for _, player := range players {
		PlayerResult := map[string]interface{} {
			"player_key": player.PlayerKey,
			"player_id": player.PlayerId,
			"name": player.Name,
			"status": player.Status,
			"status_full": player.StatusFull,
			"injury_note": player.InjuryNote,
			"display_position": player.DisplayPosition,
			"position_type": player.PositionType,
			"team_name": player.TeamName,
			"image_url": player.ImageUrl,
		}
		playerList = append(playerList, PlayerResult)
	}

	Result := map[string]interface{}{
		"roster": playerList,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}

func RosterPositionSerializer(positionTypes []PositionType) ([]byte, error) {
	Result := map[string]interface{}{
		"position_types":  positionTypes,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}

func StatCategorySerializer(statCategories []StatCategory) ([]byte, error) {
	Result := map[string]interface{}{
		"stat_categories":  statCategories,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}