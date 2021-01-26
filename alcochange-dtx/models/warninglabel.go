package models

import "time"

//Country provides entities for an Country structure
type WarningLabel struct {
	ID               int64     `json:"id"`
	EuRepresentative string    `json:"euRepresentative"`
	RefVersion       string    `json:"refVersion"`
	Logo             string    `json:"logo"`
	WarningLink      string    `json:"warningLink"`
	IndicationsLink  string    `json:"indicationsLink"`
	ManufacturerDate time.Time `json:"manufacturerDate" sql:",default:now()"`
	Version          int64     `json:"version"`
	IsActive         bool      `json:"isActive"`
	CreatedAt        time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt        time.Time `json:"updatedAt" sql:",default:now()"`
}
