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
	Points                       float64                     `json:"points" sql:",notnull"`
	AvgPonits                    float64                     `json:"avgPonits" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
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
	Points                       float64                     `json:"points" sql:",notnull"`
	AvgPonits                    float64                     `json:"avgPonits" sql:",notnull,default:0"`
	MaxPoints                    int                         `json:"maxPoints" sql:",notnull,default:0"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}
