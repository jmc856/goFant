package transactions

import (
	"time"
)


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
