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

	drinkProfileIns := []dtos.DrinkProfiles{}

	drinkProfileResponse, dPErr := da.drinkHabitAssessmentDao.GetDrinkProfile()
	if err != nil {
		da.l.Error("GetDrinkHabitAssessmentMessage Error - ", dPErr)
		sentryaccounts.SentryLogExceptions(dPErr)
		return nil, dPErr
	}

	for _, drinkPro := range drinkProfileResponse {
		drinkProfile := dtos.DrinkProfiles{}
		drinkProfile.ID = drinkPro.ID
		drinkProfile.DrinkID = drinkPro.DrinkID
		drinkProfile.Name = drinkPro.Name
		drinkProfile.QuantityText = drinkPro.QuantityUnit.QuantityText
		drinkProfile.QuantityUnitID = drinkPro.QuantityUnitID
		drinkProfile.Strength = drinkPro.Strength
		drinkProfile.Cost = drinkPro.QuantityUnit.Cost

		drinkProfileIns = append(drinkProfileIns, drinkProfile)
	}

	// for testing purpose hardcoded id is 1
	var id int64
	id = 1

	userAccessIns, uaErr := da.drinkHabitAssessmentDao.GetUserAccess(id)
	if err != nil {
		da.l.Error("GetDrinkHabitAssessmentMessage Error - ", uaErr)
		sentryaccounts.SentryLogExceptions(uaErr)
		return nil, uaErr
	}

	countryIns, cErr := da.drinkHabitAssessmentDao.GetCountry(userAccessIns.Timezone)
	if err != nil {
		da.l.Error("GetDrinkHabitAssessmentMessage Error - ", cErr)
		sentryaccounts.SentryLogExceptions(cErr)
		return nil, cErr
	}

	drinkHabitResponse := dtos.DrinkHabitAssessmentResponse{
		CountryID:           countryIns.ID,
		CurrencySymbol:      countryIns.CurrencySymbol,
		CurrenyName:         countryIns.CurrenyName,
		DrinkProfiles:       drinkProfileIns,
		DrinkHabitQuestions: drinkHabit,
	}

	return &drinkHabitResponse, nil
}
