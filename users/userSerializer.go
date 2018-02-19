package users

import (
	"encoding/json"
	"gofant/models"
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
	user_profile_result := map[string]interface{}{
		"user_profile": up,
		"user_credentials": uc,
	}

	result := models.ApiResult{
		Status: "0",
		Result: user_profile_result,
	}
	return json.Marshal(result)
}

func getFullUserSerializer(u User) ([]byte, error) {

	db := models.OpenPostgresDataBase()
	defer db.Close()

	uc, _ := getUserCredentialsFromUser(db, u)
	up, _ := getUserProfileFromUser(db, u)
	user_result := map[string]interface{}{
		"id": u.ID,
		"created_at": u.CreatedAt,
		"modified_at": u.UpdatedAt,
		"username": u.Username,
		"user_profile": up,
		"user_credentials": uc,
	}

	result := models.ApiResult{
		Status: "0",
		Result: user_result,
	}

	return json.Marshal(result)
}