package healthConditionService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type HealthConditionAssessment struct {
	dbConn                       *pg.DB
	l                            *log.Logger
	healthConditionAssessmentDao daos.HealthConditionAssessmentDao
}

func NewHealthConditionAssessment(l *log.Logger, dbConn *pg.DB) *HealthConditionAssessment {
	return &HealthConditionAssessment{
		l:                            l,
		dbConn:                       dbConn,
		healthConditionAssessmentDao: daos.NewHealthConditionAssessmentDB(l, dbConn),
	}
}

// GetHealthConditionAssessmentMessage service for logic
func (hc *HealthConditionAssessment) GetHealthConditionAssessmentMessage() (*dtos.HealthConditionAssessmentResponse, error) {
	hcaIns := dtos.HealthConditionAssessmentResponse{}
	options := dtos.HealthConditionAssessmentOption{}

	hCAQuestionQuestionResponse, err := hc.healthConditionAssessmentDao.HealthConditionAssessmentQuestion()
	if err != nil {
		hc.l.Error("GetHealthConditionAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	hcaIns.ID = hCAQuestionQuestionResponse.ID
	hcaIns.Question = hCAQuestionQuestionResponse.Question
	hcaIns.QuestionNo = hCAQuestionQuestionResponse.QuestionNo
	hcaIns.QuestionOptionTypeID = hCAQuestionQuestionResponse.QuestionOptionTypeID
	hcaIns.SequenceOrder = hCAQuestionQuestionResponse.SequenceOrder

	hcaOptionResponse, err := hc.healthConditionAssessmentDao.HealthConditionAssessmentOption()
	if err != nil {
		hc.l.Error("GetHealthConditionAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, hcaOption := range hcaOptionResponse {
		options.ID = hcaOption.ID
		options.Name = hcaOption.Name
		options.Points = hcaOption.Points
		options.QuestionID = hcaOption.AldHealthConditionQuestionID
		options.SequenceOrder = hcaOption.SequenceOrder
		hcaIns.Options = append(hcaIns.Options, options)
	}

	return &hcaIns, nil
}
