package warniglableservice

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"errors"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

//DFMMMyyyyNoSep := "Jan2006"

type SignUp struct {
	dbConn    *pg.DB
	l         *log.Logger
	signUpDao daos.SignUpDao
}

func NewSignUp(l *log.Logger, dbConn *pg.DB) *SignUp {
	return &SignUp{
		l:         l,
		dbConn:    dbConn,
		signUpDao: daos.NewSignUp(l, dbConn),
	}
}

func (sp *SignUp) UserSignUp(signReq dtos.SignUpRequest) (*dtos.SignUpResponse, error) {

	isExists := sp.signUpDao.EmailIDExists(signReq.RegisterUserRequest.EmailID)
	if isExists {
		return nil, errors.New("EmailID Already Exists")
	}

	signUp := dtos.SignUpResponse{}

	return &signUp, nil
}
