package saveAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SaveAssessment struct {
	dbConn            *pg.DB
	l                 *log.Logger
	saveAssessmentDao daos.SaveAssessmentDao
}

func NewSaveAssessment(l *log.Logger, dbConn *pg.DB) *SaveAssessment {
	return &SaveAssessment{
		l:                 l,
		dbConn:            dbConn,
		saveAssessmentDao: daos.NewSaveAssessmentDB(l, dbConn),
	}
}

// SaveAssessmentDetails service for logic
func (s *SaveAssessment) SaveAssessmentDetails(req dtos.SaveAssessmentRequest, userID int64) error {

	hErr := s.HealthConditionAssessment(req.HealthConditionAssessmentAnswer, userID)
	if hErr != nil {
		s.l.Error("SaveAssessmentDetails Error--", hErr)
		sentryaccounts.SentryLogExceptions(hErr)
		return hErr
	}

	aErr := s.AuditAssessment(req.AuditAssessmentAnswer, userID)
	if hErr != nil {
		s.l.Error("SaveAssessmentDetails Error--", aErr)
		sentryaccounts.SentryLogExceptions(aErr)
		return hErr
	}

	gErr := s.GoalSettingAssessment(req.GoalSettingAssessmentAnswer, userID)
	if gErr != nil {
		s.l.Error("SaveAssessmentDetails Error--", gErr)
		sentryaccounts.SentryLogExceptions(gErr)
		return hErr
	}

	dErr := s.DrinkAssessment(req.DrinkHabitAssessmentAnswer, userID)
	if gErr != nil {
		s.l.Error("SaveAssessmentDetails Error--", dErr)
		sentryaccounts.SentryLogExceptions(dErr)
		return dErr
	}

	sErr := s.SupportiveContactAssessment(req.SupportiveContact, userID)
	if sErr != nil {
		s.l.Error("SaveAssessmentDetails Error--", sErr)
		sentryaccounts.SentryLogExceptions(sErr)
		return sErr
	}

	return nil
}
