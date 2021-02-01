package dtos

// ReasonAssessmentResponse Response struct send to client
type ReasonAssessmentResponse struct {
	ID                   int64                    `json:"id"`
	QuestionOptionTypeID int64                    `json:"questionOptionTypeId"`
	Question             string                   `json:"question"`
	QuestionNo           int                      `json:"questionNo"`
	SequenceOrder        int                      `json:"sequenceOrder"`
	HeaderNote           string                   `json:"headerNote"`
	Options              []ReasonAssessmentOption `json:"options"`
}

type ReasonAssessmentOption struct {
	ID            int64  `json:"id"`
	QuestionID    int64  `json:"questionId"`
	Name          string `json:"name"`
	Points        int    `json:"points"`
	SequenceOrder int    `json:"sequenceOrder"`
}
