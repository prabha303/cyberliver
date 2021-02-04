package models

import (
	"time"
)

type AldGoalSettingAssessmentHeader struct {
	ID                                 int64                             `json:"id"`
	UserID                             int64                             `json:"userID"`
	User                               *User                             `json:"user" pg:"joinFK:id"`
	AldGoalSettingAssessmentQuestionID int64                             `json:"aldGoalSettingAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentQuestion   *AldGoalSettingAssessmentQuestion `json:"aldGoalSettingAssessmentQuestion" pg:"joinFK:id"`
	AldGoalSettingAssessmentOptionID   int64                             `json:"aldGoalSettingAssessmentOptionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentOption     *AldGoalSettingAssessmentOption   `json:"aldGoalSettingAssessmentOption" pg:"joinFK:id"`
	Points                             float64                           `json:"points" sql:",notnull,default:0"`
	AvgPoints                          float64                           `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                          int                               `json:"maxPoints" sql:",notnull,default:0"`
	Version                            int64                             `json:"version"`
	IsActive                           bool                              `json:"isActive"`
	CreatedAt                          time.Time                         `json:"createdAt" sql:",default:now()"`
	UpdatedAt                          time.Time                         `json:"updatedAt" sql:",default:now()"`
}

func (r *AldGoalSettingAssessmentHeader) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *AldGoalSettingAssessmentHeader) BeforeUpdate() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

type AldGoalSettingAssessmentLog struct {
	ID                                 int64                             `json:"id"`
	AldGoalSettingAssessmentHeaderID   int64                             `json:"aldGoalSettingAssessmentHeaderID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentHeader     *AldGoalSettingAssessmentHeader   `json:"aldGoalSettingAssessmentHeader" pg:"joinFK:id"`
	UserID                             int64                             `json:"userID"`
	User                               *User                             `json:"user" pg:"joinFK:id"`
	AldGoalSettingAssessmentQuestionID int64                             `json:"aldGoalSettingAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentQuestion   *AldGoalSettingAssessmentQuestion `json:"aldGoalSettingAssessmentQuestion" pg:"joinFK:id"`
	AldGoalSettingAssessmentOptionID   int64                             `json:"aldGoalSettingAssessmentOptionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentOption     *AldGoalSettingAssessmentOption   `json:"aldGoalSettingAssessmentOption" pg:"joinFK:id"`
	Points                             float64                           `json:"points" sql:",notnull,default:0"`
	AvgPoints                          float64                           `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                          int                               `json:"maxPoints" sql:",notnull,default:0"`
	Version                            int64                             `json:"version"`
	IsActive                           bool                              `json:"isActive"`
	CreatedAt                          time.Time                         `json:"createdAt" sql:",default:now()"`
	UpdatedAt                          time.Time                         `json:"updatedAt" sql:",default:now()"`
}

func (r *AldGoalSettingAssessmentLog) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
