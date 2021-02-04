package copingStrategyAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type CopingStrategyAssessment struct {
	dbConn                      *pg.DB
	l                           *log.Logger
	copingStrategyAssessmentDao daos.CopingStrategyAssessmentDao
}

func NewCopingStrategyAssessment(l *log.Logger, dbConn *pg.DB) *CopingStrategyAssessment {
	return &CopingStrategyAssessment{
		l:                           l,
		dbConn:                      dbConn,
		copingStrategyAssessmentDao: daos.NewCopingStrategyAssessmentDB(l, dbConn),
	}
}

// GetCopingStrategyAssessmentMessage service for logic
func (cs *CopingStrategyAssessment) GetCopingStrategyAssessmentMessage() (*[]dtos.CopingStrategyAssessmentResponse, error) {
	copyStrIns := []dtos.CopingStrategyAssessmentResponse{}
	options := dtos.CopingStrategyAssessmentOption{}

	copingStrategyQuestionResponse, err := cs.copingStrategyAssessmentDao.CopingStrategyAssessmentQuestion()
	if err != nil {
		cs.l.Error("GetCopingStrategyAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, csQuestion := range copingStrategyQuestionResponse {
		csIns := dtos.CopingStrategyAssessmentResponse{}

		csIns.ID = csQuestion.ID
		csIns.Question = csQuestion.Question
		csIns.QuestionNo = csQuestion.QuestionNo
		csIns.QuestionOptionTypeID = csQuestion.QuestionOptionTypeID
		csIns.SequenceOrder = csQuestion.SequenceOrder

		csOptionResponse, err := cs.copingStrategyAssessmentDao.CopingStrategyAssessmentOption(csQuestion.QuestionNo)
		if err != nil {
			cs.l.Error("GetCopingStrategyAssessmentMessage Error - ", err)
			sentryaccounts.SentryLogExceptions(err)
			return nil, err
		}

		for _, csOption := range csOptionResponse {
			options.ID = csOption.ID
			options.Name = csOption.Name
			options.Points = csOption.Points
			options.MaxPoints = csOption.MaxPoints
			options.QuestionID = csOption.AldCopingStrategyAssessmentQuestionID
			options.SequenceOrder = csOption.SequenceOrder
			csIns.Options = append(csIns.Options, options)
		}

		copyStrIns = append(copyStrIns, csIns)

	}

	return &copyStrIns, nil
}
