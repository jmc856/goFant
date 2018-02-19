package rosters

import (
	"time"
)

// Structs for Yahoo API Responses - https://fantasysports.yahooapis.com/fantasy/v2/team/371.l.102542.t.7/roster/players?format=json

// XML Struct for RosterApi
type RosterApiXml struct {
	Team struct {
		TeamKey string `xml:"team_key"`
		Roster struct {
			CoverageType string `xml:"coverage_type"`
			Players []struct {
				PlayerKey string `xml:"player_key"`
				PlayerId string `xml:"player_id"`
				Name struct {
					Full string `xml:"full"`
					First string `xml:"first"`
					Last string `xml:"last"`
				} `xml:"name"`
				Status string `xml:"status"`
				StatusFull string `xml:"status_full"`
				InjuryNote string `xml:"injury_note"`
				TeamName string `xml:"editorial_team_full_name"`
				DisplayPosition string `xml:"display_position"`
				Headshot struct {
					Url string `xml:"url"`
				} `xml:"headshot"`
				ImageUrl string `xml:"image_url"`
				PositionType string `xml:"position_type"`
			} `xml:"players>player"`
		} `xml:"roster"`
	} `xml:"team"`
}

type PositionTypeApiXml struct {
	Game struct {
		GameKey string `xml:"game_key"`
		GameId string `xml:"game_id"`
		Season string `xml:"season"`
		PositionTypes []struct {
			Type string `xml:"type"`
			DisplayName string `xml:"display_name"`
		} `xml:"position_types>position_type"`
	} `xml:"game"`
}

type StatCategoryApiXml struct {
	Game struct {
		GameKey string `xml:"game_key"`
		GameId string `xml:"game_id"`
		Season string `xml:"season"`
		StatCategories struct {
			Stats []struct {
				StatId string `xml:"stat_id"`
				Name string `xml:"name"`
				DisplayName string `xml:"display_name"`
				SortOrder   string `xml:"sort_order"`
				PositionTypes []string `xml:"position_types>position_type"`
			} `xml:"stats>stat"`
		} `xml:"stat_categories"`
	} `xml:"game"`
}


type Player struct {
	PlayerKey 		string
	PlayerId 		string
	Name struct {
		Full 		string
		First 		string
		Last 		string
	}
	Status 			string
	StatusFull 		string
	InjuryNote 		string
	TeamName 		string
	DisplayPosition string
	Headshot struct {
		Url 		string
	}
	ImageUrl 		string
	PositionType 	string
}



// STATS RELATED TABLES AND CORRESPONDING STRUCTS

type PositionType struct {
	ID              int
	CreatedAt		time.Time       `db:"created_at" json:"created_at"`
	GameKey			string       	`db:"game_key" json:"game_key"`
	GameId			string       	`db:"game_id" json:"game_id"`
	Season          string          `db:"season" json:"season"`
	Type            string          `db:"type" json:"type"`
	DisplayName     string          `db:"display_name" json:"display_name"`
}

type StatCategory struct {
	ID              int
	CreatedAt		time.Time       `db:"created_at" json:"created_at"`
	GameKey			string       	`db:"game_key" json:"game_key"`
	GameId			string       	`db:"game_id" json:"game_id"`
	Season          string          `db:"season" json:"season"`
	StatId          string          `db:"stat_id" json:"stat_id"`
	Name            string          `db:"name" json:"name"`
	DisplayName     string          `db:"display_name" json:"display_name"`
	SortOrder       string          `db:"sort_order" json:"sort_order"`
	PositionTypeId  int             `db:"position_type_id" json:"position_type_id"`
}


type RosterPosition struct {
	ID              int
	CreatedAt		time.Time       `db:"created_at" json:"created_at"`
	GameKey			string       	`db:"game_key" json:"game_key"`
	Season          string          `db:"season" json:"season"`
	Position	    string          `db:"position" json:"position"`
	Abbreviation    string          `db:"abbreviation" json:"abbreviation"`
	DisplayName     string          `db:"display_name" json:"display_name"`
	PositionTypeId  int	            `db:"position_type_id" json:"position_type_id"`
}

type StatCategories struct {
	ID              int
	CreatedAt		time.Time       `db:"created_at" json:"created_at"`
	GameKey			string       	`db:"game_key" json:"game_key"`
	Season          string          `db:"season" json:"season"`
	StatId		    string          `db:"stat_id" json:"stat_id"`
	Name	        string          `db:"name" json:"name"`
	DisplayName     string          `db:"display_name" json:"display_name"`
	SortOrder	    string          `db:"sort_order" json:"sort_order"`
	PositionTypeId  int	            `db:"position_type_id" json:"position_type_id"`
}
