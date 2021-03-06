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


const (
	Host     = "db"
	Port     = 5432
	Username = "gofant"
	Password = "trIbe19t"
	Dbname   = "gofant"

	testHost = "localhost"
	testUsername = "postgres"
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
		testHost, Port, testUsername, Password, testDbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
