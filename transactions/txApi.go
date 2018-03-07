package transactions

import (
	"gofant/users"
	"gofant/api"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func CreateTransaction(db *sqlx.DB, params map[string]string) ([]byte, error) {
	user, userErr :=users.GetUserFromPassword(db, params["username"], params["password"])
	if userErr != nil {
		return nil, userErr
	}

	var userTeamSendId int
	var userTeamReceiveId int
	db.QueryRowx(GET_USER_TEAM_BY_TEAM_KEY_AND_USER_ID, params["team_key_send"], user.ID).Scan(&userTeamSendId)

	var userReceive users.User
	db.QueryRowx(GET_USER_BY_GUID, params["user_receive_guid"]).StructScan(&userReceive)
	db.QueryRowx(GET_USER_TEAM_BY_TEAM_KEY_AND_USER_ID, params["team_key_receive"], userReceive.ID).Scan(&userTeamReceiveId)

	var tx Tx
	err := db.QueryRowx(CREATE_TX, userTeamSendId, userTeamReceiveId, params["amount"], params["odds"]).StructScan(&tx)
	if err != nil {
		return nil, api.ApiError{
			Status: api.UnknownError,
			Message: "Error creating transactions",
		}
	}

	var txd TxDeail
	errTxDetail := db.QueryRowx(CREATE_TX_DETAIL, tx.Id, params["description"], params["week"],
		         params["player_id_send"], params["player_name_send"],
				 params["player_id_receive"], params["player_name_receive"],
		         params["action"]).StructScan(&txd)
	if errTxDetail != nil {
		return nil, api.ApiError{
			Status: api.UnknownError,
			Message: "Error creating transactions",
		}
	}
	return TxDetailSerializer(tx, txd)
}


func GetTransaction(db *sqlx.DB, params map[string]string) ([]byte, error) {
	var tx Tx
	var detail TxDeail

	rows, _ := db.Queryx(GET_TX, params["transaction_id"])
	if rows.Next() == false {
		return nil, api.ApiError{
			Status: api.UnknownError,
			Message: "No transaction found",
		}
	} else {
		dbErr := rows.StructScan(&tx)
		if dbErr != nil {
			return nil, dbErr
		}
	}

	dbErr2 :=db.Get(&detail, get_tx_detail, params["transaction_id"])
	if dbErr2 != nil {
		return nil, dbErr2
	}

	return TxDetailSerializer(tx, detail)
}


func ListTransactions(db *sqlx.DB, params map[string]string) ([]byte, error) {
	user, userErr :=users.GetUserFromPassword(db, params["username"], params["password"])
	if userErr != nil {
		return nil, userErr
	}

	var txList []CombinedTx

	dbErr := db.Select(&txList, GET_TX_FOR_USER, user.ID)
	if dbErr != nil {
		return nil, dbErr
	}

	// Check and update status of transactions

	return ListTxSerializer(txList)
}


func AcceptTransaction(db *sqlx.DB, params map[string]string) ([]byte, error) {
	user, userErr :=users.GetUserFromPassword(db, params["username"], params["password"])
	if userErr != nil {
		return nil, userErr
	}

	var userTeamReceiveId int
	txId, _ := strconv.Atoi(params["transaction_id"])

	rows, _ := db.Queryx(GET_USER_TEAM_RECEIVE_ID, txId)
	if rows.Next() == false {
		return nil, api.ApiError{
			Status: api.UnknownError,
			Message: "No transaction found",
		}
	} else {
		dbErr := rows.Scan(&userTeamReceiveId)
		if dbErr != nil {
			return nil, dbErr
		}
	}

	var userTeamIds []int
	var myTx bool
	db.Select(&userTeamIds, GET_USER_TEAMS_FROM_USER_ID, int(user.ID))

	for _, b := range userTeamIds {
		if b == userTeamReceiveId {
			myTx = true
		}
	}
	if !myTx {
		return nil, api.ApiError{
		}
	}

	var tx Tx
	db.QueryRowx(ACCEPT_TX, txId).StructScan(&tx)

	return TxSerializer(tx)
}
