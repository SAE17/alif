package models

// GetQuotesResponse is
type GetQuotesResponse struct {
	Quotes []Quote `json:"quotes"`
}

// Quote -
type Quote struct {
	ID        int32  `json:"id"`
	Category  string `json:"category"`
	Author    string `json:"author"`
	QuoteText string `json:"quoteText"`
}

// UpdateQuoteRequest is
type UpdateQuoteRequest struct {
	QuoteText    string `json:"quoteText"`
	AuthorName   string `json:"authorName"`
	CategoryName string `json:"categoryName"`
}

// UpdateQuoteResponse is
type UpdateQuoteResponse struct {
	Quote Quote `json:"quotes"`
}
