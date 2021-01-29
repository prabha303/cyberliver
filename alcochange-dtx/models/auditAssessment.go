package models

import "time"

type AldAuditAssessmentQuestion struct {
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

type AldAuditAssessmentOption struct {
	ID                           int64                       `json:"id"`
	Name                         string                      `json:"name"`
	Points                       int                         `json:"points"`
	AldAuditAssessmentQuestionID int64                       `json:"aldAuditAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldAuditAssessmentQuestion   *AldAuditAssessmentQuestion `json:"aldAuditAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                int                         `json:"sequenceOrder"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}
