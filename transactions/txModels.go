package transactions

import (
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)

func MigrateTransactions(db *sqlx.DB) error {
	fmt.Println("Creating table user_teams")
	_, err := db.Exec(CREATE_TABLE_USER_TEAMS)
	if err != nil {
		return err
	}
	fmt.Println("Creating table tx")
	_, err3 := db.Exec(CREATE_TABLE_TX)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("Creating table tx_detail")
	_, err2 := db.Exec(CREATE_TABLE_TX_DETAIL)
	if err2 != nil {
		fmt.Println(err2)
	}

	return nil
}

func RemoveTransactionsTables(db *sqlx.DB) error {
	fmt.Println("Clearing all transaction tables")
	_, err := db.Exec(CLEAR_ALL_TABLES)
	if err != nil {
		return err
	}
	return nil
}


/*
 many to many: many users have many teams
*/

const CREATE_TABLE_USER_TEAMS = `CREATE TABLE IF NOT EXISTS user_teams (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
user_id int NOT NULL,
team_id int NOT NULL,
PRIMARY KEY (id),
UNIQUE (user_id, team_id),
FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE,
FOREIGN KEY (team_id) REFERENCES teams(id) ON UPDATE CASCADE
);`


const CREATE_TABLE_TX_DETAIL = `CREATE TABLE IF NOT EXISTS tx_detail (
id serial,
tx_id INT NOT NULL,
created_at timestamp NOT NULL DEFAULT NOW(),
description VARCHAR(255) NOT NULL DEFAULT '',
week VARCHAR(15) NOT NULL,
player_id_send VARCHAR(15) NOT NULL,
player_name_send VARCHAR(60) NOT NULL,
player_id_receive VARCHAR(15) NOT NULL,
player_name_receive VARCHAR(60) NOT NULL,
action VARCHAR(30) NOT NULL,
UNIQUE (tx_id),
PRIMARY KEY (id),
FOREIGN KEY (tx_id) REFERENCES tx(id) ON UPDATE CASCADE
);`

/*
 one to many: tx has many user_teams
 one to one:  tx has one tx_detail
*/

const CREATE_TABLE_TX = `CREATE TABLE IF NOT EXISTS tx (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
action_date timestamp DEFAULT '1970-1-1'::TIMESTAMP,
user_team_send_id INT NOT NULL,
user_team_receive_id INT NOT NULL,
user_team_winner_id INT DEFAULT NULL,
user_team_loser_id INT DEFAULT NULL,
amount VARCHAR(15) NOT NULL,
odds VARCHAR(15) NOT NULL,
amount_won INT NOT NULL DEFAULT 0,
status VARCHAR(20) DEFAULT 'waiting',

PRIMARY KEY (id),
FOREIGN KEY (user_team_send_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
FOREIGN KEY (user_team_receive_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
FOREIGN KEY (user_team_winner_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
FOREIGN KEY (user_team_loser_id) REFERENCES user_teams (id) ON UPDATE CASCADE
);`


const CREATE_TABLE_TX_TYPE = `CREATE TABLE IF NOT EXISTS tx_type (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
name timestamp DEFAULT NULL,
PRIMARY KEY (id),
);`

const CLEAR_ALL_TABLES = `DELETE FROM user_teams; DELETE FROM tx; DELETE FROM tx_detail;`


type Tx struct {
	Id                 		int			`db:"id" json:"id"`
	CreatedAt               time.Time	`db:"created_at" json:"created_at"`
	ActionDate				time.Time	`db:"action_date" json:"action_date"`
	UserTeamSendId 			int			`db:"user_team_send_id" json:"user_team_send_id"`
	UserTeamReceiveId 		int			`db:"user_team_receive_id" json:"user_team_receive_id"`
	UserTeamWinnerId 		int			`db:"user_team_winner_id" json:"user_team_winner_id"`
	UserTeamLoserId 		int			`db:"user_team_loser_id" json:"user_team_loser_id"`
	Amount 					int			`db:"amount" json:"amount"`
	Odds 					int			`db:"odds" json:"odds"`
	Status 					string		`db:"status" json:"status"`
}

type TxDeail struct {
	Id									int			`db:"id" json:"id"`
	CreatedAt             				time.Time	`db:"created_at" json:"created_at"`
	TxId 							 	int			`db:"tx_id" json:"tx_id"`
	Description 						string		`db:"description" json:"description"`
	Week            					string		`db:"week" json:"week"`
	PlayerIdSend 		  				string		`db:"player_id_send" json:"player_id_send"`
	PlayerNameSend	 					string		`db:"player_name_send" json:"player_name_send"`
	PlayerIdReceive   					string		`db:"player_id_receive" json:"player_id_receive"`
	PlayerNameReceive 					string		`db:"player_name_receive" json:"player_name_receive"`
	Action                      		string		`db:"action" json:"action"`
}

type CombinedTx struct {
	Id                 		int			`db:"id" json:"id"`
	CreatedAt               time.Time	`db:"created_at" json:"created_at"`
	ActionDate				time.Time	`db:"action_date" json:"action_date"`
	UserTeamSendId 			int			`db:"user_team_send_id" json:"user_team_send_id"`
	UserTeamReceiveId 		int			`db:"user_team_receive_id" json:"user_team_receive_id"`
	UserTeamWinnerId 		int			`db:"user_team_winner_id" json:"user_team_winner_id"`
	UserTeamLoserId 		int			`db:"user_team_loser_id" json:"user_team_loser_id"`
	Amount 					int			`db:"amount" json:"amount"`
	Odds 					int			`db:"odds" json:"odds"`
	Status 					string		`db:"status" json:"status"`
	Description 			string		`db:"description" json:"description"`
	Week            		string		`db:"week" json:"week"`
	PlayerIdSend 		  	string		`db:"player_id_send" json:"player_id_send"`
	PlayerNameSend	 		string		`db:"player_name_send" json:"player_name_send"`
	PlayerIdReceive   		string		`db:"player_id_receive" json:"player_id_receive"`
	PlayerNameReceive 		string		`db:"player_name_receive" json:"player_name_receive"`
	Action                  string		`db:"action" json:"action"`
}
