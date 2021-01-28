package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type AuditAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewAuditAssessmentDB(l *log.Logger, dbConn *pg.DB) *AuditAssessment {
	return &AuditAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// AuditAssessmentDao interface
type AuditAssessmentDao interface {
	AuditAssessmentQuestion() ([]models.AldAuditAssessmentQuestion, error)
	AuditAssessmentOption(id int64) ([]models.AldAuditAssessmentOption, error)
}

// AuditAssessmentQuestion get the questions from Database
func (a *AuditAssessment) AuditAssessmentQuestion() ([]models.AldAuditAssessmentQuestion, error) {
	auditIns := []models.AldAuditAssessmentQuestion{}

	err := a.dbConn.Model(&auditIns).Select()
	if err != nil {
		a.l.Error("AuditAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return auditIns, nil

}

// AuditAssessmentOption get the options from Database
func (a *AuditAssessment) AuditAssessmentOption(id int64) ([]models.AldAuditAssessmentOption, error) {
	auditIns := []models.AldAuditAssessmentOption{}

	err := a.dbConn.Model(&auditIns).Where("ald_audit_assessment_question_id = '%d'", id).Select()
	if err != nil {
		a.l.Error("AuditAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return auditIns, nil

}
