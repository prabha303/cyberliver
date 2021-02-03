package models

import "time"

type AldCopingStrategyAssessmentQuestion struct {
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

type AldCopingStrategyAssessmentOption struct {
	ID                                    int64                                `json:"id"`
	Name                                  string                               `json:"name"`
	Points                                float64                              `json:"points" sql:",default:0.0"`
	MaxPoints                             int                                  `json:"maxPoints"`
	AldCopingStrategyAssessmentQuestionID int64                                `json:"aldCopingStrategyAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldCopingStrategyAssessmentQuestion   *AldCopingStrategyAssessmentQuestion `json:"aldCopingStrategyAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                         int                                  `json:"sequenceOrder"`
	Version                               int64                                `json:"version"`
	IsActive                              bool                                 `json:"isActive"`
	CreatedAt                             time.Time                            `json:"createdAt" sql:",default:now()"`
	UpdatedAt                             time.Time                            `json:"updatedAt" sql:",default:now()"`
}
