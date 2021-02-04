package drinkHabitAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type DrinkHabitAssessment struct {
	dbConn                  *pg.DB
	l                       *log.Logger
	drinkHabitAssessmentDao daos.DrinkHabitAssessmentDao
}

func NewDrinkHabitAssessment(l *log.Logger, dbConn *pg.DB) *DrinkHabitAssessment {
	return &DrinkHabitAssessment{
		l:                       l,
		dbConn:                  dbConn,
		drinkHabitAssessmentDao: daos.NewDrinkHabitAssessmentDB(l, dbConn),
	}
}

// GetDrinkHabitAssessmentMessage service for logic
func (da *DrinkHabitAssessment) GetDrinkHabitAssessmentMessage() (*dtos.DrinkHabitAssessmentResponse, error) {
	drinkHabit := []dtos.DrinkHabitQuestions{}
	options := dtos.DrinkHabitAssessmentOption{}

	drinkHabitQuestionResponse, err := da.drinkHabitAssessmentDao.DrinkHabitAssessmentQuestion()
	if err != nil {
		da.l.Error("GetDrinkHabitAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, drinkHabitQuestion := range drinkHabitQuestionResponse {
		drinkHabitIns := dtos.DrinkHabitQuestions{}
		drinkHabitIns.ID = drinkHabitQuestion.ID
		drinkHabitIns.Question = drinkHabitQuestion.Question
		drinkHabitIns.QuestionNo = drinkHabitQuestion.QuestionNo
		drinkHabitIns.QuestionOptionTypeID = drinkHabitQuestion.QuestionOptionTypeID
		drinkHabitIns.SequenceOrder = drinkHabitQuestion.SequenceOrder

		drinkHabitOptionResponse, err := da.drinkHabitAssessmentDao.DrinkHabitAssessmentOption(drinkHabitIns.ID)
		if err != nil {
			da.l.Error("GetDrinkHabitAssessmentMessage Error - ", err)
			sentryaccounts.SentryLogExceptions(err)
			return nil, err
		}

		for _, drinkHabitOption := range drinkHabitOptionResponse {
			options.ID = drinkHabitOption.ID
			options.Name = drinkHabitOption.Name
			options.Points = drinkHabitOption.Points
			options.MaxPoints = drinkHabitOption.MaxPoints
			options.QuestionID = drinkHabitOption.AldDrinkHabitAssessmentQuestionID
			options.SequenceOrder = drinkHabitOption.SequenceOrder
			drinkHabitIns.Options = append(drinkHabitIns.Options, options)
		}

		drinkHabit = append(drinkHabit, drinkHabitIns)

	}

	asd := []dtos.DrinkProfiles{}

	drinkHabitResponse := dtos.DrinkHabitAssessmentResponse{
		DrinkProfiles:       asd,
		DrinkHabitQuestions: drinkHabit,
	}

	return &drinkHabitResponse, nil
}
