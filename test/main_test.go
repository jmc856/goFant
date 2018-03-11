package test

import (
	"net/http"
	"net/http/httptest"
	"io"
	"gofant/handlers"
	"fmt"
	"testing"
	"strings"
	"io/ioutil"
	"gofant/models"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	server *httptest.Server
	reader io.Reader
	userUrl string
	getTokenUrl string
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

	sqlFiles := []string{"../database/01_users.sql", "../database/02_leagues.sql", "../database/03_tx.sql", "../database/04_roster_stat.sql"}
	for _, fileName := range sqlFiles {
		file, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Println(err)
		}

		requests := strings.Split(string(file), ";")

		for _, request := range requests {
			_, err := ta.DB.Exec(request)
			if err != nil {
				fmt.Println(err)
			}
		}}
}


func (ta testApp) tearDown() {
	ta.DB.Exec("DROP SCHEMA public CASCADE;")
	ta.DB.Exec("CREATE SCHEMA public;")
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
	userUrl = fmt.Sprintf("%s/users", server.URL)
	getTokenUrl = fmt.Sprintf("%s/get-token", server.URL)
}

func getToken() string {
	request, _ := http.NewRequest("GET", getTokenUrl, nil)
	res, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(res.Body)
	token := string(body)
	return token
}
