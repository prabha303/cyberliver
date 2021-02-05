package models

import "time"

type QuantityUnit struct {
	ID           int64     `json:"id"`
	QuantityText string    `json:"quantityText"`
	Cost         int       `json:"cost"`
	Version      int64     `json:"version"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt    time.Time `json:"updatedAt" sql:",default:now()"`
}
