package transactions

import (
	"gofant/models"
	"encoding/json"
)

func TxSerializer(tx Tx) ([]byte, error) {
	Result := map[string]interface{}{
		"transaction": tx,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}


func TxDetailSerializer(tx Tx, detail TxDeail) ([]byte, error) {
	Result := map[string]interface{}{
		"transaction": tx,
		"transaction_detail": detail,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}


func ListTxSerializer(txs []CombinedTx) ([]byte, error) {
	txList := make([]interface{}, 0)

	for _, tx := range txs {
		txList = append(txList, tx)
	}

	Result := map[string]interface{}{
		"transactions": txList,
	}

	result := models.ApiResult{
		Status: "0",
		Result: Result,
	}

	return json.Marshal(result)
}
