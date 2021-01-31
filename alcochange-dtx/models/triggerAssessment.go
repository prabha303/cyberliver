package models

import "time"

type AldTriggerAssessmentQuestion struct {
	ID              int64     `json:"id"`
	Question        string    `json:"question"`
	QuestionNo      int       `json:"questionNo"`
	OptionType      string    `json:"optionType"`
	OptionTypeLabel string    `json:"optionTypeLabel"`
	SequenceOrder   int       `json:"sequenceOrder"`
	Version         int64     `json:"version"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt       time.Time `json:"updatedAt" sql:",default:now()"`
}

type AldTriggerAssessmentOption struct {
	ID                             int64                         `json:"id"`
	Name                           string                        `json:"name"`
	Points                         int                           `json:"points"`
	AldTriggerAssessmentQuestionID int64                         `json:"aldTriggerAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldTriggerAssessmentQuestion   *AldTriggerAssessmentQuestion `json:"aldTriggerAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                  int                           `json:"sequenceOrder"`
	Version                        int64                         `json:"version"`
	IsActive                       bool                          `json:"isActive"`
	CreatedAt                      time.Time                     `json:"createdAt" sql:",default:now()"`
	UpdatedAt                      time.Time                     `json:"updatedAt" sql:",default:now()"`
}
