package models

import (
	"time"
)

type AldHealthAssessmentHeader struct {
	ID                           int64                       `json:"id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldHealthConditionQuestionID int64                       `json:"aldHealthConditionQuestionID" validate:"required" sql:",notnull"`
	AldHealthConditionQuestion   *AldHealthConditionQuestion `json:"aldHealthConditionQuestion" pg:"joinFK:id"`
	AldHealthConditionOptionID   int64                       `json:"aldHealthConditionOptionID" validate:"required" sql:",notnull"`
	AldHealthConditionOption     *AldHealthConditionOption   `json:"aldHealthConditionOption" pg:"joinFK:id"`
	Points                       float64                     `json:"points" sql:",notnull"`
	AvgPonits                    float64                     `json:"avgPonits" sql:",notnull,default:0"`
	TotalPoints                  int                         `json:"totalPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

type AldHealthAssessmentLog struct {
	ID                           int64                       `json:"id"`
	AldHealthAssessmentHeaderID  int64                       `json:"aldHealthAssessmentHeaderID"`
	AldHealthAssessmentHeader    *AldHealthAssessmentHeader  `json:"aldHealthAssessmentHeader" pg:"joinFK:id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldHealthConditionQuestionID int64                       `json:"aldHealthConditionQuestionID" validate:"required" sql:",notnull"`
	AldHealthConditionQuestion   *AldHealthConditionQuestion `json:"aldHealthConditionQuestion" pg:"joinFK:id"`
	AldHealthConditionOptionID   int64                       `json:"aldHealthConditionOptionID" validate:"required" sql:",notnull"`
	AldHealthConditionOption     *AldHealthConditionOption   `json:"aldHealthConditionOption" pg:"joinFK:id"`
	Points                       float64                     `json:"points" sql:",notnull"`
	AvgPonits                    float64                     `json:"avgPonits" sql:",notnull,default:0"`
	TotalPoints                  int                         `json:"totalPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}
