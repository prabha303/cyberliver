package supportiveContactService

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SupportiveContact struct {
	dbConn               *pg.DB
	l                    *log.Logger
	supportiveContactDao daos.SupportiveContactDao
}

func NewSupportiveContact(l *log.Logger, dbConn *pg.DB) *SupportiveContact {
	return &SupportiveContact{
		l:                    l,
		dbConn:               dbConn,
		supportiveContactDao: daos.NewSupportiveContactDB(l, dbConn),
	}
}

// GetSupportiveContact service for logic
func (sc *SupportiveContact) GetSupportiveContact() (*dtos.RelationShipResponse, error) {
	relationShipIns := dtos.RelationShipResponse{}

	supportiveContactResponse, err := sc.supportiveContactDao.GetSupportiveContactList()
	if err != nil {
		sc.l.Error("GetSupportiveContact Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, supportiveContact := range supportiveContactResponse {

		relationShipIns.ID = supportiveContact.ID
		relationShipIns.Name = supportiveContact.Name
	}

	return &relationShipIns, nil
}
