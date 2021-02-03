package models

import (
	"cyberliver/alcochange-dtx/utils"
	"time"
)

type AldAuditAssessmentQuestion struct {
	ID                   int64               `json:"id"`
	Question             string              `json:"question"`
	QuestionNo           int                 `json:"questionNo" validate:"required" sql:",notnull,unique=idx_qid_asm"`
	QuestionOptionTypeID int64               `json:"questionOptionTypeID" validate:"required" sql:",notnull"`
	QuestionOptionType   *QuestionOptionType `json:"questionOptionType" pg:"joinFK:id"`
	SequenceOrder        int                 `json:"sequenceOrder" validate:"required" sql:",notnull,unique=idx_qid_asm"`
	Version              int64               `json:"version"`
	IsActive             bool                `json:"isActive"`
	CreatedAt            time.Time           `json:"createdAt" sql:",default:now()"`
	UpdatedAt            time.Time           `json:"updatedAt" sql:",default:now()"`
}

func (r *AldAuditAssessmentQuestion) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}

type AldAuditAssessmentOption struct {
	ID                           int64                       `json:"id"`
	Name                         string                      `json:"name"`
	AldAuditAssessmentQuestionID int64                       `json:"aldAuditAssessmentQuestionID" validate:"required" sql:",notnull,unique=idx_qid_sid_asm"`
	AldAuditAssessmentQuestion   *AldAuditAssessmentQuestion `json:"aldAuditAssessmentQuestion" pg:"joinFK:id"`
	SequenceOrder                int                         `json:"sequenceOrder" validate:"required" sql:",notnull,unique=idx_qid_sid_asm"`
	Points                       float64                     `json:"points" sql:",default:0"`
	MaxPoints                    int                         `json:"maxPoints"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldAuditAssessmentOption) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}
