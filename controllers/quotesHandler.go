package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/alif/db"
	"github.com/alif/models"
	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetQuotesHandler is
func GetQuotesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	res, err := db.GetAllQuotes()
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// GetQuoteByIDHandler is
func GetQuoteByIDHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	quoteID, _ := strconv.Atoi(p.ByName("id"))

	res, err := db.GetQuoteByID(uint32(quoteID))
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// AddQuoteHandler is
func AddQuoteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req *models.AddQuoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
		return
	}
	res, err := db.AddQuote(req)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}
