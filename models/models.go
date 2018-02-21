package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)


// Structs for database models and internal API

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

const (
	Host     = "db"
	Port     = 5432
	Username = "gofant"
	Password = "trIbe19t"
	Dbname   = "gofant"
	testDbname   = "gofant_test"
)

func OpenPostgresDataBase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		Host, Port, Username, Dbname)
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
