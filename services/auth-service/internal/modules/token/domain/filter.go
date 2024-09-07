package domain

import "github.com/golangid/candi/candishared"

// FilterToken model
type FilterToken struct {
	candishared.Filter
	ID        *int     `json:"id"`
	UserID    string   `json:"user_id"`
	DeviceID  string   `json:"device_id"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}
