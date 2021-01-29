package models

import "time"

type AldCopingStrategyAssessmentQuestion struct {
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

type AldCopingStrategyAssessmentOption struct {
	ID                                    int64                                `json:"id"`
	Name                                  string                               `json:"name"`
	Points                                int                                  `json:"points"`
	AldCopingStrategyAssessmentQuestionID int64                                `json:"aldCopingStrategyAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldCopingStrategyAssessmentQuestion   *AldCopingStrategyAssessmentQuestion `json:"aldCopingStrategyAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                         int                                  `json:"sequenceOrder"`
	Version                               int64                                `json:"version"`
	IsActive                              bool                                 `json:"isActive"`
	CreatedAt                             time.Time                            `json:"createdAt" sql:",default:now()"`
	UpdatedAt                             time.Time                            `json:"updatedAt" sql:",default:now()"`
}
