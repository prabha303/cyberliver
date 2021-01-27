package daos

import (
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type TermsAndPrivacy struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewTermsAndPrivacyDB(l *log.Logger, dbConn *pg.DB) *TermsAndPrivacy {
	return &TermsAndPrivacy{
		l:      l,
		dbConn: dbConn,
	}
}

// TermsAndPrivacyDao interface
type TermsAndPrivacyDao interface {
	TermsAndPrivacyMessage() (*models.AlcoChangeTermsAndPrivacy, error)
}

// TermsAndPrivacyMessage get the terms and privacy from Database
func (tp *TermsAndPrivacy) TermsAndPrivacyMessage() (*models.AlcoChangeTermsAndPrivacy, error) {
	tpIns := models.AlcoChangeTermsAndPrivacy{}
	err := tp.dbConn.Model(&tpIns).Select()
	if err != nil {
		tp.l.Error("TermsAndPrivacyMessage Error", err.Error())
		sentryaccounts.SentryLogExceptions(err)
		// return nil, err
	}
	return &tpIns, nil

}
