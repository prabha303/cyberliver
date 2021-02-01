package triggerAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type TriggerAssessment struct {
	dbConn               *pg.DB
	l                    *log.Logger
	triggerAssessmentDao daos.TriggerAssessmentDao
}

func NewTriggerAssessment(l *log.Logger, dbConn *pg.DB) *TriggerAssessment {
	return &TriggerAssessment{
		l:                    l,
		dbConn:               dbConn,
		triggerAssessmentDao: daos.NewTriggerAssessmentDB(l, dbConn),
	}
}

// GetTriggerAssessmentMessage service for logic
func (t *TriggerAssessment) GetTriggerAssessmentMessage() (*dtos.TriggerAssessmentResponse, error) {
	triggerIns := dtos.TriggerAssessmentResponse{}
	options := dtos.TriggerAssessmentOption{}

	triggerQuestionResponse, err := t.triggerAssessmentDao.TriggerAssessmentQuestion()
	if err != nil {
		t.l.Error("GetTriggerAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, triggerQuestion := range triggerQuestionResponse {

		triggerIns.ID = triggerQuestion.ID
		triggerIns.Question = triggerQuestion.Question
		triggerIns.QuestionNo = triggerQuestion.QuestionNo
		// triggerIns.OptionType = triggerQuestion.OptionType
		// triggerIns.OptionTypeLabel = triggerQuestion.OptionTypeLabel
		triggerIns.SequenceOrder = triggerQuestion.SequenceOrder

		triggerOptionResponse, err := t.triggerAssessmentDao.TriggerAssessmentOption(triggerIns.ID)
		if err != nil {
			t.l.Error("GetTriggerAssessmentMessage Error - ", err)
			sentryaccounts.SentryLogExceptions(err)
			return nil, err
		}

		for _, triggerOption := range triggerOptionResponse {
			options.ID = triggerOption.ID
			options.Name = triggerOption.Name
			options.Points = triggerOption.Points
			options.QuestionID = triggerOption.AldTriggerAssessmentQuestionID
			options.SequenceOrder = triggerOption.SequenceOrder
			triggerIns.Options = append(triggerIns.Options, options)
		}
	}

	return &triggerIns, nil
}
