package dtos

import "time"

type SaveAssessmentRequest struct {
	HealthConditionAssessmentAnswer HealthConditionAssessmentAnswer `json:"healthConditionAssessmentAnswer"`
	AuditAssessmentAnswer           AuditAssessmentAnswer           `json:"auditAssessmentAnswer"`
	GoalSettingAssessmentAnswer     GoalSettingAssessmentAnswer     `json:"goalSettingAssessmentAnswer"`
}

type HealthConditionAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type AuditAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type GoalSettingAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type UserAnswer struct {
	QuestionID int64   `json:"questionId"`
	OptionID   int64   `json:"optionId"`
	Points     float64 `json:"points"`
	MaxPoints  int     `json:"maxPoints"`
}
