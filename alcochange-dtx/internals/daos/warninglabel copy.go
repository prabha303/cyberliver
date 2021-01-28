package daos

import (
	"ecargoware/alcochange-dtx/models"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type WarningLabel struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewWarningLabelDB(l *log.Logger, dbConn *pg.DB) *WarningLabel {
	return &WarningLabel{
		l:      l,
		dbConn: dbConn,
	}
}

type WarningLabelDao interface {
	WarningLabelMessage() (*models.WarningLabel, error)
}

func (w *WarningLabel) WarningLabelMessage() (*models.WarningLabel, error) {
	warningLabel := models.WarningLabel{}
	err := w.dbConn.Model(&warningLabel).Select()
	if err != nil {
		w.l.Error("WarningLabelMessage Error", err.Error())
		//return nil, err
	}
	return &warningLabel, nil

}
