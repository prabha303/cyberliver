package daos

import (
	"cyberliver/alcochange-dtx/models"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type TriggerAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewTriggerAssessmentDB(l *log.Logger, dbConn *pg.DB) *TriggerAssessment {
	return &TriggerAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// TriggerAssessmentDao interface
type TriggerAssessmentDao interface {
	TriggerAssessmentQuestion() ([]models.AldTriggerAssessmentQuestion, error)
	TriggerAssessmentOption(id int64) ([]models.AldTriggerAssessmentOption, error)
}

// TriggerAssessmentQuestion get the questions from Database
func (t *TriggerAssessment) TriggerAssessmentQuestion() ([]models.AldTriggerAssessmentQuestion, error) {
	triggerIns := []models.AldTriggerAssessmentQuestion{}

	err := t.dbConn.Model(&triggerIns).Select()
	if err != nil {
		t.l.Error("TriggerAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return triggerIns, nil

}

// TriggerAssessmentOption get the options from Database
func (t *TriggerAssessment) TriggerAssessmentOption(id int64) ([]models.AldTriggerAssessmentOption, error) {
	triggerIns := []models.AldTriggerAssessmentOption{}

	err := t.dbConn.Model(&triggerIns).Where("ald_trigger_assessment_question_id = '%d'", id).Select()
	if err != nil {
		t.l.Error("TriggerAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return triggerIns, nil

}
