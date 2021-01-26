package models

import "time"

type PatientAccessCode struct {
	ID           int64     `json:"id"`
	AccessCode   string    `json:"accessCode"`
	SolutionType string    `json:"solutionType"`
	IsRedeemed   bool      `json:"isRedeemed"`
	Version      int64     `json:"version"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt    time.Time `json:"updatedAt" sql:",default:now()"`
}
