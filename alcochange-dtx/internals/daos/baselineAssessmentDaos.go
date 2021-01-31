package daos

import (
	"cyberliver/alcochange-dtx/dtos"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type BaselineAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewBaselineAssessmentDB(l *log.Logger, dbConn *pg.DB) *BaselineAssessment {
	return &BaselineAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// BaselineAssessmentDao interface
type BaselineAssessmentDao interface {
	BaselineAssessmentMessage() (*dtos.BaselineAssessmentResponse, error)
}

// BaselineAssessmentMessage get the terms and privacy from Database
func (ba *BaselineAssessment) BaselineAssessmentMessage() (*dtos.BaselineAssessmentResponse, error) {
	baIns := dtos.BaselineAssessmentResponse{}
	ba.dbConn.Query(&baIns.Logo, "SELECT text from constants where code = 'BLA_1' AND is_active = true")
	ba.dbConn.Query(&baIns.WelcomeNote, "SELECT text from constants where code = 'BLA_2' AND is_active = true")
	ba.dbConn.Query(&baIns.Header, "SELECT text from constants where code = 'BLA_3' AND is_active = true")
	ba.dbConn.Query(&baIns.HeaderNote, "SELECT text from constants where code = 'BLA_4' AND is_active = true")
	ba.dbConn.Query(&baIns.PatientHealth, "SELECT text from constants where code = 'BLA_5' AND is_active = true")
	ba.dbConn.Query(&baIns.Audit, "SELECT text from constants where code = 'BLA_6' AND is_active = true")
	ba.dbConn.Query(&baIns.Habit, "SELECT text from constants where code = 'BLA_7' AND is_active = true")
	ba.dbConn.Query(&baIns.Goal, "SELECT text from constants where code = 'BLA_8' AND is_active = true")
	ba.dbConn.Query(&baIns.Strategy, "SELECT text from constants where code = 'BLA_9' AND is_active = true")
	ba.dbConn.Query(&baIns.SupportiveContacts, "SELECT text from constants where code = 'BLA_10' AND is_active = true")
	ba.dbConn.Query(&baIns.EndNote, "SELECT text from constants where code = 'BLA_11' AND is_active = true")

	return &baIns, nil

}
