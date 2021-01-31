package daos

import (
	"cyberliver/alcochange-dtx/models"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type HealthConditionAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewHealthConditionAssessmentDB(l *log.Logger, dbConn *pg.DB) *HealthConditionAssessment {
	return &HealthConditionAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// HealthConditionAssessmentDao interface
type HealthConditionAssessmentDao interface {
	HealthConditionAssessmentQuestion() (*models.AldHealthConditionQuestion, error)
	HealthConditionAssessmentOption() ([]models.AldHealthConditionOption, error)
}

// HealthConditionAssessmentQuestion get the questions from Database
func (ba *HealthConditionAssessment) HealthConditionAssessmentQuestion() (*models.AldHealthConditionQuestion, error) {
	hcIns := models.AldHealthConditionQuestion{}

	err := ba.dbConn.Model(&hcIns).Select()
	if err != nil {
		ba.l.Error("HealthConditionAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return &hcIns, nil

}

// HealthConditionAssessmentOption get the options from Database
func (ba *HealthConditionAssessment) HealthConditionAssessmentOption() ([]models.AldHealthConditionOption, error) {
	hcIns := []models.AldHealthConditionOption{}

	err := ba.dbConn.Model(&hcIns).Select()
	if err != nil {
		ba.l.Error("HealthConditionAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return hcIns, nil

}
