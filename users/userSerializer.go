package users

import (
	"encoding/json"
	"gofant/models"
	"gofant/authorization"
)

func UserSerializer(u User) ([]byte, error) {
	user_result := map[string]interface{}{
		"id": u.ID,
		"created_at": u.CreatedAt,
		"modified_at": u.UpdatedAt,
		"username": u.Username,
		"password": u.Password,
	}

	result := models.ApiResult{
		Status: "0",
		Result: user_result,
	}

	return json.Marshal(result)
}

func getUserProfileAndCredentialsSerializer(up UserProfile, uc UserCredential) ([]byte, error) {
	upResult := map[string]interface{}{
		"user_profile": up,
		"user_credentials": uc,
		"access_token": authorization.CreateJWT(),
	}

	result := models.ApiResult{
		Status: "0",
		Result: upResult,
	}
	return json.Marshal(result)
}

func getFullUserSerializer(u User) ([]byte, error) {

	db := models.OpenPostgresDataBase()
	defer db.Close()

	uc, _ := getUserCredentialsFromUser(db, u)
	up, _ := getUserProfileFromUser(db, u)
	userResult := map[string]interface{}{
		"id": u.ID,
		"created_at": u.CreatedAt,
		"modified_at": u.UpdatedAt,
		"username": u.Username,
		"user_profile": up,
		"user_credentials": uc,
	}

	result := models.ApiResult{
		Status: "0",
		Result: userResult,
	}

	return json.Marshal(result)
}