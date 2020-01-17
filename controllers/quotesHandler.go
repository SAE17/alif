package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/alif/util"

	models "github.com/alif/models"
	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetQuotesHandler is
func GetQuotesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var res models.GetQuotesResponse

	res.Quotes = util.GetAllWuotes()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// GetQuotesByCategoryHandler is
func GetQuotesByCategoryHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryName := p.ByName("category_name")

	res := models.GetQuotesResponse{}

	for _, quote := range util.GetAllWuotes() {
		if quote.Category == categoryName {
			res.Quotes = append(res.Quotes, quote)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

//DeleteQuotesHandler is
func DeleteQuotesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	quoteID, _ := strconv.Atoi(p.ByName("id"))
	res := models.GetQuotesResponse{}

	util.DeleteQuote(quoteID)

	res.Quotes = util.GetAllWuotes()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

//UpdateQuoteHandler is
func UpdateQuoteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		req *models.UpdateQuoteRequest
	)
	quoteID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Println("Ошибка парсинка id", err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Ошибка декодирования ", err)
		return
	}
	if req.AuthorName == "" {
		w.Write([]byte("Путое поле"))
		return
	}
	if req.CategoryName == "" {
		w.Write([]byte("Путое поле"))
		return
	}
	if req.QuoteText == "" {
		w.Write([]byte("Путое поле"))
		return
	}

	val, found := util.GetQuote(quoteID)
	if !found {
		w.Write([]byte("Not found with whis id"))
		return
	}
	quote := val.(models.Quote)
	quote.Author = req.AuthorName
	quote.Category = req.CategoryName
	quote.QuoteText = req.QuoteText
	util.SaveQuote(quoteID, quote)

	res := &models.GetQuotesResponse{}
	res.Quotes = util.GetAllWuotes()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// GetRandomQuoteHandler is
func GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	randomID := rand.Intn(util.GetQuoteCount())
	util.DeleteQuote(randomID)
	res := &models.GetQuotesResponse{}
	res.Quotes = util.GetAllWuotes()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}
