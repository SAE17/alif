package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	models "github.com/alif/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	q *models.GetQuotesResponse
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.Unmarshal(b, &q); err != nil {
		log.Println(err)
	}
	log.Println("-------- * ------- Starting Logging -------- * -------")
}

//GetQuotesHandler is
func GetQuotesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var res *models.GetQuotesResponse
	res = q

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// GetQuotesByCategoryHandler is
func GetQuotesByCategoryHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryName := p.ByName("category_name")

	res := models.GetQuotesResponse{}

	for _, quote := range q.Quotes {
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
	index := 0
	for _, quote := range q.Quotes {
		if quote.ID == int32(quoteID) {
			q.Quotes = append(q.Quotes[:index], q.Quotes[index+1:]...)
		}
		index++
	}
}

//UpdateQuoteHandler is
func UpdateQuoteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		req *models.UpdateQuoteRequest
	)
	quoteID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Println("После парсинга id", err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	if err != nil {
		log.Println("После декодирования ", err)
		return
	}
	res := &models.GetQuotesResponse{}
	for _, quote := range q.Quotes {
		if quote.ID == int32(quoteID) {
			if req.AuthorName != "" {
				quote.Author = req.AuthorName
			}
			if req.CategoryName != "" {
				quote.Category = req.CategoryName
			}
			if req.QuoteText != "" {
				quote.QuoteText = req.QuoteText
			}
		}
		res.Quotes = append(res.Quotes, quote)
	}
	q = res
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(q)
}

// GetRandomQuoteHandler is
func GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	randomID := rand.Intn(len(q.Quotes))
	res := q.Quotes[randomID]
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}


