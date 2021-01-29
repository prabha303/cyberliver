package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type ReasonAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewReasonAssessmentDB(l *log.Logger, dbConn *pg.DB) *ReasonAssessment {
	return &ReasonAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// ReasonAssessmentDao interface
type ReasonAssessmentDao interface {
	ReasonAssessmentQuestion() (models.AldReasonAssessmentQuestion, error)
	ReasonAssessmentOption() ([]models.AldReasonAssessmentOption, error)
}

// ReasonAssessmentQuestion get the questions from Database
func (r *ReasonAssessment) ReasonAssessmentQuestion() (models.AldReasonAssessmentQuestion, error) {
	reasonIns := models.AldReasonAssessmentQuestion{}

	err := r.dbConn.Model(&reasonIns).Select()
	if err != nil {
		r.l.Error("ReasonAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return reasonIns, nil

}

// ReasonAssessmentOption get the options from Database
func (r *ReasonAssessment) ReasonAssessmentOption() ([]models.AldReasonAssessmentOption, error) {
	reasonIns := []models.AldReasonAssessmentOption{}

	err := r.dbConn.Model(&reasonIns).Select()
	if err != nil {
		r.l.Error("ReasonAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return reasonIns, nil

}
