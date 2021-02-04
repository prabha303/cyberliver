package models

import (
	"time"
)

type AldAuditAssessmentHeader struct {
	ID                           int64                       `json:"id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldAuditAssessmentQuestionID int64                       `json:"aldAuditAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldAuditAssessmentQuestion   *AldAuditAssessmentQuestion `json:"aldAuditAssessmentQuestion" pg:"joinFK:id"`
	AldAuditAssessmentOptionID   int64                       `json:"aldAuditAssessmentOptionID" validate:"required" sql:",notnull"`
	AldAuditAssessmentOption     *AldAuditAssessmentOption   `json:"aldAuditAssessmentOption" pg:"joinFK:id"`
	Points                       float64                     `json:"points" sql:",notnull,default:0"`
	AvgPoints                    float64                     `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldAuditAssessmentHeader) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *AldAuditAssessmentHeader) BeforeUpdate() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

type AldAuditAssessmentLog struct {
	ID                           int64                       `json:"id"`
	AldAuditAssessmentHeaderID   int64                       `json:"aldAuditAssessmentHeaderID" validate:"required" sql:",notnull"`
	AldAuditAssessmentHeader     *AldAuditAssessmentHeader   `json:"aldAuditAssessmentHeader" pg:"joinFK:id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldAuditAssessmentQuestionID int64                       `json:"aldAuditAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldAuditAssessmentQuestion   *AldAuditAssessmentQuestion `json:"aldAuditAssessmentQuestion" pg:"joinFK:id"`
	AldAuditAssessmentOptionID   int64                       `json:"aldAuditAssessmentOptionID" validate:"required" sql:",notnull"`
	AldAuditAssessmentOption     *AldAuditAssessmentOption   `json:"aldAuditAssessmentOption" pg:"joinFK:id"`
	Points                       float64                     `json:"points" sql:",notnull,default:0"`
	AvgPoints                    float64                     `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldAuditAssessmentLog) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
