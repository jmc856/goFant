package users


import (
	"gopkg.in/validator.v2"
	"net/http"
	"encoding/json"
)

// Validates and parses request parameters of /login
// Returns map of request payload or error in bytes
func ValidateCreateUser(r *http.Request) (map[string]string, error) {
	type CreateUserRequest struct {
		Username		string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password		string			`validate:"min=8"`
		Email           string
	}
	var request CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if errs := validator.Validate(request); errs != nil {
		return nil, errs
	}
	return map[string]string{"Username": request.Username, "Password": request.Password}, nil
}

func ValidateUpdateUser(r *http.Request) (map[string]string, error) {
	type UpdateeUserRequest struct {
		Username		string			`validate:"min=3,max=40,regexp=^[a-zA-Z]"`
		Password		string			`validate:"min=8"`
		NewUsername     string
		NewPassword     string			`validate:"min=8"`
		NewEmail        string

	}
	var request UpdateeUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"username": request.Username,
		"password": request.Password,
		"new_username": request.NewUsername,
		"new_password": request.NewPassword,
		"new_email": request.NewEmail}, nil
}