package models

// GetQuotesResponse is
type GetQuotesResponse struct {
	Quotes []Quote `json:"quotes"`
}

// Quote -
type Quote struct {
	ID       int32    `json:"id"`
	Category Category `json:"category"`
	Author   Author   `json:"author"`
	Title    string   `json:"title"`
}

// Author is
type Author struct {
	ID   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Category is
type Category struct {
	ID    int32  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// AddQuoteRequest is
type AddQuoteRequest struct {
	Title      string `json:"title"`
	AuthorID   uint32 `json:"author_id"`
	CategoryID uint32 `json:"category_id"`
}

//AddQuoteResponse is
type AddQuoteResponse struct {
	Quote Quote `json:"quote"`
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
