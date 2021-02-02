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
	Points                             float64                           `json:"points" sql:",notnull"`
	AvgPonits                          float64                           `json:"avgPonits" sql:",notnull,default:0"`
	MaxPoints                          int                               `json:"maxPoints" sql:",notnull,default:0"`
	Version                            int64                             `json:"version"`
	IsActive                           bool                              `json:"isActive"`
	CreatedAt                          time.Time                         `json:"createdAt" sql:",default:now()"`
	UpdatedAt                          time.Time                         `json:"updatedAt" sql:",default:now()"`
}

type AldGoalSettingAssessmentLog struct {
	ID                                 int64                             `json:"id"`
<<<<<<< Updated upstream
	Name                               string                            `json:"name"`
	Points                             float64                           `json:"points"`
	MaxPoints                          int                               `json:"maxPoints"`
=======
	AldGoalSettingAssessmentHeaderID   int64                             `json:"aldGoalSettingAssessmentHeaderID"`
	AldGoalSettingAssessmentHeader     *AldGoalSettingAssessmentHeader   `json:"aldGoalSettingAssessmentHeader" pg:"joinFK:id"`
	UserID                             int64                             `json:"userID"`
	User                               *User                             `json:"user" pg:"joinFK:id"`
>>>>>>> Stashed changes
	AldGoalSettingAssessmentQuestionID int64                             `json:"aldGoalSettingAssessmentQuestionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentQuestion   *AldGoalSettingAssessmentQuestion `json:"aldGoalSettingAssessmentQuestion" pg:"joinFK:id"`
	AldGoalSettingAssessmentOptionID   int64                             `json:"aldGoalSettingAssessmentOptionID" validate:"required" sql:",notnull"`
	AldGoalSettingAssessmentOption     *AldGoalSettingAssessmentOption   `json:"aldGoalSettingAssessmentOption" pg:"joinFK:id"`
	Points                             float64                           `json:"points" sql:",notnull"`
	AvgPonits                          float64                           `json:"avgPonits" sql:",notnull,default:0"`
	TotalPoints                        int                               `json:"totalPoints" sql:",notnull,default:0"`
	Version                            int64                             `json:"version"`
	IsActive                           bool                              `json:"isActive"`
	CreatedAt                          time.Time                         `json:"createdAt" sql:",default:now()"`
	UpdatedAt                          time.Time                         `json:"updatedAt" sql:",default:now()"`
}
