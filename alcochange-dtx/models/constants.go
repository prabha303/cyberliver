package models

import "time"

type Constants struct {
	ID          int64     `json:"id"`
	Code        string    `json:"code"`
	Text        string    `json:"text"`
	ProductCode string    `json:"productCode"`
	Version     int64     `json:"version"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt   time.Time `json:"updatedAt" sql:",default:now()"`
}
