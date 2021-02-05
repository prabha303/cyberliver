package models

import "time"

type DrinkCategory struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	DrinkID        int           `json:"drinkId" sql:",notnull"`
	QuantityUnitID int64         `json:"quantityUnitID" validate:"required" sql:",notnull"`
	QuantityUnit   *QuantityUnit `json:"quantityUnit" pg:"joinFK:id"`
	Strength       string        `json:"strength"`
	Version        int64         `json:"version"`
	IsActive       bool          `json:"isActive"`
	CreatedAt      time.Time     `json:"createdAt" sql:",default:now()"`
	UpdatedAt      time.Time     `json:"updatedAt" sql:",default:now()"`
}
