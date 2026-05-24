// Package main is the start of the API, creates all the routes and sets a handlers for each one.
package main

import (
	"fmt"
	"log"
	"net/http"
	"restedancestor/database"
	"restedancestor/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// 1. Initialize DB from database.go
	database.Database("database/database.db")
	defer database.DB.Close()
	handlers.DB = database.DB
	//Define router to handle all routes

	router := httprouter.New()
	//list of routes
	router.GET("/all", handlers.GetQuotesHandler)
	router.GET("/", handlers.Random)
	router.GET("/quotes/:id", handlers.GetQuoteByIDHandler)

	//Server initialization message
	fmt.Println("Welcome to restedancestor, the API is running in a maddening fashion!")
	fmt.Println("The Ancestor is waiting and listening on localhost:8080")

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
