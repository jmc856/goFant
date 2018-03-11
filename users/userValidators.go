package users


import (
	"gopkg.in/validator.v2"
	"net/http"
	"encoding/json"
	"fmt"
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
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	if err := validator.Validate(request); err != nil {
		return nil, err
	}
	return map[string]string{"Username": request.Username, "Password": request.Password}, nil
}

func ValidateUpdateUser(r *http.Request) (map[string]string, error) {
	type UpdateUserRequest struct {
		NewUsername     string			`validate:"min=8"   json:"new_username"`
		NewPassword     string			`validate:"min=8"   json:"new_password"`
		NewEmail        string			`validate:"min=8"   json:"new_email"`

	}
	var request UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return map[string]string{
		"new_username": request.NewUsername,
		"new_password": request.NewPassword,
		"new_email": request.NewEmail}, nil
}