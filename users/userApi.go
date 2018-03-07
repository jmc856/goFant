package users

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"golang.org/x/oauth2"
	"crypto/rand"
	"github.com/jmoiron/sqlx"
	"gofant/authorization"
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


func GetUser(db *sqlx.DB, params map[string]string) ([]byte, error) {
	fmt.Println(params)
	return nil, nil
}

func CreateUser(db *sqlx.DB, params map[string]string) ([]byte, error) {

	// encrypt password
	password, err := encryptPassword(params["Password"])
	if err != nil {
		return nil, err
	}
	state := generateRandom()

	// store user
	userNew, userErr := insertUser(db, params["Username"], string(password), state)
	if userErr != nil {
		return nil, userErr
	}
	return UserSerializer(userNew)
}

func UpdateUser(db *sqlx.DB, params map[string]string) ([]byte, error) {

	u, err := GetUserFromPassword(db, params["username"], params["password"])

	if err != nil {
		return nil, err
		//not_found_user := User{Username: u.NewUsername}
		//return notFoundUserErrorSerializer(not_found_user)
	}
	// Save new values
	userNew, err := updateUser(db, u, params["new_username"], params["new_password"], params["new_email"])

	return UserSerializer(userNew)

}
func CreateUserProfileAndCredentials(db *sqlx.DB,user User, userContents []byte, token *oauth2.Token) ([]byte, error) {
	var ur UserResult
	json.Unmarshal(userContents, &ur)

	up := UserProfile{
		Firstname: ur.Profile.Firstname,
		Lastname: ur.Profile.Lastname,
		Guid: ur.Profile.Guid,
		Nickname: ur.Profile.Nickname,
	}
	creds := UserCredential{
		AccessToken: authorization.CreateJWT(),
		YahooAccessToken: token.AccessToken,
		YahooRefreshToken: token.RefreshToken,
		Expiration: token.Expiry,
		Type: token.TokenType,
	}

	upNew, dbErr := insertUserProfile(db, user, up)
	if dbErr != nil {
		return nil, dbErr
	}

	ucNew, dbErr := insertUserCredentials(db, user, creds)
	if dbErr != nil {
		return nil, dbErr
	}

	return getUserProfileAndCredentialsSerializer(upNew, ucNew)
}

//Returns a users information according to guid data
func GetUserProfileFromYahoo(guid string, accessToken string) (*http.Response, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://social.yahooapis.com/v1/user/%s/profile?format=json", guid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + accessToken)

	response, err := client.Do(req)
	return response, err
}

func Login(db *sqlx.DB, params map[string]string) ([]byte, error) {
	u, err := GetUserFromPassword(db, params["Username"], params["Password"])

	if err != nil {
		return nil, err
	}

	uc, err := getUserCredentials(db, u)

	// Check access token and refresh if expired
	if uc.AccessToken != "" {
		tokenFresh := uc.checkToken()
		if !tokenFresh {
			_, refreshErr := uc.refreshToken(db)
			if refreshErr != nil {
				return nil, refreshErr
			}
		}
	}

	return getFullUserSerializer(u)
}
