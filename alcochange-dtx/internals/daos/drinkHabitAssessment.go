package daos

import (
	"cyberliver/alcochange-dtx/models"
	"cyberliver/alcochange-dtx/sentryaccounts"

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
	GetDrinkProfile() ([]models.DrinkCategory, error)
	GetUserAccess(id int64) (models.UserAccess, error)
	GetCountry(timeZone string) (models.Country, error)
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

// GetDrinkProfile get the questions from Database
func (da *DrinkHabitAssessment) GetDrinkProfile() ([]models.DrinkCategory, error) {
	drinkProfileIns := []models.DrinkCategory{}

	err := da.dbConn.Model(&drinkProfileIns).Select()
	if err != nil {
		da.l.Error("DrinkHabitAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return drinkProfileIns, nil

}

// GetUserAccess get the questions from Database
func (da *DrinkHabitAssessment) GetUserAccess(id int64) (models.UserAccess, error) {
	userAccess := models.UserAccess{}
	err := da.dbConn.Model(&userAccess).Where("user_id = ?", id).Select()
	if err != nil {
		da.l.Error("GetUserAccess Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return userAccess, nil

}

// GetCountry get the questions from Database
func (da *DrinkHabitAssessment) GetCountry(timeZone string) (models.Country, error) {
	country := models.Country{}
	err := da.dbConn.Model(&country).Where("time_zone = ?", timeZone).Select()
	if err != nil {
		da.l.Error("GetCountry Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return country, nil

}
