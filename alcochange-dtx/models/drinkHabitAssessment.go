package models

import "time"

type AldDrinkHabitAssessmentQuestion struct {
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

type AldDrinkHabitAssessmentOption struct {
	ID                                int64                            `json:"id"`
	Name                              string                           `json:"name"`
	Points                            int                              `json:"points"`
	AldDrinkHabitAssessmentQuestionID int64                            `json:"aldDrinkHabitAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentQuestion   *AldDrinkHabitAssessmentQuestion `json:"aldDrinkHabitAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                     int                              `json:"sequenceOrder"`
	Version                           int64                            `json:"version"`
	IsActive                          bool                             `json:"isActive"`
	CreatedAt                         time.Time                        `json:"createdAt" sql:",default:now()"`
	UpdatedAt                         time.Time                        `json:"updatedAt" sql:",default:now()"`
}
