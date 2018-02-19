package main

import (
	"fmt"
	"net/http"
	"gofant/handlers"
	"gofant/models"
)


func main() {
	// Run Migrations for database Models
	db := models.OpenPostgresDataBase()
	fmt.Println(http.ListenAndServe(":8080", handlers.Handlers(db)))
}
