package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/jmoiron/sqlx"
)


// Structs for database models and internal API
type MasterCreds struct {
	gorm.Model
	AppName		  	string			`gorm:"size:40;unique"`
	ClientId      	string			`gorm:"size:100;unique"`
	ClientSecret  	string			`gorm:"size:50;unique"`
}

// Generic api result
type ApiResult struct {
	Status string				  	`json:"status"`
	Result map[string]interface{} 	`json:"result"`
}

type GenericApiError struct {
	Status string					`json:"status"`
	Error string					`json:"error"`
	Result map[string]interface{}	`json:"result"`
}

type ApiError struct {
	Status string					`json:"status"`
	Error string					`json:"error"`
	Message string					`json:"message"`
}

type YahooApiError struct {
	Status 		string					`json:"status"`
	YahooStatus int						`json:"yahoo_status"`
	Error       string					`json:"error"`
	Result      map[string]interface{}	`json:"result"`
}

type GameIds struct {
	gorm.Model
	Season            int
	Type              string
	GameId            int
}


const (
	Host     = "localhost"
	Port     = 5432
	Username = "jmc856"
	Password = "trIbe19t"
	Dbname   = "gofant"
	testDbname   = "gofant_test"
)



func OpenDataBase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		Host, Port, Username, Password, Dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func OpenPostgresDataBase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, Username, Password, Dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func OpenTestPostgresDataBase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, Username, Password, testDbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
