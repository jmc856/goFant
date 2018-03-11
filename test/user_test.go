package test

import (
	"io/ioutil"
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"fmt"
	"net/http"
	"gofant/api"
	"gofant/models"
	"gofant/users"
	"strconv"
	"time"
)

const insertUser = `INSERT INTO users
				  (username, password, state)
				   VALUES ($1, $2, $3)
				   RETURNING id, created_at, username, password, state`

const createUserCredentials = `INSERT INTO user_credentials
				  (updated_at, user_id, access_token, yahoo_access_token, yahoo_refresh_token, expiration, type)
				   VALUES ($1, $2, $3, $4, $5, $6, $7)
				   RETURNING id, user_id, created_at, updated_at, access_token, yahoo_access_token, yahoo_refresh_token, expiration, type`

func TestGetToken(t *testing.T) {
	request, _ := http.NewRequest("GET", getTokenUrl, nil)
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	token := string(body)
	assert.NotNil(t, token, "Token should not be nil")
}

func createUser(username string) users.User {
	var userNew users.User
	testDb := models.OpenTestPostgresDataBase()
	defer testDb.Close()
	if dbErr := testDb.QueryRowx(insertUser, username, "Mock Password", users.GenerateRandom()).StructScan(&userNew); dbErr != nil {
		fmt.Println("USER CREATED")
		fmt.Println(dbErr)
	}
	return userNew
}

func createUserCredential(userId int, accessToken string) users.UserCredential {
	var ucNew users.UserCredential
	testDb := models.OpenTestPostgresDataBase()
	defer testDb.Close()
	if dbErr := testDb.QueryRowx(createUserCredentials,
		time.Now(), userId, accessToken, "Mock Yahoo Access Token", "Mock Yahoo Refresh Token", time.Now(),
			"Mock Type").StructScan(&ucNew); dbErr != nil {
		fmt.Println(dbErr)
	}
	return ucNew
}

func TestEndpointWithNoToken(t *testing.T) {
	userJson := `{"username": "s", "password": "guzz"}`
	reader = strings.NewReader(userJson)
	request, _ := http.NewRequest("POST", userUrl, reader)
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, string(body), "Required authorization token not found\n", "Returns no token found")
	assert.Equal(t, res.StatusCode, 401, "Should return a 401")
}

func TestCreateUserUserTooShort(t *testing.T) {
	var createUserError api.ApiError
	userJson := `{"username": "s", "password": "guzz"}`
	token := getToken()
	reader = strings.NewReader(userJson)
	request, _ := http.NewRequest("POST", userUrl, reader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &createUserError)
	assert.Equal(t, createUserError.Status, "1", "Status should be 1")
}

func TestCreateUser(t *testing.T) {
	var apiResult models.ApiResult
	userJson := `{"username": "new_user", "password": "test_password"}`
	token := getToken()
	reader = strings.NewReader(userJson)
	request, _ := http.NewRequest("POST", userUrl, reader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &apiResult)
	assert.Equal(t, apiResult.Status, "0", "Api status should be 0")
	assert.Equal(t, res.StatusCode, 201, "Should return 201 for created object")
	assert.Equal(t, apiResult.Result["username"], "new_user")
	assert.NotEqualf(t, apiResult.Result["password"], "test_password", "Password now encrypted")
}

func TestGetUser(t *testing.T) {
	var apiResult models.ApiResult
	username := "Mock Get User"
	u := createUser(username)
	token := getToken()
	userId := strconv.Itoa(u.ID)
	getUserUrl := fmt.Sprintf("%s/%s", userUrl, userId)
	request, _ := http.NewRequest("GET", getUserUrl, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &apiResult)
	assert.Equal(t, apiResult.Status, "0","Api status should be 0")
	assert.Equal(t, res.StatusCode, 200,"Should return 200 for get request")
	assert.Equal(t, apiResult.Result["username"], username)
}

func TestUpdateUser(t *testing.T) {
	var apiResult models.ApiResult
	u := createUser("Mock Update User")
	userJson := `{"new_username": "modified_user"}`
	token := getToken()
	createUserCredential(u.ID, token)
	reader = strings.NewReader(userJson)
	userId := strconv.Itoa(u.ID)
	updateUserUrl := fmt.Sprintf("%s/%s", userUrl, userId)
	request, _ := http.NewRequest("PUT", updateUserUrl, reader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &apiResult)
	assert.Equal(t, apiResult.Status, "0")
	assert.Equal(t, res.StatusCode, 200, "Should return 200 with updated object")
	assert.Equal(t, apiResult.Result["username"], "modified_user")
}
