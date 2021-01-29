package models

import "time"

type AldReasonAssessmentQuestion struct {
	ID              int64     `json:"id"`
	Question        string    `json:"question"`
	QuestionNo      int       `json:"questionNo"`
	OptionType      string    `json:"optionType"`
	OptionTypeLabel string    `json:"optionTypeLabel"`
	SequenceOrder   int       `json:"sequenceOrder"`
	HeaderNote      string    `json:"headerNote"`
	Version         int64     `json:"version"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt       time.Time `json:"updatedAt" sql:",default:now()"`
}

type AldReasonAssessmentOption struct {
	ID                            int64                        `json:"id"`
	Name                          string                       `json:"name"`
	Points                        int                          `json:"points"`
	AldReasonAssessmentQuestionID int64                        `json:"aldReasonAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldReasonAssessmentQuestion   *AldReasonAssessmentQuestion `json:"aldReasonAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                 int                          `json:"sequenceOrder"`
	Version                       int64                        `json:"version"`
	IsActive                      bool                         `json:"isActive"`
	CreatedAt                     time.Time                    `json:"createdAt" sql:",default:now()"`
	UpdatedAt                     time.Time                    `json:"updatedAt" sql:",default:now()"`
}
