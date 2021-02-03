package models

import (
	"cyberliver/alcochange-dtx/utils"
	"time"
)

type AldHealthConditionQuestion struct {
	ID                   int64               `json:"id"`
	Question             string              `json:"question"`
	QuestionNo           int                 `json:"questionNo" validate:"required" sql:",notnull,unique=idx_qid_soid_helth_q"`
	QuestionOptionTypeID int64               `json:"questionOptionTypeID" validate:"required" sql:",notnull"`
	QuestionOptionType   *QuestionOptionType `json:"questionOptionType" pg:"joinFK:id"`
	SequenceOrder        int                 `json:"sequenceOrder" validate:"required" sql:",notnull,unique=idx_qid_soid_helth_q"`
	Version              int64               `json:"version"`
	IsActive             bool                `json:"isActive"`
	CreatedAt            time.Time           `json:"createdAt" sql:",default:now()"`
	UpdatedAt            time.Time           `json:"updatedAt" sql:",default:now()"`
}

func (r *AldHealthConditionQuestion) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}

type AldHealthConditionOption struct {
	ID                           int64                       `json:"id"`
	Name                         string                      `json:"name"`
	AldHealthConditionQuestionID int64                       `json:"aldHealthConditionQuestionID" validate:"required" sql:",notnull,unique=idx_qid_sid"`
	AldHealthConditionQuestion   *AldHealthConditionQuestion `json:"aldHealthConditionQuestion" pg:"joinFK:id"`
	SequenceOrder                int                         `json:"sequenceOrder" validate:"required" sql:",notnull,unique=idx_qid_sid"`
	Points                       float64                     `json:"points" sql:",default:0.0"`
	MaxPoints                    int                         `json:"maxPoints"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}

func (r *AldHealthConditionOption) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}
