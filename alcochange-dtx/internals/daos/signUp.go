package daos

import (
	"ecargoware/alcochange-dtx/models"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SignUp struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewSignUp(l *log.Logger, dbConn *pg.DB) *SignUp {
	return &SignUp{
		l:      l,
		dbConn: dbConn,
	}
}

type SignUpDao interface {
	EmailIDExists(email string) bool
}

func (sp *SignUp) EmailIDExists(email string) bool {
	users := []models.Users{}
	sp.dbConn.Model(&users).Where("LOWER(email_id) = LOWER(?)", email).Select()
	return len(users) > 0
}
