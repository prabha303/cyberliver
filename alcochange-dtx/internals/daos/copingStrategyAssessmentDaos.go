package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type CopingStrategyAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewCopingStrategyAssessmentDB(l *log.Logger, dbConn *pg.DB) *CopingStrategyAssessment {
	return &CopingStrategyAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// CopingStrategyAssessmentDao interface
type CopingStrategyAssessmentDao interface {
	CopingStrategyAssessmentQuestion() ([]models.AldCopingStrategyAssessmentQuestion, error)
	CopingStrategyAssessmentOption(id int64) ([]models.AldCopingStrategyAssessmentOption, error)
}

// CopingStrategyAssessmentQuestion get the questions from Database
func (cs *CopingStrategyAssessment) CopingStrategyAssessmentQuestion() ([]models.AldCopingStrategyAssessmentQuestion, error) {
	copingStrategyIns := []models.AldCopingStrategyAssessmentQuestion{}

	err := cs.dbConn.Model(&copingStrategyIns).Select()
	if err != nil {
		cs.l.Error("CopingStrategyAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return copingStrategyIns, nil

}

// CopingStrategyAssessmentOption get the options from Database
func (cs *CopingStrategyAssessment) CopingStrategyAssessmentOption(id int64) ([]models.AldCopingStrategyAssessmentOption, error) {
	copingStrategyIns := []models.AldCopingStrategyAssessmentOption{}

	err := cs.dbConn.Model(&copingStrategyIns).Where("ald_coping_strategy_assessment_question_id = '%d'", id).Select()
	if err != nil {
		cs.l.Error("CopingStrategyAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return copingStrategyIns, nil

}
