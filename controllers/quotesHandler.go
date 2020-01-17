package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	if err := json.Unmarshal(b, &q); err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("-------- * ------- Starting Logging -------- * -------")
}

//GetQuotesHandler is
func GetQuotesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	res := models.GetQuotesResponse{}
	for _, quote := range q.Quotes {
		res.Quotes = append(res.Quotes, quote)
	}

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
		log.Println("После кодирования ", err)
		return
	}
	res := &models.GetQuotesResponse{}
	for _, quote := range q.Quotes {
		if quote.ID == int32(quoteID) {
			quote.Author = req.AuthorName
		}
		res.Quotes = append(res.Quotes, quote)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	q = res
	json.NewEncoder(w).Encode(q)
}
