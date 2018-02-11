package leagues

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)


/// https://fantasysports.yahooapis.com/fantasy/v2/users;use_login=1/games;season='%s';game_keys=%s/teams?format=json

//// Data structures for receiving Yahoo JSON api calls
type TeamApi struct {
	TeamKey string `json:"team_key"`
	TeamId  string `json:"team_id"`
	Name    string `json:"name"`

}


//// Data structures for receiving Yahoo XML api calls

type UserTeamApiXml struct {
	Users []struct {
		Guid string `xml:"guid"`
		Games []struct {
			GameKey    string  `xml:"game_key"`
			GameId     string  `xml:"game_id"`
			Name       string  `xml:"name"`
			Code       string  `xml:"code"`
			Season     string  `xml:"season"`

			Teams []struct {
				TeamKey    				string  `xml:"team_key"`
				TeamId     				string  `xml:"team_id"`
				Name       			    string  `xml:"name"`
				IsOwnedByCurrentLogin   string  `xml:"is_owned_by_current_login"`
				WaiverPriority 			string `xml:"waiver_priority"`
				FaabBalance 			string `xml:"faab_balance"`
				NumOfMoves 				string `xml:"number_of_moves"`
				NumOfTrades      	    string `xml:"number_of_trades"`

				TeamLogos  []struct {
					Size 	string `xml:"size"`
					Url 	string  `xml:"url"`
				} `xml:"team_logos>team_logo"`

				Managers  []struct {
					ManagerId      string  `xml:"manager_id"`
					Nickname       string  `xml:"nickname"`
					Guid           string  `xml:"guid"`
					Comanager      string  `xml:"is_comanager"`
					Email          string  `xml:"email"`
					Image_url      string  `xml:"image_url"`
				} `xml:"managers>manager"`
			} `xml:"teams>team"`
		} `xml:"games>game"`
	} `xml:"users>user"`
}


type LeagueTeamsApiXml struct {
	League struct {
		LeagueKey string `xml:"league_key"`
		CurrentWeek string `xml:"current_week"`
		Name string `xml:"name"`
		Season string `xml:"season"`
		Teams []struct {
			Name    string `xml:"name"`
			TeamKey string `xml:"team_key"`
			TeamLogos []struct {
				Size string `xml:"large"`
				Url  string `xml:"url"`

			} `xml:"team_logos>team_logo"`

			WaiverPriority string `xml:"waiver_priority"`
			FaabBalance    string `xml:"faab_balance"`
			NumOfMoves     string `xml:"number_of_moves"`
			NumOfTrades    string `xml:"number_of_trades"`

			Managers []struct {
				ManagerId  string  `xml:"manager_id"`
				Nickname   string  `xml:"nickname"`
				Guid       string  `xml:"guid"`
				Email      string  `xml:"email"`
				ImageUrl   string  `xml:"image_url"`

			} `xml:"managers>manager"`

		} `xml:"teams>team"`

	} `xml:"league"`
}

type LeagueApiXml struct {
	League struct {
		CurrentWeek string `xml:"current_week"`
		DraftStatus string `xml:"draft_status"`
		Name        string `xml:"name"`
		LeagueKey   string `xml:"league_key"`
		LeagueType  string `xml:"league_type"`
		NumTeams    int    `xml:"num_teams"`
		Season      string  `xml:"season"`
		GameCode    string  `xml:"game_code"`
		LeagueId	string  `xml:"league_id"`
	} `xml:"league"`
}

// Database Objects for storage
type League struct {
	ID              int
	CreatedAt		time.Time       `db:"created_at" json:"created_at"`
	Name            string          `db:"name" json:"name"`
	LeagueKey       string          `db:"league_key" json:"league_key"`
	CurrentWeek     int				`db:"current_week" json:"current_week"`
	DraftStatus     string          `db:"draft_status" json:"draft_status"`
	NumTeams        int             `db:"num_teams" json:"num_teams"`
	LeagueType      string          `db:"league_type" json:"league_type"`
	Season          string			`db:"season" json:"season"`
	GameId          string          `db:"game_id" `
	LeagueId        int				`db:"league_id" json:"league_id"`
}

type Team struct {
	ID          int
	CreatedAt   time.Time  	`db:"created_at" json:"created_at"`
	Name		string  	`db:"name" json:"name"`
	LeagueId    string  	`db:"league_id" json:"league_id"`
	TeamId		string		`db:"team_id" json:"team_id"`
	TeamKey		string		`db:"team_key" json:"team_key"`
	GameId		string		`db:"game_id" json:"game_id"`
}

const CREATE_TABLE_LEAGUES = `CREATE TABLE IF NOT EXISTS leagues (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
name VARCHAR(50) NOT NULL,
league_key VARCHAR(100) NOT NULL,
current_week INTEGER NOT NULL,
draft_status VARCHAR(50) NOT NULL,
league_type  VARCHAR(50) NOT NULL,
num_teams    INTEGER NOT NULL,
season VARCHAR(10) NOT NULL,
game_id VARCHAR(30) NOT NULL,
PRIMARY KEY (id),
UNIQUE (league_key)
);`

const CREATE_TABLE_TEAMS = `CREATE TABLE IF NOT EXISTS teams (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
name VARCHAR(50) NOT NULL,
league_id INT NOT NULL,
team_id INT NOT NULL,
team_key VARCHAR(30) NOT NULL,
-- manager_id INT NOT NULL,
game_id VARCHAR(30) NOT NULL,
PRIMARY KEY (id),
CONSTRAINT fk_league_id FOREIGN KEY (league_id) REFERENCES leagues (id),
UNIQUE (team_key)
);`

const CLEAR_ALL_TABLES = `DELETE from teams; DELETE FROM leagues;`


func MigrateLeagues(db *sqlx.DB) error {
	fmt.Println("Creating table leagues")
	_, err := db.Exec(CREATE_TABLE_LEAGUES)
	if err != nil {
		return err
	}

	fmt.Println("Creating table teams")
	_, err2 := db.Exec(CREATE_TABLE_TEAMS)
	if err2 != nil {
		return err2
	}
	return nil
}

func RemoveLeaguesTables(db *sqlx.DB) error {
	fmt.Println("Clearing all leagues tables")
	_, err := db.Exec(CLEAR_ALL_TABLES)
	if err != nil {
		return err
	}
	return nil
}
