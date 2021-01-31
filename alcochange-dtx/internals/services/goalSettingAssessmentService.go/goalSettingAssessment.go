package goalSettingAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type GoalSettingAssessment struct {
	dbConn                   *pg.DB
	l                        *log.Logger
	goalSettingAssessmentDao daos.GoalSettingAssessmentDao
}

func NewGoalSettingAssessment(l *log.Logger, dbConn *pg.DB) *GoalSettingAssessment {
	return &GoalSettingAssessment{
		l:                        l,
		dbConn:                   dbConn,
		goalSettingAssessmentDao: daos.NewGoalSettingAssessmentDB(l, dbConn),
	}
}

// GetGoalSettingAssessmentMessage service for logic
func (gs *GoalSettingAssessment) GetGoalSettingAssessmentMessage() (*dtos.GoalSettingAssessmentResponse, error) {
	gsIns := dtos.GoalSettingAssessmentResponse{}
	options := dtos.GoalSettingAssessmentOption{}

	goalSettingQuestionResponse, err := gs.goalSettingAssessmentDao.GoalSettingAssessmentQuestion()
	if err != nil {
		gs.l.Error("GetGoalSettingAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, gsQuestion := range goalSettingQuestionResponse {

		gsIns.ID = gsQuestion.ID
		gsIns.Question = gsQuestion.Question
		gsIns.QuestionNo = gsQuestion.QuestionNo
		gsIns.OptionType = gsQuestion.OptionType
		gsIns.OptionTypeLabel = gsQuestion.OptionTypeLabel
		gsIns.SequenceOrder = gsQuestion.SequenceOrder

		gsOptionResponse, err := gs.goalSettingAssessmentDao.GoalSettingAssessmentOption(gsIns.ID)
		if err != nil {
			gs.l.Error("GetGoalSettingAssessmentMessage Error - ", err)
			sentryaccounts.SentryLogExceptions(err)
			return nil, err
		}

		for _, gsOption := range gsOptionResponse {
			options.ID = gsOption.ID
			options.Name = gsOption.Name
			options.Points = gsOption.Points
			options.QuestionID = gsOption.AldGoalSettingAssessmentQuestionID
			options.SequenceOrder = gsOption.SequenceOrder
			gsIns.Options = append(gsIns.Options, options)
		}
	}

	return &gsIns, nil
}
