package daos

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type PatientAccessCode struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewPatientAccessCodeDB(l *log.Logger, dbConn *pg.DB) *PatientAccessCode {
	return &PatientAccessCode{
		l:      l,
		dbConn: dbConn,
	}
}

// PatientAccessCodeDao interface
type PatientAccessCodeDao interface {
	GetPatientByAccessCode(req dtos.PatientAccessCodeReq) error
}

func (pac *PatientAccessCode) GetPatientByAccessCode(req dtos.PatientAccessCodeReq) error {

	patAccCodeIns := &models.PatientAccessCode{}
	if err := pac.dbConn.Model(patAccCodeIns).Where("access_code = '%v' AND solution_type = '%v' AND is_redeemed = false", req.AccessCode, req.SolutionType).Select(); err != nil {
		pac.l.Error("GetPatientByAccessCode Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return err
	}

	return nil
}
