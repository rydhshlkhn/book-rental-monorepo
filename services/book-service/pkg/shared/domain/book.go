package domain

import (
	"time"
)

type Book struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	ISBN         string     `gorm:"column:isbn;type:varchar(13)" json:"isbn"`
	Title        string     `gorm:"column:title;type:varchar(255)" json:"title"`
	Subject      string     `gorm:"column:subject;type:varchar(255)" json:"subject"`
	Publisher    string     `gorm:"column:publisher;type:varchar(255)" json:"publisher"`
	Language     string     `gorm:"column:language;type:varchar(100)" json:"language"`
	NumberOfPage int        `gorm:"column:number_of_page" json:"number_of_page"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	BookItems    []BookItem `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book_items"`
}

func (Book) TableName() string {
	return "books"
}

type BookItem struct {
	ID              int       `gorm:"column:id;primary_key" json:"id"`
	Barcode         string    `gorm:"column:barcode;type:varchar(255)" json:"barcode"`
	IsReferenceOnly bool      `gorm:"column:is_reference_only" json:"is_reference_only"`
	Borrowed        time.Time `gorm:"column:borrowed" json:"borrowed"`
	DueDate         time.Time `gorm:"column:due_date" json:"due_date"`
	FormatID        int       `gorm:"column:format_id" json:"format_id"`
	StatusID        int       `gorm:"column:status_id" json:"status_id"`
	DateOfPurchase  time.Time `gorm:"column:date_of_purchase" json:"date_of_purchase"`
	PublicationDate time.Time `gorm:"column:publication_date" json:"publication_date"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	BookID          int `gorm:"column:book_id" json:"book_id"`
}

func (BookItem) TableName() string {
	return "book_items"
}

type BookFormat struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BookFormat) TableName() string {
	return "book_formats"
}

type BookStatus struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BookStatus) TableName() string {
	return "book_statuses"
}
