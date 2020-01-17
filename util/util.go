package util

import (
	"strconv"

	"github.com/alif/models"
	"github.com/patrickmn/go-cache"
)

var (
	quoteCache *cache.Cache
)

func init() {
	quoteCache = cache.New(cache.NoExpiration, cache.NoExpiration)

	// add to cache some fake date for start
	SaveQuote(1, models.Quote{
		ID:        1,
		Category:  "Science",
		Author:    "Akmal Shaimardonov",
		QuoteText: "Hello World, Alif!!!"})
	SaveQuote(2, models.Quote{
		ID:        2,
		Category:  "Sport",
		Author:    "Jahongir Anvarov",
		QuoteText: "Highway to hell",
	})

}

// SaveQuote сохраняет цитату в кэш
func SaveQuote(id int, value models.Quote) {
	quoteCache.Set(strconv.Itoa(id), value, cache.NoExpiration)
}

//GetAllWuotes получает все цитаты
func GetAllWuotes() []models.Quote {
	elements := quoteCache.Items()

	val := models.Quote{}
	quotes := []models.Quote{}
	for _, element := range elements {
		val = element.Object.(models.Quote)
		quotes = append(quotes, val)
	}
	return quotes
}

// GetQuote получает цитату по id
func GetQuote(id int) (interface{}, bool) {
	return quoteCache.Get(strconv.Itoa(id))
}

// GetQuoteCount возвращает количество цитат
func GetQuoteCount() int {
	return quoteCache.ItemCount()
}

// DeleteQuote удаляет данные по id
func DeleteQuote(id int) {
	quoteCache.Delete(strconv.Itoa(id))
}
