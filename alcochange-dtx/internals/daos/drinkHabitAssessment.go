package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type DrinkHabitAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewDrinkHabitAssessmentDB(l *log.Logger, dbConn *pg.DB) *DrinkHabitAssessment {
	return &DrinkHabitAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// DrinkHabitAssessmentDao interface
type DrinkHabitAssessmentDao interface {
	DrinkHabitAssessmentQuestion() ([]models.AldDrinkHabitAssessmentQuestion, error)
	DrinkHabitAssessmentOption(id int64) ([]models.AldDrinkHabitAssessmentOption, error)
}

// DrinkHabitAssessmentQuestion get the questions from Database
func (da *DrinkHabitAssessment) DrinkHabitAssessmentQuestion() ([]models.AldDrinkHabitAssessmentQuestion, error) {
	drinkHabitIns := []models.AldDrinkHabitAssessmentQuestion{}

	err := da.dbConn.Model(&drinkHabitIns).Select()
	if err != nil {
		da.l.Error("DrinkHabitAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return drinkHabitIns, nil

}

// DrinkHabitAssessmentOption get the options from Database
func (da *DrinkHabitAssessment) DrinkHabitAssessmentOption(id int64) ([]models.AldDrinkHabitAssessmentOption, error) {
	drinkHabitIns := []models.AldDrinkHabitAssessmentOption{}

	err := da.dbConn.Model(&drinkHabitIns).Where("ald_drink_habit_assessment_question_id = '%d'", id).Select()
	if err != nil {
		da.l.Error("DrinkHabitAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return drinkHabitIns, nil

}
