package main

import (
	"log"
	"net/http"

	"github.com/alif/controllers"
	"github.com/alif/db"

	"github.com/julienschmidt/httprouter"
)

var (
	router = httprouter.New()
)

func main() {
	db.Connect()
	initRoutes()
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
	defer db.Close()
}

func initRoutes() {
	router.GET("/alif/quotes", controllers.GetQuotesHandler)
	router.GET("/alif/quotes/:id", controllers.GetQuoteByIDHandler)
	router.POST("/alif/quotes", controllers.AddQuoteHandler)
	router.DELETE("/alif/quotes/:id", controllers.DeleteQuoteHandle)
}

// Run is
func Run() {

}
