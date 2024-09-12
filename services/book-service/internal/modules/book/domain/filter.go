package domain

import "github.com/golangid/candi/candishared"

// FilterBook model
type FilterBook struct {
	candishared.Filter
	ID        *int     `json:"id"`
	Title     string   `json:"title"`
	Subject   string   `json:"subject"`
	Publisher string   `json:"publisher"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}

type FilterBookItem struct {
	candishared.Filter
	ID        *int     `json:"id"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}