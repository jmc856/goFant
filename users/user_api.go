package users

import (
	"gofant/api"
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"golang.org/x/oauth2"
	"crypto/rand"
	"github.com/jmoiron/sqlx"
)

func generateRandom() string {
	n := 15
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}


func CreateUser(env *api.Env, params map[string]string) ([]byte, error) {

	// encrypt password
	password_encrypted, err := encryptPassword(params["Password"])
	if err != nil {
		return nil, err
	}
	state := generateRandom()

	// store user
	new_user, user_error := insertUser(env.DB, params["Username"], string(password_encrypted), state)
	if user_error != nil {
		return nil, user_error
	}
	return UserSerializer(new_user)
}

func UpdateUser(env *api.Env, params map[string]string) ([]byte, error) {
	u, err := GetUserFromPassword(env.DB, params["username"], params["password"])

	if err != nil {
		return nil, err
		//not_found_user := User{Username: u.NewUsername}
		//return notFoundUserErrorSerializer(not_found_user)
	}
	// Save new values
	new_user, err := updateUser(env.DB, u, params["new_username"], params["new_password"], params["new_email"])

	return UserSerializer(new_user)

}
func CreateUserProfileAndCredentials(db *sqlx.DB,user User, user_contents []byte, token *oauth2.Token) ([]byte, error) {
	var user_result UserResult
	json.Unmarshal(user_contents, &user_result)

	up := UserProfile{
		Firstname: user_result.Profile.Firstname,
		Lastname: user_result.Profile.Lastname,
		Guid: user_result.Profile.Guid,
		Nickname: user_result.Profile.Nickname,
	}
	creds := UserCredential{
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiration: token.Expiry,
		Type: token.TokenType,
	}

	up_new, up_error := insertUserProfile(db, user, up)
	if up_error != nil {
		return nil, up_error
	}

	uc_new, uc_error := insertUserCredentials(db, user, creds)
	if uc_error != nil {
		return nil, uc_error
	}

	return getUserProfileAndCredentialsSerializer(up_new, uc_new)
}

//Returns a users information according to guid data
func GetUserProfileFromYahoo(guid string, access_token string) (*http.Response, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://social.yahooapis.com/v1/user/%s/profile?format=json", guid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + access_token)

	response, err := client.Do(req)
	return response, err
}

func LoginUsername(db *sqlx.DB,params map[string]string) ([]byte, error) {
	u, err := GetUserFromPassword(db, params["Username"], params["Password"])

	if err != nil {
		return nil, err
	}

	user_credentials, err := getUserCredentialsFromUser(db, u)

	// Check access token and refresh if expired
	if user_credentials.AccessToken != "" {
		token_fresh := user_credentials.checkToken()
		if !token_fresh {
			_, refresh_error := user_credentials.refreshToken(db)
			if refresh_error != nil {
				return nil, refresh_error
			}
		}
	}

	return getFullUserSerializer(u)
}
