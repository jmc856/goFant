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
)

func TestGetToken(t *testing.T) {
	request, _ := http.NewRequest("GET", getTokenUrl, nil)
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	token := string(body)
	assert.NotNil(t, token, "Token should not be nil")
}

func TestCreateUserUserTooShort(t *testing.T) {
	var createUserError api.ApiError
	userJson := `{"username": "s", "password": "guzz"}`
	token := getToken()
	reader = strings.NewReader(userJson)
	request, _ := http.NewRequest("POST", createUserUrl, reader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &createUserError)
	assert.Equal(t, createUserError.Status, "1", "Status should be 1")
}

func TestCreateUser(t *testing.T) {
	var apiResult models.ApiResult
	var apiResultDup models.ApiError
	userJson := `{"username": "new_user", "password": "test_password"}`
	token := getToken()
	reader = strings.NewReader(userJson)
	request, _ := http.NewRequest("POST", createUserUrl, reader)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &apiResult)
	assert.Equal(t, apiResult.Status, "0", "Status should be 0")
	assert.Equal(t, apiResult.Result["username"], "new_user")
	assert.NotEqualf(t, apiResult.Result["password"], "test_password", "Password now encrypted")

	t.Skip("skipping test; Duplicate username error code not finished")
	request_dup, _ := http.NewRequest("POST", createUserUrl, reader)
	res_dup, _ := http.DefaultClient.Do(request_dup)
	body_dup, _ := ioutil.ReadAll(res_dup.Body)
	json.Unmarshal(body_dup, &apiResultDup)
	assert.Equal(t, apiResultDup.Status, "1", "Status should be 0")
	assert.Equal(t, apiResultDup.Message, "Duplicate Username", "Status should be 0")
}
