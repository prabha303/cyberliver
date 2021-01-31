package baselineAssessmentService

import (
	"cyberliver/alcochange-dtx/conf"
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type BaselineAssessment struct {
	dbConn                *pg.DB
	l                     *log.Logger
	baselineAssessmentDao daos.BaselineAssessmentDao
}

func NewBaselineAssessment(l *log.Logger, dbConn *pg.DB) *BaselineAssessment {
	return &BaselineAssessment{
		l:                     l,
		dbConn:                dbConn,
		baselineAssessmentDao: daos.NewBaselineAssessmentDB(l, dbConn),
	}
}

// GetBaselineAssessmentMessage service for logic
func (ba *BaselineAssessment) GetBaselineAssessmentMessage() (*dtos.BaselineAssessmentResponse, error) {
	baselineAssessmentResponse, err := ba.baselineAssessmentDao.BaselineAssessmentMessage()
	if err != nil {
		ba.l.Error("GetBaselineAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	baselineAssessmentResponse.ButtonText = conf.BaselineAssessmentButtonText

	return baselineAssessmentResponse, nil
}
