package models

import "time"

type AldReasonAssessmentQuestion struct {
	ID                   int64               `json:"id"`
	Question             string              `json:"question"`
	QuestionNo           int                 `json:"questionNo"`
	QuestionOptionTypeID int64               `json:"questionOptionTypeID" validate:"required" sql:",notnull"`
	QuestionOptionType   *QuestionOptionType `json:"questionOptionType" pg:"joinFK:id"`
	SequenceOrder        int                 `json:"sequenceOrder"`
	HeaderNote           string              `json:"headerNote"`
	Version              int64               `json:"version"`
	IsActive             bool                `json:"isActive"`
	CreatedAt            time.Time           `json:"createdAt" sql:",default:now()"`
	UpdatedAt            time.Time           `json:"updatedAt" sql:",default:now()"`
}

type AldReasonAssessmentOption struct {
	ID                            int64                        `json:"id"`
	Name                          string                       `json:"name"`
	Points                        float64                      `json:"points"`
	MaxPoints                     int                          `json:"maxPoints"`
	AldReasonAssessmentQuestionID int64                        `json:"aldReasonAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldReasonAssessmentQuestion   *AldReasonAssessmentQuestion `json:"aldReasonAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                 int                          `json:"sequenceOrder"`
	Version                       int64                        `json:"version"`
	IsActive                      bool                         `json:"isActive"`
	CreatedAt                     time.Time                    `json:"createdAt" sql:",default:now()"`
	UpdatedAt                     time.Time                    `json:"updatedAt" sql:",default:now()"`
}
