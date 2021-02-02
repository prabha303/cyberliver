package daos

import (
	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SaveAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewSaveAssessmentDB(l *log.Logger, dbConn *pg.DB) *SaveAssessment {
	return &SaveAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

type SaveAssessmentDao interface {
}
