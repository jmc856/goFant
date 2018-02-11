package main

import (
	"fmt"
	"net/http"
	"gofant/users"
	"gofant/leagues"
	"gofant/transactions"
	"gofant/handlers"
	"gofant/models"
	"gofant/rosters"
)


func main() {
	// Run Migrations for database Models
	db := models.OpenPostgresDataBase()
	leagues.MigrateLeagues(db)
	users.MigrateUsers(db)
	transactions.MigrateTransactions(db)
	rosters.MigrateRostersAndStats(db)

	fmt.Println(http.ListenAndServe(":8001", handlers.Handlers(db)))
}
