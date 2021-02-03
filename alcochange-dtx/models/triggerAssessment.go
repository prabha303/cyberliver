package models

import "time"

type AldTriggerAssessmentQuestion struct {
	ID                   int64               `json:"id"`
	Question             string              `json:"question"`
	QuestionNo           int                 `json:"questionNo"`
	QuestionOptionTypeID int64               `json:"questionOptionTypeID" validate:"required" sql:",notnull"`
	QuestionOptionType   *QuestionOptionType `json:"questionOptionType" pg:"joinFK:id"`
	SequenceOrder        int                 `json:"sequenceOrder"`
	Version              int64               `json:"version"`
	IsActive             bool                `json:"isActive"`
	CreatedAt            time.Time           `json:"createdAt" sql:",default:now()"`
	UpdatedAt            time.Time           `json:"updatedAt" sql:",default:now()"`
}

type AldTriggerAssessmentOption struct {
	ID                             int64                         `json:"id"`
	Name                           string                        `json:"name"`
	Points                         float64                       `json:"points" sql:",notnull,default:0.0"`
	AldTriggerAssessmentQuestionID int64                         `json:"aldTriggerAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldTriggerAssessmentQuestion   *AldTriggerAssessmentQuestion `json:"aldTriggerAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                  int                           `json:"sequenceOrder"`
	Version                        int64                         `json:"version"`
	IsActive                       bool                          `json:"isActive"`
	CreatedAt                      time.Time                     `json:"createdAt" sql:",default:now()"`
	UpdatedAt                      time.Time                     `json:"updatedAt" sql:",default:now()"`
}
