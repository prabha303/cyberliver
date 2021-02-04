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
	Points                       float64                     `json:"points" sql:",notnull,default:0"`
	AvgPoints                    float64                     `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldHealthAssessmentHeader) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *AldHealthAssessmentHeader) BeforeUpdate() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

type AldHealthAssessmentLog struct {
	ID                           int64                       `json:"id"`
	AldHealthAssessmentHeaderID  int64                       `json:"aldHealthAssessmentHeaderID" validate:"required" sql:",notnull"`
	AldHealthAssessmentHeader    *AldHealthAssessmentHeader  `json:"aldHealthAssessmentHeader" pg:"joinFK:id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldHealthConditionQuestionID int64                       `json:"aldHealthConditionQuestionID" validate:"required" sql:",notnull"`
	AldHealthConditionQuestion   *AldHealthConditionQuestion `json:"aldHealthConditionQuestion" pg:"joinFK:id"`
	AldHealthConditionOptionID   int64                       `json:"aldHealthConditionOptionID" validate:"required" sql:",notnull"`
	AldHealthConditionOption     *AldHealthConditionOption   `json:"aldHealthConditionOption" pg:"joinFK:id"`
	Points                       float64                     `json:"points" sql:",notnull,default:0"`
	AvgPoints                    float64                     `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldHealthAssessmentLog) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
