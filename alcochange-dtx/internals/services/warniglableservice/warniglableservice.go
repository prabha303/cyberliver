package warniglableservice

import (
	"errors"

	"ecargoware/alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

//DFMMMyyyyNoSep := "Jan2006"

type WarningLabel struct {
	dbConn          *pg.DB
	l               *log.Logger
	warningLabelDao daos.WarningLabelDao
}

type WarniglableResponse struct {
	EuRepresentative string `json:"euRepresentative"`
	RefVersion       string `json:"version"`
	Logo             string `json:"logo"`
	Manufacturer     string `json:"manufacturer"`
	WarningLink      string `json:"warningLink"`
	IndicationsLink  string `json:"indicationsLink"`
	ManufacturerDate string `json:"manufacturerDate"`
	UpdatedDate      string `json:"updatedDate"`
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func NewWarning(l *log.Logger, dbConn *pg.DB) *WarningLabel {
	return &WarningLabel{
		l:               l,
		dbConn:          dbConn,
		warningLabelDao: daos.NewWarningLabelDB(l, dbConn),
	}
}

func (w *WarningLabel) GetWarniglableMessage() (*WarniglableResponse, error) {
	warniglableResponse := WarniglableResponse{}
	response, err := w.warningLabelDao.WarningLabelMessage()
	if err != nil {
		w.l.Error("WarniglableMessage Error - ", err)
		return nil, err
	}
	warniglableResponse.EuRepresentative = response.EuRepresentative
	warniglableResponse.IndicationsLink = response.IndicationsLink
	warniglableResponse.Logo = response.Logo
	warniglableResponse.Manufacturer = response.Manufacturer
	warniglableResponse.ManufacturerDate = ""
	warniglableResponse.UpdatedDate = ""
	warniglableResponse.RefVersion = response.RefVersion
	warniglableResponse.WarningLink = response.WarningLink
	return &warniglableResponse, nil
}
