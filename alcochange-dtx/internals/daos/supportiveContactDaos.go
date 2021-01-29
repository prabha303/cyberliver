package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SupportiveContactAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewSupportiveContactDB(l *log.Logger, dbConn *pg.DB) *SupportiveContactAssessment {
	return &SupportiveContactAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

// SupportiveContactDao interface
type SupportiveContactDao interface {
	GetSupportiveContactList() ([]models.AldRelationShip, error)
}

// GetSupportiveContactList get the questions from Database
func (sc *SupportiveContactAssessment) GetSupportiveContactList() ([]models.AldRelationShip, error) {
	relationShipIns := []models.AldRelationShip{}

	err := sc.dbConn.Model(&relationShipIns).Select()
	if err != nil {
		sc.l.Error("GetSupportiveContactList Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		//return nil, err
	}

	return relationShipIns, nil

}
