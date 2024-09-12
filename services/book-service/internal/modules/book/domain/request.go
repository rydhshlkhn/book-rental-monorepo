package domain

import (
	shareddomain "monorepo/services/book-service/pkg/shared/domain"
	"time"
)

// BookItem struct
type RequestBookItem struct {
	ID              int       `json:"id"`
	Barcode         string    `json:"barcode"`
	IsReferenceOnly bool      `json:"is_reference_only"`
	Borrowed        time.Time `json:"borrowed"`
	DueDate         time.Time `json:"due_date"`
	FormatID        int       `json:"format_id"`
	StatusID        int       `json:"status_id"`
	DateOfPurchase  time.Time `json:"date_of_purchase"`
	PublicationDate time.Time `json:"publication_date"`
	BookID          int       `json:"book_id"`
}

// Deserialize to db model
func (r *RequestBookItem) Deserialize() (res shareddomain.BookItem) {
	res.Barcode = r.Barcode
	res.IsReferenceOnly = r.IsReferenceOnly
	res.Borrowed = r.Borrowed
	res.DueDate = r.DueDate
	res.FormatID = r.FormatID
	res.StatusID = r.StatusID
	res.DateOfPurchase = r.DateOfPurchase
	res.PublicationDate = r.PublicationDate
	res.BookID = r.BookID
	return
}

// Book struct
type RequestBook struct {
	ID           int               `json:"id"`
	ISBN         string            `json:"isbn"`
	Title        string            `json:"title"`
	Subject      string            `json:"subject"`
	Publisher    string            `json:"publisher"`
	Language     string            `json:"language"`
	NumberOfPage int               `json:"numberOfPage"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	BookItems    []RequestBookItem `json:"book_items"`
}

// Deserialize to db model
func (r *RequestBook) Deserialize() (res shareddomain.Book) {
	var bookItems []shareddomain.BookItem
	for _, bi := range r.BookItems {
		bookItems = append(bookItems, shareddomain.BookItem{
			Barcode:         bi.Barcode,
			IsReferenceOnly: bi.IsReferenceOnly,
			Borrowed:        bi.Borrowed,
			DueDate:         bi.DueDate,
			FormatID:        bi.FormatID,
			StatusID:        bi.StatusID,
			DateOfPurchase:  bi.DateOfPurchase,
			PublicationDate: bi.PublicationDate,
		})
	}

	res.Title = r.Title
	res.ISBN = r.ISBN
	res.Subject = r.Subject
	res.Publisher = r.Publisher
	res.Language = r.Language
	res.NumberOfPage = r.NumberOfPage
	res.CreatedAt = r.CreatedAt
	res.UpdatedAt = r.UpdatedAt
	res.BookItems = bookItems
	return
}
