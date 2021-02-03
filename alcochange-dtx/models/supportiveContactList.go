package models

import "time"

type AldRelationShip struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Version   int64     `json:"version"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt time.Time `json:"updatedAt" sql:",default:now()"`
}
