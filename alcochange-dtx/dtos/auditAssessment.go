package dtos

// AuditAssessmentResponse Response struct send to client
type AuditAssessmentResponse struct {
	ID              int64                   `json:"id"`
	OptionType      string                  `json:"optionType"`
	OptionTypeLabel string                  `json:"optionTypeLabel"`
	Question        string                  `json:"question"`
	QuestionNo      int                     `json:"questionNo"`
	SequenceOrder   int                     `json:"sequenceOrder"`
	Options         []AuditAssessmentOption `json:"options"`
}

type AuditAssessmentOption struct {
	ID            int64  `json:"id"`
	QuestionID    int64  `json:"questionId"`
	Name          string `json:"name"`
	Points        int    `json:"points"`
	SequenceOrder int    `json:"sequenceOrder"`
}
