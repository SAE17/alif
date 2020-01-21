package db

import (
	"log"

	"github.com/alif/models"
)

// GetAllQuotes получает все цитаты
func GetAllQuotes() (quotes []*models.Quote, err error) {
	query := ` SELECT q.id, q.title, c.title, a.name FROM "Quotes" q
			INNER JOIN "Categories" c ON q.category_id = c.id
			INNER JOIN "Authors" a ON a.id = q.author_id `
	query += " ORDER BY q.id "
	rows, err := pgPool.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		quote := models.Quote{
			Category: models.Category{},
			Author:   models.Author{},
		}
		err = rows.Scan(
			&quote.ID,
			&quote.Title,
			&quote.Category.Title,
			&quote.Author.Name,
		)
		if err != nil {
			log.Println(err)
			return
		}
		quotes = append(quotes, &quote)
	}
	return quotes, nil
}

// GetQuoteByID получет Цитату по его ID
func GetQuoteByID(id uint32) (*models.Quote, error) {

	quote := models.Quote{
		Author:   models.Author{},
		Category: models.Category{},
	}
	query := ` SELECT q.id, q.title, c.title, a.name FROM "Quotes" q
		INNER JOIN "Categories" c ON q.category_id = c.id
		INNER JOIN "Authors" a ON a.id = q.author_id
		WHERE q.id = $1 `

	err := pgPool.QueryRow(query, id).Scan(
		&quote.ID,
		&quote.Title,
		&quote.Category.Title,
		&quote.Author.Name,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &quote, nil
}

//DeleteQuote is
func DeleteQuote(id uint32) ([]*models.Quote, error) {
	query := ` DELETE FROM "Quotes" q WHERE q.id = $1`
	_, err := pgPool.Exec(query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return GetAllQuotes()
}

// AddQuote добавляет новую запись в БД
func AddQuote(quote *models.AddQuoteRequest) (*models.Quote, error) {
	// addQuote, _ := pgPool.Begin()
	query := `INSERT INTO "Quotes" (title, category_id, author_id) VALUES 
				($1, $2, $3) `
	_, err := pgPool.Exec(query,
		quote.Title,
		quote.AuthorID,
		quote.CategoryID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	lastID := lastQuoteID()
	quoteFromDB, err := GetQuoteByID(lastID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return quoteFromDB, nil
}

//Возвращает ID последней записи в БД
func lastQuoteID() uint32 {
	quotes, _ := GetAllQuotes()
	var quantity uint32 = 0
	for _, quote := range quotes {
		quantity = uint32(quote.ID)
	}
	return quantity
}
