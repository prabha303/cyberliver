package models

import (
	"time"
)

type AldDrinkProfileHeader struct {
	ID             int64         `json:"id"`
	DrinkID        int64         `json:"drinkId"`
	UserID         int64         `json:"userID"`
	User           *User         `json:"user" pg:"joinFK:id"`
	Name           string        `json:"name"`
	DrinkCount     int           `json:"drinkCount"`
	Quantity       int           `json:"quantity"`
	QuantityUnitID int64         `json:"quantityUnitId" validate:"required" sql:",notnull"`
	QuantityUnit   *QuantityUnit `json:"quantityUnit" pg:"joinFK:id"`
	CountryID      int64         `json:"countryId" validate:"required" sql:",notnull"`
	Country        *Country      `json:"country" pg:"joinFK:id"`
	Cost           int           `json:"cost"`
	Calories       int           `json:"calories"`
	Version        int64         `json:"version"`
	IsActive       bool          `json:"isActive"`
	CreatedAt      time.Time     `json:"createdAt" sql:",default:now()"`
	UpdatedAt      time.Time     `json:"updatedAt" sql:",default:now()"`
}

func (r *AldDrinkProfileHeader) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *AldDrinkProfileHeader) BeforeUpdate() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

type AldDrinkProfileLog struct {
	ID                      int64                  `json:"id"`
	DrinkID                 int64                  `json:"drinkId"`
	Name                    string                 `json:"name"`
	DrinkCount              int                    `json:"drinkCount"`
	UserID                  int64                  `json:"userID"`
	User                    *User                  `json:"user" pg:"joinFK:id"`
	Quantity                int                    `json:"quantity"`
	QuantityUnitID          int64                  `json:"quantityUnitId" validate:"required" sql:",notnull"`
	QuantityUnit            *QuantityUnit          `json:"quantityUnit" pg:"joinFK:id"`
	CountryID               int64                  `json:"countryId" validate:"required" sql:",notnull"`
	Country                 *Country               `json:"country" pg:"joinFK:id"`
	AldDrinkProfileHeaderID int64                  `json:"aldDrinkProfileHeaderId" validate:"required" sql:",notnull"`
	AldDrinkProfileHeader   *AldDrinkProfileHeader `json:"aldDrinkProfileHeader" pg:"joinFK:id"`
	Cost                    int                    `json:"cost"`
	Calories                int                    `json:"calories"`
	Version                 int64                  `json:"version"`
	IsActive                bool                   `json:"isActive"`
	CreatedAt               time.Time              `json:"createdAt" sql:",default:now()"`
	UpdatedAt               time.Time              `json:"updatedAt" sql:",default:now()"`
}

func (r *AldDrinkProfileLog) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

type AldDrinkHabitAssessmentHeader struct {
	ID                                int64                            `json:"id"`
	UserID                            int64                            `json:"userId"`
	User                              *User                            `json:"user" pg:"joinFK:id"`
	AldDrinkHabitAssessmentQuestionID int64                            `json:"aldDrinkHabitAssessmentQuestionId" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentQuestion   *AldDrinkHabitAssessmentQuestion `json:"aldDrinkHabitAssessmentQuestion" pg:"joinFK:id"`
	AldDrinkHabitAssessmentOptionID   int64                            `json:"aldDrinkHabitAssessmentOptionId" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentOption     *AldDrinkHabitAssessmentOption   `json:"aldDrinkHabitAssessmentOption" pg:"joinFK:id"`
	Points                            float64                          `json:"points" sql:",notnull,default:0"`
	AvgPoints                         float64                          `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                         int                              `json:"maxPoints" sql:",notnull,default:0"`
	Version                           int64                            `json:"version"`
	IsActive                          bool                             `json:"isActive"`
	CreatedAt                         time.Time                        `json:"createdAt" sql:",default:now()"`
	UpdatedAt                         time.Time                        `json:"updatedAt" sql:",default:now()"`
}

func (r *AldDrinkHabitAssessmentHeader) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *AldDrinkHabitAssessmentHeader) BeforeUpdate() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

type AldDrinkHabitAssessmentLog struct {
	ID                                int64                            `json:"id"`
	AldDrinkHabitAssessmentHeaderID   int64                            `json:"aldDrinkHabitAssessmentHeaderId" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentHeader     *AldDrinkHabitAssessmentHeader   `json:"aldDrinkHabitAssessmentHeader" pg:"joinFK:id"`
	UserID                            int64                            `json:"userID"`
	User                              *User                            `json:"user" pg:"joinFK:id"`
	AldDrinkHabitAssessmentQuestionID int64                            `json:"aldDrinkHabitAssessmentQuestionId" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentQuestion   *AldDrinkHabitAssessmentQuestion `json:"aldDrinkHabitAssessmentQuestion" pg:"joinFK:id"`
	AldDrinkHabitAssessmentOptionID   int64                            `json:"aldDrinkHabitAssessmentOptionId" validate:"required" sql:",notnull"`
	AldDrinkHabitAssessmentOption     *AldDrinkHabitAssessmentOption   `json:"aldDrinkHabitAssessmentOption" pg:"joinFK:id"`
	Points                            float64                          `json:"points" sql:",notnull,default:0"`
	AvgPoints                         float64                          `json:"AvgPoints" sql:",notnull,default:0"`
	MaxPoints                         int                              `json:"maxPoints" sql:",notnull,default:0"`
	Version                           int64                            `json:"version"`
	IsActive                          bool                             `json:"isActive"`
	CreatedAt                         time.Time                        `json:"createdAt" sql:",default:now()"`
	UpdatedAt                         time.Time                        `json:"updatedAt" sql:",default:now()"`
}

func (r *AldDrinkHabitAssessmentLog) BeforeInsert() {
	// currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
