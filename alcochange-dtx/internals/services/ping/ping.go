package ping

import (
	"errors"

	"ecargoware/alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type Ping struct {
	dbConn *pg.DB
	l      *log.Logger
	ping   *daos.Ping
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func New(l *log.Logger, dbConn *pg.DB) *Ping {
	return &Ping{
		l:      l,
		dbConn: dbConn,
		ping:   daos.NewPing(l, dbConn),
	}
}

func (p *Ping) Ping() error {
	ok, err := p.ping.Ping()
	if err != nil {
		return err
	}

	if !ok {
		return ErrUnableToPingDB
	}
	return nil
}
