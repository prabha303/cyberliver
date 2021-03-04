package dtos

// DrinkHabitAssessmentResponse Response struct send to client
type DrinkHabitAssessmentResponse struct {
	CountryID           int64                 `json:"countryId"`
	CurrencySymbol      string                `json:"currencySymbol"`
	CurrenyName         string                `json:"currenyName"`
	DrinkProfiles       []DrinkProfiles       `json:"drinkProfiles"`
	DrinkHabitQuestions []DrinkHabitQuestions `json:"drinkHabitQuestions"`
}

type DrinkProfiles struct {
	ID             int64  `json:"id"`
	DrinkID        int    `json:"drinkId"`
	Name           string `json:"name"`
	QuantityUnitID int64  `json:"quantityUnitId"`
	QuantityText   string `json:"quantityText"`
	Strength       string `json:"strength"`
	Cost           int    `json:"cost"`
}

type DrinkHabitQuestions struct {
	ID                   int64                        `json:"id"`
	QuestionOptionTypeID int64                        `json:"questionOptionTypeId"`
	Question             string                       `json:"question"`
	QuestionNo           int                          `json:"questionNo"`
	SequenceOrder        int                          `json:"sequenceOrder"`
	Options              []DrinkHabitAssessmentOption `json:"options"`
}

type DrinkHabitAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	MaxPoints     int     `json:"maxPoints"`
	SequenceOrder int     `json:"sequenceOrder"`
}
