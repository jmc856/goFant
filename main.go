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
	port := "8080"
	fmt.Println("API server runing on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", port), handlers.Handlers(db)))
}
