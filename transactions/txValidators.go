package transactions

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"net/http"
	"bytes"
	"io"
	"fmt"
)

func validateApiCreds(r io.Reader) (map[string]string, error) {
	var creds ApiCreds
	decodeErr := json.NewDecoder(r).Decode(&creds)
	if decodeErr != nil {
		return nil, decodeErr
	}
	if validationErr := validator.Validate(creds); validationErr != nil {
		return nil, validationErr
	}
	return map[string]string{
		"access_token": creds.AccessToken,
		"username":     creds.Username,
		"password":     creds.Password,
	}, nil
}

type ApiCreds struct {
	AccessToken 		string			`validate:"nonzero"`
	Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
	Password 			string			`validate:"min=8"`
}

func ValidateCreateTransaction(r *http.Request) (map[string]string, error) {
	type TransactionRequest struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"   json:"password"`
		AccessToken			string			`validate:"nonzero" json:"access_token"`
		LeagueKey  			string			`validate:"nonzero" json:"league_key"`
		TeamKeySender		string       	`validate:"nonzero" json:"team_key_send"`
		TeamKeyRecipient	string    		`validate:"nonzero" json:"team_key_receive"`
		GuidRecipient       string          `validate:"nonzero" json:"user_receive_guid"`
		PlayerIdSend 		string     		`validate:"nonzero" json:"player_id_send"`
		PlayerIdReceive 	string  		`validate:"nonzero" json:"player_id_receive"`
		PlayerNameSend 		string     		`validate:"nonzero" json:"player_name_send"`
		PlayerNameReceive 	string  		`validate:"nonzero" json:"player_name_receive"`
		Week 				string  		`validate:"nonzero" json:"week"`
		Amount 				string  		`validate:"nonzero" json:"amount"`
		Odds 				string  		`validate:"nonzero" json:"odds"`
		Action				string  		`validate:"nonzero" json:"action"`
	}

	var req TransactionRequest

	decodeErr := json.NewDecoder(r.Body).Decode(&req)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		return nil, decodeErr
	}

	if validationErr := validator.Validate(req); validationErr != nil {
		fmt.Println(validationErr)
		return nil, validationErr
	}
	params := map[string]string{
		"username":             req.Username,
		"password":				req.Password,
		"access_token": 		req.AccessToken,
		"league_key":   		req.LeagueKey,
		"team_key_send":    	req.TeamKeySender,
		"team_key_receive":  	req.TeamKeyRecipient,
		"user_receive_guid":	req.GuidRecipient,
		"player_id_send":	  	req.PlayerIdSend,
		"player_name_send":	  	req.PlayerNameSend,
		"player_id_receive": 	req.PlayerIdReceive,
		"player_name_receive":	req.PlayerNameReceive,
		"week":             	req.Week,
		"amount":             	req.Amount,
		"odds":              	req.Odds,
		"action":            	req.Action,

	}
	return params, nil
}

func ValidateGetTransaction(r *http.Request) (map[string]string, error) {
	type TxGetRequest struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"   json:"password"`
		AccessToken			string			`validate:"nonzero" json:"access_token"`
		TransactionId  		string			`validate:"nonzero" json:"transaction_id"`
	}

	var req TxGetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(req); errs != nil {
		return nil, errs
	}
	params := map[string]string{
		"access_token": req.AccessToken,
		"username":     req.Username,
		"password":     req.Password,
		"transaction_id": req.TransactionId,
	}
	return params, nil
}

func ValidateAcceptTransaction(r *http.Request) (map[string]string, error) {
	type TransactionRequest struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"   json:"password"`
		AccessToken			string			`validate:"nonzero" json:"access_token"`
		TransactionId  		string				`validate:"nonzero" json:"transaction_id"`
	}

	var req TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(req); errs != nil {
		return nil, errs
	}
	params := map[string]string{
		"access_token": req.AccessToken,
		"username":     req.Username,
		"password":     req.Password,
		"transaction_id": req.TransactionId,
	}
	return params, nil
}


func ValidateListTransaction(r *http.Request) (map[string]string, error) {
	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)
	params, err := validateApiCreds(reader)
	if err != nil {
		return nil, err
	}
	return params, nil
}
