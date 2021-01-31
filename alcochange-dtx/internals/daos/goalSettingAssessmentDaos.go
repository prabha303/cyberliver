package daos

import (
	"cyberliver/alcochange-dtx/models"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type GoalSettingAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewGoalSettingAssessmentDB(l *log.Logger, dbConn *pg.DB) *GoalSettingAssessment {
	return &GoalSettingAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// GoalSettingAssessmentDao interface
type GoalSettingAssessmentDao interface {
	GoalSettingAssessmentQuestion() ([]models.AldGoalSettingAssessmentQuestion, error)
	GoalSettingAssessmentOption(id int64) ([]models.AldGoalSettingAssessmentOption, error)
}

// GoalSettingAssessmentQuestion get the questions from Database
func (a *GoalSettingAssessment) GoalSettingAssessmentQuestion() ([]models.AldGoalSettingAssessmentQuestion, error) {
	gsIns := []models.AldGoalSettingAssessmentQuestion{}

	err := a.dbConn.Model(&gsIns).Select()
	if err != nil {
		a.l.Error("GoalSettingAssessmentQuestion Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return gsIns, nil

}

// GoalSettingAssessmentOption get the options from Database
func (a *GoalSettingAssessment) GoalSettingAssessmentOption(id int64) ([]models.AldGoalSettingAssessmentOption, error) {
	gsIns := []models.AldGoalSettingAssessmentOption{}

	err := a.dbConn.Model(&gsIns).Where("ald_goal_setting__assessment_question_id = '%d'", id).Select()
	if err != nil {
		a.l.Error("GoalSettingAssessmentOption Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return gsIns, nil

}
