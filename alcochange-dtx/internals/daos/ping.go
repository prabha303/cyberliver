package daos

import (
	"fmt"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type PingCheckStruct struct {
	Count int32 `json:"count"`
}

type Ping struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewPing(l *log.Logger, dbConn *pg.DB) *Ping {
	return &Ping{
		l:      l,
		dbConn: dbConn,
	}
}

func (p *Ping) Ping() (bool, error) {
	pingModel := PingCheckStruct{}
	//start := time.Now()
	checkDB := fmt.Sprintf(`SELECT count(1)`)
	_, err := p.dbConn.Query(&pingModel, checkDB)
	//spRuntime := time.Since(start).Seconds()
	//p.l.Debug("accounts Ping run time", spRuntime)
	return (pingModel.Count == 1), err
	//return true, nil
}
