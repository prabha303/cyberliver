package patientaccesscodeservice

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/daos"
	"cyberliver/alcochange-dtx/sentryaccounts"
	"errors"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type PatientAccessCode struct {
	dbConn               *pg.DB
	l                    *log.Logger
	patientAccessCodeDao daos.PatientAccessCodeDao
}

var (
	ErrUnableToPingDB               = errors.New("Unable to ping database")
	PatientAccessCodeSuccessMessage = "Access Code Successfully Verified..."
)

func NewPatientAccessCode(l *log.Logger, dbConn *pg.DB) *PatientAccessCode {
	return &PatientAccessCode{
		l:                    l,
		dbConn:               dbConn,
		patientAccessCodeDao: daos.NewPatientAccessCodeDB(l, dbConn),
	}
}

func (pac *PatientAccessCode) VerifyPatientByAccessCode(req dtos.PatientAccessCodeReq) (*dtos.PatientAccessCodeResponse, error) {
	pacResIns := dtos.PatientAccessCodeResponse{}
	if req.AccessCode == "" || req.SolutionType == "" {
		return nil, errors.New("Invalid json")
	}
	err := pac.patientAccessCodeDao.GetPatientByAccessCode(req)
	if err != nil {
		pac.l.Error("VerifyPatientByAccessCode Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}
	pacResIns.Message = PatientAccessCodeSuccessMessage
	return &pacResIns, nil
}
