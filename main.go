// Package main is the start of the API, creates all the routes and sets a handlers for each one.
package main

import (
	//"fmt"
	"log"
	"net/http"
	"restedancestor/database"

	//"restedancestor/handlers"

	//"github.com/julienschmidt/httprouter"
)

func main() {
	// 1. Initialize DB from database.go
	database.Database()
	defer DB.Close()

	// 2. Map routes to handlers from handlers.go
	http.HandleFunc("GET /api/quotes", GetQuotesHandler)
	http.HandleFunc("GET /api/quotes/{id}", GetQuoteByIDHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//func main() {
	// initiates router
//	router := httprouter.New()

	// lists routes
//	router.GET("/random", handlers.Random)
//	router.GET("/all", handlers.All)
//	router.GET("/senile", handlers.Senile)
//	router.GET("/search/:word", handlers.Search)
//	router.GET("/top", handlers.Top)
	//commented the error handlers length
	//router.GET("/length/:len", handlers.Length)
	//uuid routes
//	router.GET("/uuid/:uuid/find", handlers.Find)
//	router.POST("/uuid/:uuid/like", handlers.Like)
//	router.POST("/uuid/:uuid/dislike", handlers.Dislike)
//	fmt.Println("Welcome to restedancestor, the API is running in a maddening fashion!")
//	fmt.Println("The Ancestor is waiting and listening on localhost:8080")
//	err := http.ListenAndServe(":8080", router)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
