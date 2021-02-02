package saveAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"

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
func (s *SaveAssessment) SaveAssessmentDetails(req dtos.SaveAssessmentRequest) error {

	return nil
}
