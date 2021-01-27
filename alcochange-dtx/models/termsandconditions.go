package models

import (
	"time"
)

// AlcoChangeTermsAndPrivacy provide terms and conditions to display
type AlcoChangeTermsAndPrivacy struct {
	ID           int64     `json:"id"`
	VersionInfo  string    `json:"versionInfo"`
	Instructions string    `json:"instructions"`
	Contents     string    `json:"contents"`
	Logo         string    `json:"logo"`
	Version      int64     `json:"version"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt    time.Time `json:"updatedAt" sql:",default:now()"`
}
