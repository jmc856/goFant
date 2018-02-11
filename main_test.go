package main_test

import (
	"net/http"
	"net/http/httptest"
	"io"
	"gofant/handlers"
	"fmt"
	"testing"
	"strings"
	"io/ioutil"
	"encoding/json"
	"gofant/api"
	"github.com/stretchr/testify/assert"
	"gofant/models"
	"gofant/leagues"
	"gofant/users"
	"gofant/transactions"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	server *httptest.Server
	reader io.Reader
	createUserUrl string
)

type setUp interface {
	runserver() httptest.Server
	buildSchema()
	tearDown()
}

type testApp struct {
	DB *sqlx.DB
}

func (ta testApp) buildSchema() {
	leagues.MigrateLeagues(ta.DB)
	users.MigrateUsers(ta.DB)
	transactions.MigrateTransactions(ta.DB)
}

func (ta testApp) tearDown() {
	leagues.RemoveLeaguesTables(ta.DB)
	users.RemoveUsersTables(ta.DB)
	transactions.RemoveTransactionsTables(ta.DB)
}

func (ta testApp) runserver() *httptest.Server {
	return httptest.NewServer(handlers.Handlers(ta.DB))
}

// Build tables on test database.  Run all tests.  Clear tables
func TestMain(m *testing.M) {
	testDb := models.OpenTestPostgresDataBase()
	testApp := testApp{DB: testDb}
	testApp.initialize()

	code := m.Run()

	testApp.tearDown()

	os.Exit(code)
}


func (ta testApp) initialize() {
	ta.buildSchema()
	server := ta.runserver()
	createUserUrl = fmt.Sprintf("%s/users/create", server.URL)
}

func TestCreateUserUserTooShort(t *testing.T) {
	var createUserError api.ApiError
	userJson := `{"username": "s", "password": "guzz"}`

	reader = strings.NewReader(userJson)

	request, _ := http.NewRequest("POST", createUserUrl, reader)
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &createUserError)
	assert.Equal(t, createUserError.Status, "1", "Status should be 1")
}

func TestCreateUser(t *testing.T) {
	var apiResult models.ApiResult
	var apiResultDup models.ApiError
	userJson := `{"username": "new_user", "password": "test_password"}`

	reader = strings.NewReader(userJson)

	request, _ := http.NewRequest("POST", createUserUrl, reader)
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
	fmt.Println(string(body_dup))
	assert.Equal(t, apiResultDup.Status, "1", "Status should be 0")
	assert.Equal(t, apiResultDup.Message, "Duplicate Username", "Status should be 0")
}
