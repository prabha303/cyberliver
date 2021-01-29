package dtos

// GoalSettingAssessmentResponse Response struct send to client
type GoalSettingAssessmentResponse struct {
	ID              int64                         `json:"id"`
	OptionType      string                        `json:"optionType"`
	OptionTypeLabel string                        `json:"optionTypeLabel"`
	Question        string                        `json:"question"`
	QuestionNo      int                           `json:"questionNo"`
	SequenceOrder   int                           `json:"sequenceOrder"`
	Options         []GoalSettingAssessmentOption `json:"options"`
}

type GoalSettingAssessmentOption struct {
	ID            int64  `json:"id"`
	QuestionID    int64  `json:"questionId"`
	Name          string `json:"name"`
	Points        int    `json:"points"`
	SequenceOrder int    `json:"sequenceOrder"`
}
