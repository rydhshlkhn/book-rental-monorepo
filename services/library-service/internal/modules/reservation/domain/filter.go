package domain

import "github.com/golangid/candi/candishared"

// FilterReservation model
type FilterReservation struct {
	candishared.Filter
	ID         *int     `json:"id"`
	BookItemID int      `json:"book_item_id"`
	UserID     int      `json:"user_id"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Preloads   []string `json:"-"`
}
