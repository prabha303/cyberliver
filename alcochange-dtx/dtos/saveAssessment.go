package dtos

import "time"

type SaveAssessmentRequest struct {
	HealthConditionAssessmentAnswer HealthConditionAssessmentAnswer `json:"healthConditionAssessmentAnswer"`
	AuditAssessmentAnswer           AuditAssessmentAnswer           `json:"auditAssessmentAnswer"`
}

type HealthConditionAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type AuditAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type UserAnswer struct {
	QuestionID int `json:"questionId"`
	OptionID   int `json:"optionId"`
	Points     int `json:"points"`
}
