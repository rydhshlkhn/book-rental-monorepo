package domain

import (
	shareddomain "monorepo/services/book-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseBookList model
type ResponseBookList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseBook   `json:"data"`
}

type ResponseBookItem struct {
	ID              int       `json:"id"`
	Barcode         string    `json:"barcode"`
	IsReferenceOnly bool      `json:"is_reference_only"`
	Borrowed        time.Time `json:"borrowed"`
	DueDate         time.Time `json:"due_date"`
	FormatID        int       `json:"format_id"`
	StatusID        int       `json:"status_id"`
	DateOfPurchase  time.Time `json:"date_of_purchase"`
	PublicationDate time.Time `json:"publication_date"`
}

// ResponseBook model
type ResponseBook struct {
	ID           int                `json:"id"`
	ISBN         string             `json:"isbn"`
	Title        string             `json:"title"`
	Subject      string             `json:"subject"`
	Publisher    string             `json:"publisher"`
	Language     string             `json:"language"`
	NumberOfPage int                `json:"numberOfPage"`
	CreatedAt    string             `json:"createdAt"`
	UpdatedAt    string             `json:"updatedAt"`
	BookItems    []ResponseBookItem `json:"book_items"`
}

// Serialize from db model
func (r *ResponseBook) Serialize(source *shareddomain.Book) {
	var bookItems []ResponseBookItem
	for _, bi := range source.BookItems {
		bookItems = append(bookItems, ResponseBookItem{
			ID:              bi.ID,
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

	r.ID = source.ID
	r.Title = source.Title
	r.ISBN = source.ISBN
	r.Subject = source.Subject
	r.Publisher = source.Publisher
	r.Language = source.Language
	r.NumberOfPage = source.NumberOfPage
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
	r.BookItems = bookItems
}
