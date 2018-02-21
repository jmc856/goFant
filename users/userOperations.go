package users

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"net/url"
	"bytes"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"gofant/api"
	"github.com/jmoiron/sqlx"
)

func encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Comparing given password string with database password hash
func ValidatePassword(hash []byte, password string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return false
	}
	return true
}

func (uc UserCredential) checkToken() bool {
	if uc.Expiration.Before(time.Now()) {
		return false
	}
	return true
}

func (uc UserCredential) refreshToken(db *sqlx.DB) (UserCredential, error) {

	apiUrl := "https://api.login.yahoo.com/oauth2/get_token"
	data := url.Values{}
	data.Set("client_id", os.Getenv("GOFANT_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("GOFANT_CLIENT_SECRET"))
	data.Add("grant_type", "refresh_token")
	data.Add("redirect_uri", "http://127.0.0.1/YahooCallback")
	data.Add("refresh_token", uc.RefreshToken)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(r)

	//defer response.Body.Close()
	if response.StatusCode != 200 {
		return UserCredential{}, api.ApiError{Status:"1", Message:"Could not refresh access token"}
	}

	var ucResult CredentialResult
	contents, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(contents, &ucResult)
	ucNew := UserCredential{
		AccessToken: ucResult.AccessToken,
		RefreshToken: ucResult.RefreshToken,
		Expiration: time.Now().Add(time.Hour * time.Duration(ucResult.ExpiresIn/3600)),
		Type: ucResult.TokenType,
	}

	if dbErr := db.QueryRowx(updateUserCredentials, time.Now(), ucNew.AccessToken, ucNew.RefreshToken, ucNew.Expiration, uc.ID).StructScan(&uc); dbErr != nil {
		return ucNew, api.ApiError{
			Status:"1",
			ErrorString: dbErr.Error(),
			Message:"Error saving user credentials",
		}
	}
	return ucNew, nil
}