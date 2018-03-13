package transactions

import (
	"gofant/users"
	"gofant/api"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func CreateTransaction(db *sqlx.DB, params map[string]string) ([]byte, error) {
	var userTeamSendId int
	var userTeamReceiveId int
	db.QueryRowx(selectUserTeamByTeamKeyAndId, params["team_key_send"], params["user_id"]).Scan(&userTeamSendId)

	var userReceive users.User
	db.QueryRowx(selectUserByGuid, params["user_receive_guid"]).StructScan(&userReceive)
	db.QueryRowx(selectUserTeamByTeamKeyAndId, params["team_key_receive"], userReceive.ID).Scan(&userTeamReceiveId)

	var tx Tx
	err := db.QueryRowx(createTx, userTeamSendId, userTeamReceiveId, params["amount"], params["odds"]).StructScan(&tx)
	if err != nil {
		return nil, api.ApiError{
			Status: api.UnknownError,
			Message: "Error creating transactions",
		}
	}

	var txd TxDeail
	errTxDetail := db.QueryRowx(createTxDetail, tx.Id, params["description"], params["week"],
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
	txId, _ := strconv.Atoi(params["txId"])

	rows, _ := db.Queryx(getTx, txId)
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

	dbErr2 := db.Get(&detail, getTxDetail, txId)
	if dbErr2 != nil {
		return nil, dbErr2
	}

	return TxDetailSerializer(tx, detail)
}


func ListTransactions(db *sqlx.DB, params map[string]string) ([]byte, error) {
	var txList []CombinedTx
	userId, _ := strconv.Atoi(params["user_id"])

	dbErr := db.Select(&txList, getTxForUser, userId)
	if dbErr != nil {
		return nil, dbErr
	}
	// Check and update status of transactions

	return ListTxSerializer(txList)
}


func EditTransaction(db *sqlx.DB, params map[string]string) ([]byte, error) {
	var userTeamReceiveId int
	txId, _ := strconv.Atoi(params["txId"])

	rows, _ := db.Queryx(getUserTeamReceiveId, txId)
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
	userId, _ := strconv.Atoi(params["user_id"])
	db.Select(&userTeamIds, getUserTeamsFromUserId, userId)

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
	db.QueryRowx(acceptTx, txId).StructScan(&tx)

	return TxSerializer(tx)
}
