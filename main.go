package main

import (
	"log"
	"net/http"

	"github.com/alif/controllers"

	"github.com/julienschmidt/httprouter"
)

var (
	router = httprouter.New()
)

func main() {
	initRoutes()
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}

func initRoutes() {
	router.GET("/alif/quotes", controllers.GetQuotesHandler)
	router.GET("/alif/quotes/:category_name", controllers.GetQuotesByCategoryHandler)
	router.GET("/alif/quote/random", controllers.GetRandomQuoteHandler)
	router.PUT("/alif/quotes/:id", controllers.UpdateQuoteHandler)
	router.DELETE("/alif/quotes/:id", controllers.DeleteQuotesHandler)
}
