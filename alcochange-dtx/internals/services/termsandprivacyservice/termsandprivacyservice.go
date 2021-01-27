package termsandprivacyservice

import (
	"ecargoware/alcochange-dtx/conf"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type TermsAndPrivacy struct {
	dbConn             *pg.DB
	l                  *log.Logger
	termsAndPrivacyDao daos.TermsAndPrivacyDao
}

func NewTermsAndPrivacy(l *log.Logger, dbConn *pg.DB) *TermsAndPrivacy {
	return &TermsAndPrivacy{
		l:                  l,
		dbConn:             dbConn,
		termsAndPrivacyDao: daos.NewTermsAndPrivacyDB(l, dbConn),
	}
}

// TermsAndPrivacyResponse Response struct
type TermsAndPrivacyResponse struct {
	VersionInfo  string `json:"versionInfo"`
	Instructions string `json:"instructions"`
	Contents     string `json:"contents"`
	ButtonText   string `json:"buttonText"`
	Logo         string `json:"logo"`
}

// GetTermsAndPrivacyMessage service for logic
func (tp *TermsAndPrivacy) GetTermsAndPrivacyMessage() (*TermsAndPrivacyResponse, error) {
	termsAndPrivacyResponse := TermsAndPrivacyResponse{}
	response, err := tp.termsAndPrivacyDao.TermsAndPrivacyMessage()
	if err != nil {
		tp.l.Error("GetTermsAndPrivacyMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	termsAndPrivacyResponse.VersionInfo = response.VersionInfo
	termsAndPrivacyResponse.Instructions = response.Instructions
	termsAndPrivacyResponse.Contents = response.Contents
	termsAndPrivacyResponse.ButtonText = conf.WarningAndPrivacyButtonText
	termsAndPrivacyResponse.Logo = response.Logo

	return &termsAndPrivacyResponse, nil
}
