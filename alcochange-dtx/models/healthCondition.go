package models

import "time"

type AldHealthConditionQuestion struct {
	ID              int64     `json:"id"`
	Question        string    `json:"question"`
	QuestionNo      int       `json:"questionNo"`
	OptionType      string    `json:"optionType"`
	OptionTypeLabel string    `json:"optionTypeLabel"`
	Version         int64     `json:"version"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt       time.Time `json:"updatedAt" sql:",default:now()"`
}

type AldHealthConditionOption struct {
	ID                           int64                       `json:"id"`
	Name                         string                      `json:"name"`
	Points                       int                         `json:"points"`
	AldHealthConditionQuestionID int64                       `json:"aldHealthConditionQuestionID" validate:"required" sql:",notnull"`
	AldHealthConditionQuestion   *AldHealthConditionQuestion `json:"aldHealthConditionQuestion" pg:"joinFK:id"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}
