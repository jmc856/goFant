package yahoo

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"net/http"
)

func ValidateYahooLogin(r *http.Request) (map[string]string, error) {
	type YahooLoginRequest struct {
		Username 			string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password 			string			`validate:"min=8"`
	}

	var yahooRequest YahooLoginRequest
	err := json.NewDecoder(r.Body).Decode(&yahooRequest)
	if err != nil {
		return nil, err
	}

	if validationErr := validator.Validate(yahooRequest); validationErr != nil {
		return nil, validationErr
	}
	yahooLoginParams := map[string]string{
		"username":     yahooRequest.Username,
		"password":     yahooRequest.Password,
	}
	return yahooLoginParams, nil
}

func ValidateYahooCallback(r *http.Request) (map[string]string, error) {
	params := make(map[string]string, 0)
	return params, nil
}